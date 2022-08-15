package interceptor

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/quickfeed/quickfeed/database"
	"github.com/quickfeed/quickfeed/qf"
	"github.com/quickfeed/quickfeed/web/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type (
	role      int
	roles     []role
	requestID interface {
		IDFor(string) uint64
	}
)

const (
	none role = iota
	// user role implies that user attempts to access information about himself.
	user
	// group role implies that the user is a course student + a member of the given group.
	group
	// student role implies that the user is enrolled in the course with any role.
	student
	// teacher: user enrolled in the course with teacher status.
	teacher
	// courseAdmin: an admin user who is also enrolled into the course.
	courseAdmin
	// admin is the user with admin privileges.
	admin
)

// If there are several roles that can call a method, a role with the least privilege must come first.
var access = map[string]roles{
	"GetUser":                 {none},
	"GetCourse":               {none},
	"GetCourses":              {none},
	"CreateEnrollment":        {user},
	"UpdateCourseVisibility":  {user},
	"GetCoursesByUser":        {user},
	"UpdateUser":              {user, admin},
	"GetEnrollmentsByUser":    {user, admin},
	"GetSubmissions":          {group, student, teacher, courseAdmin},
	"GetGroupByUserAndCourse": {group, teacher},
	"CreateGroup":             {group, teacher},
	"GetGroup":                {group, teacher},
	"GetAssignments":          {student, teacher},
	"GetEnrollmentsByCourse":  {student, teacher},
	"GetRepositories":         {student, teacher},
	"UpdateGroup":             {teacher},
	"DeleteGroup":             {teacher},
	"GetGroupsByCourse":       {teacher},
	"UpdateCourse":            {teacher},
	"UpdateEnrollments":       {teacher},
	"UpdateAssignments":       {teacher},
	"UpdateSubmission":        {teacher},
	"UpdateSubmissions":       {teacher},
	"RebuildSubmissions":      {teacher},
	"CreateBenchmark":         {teacher},
	"UpdateBenchmark":         {teacher},
	"DeleteBenchmark":         {teacher},
	"CreateCriterion":         {teacher},
	"UpdateCriterion":         {teacher},
	"DeleteCriterion":         {teacher},
	"CreateReview":            {teacher},
	"UpdateReview":            {teacher},
	"GetReviewers":            {teacher},
	"IsEmptyRepo":             {teacher},
	"GetSubmissionsByCourse":  {courseAdmin},
	"GetUserByCourse":         {teacher, admin},
	"GetUsers":                {admin},
	"GetOrganization":         {admin},
	"CreateCourse":            {admin},
}

// AccessControl checks user information stored in the JWT claims agains the list of roles required to call the method.
func AccessControl(logger *zap.SugaredLogger, tm *auth.TokenManager) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod[strings.LastIndex(info.FullMethod, "/")+1:]
		logger.Debugf("ACCESS CONTROL for method %s", method) // tmp
		req, ok := request.(requestID)
		// The GetUserByCourse method sends a CourseUserRequest which has no IDs and needs a database query.
		if !ok && method != "GetUserByCourse" {
			logger.Errorf("%s failed: message type '%s' does not implement IDFor interface",
				method, reflect.TypeOf(request).String())
			return nil, ErrAccessDenied
		}
		roles, ok := access[method]
		if !ok {
			logger.Errorf("No access roles defined for %s", method)
			return nil, ErrAccessDenied
		}
		logger.Debug("Got roles: ", roles) // tmp
		claims, err := tm.GetClaims(ctx)
		if err != nil {
			logger.Errorf("Access control: failed to get claims from request context: %v", err)
			return handler(ctx, request)
		}
		logger.Debug("Got user claims: ", claims) // tmp
		for _, role := range roles {
			switch role {
			case none:
				return handler(ctx, request)
			case user:
				if req.IDFor("user") == claims.UserID {
					// Make sure the user is not updating own admin status.
					if method == "UpdateUser" {
						if req.(*qf.User).GetIsAdmin() != claims.Admin {
							logger.Errorf("Access control: user %d attempted to change admin status from %v to %v",
								claims.UserID, claims.Admin, req.(*qf.User).GetIsAdmin())
						}
					}
					return handler(ctx, request)
				}
			case student:
				courseID := req.IDFor("course")
				status, ok := claims.Courses[courseID]
				if ok && status == qf.Enrollment_STUDENT {
					return handler(ctx, request)
				}
			case group:
				groupID := req.IDFor("group")
				for _, group := range claims.Groups {
					if group == groupID {
						return handler(ctx, request)
					}
				}
			case teacher:
				courseID := req.IDFor("course")
				status, ok := claims.Courses[courseID]
				if ok && status == qf.Enrollment_TEACHER {
					return handler(ctx, request)
				}
			// This role has a single use case: the GetUserByCourse method wich sends a CourseUserRequest
			// that does not include any IDs.
			case courseAdmin:
				if claims.Admin {
					if err := isCourseTeacher(tm.Database(), request.(*qf.CourseUserRequest), claims.Courses); err != nil {
						logger.Errorf("AccessControl: %v", err)
						return nil, ErrAccessDenied
					}
					return handler(ctx, request)
				}
			case admin:
				if claims.Admin {
					return handler(ctx, request)
				}
			}
		}
		logger.Errorf("%f failed for user %d: insufficient user privileges. Required roles: %v", method, claims.UserID, roles)
		return nil, ErrAccessDenied
	}
}

// isCourseTeacher checks if the user is a teacher in the course in the CourseUserRequest.
func isCourseTeacher(db database.Database, request *qf.CourseUserRequest, courses map[uint64]qf.Enrollment_UserStatus) error {
	for courseID, status := range courses {
		if status == qf.Enrollment_TEACHER {
			course, err := db.GetCourse(courseID, false)
			if err != nil {
				return err
			}
			if course.GetCode() == request.GetCourseCode() && course.GetYear() == request.GetCourseYear() {
				return nil
			}
		}
	}
	return fmt.Errorf("user is not teacher of the %s course", request.GetCourseCode())
}
