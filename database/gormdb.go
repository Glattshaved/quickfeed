package database

import (
	"errors"

	"github.com/autograde/aguis/models"
	"github.com/jinzhu/gorm"
)

// GormDB implements the Database interface.
type GormDB struct {
	conn *gorm.DB
}

// NewGormDB creates a new gorm database using the provided driver.
func NewGormDB(driver, path string, logger GormLogger) (*GormDB, error) {
	conn, err := gorm.Open(driver, path)
	if err != nil {
		return nil, err
	}

	if logger != nil {
		conn.LogMode(true)
		conn.SetLogger(logger)
	}

	conn.AutoMigrate(
		&models.User{},
		&models.RemoteIdentity{},
		&models.Course{},
		&models.Enrollment{},
		&models.Assignment{},
		&models.Submission{},
		&models.Group{},
	)

	return &GormDB{conn}, nil
}

// GetUser implements the Database interface.
func (db *GormDB) GetUser(id uint64) (*models.User, error) {
	var user models.User
	if err := db.conn.Preload("RemoteIdentities").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByRemoteIdentity implements the Database interface.
func (db *GormDB) GetUserByRemoteIdentity(provider string, remoteID uint64, accessToken string) (*models.User, error) {
	tx := db.conn.Begin()

	// Get the remote identity.
	var remoteIdentity models.RemoteIdentity
	if err := tx.
		Where(&models.RemoteIdentity{
			Provider: provider,
			RemoteID: remoteID,
		}).
		First(&remoteIdentity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update the access token.
	if err := tx.Model(&remoteIdentity).Update("access_token", accessToken).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Get the user.
	var user models.User
	if err := tx.Preload("RemoteIdentities").First(&user, remoteIdentity.UserID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers implements the Database interface.
func (db *GormDB) GetUsers(ids ...uint64) ([]*models.User, error) {
	m := db.conn
	if len(ids) > 0 {
		m = m.Where(ids)
	}

	var users []*models.User
	if err := m.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// SetAdmin implements the Database interface.
func (db *GormDB) SetAdmin(id uint64) error {
	var user models.User
	if err := db.conn.First(&user, id).Error; err != nil {
		return err
	}
	user.IsAdmin = true
	return db.conn.Save(&user).Error
}

// GetRemoteIdentity implements the Database interface.
func (db *GormDB) GetRemoteIdentity(provider string, remoteID uint64) (*models.RemoteIdentity, error) {
	var remoteIdentity models.RemoteIdentity
	if err := db.conn.Model(&models.RemoteIdentity{}).
		Where(&models.RemoteIdentity{
			Provider: provider,
			RemoteID: remoteID,
		}).
		First(&remoteIdentity).Error; err != nil {
		return nil, err
	}
	return &remoteIdentity, nil
}

// CreateUserFromRemoteIdentity implements the Database interface.
func (db *GormDB) CreateUserFromRemoteIdentity(provider string, remoteID uint64, accessToken string) (*models.User, error) {
	var count uint64
	if err := db.conn.
		Model(&models.RemoteIdentity{}).
		Where(&models.RemoteIdentity{
			Provider: provider,
			RemoteID: remoteID,
		}).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, ErrDuplicateIdentity
	}

	user := models.User{
		RemoteIdentities: []*models.RemoteIdentity{{
			Provider:    provider,
			RemoteID:    remoteID,
			AccessToken: accessToken,
		}},
	}
	if err := db.conn.Create(&user).Error; err != nil {
		return nil, err
	}
	// The first user defaults to admin user.
	if user.ID == 1 {
		if err := db.SetAdmin(1); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

// ErrDuplicateIdentity is returned when trying to associate a remote identity
// with a user account and the identity is already in use.
var ErrDuplicateIdentity = errors.New("remote identity register with another user")

// AssociateUserWithRemoteIdentity implements the Database interface.
func (db *GormDB) AssociateUserWithRemoteIdentity(userID uint64, provider string, remoteID uint64, accessToken string) error {
	var count uint64
	if err := db.conn.
		Model(&models.RemoteIdentity{}).
		Where(&models.RemoteIdentity{
			Provider: provider,
			RemoteID: remoteID,
		}).
		Not(&models.RemoteIdentity{
			UserID: userID,
		}).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return ErrDuplicateIdentity
	}

	var remoteIdentity models.RemoteIdentity
	return db.conn.
		Where(models.RemoteIdentity{Provider: provider, RemoteID: remoteID, UserID: userID}).
		Assign(models.RemoteIdentity{AccessToken: accessToken}).
		FirstOrCreate(&remoteIdentity).Error
}

// CreateCourse implements the Database interface.
func (db *GormDB) CreateCourse(course *models.Course) error {
	return db.conn.Create(course).Error
}

// GetCourses implements the Database interface.
// If one or more course ids are provided, the corresponding courses
// are returned. Otherwise, all courses are returned.
func (db *GormDB) GetCourses(ids ...uint64) ([]*models.Course, error) {
	m := db.conn
	if len(ids) > 0 {
		m = m.Where(ids)
	}
	var courses []*models.Course
	if err := m.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

// GetAssignmentsByCourse implements the Database interface
func (db *GormDB) GetAssignmentsByCourse(id uint64) ([]*models.Assignment, error) {
	var course models.Course
	if err := db.conn.Preload("Assignments").First(&course, id).Error; err != nil {
		return nil, err
	}
	return course.Assignments, nil
}

// CreateSubmission implements the Database interface
// TODO: Also check enrollment to see if the user is
// enrolled in the course the assignment belongs to
func (db *GormDB) CreateSubmission(submission *models.Submission) error {
	var course uint64
	// Checks that the course exists
	if err := db.conn.Model(&models.Assignment{}).Where(&models.Assignment{
		ID: submission.AssignmentID,
	}).Count(&course).Error; err != nil {
		return err
	}
	var user uint64
	// Checks that the user exists
	if err := db.conn.Model(&models.User{}).Where(&models.User{
		ID: submission.UserID,
	}).Count(&user).Error; err != nil {
		return err
	}
	if course+user != 2 {
		return gorm.ErrRecordNotFound
	}
	return db.conn.Create(&submission).Error
}

// GetSubmissionForUser implements the Database interface
func (db *GormDB) GetSubmissionForUser(assignmentID uint64, userID uint64) (*models.Submission, error) {
	var submission models.Submission
	if err := db.conn.Where(&models.Submission{AssignmentID: assignmentID, UserID: userID}).Last(&submission).Error; err != nil {
		return nil, err
	}
	return &submission, nil
}

// GetSubmissions implements the Database interface
func (db *GormDB) GetSubmissions(courseID uint64, userID uint64) ([]*models.Submission, error) {
	var course models.Course
	if err := db.conn.Preload("Assignments").First(&course, courseID).Error; err != nil {
		return nil, err
	}

	latestSubs := make([]*models.Submission, 0)
	for _, v := range course.Assignments {
		temp, err := db.GetSubmissionForUser(v.ID, userID)
		if err == nil {
			latestSubs = append(latestSubs, temp)
		} else if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return latestSubs, nil
}

// CreateAssignment implements the Database interface
func (db *GormDB) CreateAssignment(assignment *models.Assignment) error {
	var course uint64
	if err := db.conn.Model(&models.Course{}).Where(&models.Course{
		ID: assignment.CourseID,
	}).Count(&course).Error; err != nil {
		return err
	}
	if course != 1 {
		return gorm.ErrRecordNotFound
	}

	return db.conn.
		Where(models.Assignment{
			CourseID: assignment.CourseID,
			Order:    assignment.Order,
		}).
		Assign(models.Assignment{
			Name:        assignment.Name,
			Language:    assignment.Language,
			Deadline:    assignment.Deadline,
			AutoApprove: assignment.AutoApprove,
		}).FirstOrCreate(assignment).Error
}

// CreateEnrollment implements the Database interface.
func (db *GormDB) CreateEnrollment(enrollment *models.Enrollment) error {
	var user, course uint64
	if err := db.conn.Model(&models.User{}).Where(&models.User{
		ID: enrollment.UserID,
	}).Count(&user).Error; err != nil {
		return err
	}
	if err := db.conn.Model(&models.Course{}).Where(&models.Course{
		ID: enrollment.CourseID,
	}).Count(&course).Error; err != nil {
		return err
	}
	if user+course != 2 {
		return gorm.ErrRecordNotFound
	}

	return db.conn.Where(enrollment).FirstOrCreate(enrollment).Error
}

// AcceptEnrollment implements the Database interface.
func (db *GormDB) AcceptEnrollment(id uint64) error {
	return db.setEnrollment(id, models.Accepted)
}

// RejectEnrollment implements the Database interface.
func (db *GormDB) RejectEnrollment(id uint64) error {
	return db.setEnrollment(id, models.Rejected)
}

// GetEnrollmentsByCourse implements the Database interface.
func (db *GormDB) GetEnrollmentsByCourse(id uint64, statuses ...uint) ([]*models.Enrollment, error) {
	return db.getEnrollments(&models.Course{ID: id}, statuses...)
}

func (db *GormDB) getEnrollments(model interface{}, statuses ...uint) ([]*models.Enrollment, error) {
	if len(statuses) == 0 {
		statuses = []uint{models.Pending, models.Rejected, models.Accepted}
	}
	var enrollments []*models.Enrollment
	if err := db.conn.Model(model).
		Where("status in (?)", statuses).
		Association("Enrollments").
		Find(&enrollments).Error; err != nil {
		return nil, err
	}

	return enrollments, nil
}

// GetEnrollmentByCourseAndUser return a record of Enrollment
func (db *GormDB) GetEnrollmentByCourseAndUser(cid uint64, uid uint64) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	if err := db.conn.
		Where(&models.Enrollment{
			CourseID: cid,
			UserID:   uid,
		}).
		First(&enrollment).Error; err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (db *GormDB) setEnrollment(id uint64, status uint) error {
	if status > models.Accepted {
		panic("invalid status")
	}
	return db.conn.
		Model(&models.Enrollment{}).
		Where(&models.Enrollment{ID: id}).
		Update(&models.Enrollment{Status: int(status)}).Error
}

// GetCoursesByUser returns all courses (with enrollment status)
// for the given user id.
// If enrollment statuses is provided, the set of courses returned
// is filtered according to these enrollment statuses.
func (db *GormDB) GetCoursesByUser(id uint64, statuses ...uint) ([]*models.Course, error) {
	enrollments, err := db.getEnrollments(&models.User{ID: id}, statuses...)
	if err != nil {
		return nil, err
	}

	var courseIDs []uint64
	m := make(map[uint64]*models.Enrollment)
	for _, enrollment := range enrollments {
		m[enrollment.CourseID] = enrollment
		courseIDs = append(courseIDs, enrollment.CourseID)
	}

	if len(statuses) == 0 {
		courseIDs = nil
	} else if len(courseIDs) == 0 {
		// No need to query database since user have no enrolled courses.
		return []*models.Course{}, nil
	}
	courses, err := db.GetCourses(courseIDs...)
	if err != nil {
		return nil, err
	}

	for _, course := range courses {
		course.Enrolled = models.None
		if enrollment, ok := m[course.ID]; ok {
			course.Enrolled = enrollment.Status
		}
	}
	return courses, nil
}

// GetCourse implements the Database interface
func (db *GormDB) GetCourse(id uint64) (*models.Course, error) {
	var course models.Course
	if err := db.conn.First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// UpdateCourse implements the Database interface
func (db *GormDB) UpdateCourse(course *models.Course) error {
	return db.conn.Model(&models.Course{}).Updates(course).Error
}

// CreateGroup creates a new group and assign users to newly created group
func (db *GormDB) CreateGroup(group *models.Group) error {
	// Uses transaction to update more than one db tables
	// Failure in one table rollbacks all
	tx := db.conn.Begin()
	if err := tx.Model(&models.Group{}).Create(group).Error; err != nil {
		tx.Rollback()
		return err
	}
	var userids []uint64
	for _, u := range group.Users {
		userids = append(userids, u.ID)
	}
	// Batch update
	if err := tx.Model(&models.Enrollment{}).
		Where("user_id IN (?) AND course_id = ?", userids, group.CourseID).
		Updates(&models.Enrollment{
			GroupID: group.ID,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// GetGroup returns a group specified by id return error if does not exits
func (db *GormDB) GetGroup(id uint64) (*models.Group, error) {
	var group models.Group
	if err := db.conn.Preload("Enrollments").First(&group, id).Error; err != nil {
		return nil, err
	}
	var userIDs []uint64
	for _, enrollment := range group.Enrollments {
		userIDs = append(userIDs, enrollment.UserID)
	}
	if len(userIDs) > 0 {
		users, err := db.GetUsers(userIDs...)
		if err != nil {
			return nil, err
		}
		group.Users = users
	}
	return &group, nil
}

// UpdateGroupStatus updates status field of a group
func (db *GormDB) UpdateGroupStatus(group *models.Group) error {
	return db.conn.Model(group).Update("status", group.Status).Error
}

// GetGroups returns a list of groups
func (db *GormDB) GetGroups(cid uint64) ([]*models.Group, error) {
	var groups []*models.Group
	if err := db.conn.
		Preload("Enrollments").
		Where(&models.Course{
			ID: cid,
		}).
		Find(&groups).Error; err != nil {
		return nil, err
	}

	for _, group := range groups {
		var userIDs []uint64
		for _, enrollment := range group.Enrollments {
			userIDs = append(userIDs, enrollment.UserID)
		}
		if len(userIDs) > 0 {
			users, err := db.GetUsers(userIDs...)
			if err != nil {
				return nil, err
			}
			group.Users = users
		}

	}
	return groups, nil
}

// Close closes the gorm database.
func (db *GormDB) Close() error {
	return db.conn.Close()
}
