package web_test

import (
	"context"
	"testing"

	"github.com/quickfeed/quickfeed/ci"
	"github.com/quickfeed/quickfeed/internal/qtest"
	"github.com/quickfeed/quickfeed/qf"
	"github.com/quickfeed/quickfeed/qlog"
	"github.com/quickfeed/quickfeed/scm"
	"github.com/quickfeed/quickfeed/web"
)

func TestMakeSCMs(t *testing.T) {
	scmConfig, err := scm.NewSCMConfig()
	if err != nil {
		t.Skip("Requires a valid SCM app")
	}
	mgr := scm.NewSCMManager(scmConfig)
	db, cleanup := qtest.TestDB(t)
	defer cleanup()
	ctx := context.Background()
	logger := qlog.Logger(t).Desugar()
	q := web.NewQuickFeedService(logger, db, mgr, web.BaseHookOptions{}, &ci.Local{})
	admin := qtest.CreateFakeUser(t, db, 1)
	course := &qf.Course{
		Name:             "Test course",
		OrganizationPath: scm.GetTestOrganization(t),
		Provider:         "fake",
	}
	if err := db.CreateCourse(admin.ID, course); err != nil {
		t.Fatal(err)
	}
	if err := q.MakeSCMs(ctx); err != nil {
		t.Fatal(err)
	}
	_, ok := mgr.GetSCM(course.OrganizationPath)
	if !ok {
		t.Fatalf("Missing scm client for organization %s", course.OrganizationPath)
	}
}
