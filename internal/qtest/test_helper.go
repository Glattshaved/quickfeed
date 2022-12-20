package qtest

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/quickfeed/quickfeed/database"
	"github.com/quickfeed/quickfeed/qf"
)

// TestDB returns a test database and close function.
// This function should only be used as a test helper.
func TestDB(t *testing.T) (database.Database, func()) {
	t.Helper()

	f, err := os.CreateTemp(t.TempDir(), "test.db")
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}

	db, err := database.NewGormDB(f.Name(), Logger(t).Desugar())
	if err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}

	return db, func() {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
		if err := os.Remove(f.Name()); err != nil {
			t.Error(err)
		}
	}
}

// CreateFakeUser is a test helper to create a user in the database
// with the given remote id and the fake scm provider.
func CreateFakeUser(t *testing.T, db database.Database, remoteID uint64) *qf.User {
	t.Helper()
	user := &qf.User{
		ScmRemoteID:  remoteID,
		RefreshToken: "token",
	}
	if err := db.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	return user
}

func CreateNamedUser(t *testing.T, db database.Database, remoteID uint64, name string) *qf.User {
	t.Helper()
	user := &qf.User{
		Name:         name,
		Login:        name,
		ScmRemoteID:  remoteID,
		RefreshToken: "refresh_token",
	}
	if err := db.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	return user
}

func CreateCourse(t *testing.T, db database.Database, user *qf.User, course *qf.Course) {
	t.Helper()
	if err := db.CreateCourse(user.ID, course); err != nil {
		t.Fatal(err)
	}
}

func EnrollStudent(t *testing.T, db database.Database, student *qf.User, course *qf.Course) {
	t.Helper()
	query := &qf.Enrollment{
		UserID:   student.ID,
		CourseID: course.ID,
	}
	if err := db.CreateEnrollment(query); err != nil {
		t.Fatal(err)
	}
	query.Status = qf.Enrollment_STUDENT
	if err := db.UpdateEnrollment(query); err != nil {
		t.Fatal(err)
	}
}

func EnrollTeacher(t *testing.T, db database.Database, student *qf.User, course *qf.Course) {
	t.Helper()
	query := &qf.Enrollment{
		UserID:   student.ID,
		CourseID: course.ID,
	}
	if err := db.CreateEnrollment(query); err != nil {
		t.Fatal(err)
	}
	query.Status = qf.Enrollment_TEACHER
	if err := db.UpdateEnrollment(query); err != nil {
		t.Fatal(err)
	}
}

func RandomString(t *testing.T) string {
	t.Helper()
	randomness := make([]byte, 10)
	if _, err := rand.Read(randomness); err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x", sha256.Sum256(randomness))[:6]
}

// AssignmentsWithTasks returns a list of test assignments with tasks for the given course.
func AssignmentsWithTasks(t *testing.T, courseID uint64) []*qf.Assignment {
	return []*qf.Assignment{
		{
			CourseID:    courseID,
			Name:        "lab1",
			Deadline:    Timestamp(t, "2022-12-01T19:00:00"),
			AutoApprove: false,
			Order:       1,
			IsGroupLab:  false,
			Tasks: []*qf.Task{
				{Title: "Fibonacci", Name: "fib", AssignmentOrder: 1, Body: "Implement fibonacci"},
				{Title: "Lucas Numbers", Name: "luc", AssignmentOrder: 1, Body: "Implement lucas numbers"},
			},
		},
		{
			CourseID:    courseID,
			Name:        "lab2",
			Deadline:    Timestamp(t, "2022-12-12T19:00:00"),
			AutoApprove: false,
			Order:       2,
			IsGroupLab:  false,
			Tasks: []*qf.Task{
				{Title: "Addition", Name: "add", AssignmentOrder: 2, Body: "Implement addition"},
				{Title: "Subtraction", Name: "sub", AssignmentOrder: 2, Body: "Implement subtraction"},
				{Title: "Multiplication", Name: "mul", AssignmentOrder: 2, Body: "Implement multiplication"},
			},
		},
	}
}

func RequestWithCookie[T any](message *T, cookie string) *connect.Request[T] {
	request := connect.NewRequest(message)
	request.Header().Set("cookie", cookie)
	return request
}
