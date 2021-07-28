package ci

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	pb "github.com/autograde/quickfeed/ag"
	"github.com/autograde/quickfeed/database"
	"github.com/autograde/quickfeed/kit/score"
	"github.com/autograde/quickfeed/log"
	"go.uber.org/zap"
)

const (
	// qf101 is a test course for experimenting with go.sh behavior.
	// The test below will run locally on the test machine, not on the QuickFeed machine.
	getURL  = "https://github.com/qf101/meling-labs.git"
	testURL = "https://github.com/qf101/tests.git"
)

func TestRunTests(t *testing.T) {
	// The access token is a 'personal access token' for the user that has access to the repos below.
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if len(accessToken) < 1 {
		t.Skip("This test requires a 'GITHUB_ACCESS_TOKEN' and access to the 'autograder-test' GitHub organization")
	}

	randomness := make([]byte, 10)
	_, err := rand.Read(randomness)
	if err != nil {
		t.Fatal(err)
	}
	randomString := fmt.Sprintf("%x", sha1.Sum(randomness))
	info := &AssignmentInfo{
		AssignmentName:     "lab1",
		Script:             "go.sh",
		CreatorAccessToken: accessToken,
		GetURL:             getURL,
		TestURL:            testURL,
		RandomSecret:       randomString,
	}
	runData := &RunData{
		Course: &pb.Course{Code: "DAT320"},
		Assignment: &pb.Assignment{
			Name:             info.AssignmentName,
			ContainerTimeout: 1,
		},
		Repo:     &pb.Repository{},
		JobOwner: "muggles",
	}

	runner, err := NewDockerCI(log.Zap(true))
	if err != nil {
		t.Fatal(err)
	}
	defer runner.Close()
	ed, err := runTests("scripts", runner, info, runData)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n%s\nExecTime: %v\nSecret: %v\n", ed.out, ed.execTime, info.RandomSecret)
}

func TestRecordResults(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "testing")
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}

	db, err := database.NewGormDB(f.Name(), zap.NewNop())
	if err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}

	// clean up the temp database file
	defer func() {
		if err := os.Remove(f.Name()); err != nil {
			t.Error(err)
		}
	}()

	var user pb.User
	if err := db.CreateUserFromRemoteIdentity(&user,
		&pb.RemoteIdentity{
			Provider:    "fake",
			RemoteID:    1,
			AccessToken: "token",
		}); err != nil {
		t.Fatal(err)
	}

	if err := db.UpdateUser(&pb.User{ID: user.ID, IsAdmin: true}); err != nil {
		t.Fatal(err)
	}

	course := &pb.Course{
		Name:            "Test",
		OrganizationID:  1,
		CourseCreatorID: user.ID,
		SlipDays:        5,
	}
	if err := db.CreateCourse(user.ID, course); err != nil {
		t.Fatal(err)
	}

	assignment := &pb.Assignment{
		CourseID:         course.ID,
		Name:             "lab1",
		ScriptFile:       "go.sh",
		Deadline:         "2022-11-11T13:00:00",
		AutoApprove:      true,
		ScoreLimit:       70,
		Order:            1,
		IsGroupLab:       false,
		ContainerTimeout: 1,
	}
	if err = db.CreateAssignment(assignment); err != nil {
		t.Fatal(err)
	}

	buildInfo := &score.BuildInfo{
		BuildDate: "2022-11-10T13:00:00",
		BuildLog:  "Testing",
		ExecTime:  33333,
	}
	testScores := []*score.Score{
		{
			Secret:   "secret",
			TestName: "Test",
			Score:    10,
			MaxScore: 15,
			Weight:   1,
		},
	}

	// Must create a new submission with correct scores and build info, not approved
	results := &score.Results{
		BuildInfo: buildInfo,
		Scores:    testScores,
	}
	runData := &RunData{
		Course:     course,
		Assignment: assignment,
		Repo: &pb.Repository{
			UserID: 1,
		},
		JobOwner: "test",
	}

	recordResults(zap.NewNop().Sugar(), db, runData, results)

	submission, err := db.GetSubmission(&pb.Submission{AssignmentID: assignment.ID, UserID: user.ID})
	if err != nil {
		t.Fatal(err)
	}
	if submission.Status == pb.Submission_APPROVED {
		t.Fatal("Submission must not be auto approved")
	}
	if !reflect.DeepEqual(testScores, submission.Scores) {
		t.Fatalf("Incorrect submission scores. Want: %+v, got %+v", testScores, submission.Scores)
	}
	if !reflect.DeepEqual(buildInfo.BuildDate, submission.BuildInfo.BuildDate) {
		t.Fatalf("Incorrect build date. Want: %s, got %s", buildInfo.BuildDate, submission.BuildInfo.BuildDate)
	}

	// Updating submission after deadline: build info and slip days must be updated
	newBuildDate := "2022-11-12T13:00:00"
	results.BuildInfo.BuildDate = newBuildDate
	recordResults(zap.NewNop().Sugar(), db, runData, results)

	enrollment, err := db.GetEnrollmentByCourseAndUser(course.ID, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	if enrollment.RemainingSlipDays(course) == int32(course.SlipDays) || len(enrollment.UsedSlipDays) < 1 {
		t.Fatal("Student must have reduced slip days")
	}
	updatedSubmission, err := db.GetSubmission(&pb.Submission{AssignmentID: assignment.ID, UserID: user.ID})
	if err != nil {
		t.Fatal(err)
	}
	if updatedSubmission.BuildInfo.BuildDate != newBuildDate {
		t.Fatalf("Incorrect build date: want %s, got %s", newBuildDate, updatedSubmission.BuildInfo.BuildDate)
	}

	// Rebuilding after deadline: delivery date and slip days must stay unchanged
	runData.Rebuild = true
	results.BuildInfo.BuildDate = "2022-11-13T13:00:00"
	slipDaysBeforeUpdate := enrollment.RemainingSlipDays(course)
	recordResults(zap.NewNop().Sugar(), db, runData, results)
	updatedEnrollment, err := db.GetEnrollmentByCourseAndUser(course.ID, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	if updatedEnrollment.RemainingSlipDays(course) != slipDaysBeforeUpdate {
		t.Fatalf("Incorrect number of slip days: expected %d, got %d", slipDaysBeforeUpdate, updatedEnrollment.RemainingSlipDays(course))
	}
}
