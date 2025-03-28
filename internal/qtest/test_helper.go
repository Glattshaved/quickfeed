package qtest

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/go-cmp/cmp"
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

// CreateFakeUser is a test helper to create a user in the database.
func CreateFakeUser(t *testing.T, db database.Database) *qf.User {
	t.Helper()
	user := &qf.User{}
	if err := db.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	return user
}

func CreateFakeCustomUser(t *testing.T, db database.Database, user *qf.User) *qf.User {
	t.Helper()
	if err := db.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	return user
}

// CreateCourse is a test helper to create a course in the database; it updates the course with the ID.
func CreateCourse(t *testing.T, db database.Database, user *qf.User, course *qf.Course) {
	t.Helper()
	if err := db.CreateCourse(user.GetID(), course); err != nil {
		t.Fatal(err)
	}
}

func CreateAssignment(t *testing.T, db database.Database, assignment *qf.Assignment) {
	t.Helper()
	if err := db.CreateAssignment(assignment); err != nil {
		t.Fatal(err)
	}
}

func CreateGroup(t *testing.T, db database.Database, course *qf.Course, groupSize int) *qf.Group {
	t.Helper()
	var users []*qf.User
	for range groupSize {
		user := CreateFakeUser(t, db)
		users = append(users, user)
		EnrollStudent(t, db, user, course)
	}
	group := &qf.Group{
		CourseID: course.ID,
		Name:     "group " + RandomString(t),
		Users:    users,
	}
	if err := db.CreateGroup(group); err != nil {
		t.Fatal(err)
	}
	return group
}

func GetEnrollment(t *testing.T, db database.Database, userID, courseID uint64) *qf.Enrollment {
	t.Helper()
	enrollment, err := db.GetEnrollmentByCourseAndUser(courseID, userID)
	if err != nil {
		t.Fatal(err)
	}
	return enrollment
}

func EnrollStudent(t *testing.T, db database.Database, student *qf.User, course *qf.Course) {
	t.Helper()
	query := &qf.Enrollment{
		UserID:   student.GetID(),
		CourseID: course.GetID(),
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
		UserID:   student.GetID(),
		CourseID: course.GetID(),
	}
	if err := db.CreateEnrollment(query); err != nil {
		t.Fatal(err)
	}
	query.Status = qf.Enrollment_TEACHER
	if err := db.UpdateEnrollment(query); err != nil {
		t.Fatal(err)
	}
}

func EnrollUser(t *testing.T, db database.Database, user *qf.User, course *qf.Course, status qf.Enrollment_UserStatus) *qf.Enrollment {
	t.Helper()
	enrollment := &qf.Enrollment{
		UserID:   user.GetID(),
		CourseID: course.GetID(),
	}
	if err := db.CreateEnrollment(enrollment); err != nil {
		t.Fatal(err)
	}
	enrollment.Status = status
	if err := db.UpdateEnrollment(enrollment); err != nil {
		t.Fatal(err)
	}
	return enrollment
}

func RandomString(t *testing.T) string {
	t.Helper()
	randomness := make([]byte, 10)
	if _, err := rand.Read(randomness); err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x", sha256.Sum256(randomness))[:6]
}

func RequestWithCookie[T any](message *T, cookie string) *connect.Request[T] {
	request := connect.NewRequest(message)
	request.Header().Set("cookie", cookie)
	return request
}

// Ptr returns a pointer to the given value.
//
// How to use:
//   - Use this function to create a pointer to a value.
//   - This function is useful when initializing a struct with a pointer field.
//
// Example:
//
//	type MyStruct struct {
//		Field *int
//	    Src   *string
//	}
//	myStruct := MyStruct{
//		Field: Ptr(10),
//		Src:   Ptr("hello"),
//	}
func Ptr[T any](t T) *T {
	return &t
}

// Diff compares the got and want values and prints a diff with the given message.
func Diff(t *testing.T, msg string, got, want any, opts ...cmp.Option) {
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Errorf("%s: (-got +want)\n%s", msg, diff)
	}
}
