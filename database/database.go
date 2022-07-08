package database

import "github.com/quickfeed/quickfeed/qf"

// Database contains methods for manipulating the database.
type Database interface {
	// CreateUserFromRemoteIdentity creates new user record from remote identity, sets user with ID 1 as admin.
	CreateUserFromRemoteIdentity(*qf.User, *qf.RemoteIdentity) error
	// AssociateUserWithRemoteIdentity associates user with the given remote identity.
	AssociateUserWithRemoteIdentity(userID uint64, provider string, remoteID uint64, accessToken string) error
	// UpdateAccessToken updates the access token for the given remote identity.
	// The supplied remote identity must contain Provider, RemoteID and AccessToken.
	UpdateAccessToken(*qf.RemoteIdentity) error
	// GetUserByRemoteIdentity returns the user for the given remote identity.
	// The supplied remote identity must contain Provider and RemoteID.
	GetUserByRemoteIdentity(*qf.RemoteIdentity) (*qf.User, error)

	// GetUser returns the given user, including remote identities.
	GetUser(uint64) (*qf.User, error)
	// GetUserByCourse returns the owner of the given login
	// with preloaded course matching the given query.
	GetUserByCourse(*qf.Course, string) (*qf.User, *qf.Course, error)
	// GetUserWithEnrollments returns the user by ID with preloaded user enrollments.
	GetUserWithEnrollments(uint64) (*qf.User, error)
	// GetUsers returns the users for the given set of user IDs.
	GetUsers(...uint64) ([]*qf.User, error)
	// UpdateUser updates the user's details, excluding remote identities.
	UpdateUser(*qf.User) error

	// CreateCourse creates a new course if user with given ID is admin, enrolls user as course teacher.
	CreateCourse(uint64, *qf.Course) error
	// GetCourse fetches course by ID. If withInfo is true, preloads course
	// assignments, active enrollments and groups.
	GetCourse(uint64, bool) (*qf.Course, error)
	// GetCourseByOrganizationID fetches course by organization ID.
	GetCourseByOrganizationID(organizationID uint64) (*qf.Course, error)
	// GetCourses returns a list of courses. If one or more course IDs are provided,
	// the corresponding courses are returned. Otherwise, all courses are returned.
	GetCourses(...uint64) ([]*qf.Course, error)
	// GetCoursesByUser returns all courses (with enrollment status)
	// for the given user id.
	// If enrollment statuses is provided, the set of courses returned
	// is filtered according to these enrollment statuses.
	GetCoursesByUser(userID uint64, statuses ...qf.Enrollment_UserStatus) ([]*qf.Course, error)
	// GetCourseTeachers returns a list of all teachers in a course.
	GetCourseTeachers(query *qf.Course) ([]*qf.User, error)
	// UpdateCourse updates course information.
	UpdateCourse(*qf.Course) error

	// CreateEnrollment creates a new pending enrollment.
	CreateEnrollment(*qf.Enrollment) error
	// RejectEnrollment removes the user enrollment from the database
	RejectEnrollment(userID, courseID uint64) error
	// UpdateEnrollmentStatus changes status of the course enrollment for the given user and course.
	UpdateEnrollment(*qf.Enrollment) error
	// GetEnrollmentByCourseAndUser returns a user enrollment for the given course ID.
	GetEnrollmentByCourseAndUser(courseID uint64, userID uint64) (*qf.Enrollment, error)
	// GetEnrollmentsByCourse fetches all course enrollments with given statuses.
	GetEnrollmentsByCourse(courseID uint64, statuses ...qf.Enrollment_UserStatus) ([]*qf.Enrollment, error)
	// GetEnrollmentsByUser fetches all enrollments for the given user
	GetEnrollmentsByUser(userID uint64, statuses ...qf.Enrollment_UserStatus) ([]*qf.Enrollment, error)

	// CreateGroup creates a new group and assign users to newly created group.
	CreateGroup(*qf.Group) error
	// UpdateGroup updates a group with the specified users and enrollments.
	UpdateGroup(group *qf.Group) error
	// UpdateGroupStatus updates status field of a group.
	UpdateGroupStatus(*qf.Group) error
	// DeleteGroup deletes a group and its corresponding enrollments.
	DeleteGroup(uint64) error
	// GetGroup returns the group with the specified group ID.
	GetGroup(uint64) (*qf.Group, error)
	// GetGroupsByCourse returns the groups for the given course.
	GetGroupsByCourse(courseID uint64, statuses ...qf.Group_GroupStatus) ([]*qf.Group, error)

	// CreateAssignment creates a new or updates an existing assignment.
	CreateAssignment(*qf.Assignment) error
	// GetAssignment returns assignment mathing the given query.
	GetAssignment(query *qf.Assignment) (*qf.Assignment, error)
	// GetAssignmentsByCourse returns a list of all assignments for the given course ID.
	GetAssignmentsByCourse(uint64, bool) ([]*qf.Assignment, error)
	// UpdateAssignments updates the specified list of assignments.
	UpdateAssignments([]*qf.Assignment) error
	// CreateBenchmark creates a new grading benchmark.
	CreateBenchmark(*qf.GradingBenchmark) error
	// UpdateBenchmark updates the given benchmark.
	UpdateBenchmark(*qf.GradingBenchmark) error
	// DeleteBenchmark deletes the given benchmark.
	DeleteBenchmark(*qf.GradingBenchmark) error
	// CreateCriterion creates a new grading criterion.
	CreateCriterion(*qf.GradingCriterion) error
	// UpdateCriterion updates the given criterion.
	UpdateCriterion(*qf.GradingCriterion) error
	// DeleteCriterion deletes the given criterion.
	DeleteCriterion(*qf.GradingCriterion) error

	// CreateSubmission creates a new submission record or updates the most
	// recent submission, as defined by the provided submissionQuery.
	// The submissionQuery must always specify the assignment, and may specify the ID of
	// either an individual student or a group, but not both.
	CreateSubmission(*qf.Submission) error
	// GetSubmission returns a single submission matching the given query.
	GetSubmission(query *qf.Submission) (*qf.Submission, error)
	// GetLastSubmissions returns a list of submission entries for the given course, matching the given query.
	GetLastSubmissions(courseID uint64, query *qf.Submission) ([]*qf.Submission, error)
	// GetSubmissions returns all submissions matching the query.
	GetSubmissions(*qf.Submission) ([]*qf.Submission, error)
	// GetAssignmentsWithSubmissions returns a list of assignments with the latest submissions for the given course.
	GetAssignmentsWithSubmissions(courseID uint64, requestType qf.SubmissionsForCourseRequest_Type, withBuildInfo bool) ([]*qf.Assignment, error)
	// UpdateSubmission updates the specified submission with approved or not approved.
	UpdateSubmission(*qf.Submission) error
	// UpdateSubmissions releases and/or approves all submissions with a certain score
	UpdateSubmissions(uint64, *qf.Submission) error
	// GetReview returns a single review matching the given query.
	GetReview(query *qf.Review) (*qf.Review, error)
	// CreateReview adds a new submission review.
	CreateReview(*qf.Review) error
	// UpdateReview updates the given review.
	UpdateReview(*qf.Review) error
	// DeleteReview removes all review records matching the query.
	DeleteReview(*qf.Review) error
	// GetBenchmarks return all benchmarks and criteria for an assignmend
	GetBenchmarks(*qf.Assignment) ([]*qf.GradingBenchmark, error)
	// CreateRepository creates a new repository.
	CreateRepository(repo *qf.Repository) error
	// GetRepositories returns repositories that match the given query.
	GetRepositories(query *qf.Repository) ([]*qf.Repository, error)
	// DeleteRepository deletes repository for the given remote provider's ID.
	DeleteRepository(remoteID uint64) error
	// GetRepositoriesWithIssues gets repositories with issues
	GetRepositoriesWithIssues(query *qf.Repository) ([]*qf.Repository, error)

	// GetTasks returns tasks that match the given query.
	GetTasks(query *qf.Task) ([]*qf.Task, error)
	// CreateIssues creates a batch of issues
	CreateIssues(issues []*qf.Issue) error
	// SynchronizeAssignmentTasks synchronizes all tasks of each assignment in a given course. Returns created, updated and deleted tasks
	SynchronizeAssignmentTasks(course *qf.Course, taskMap map[uint32]map[string]*qf.Task) ([]*qf.Task, []*qf.Task, error)
	// CreatePullRequest creates a pull request
	CreatePullRequest(pullRequest *qf.PullRequest) error
	// GetPullRequest returns the pull request matching the given query
	GetPullRequest(query *qf.PullRequest) (*qf.PullRequest, error)
	// HandleMergingPR handles merging a pull request
	HandleMergingPR(query *qf.PullRequest) error
	// DeletePullRequest updates the pull request matching the given query
	UpdatePullRequest(pullRequest *qf.PullRequest) error

	// UpdateSlipDays updates used slipdays for the given course enrollment
	UpdateSlipDays([]*qf.UsedSlipDays) error
}
