package ag_test

import (
	"testing"

	pb "github.com/quickfeed/quickfeed/ag"
)

func TestIsValid(t *testing.T) {
	tests := map[string]struct {
		pr   *pb.PullRequest
		want bool
	}{
		"Valid":                   {pr: &pb.PullRequest{ScmRepositoryID: 1, TaskID: 1, IssueID: 1, UserID: 1, SourceBranch: "A", Number: 1}, want: true},
		"Invalid ScmRepositoryID": {pr: &pb.PullRequest{ScmRepositoryID: 0, TaskID: 1, IssueID: 1, UserID: 1, SourceBranch: "A", Number: 1}, want: false},
		"Invalid TaskID":          {pr: &pb.PullRequest{ScmRepositoryID: 1, TaskID: 0, IssueID: 1, UserID: 1, SourceBranch: "A", Number: 1}, want: false},
		"Invalid IssueID":         {pr: &pb.PullRequest{ScmRepositoryID: 1, TaskID: 1, IssueID: 0, UserID: 1, SourceBranch: "A", Number: 1}, want: false},
		"Invalid UserID":          {pr: &pb.PullRequest{ScmRepositoryID: 1, TaskID: 1, IssueID: 1, UserID: 0, SourceBranch: "A", Number: 1}, want: false},
		"Invalid BranchName":      {pr: &pb.PullRequest{ScmRepositoryID: 1, TaskID: 1, IssueID: 1, UserID: 1, SourceBranch: "", Number: 1}, want: false},
		"Invalid Number":          {pr: &pb.PullRequest{ScmRepositoryID: 1, TaskID: 1, IssueID: 1, UserID: 1, SourceBranch: "A", Number: 0}, want: false},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.pr.Valid()
			if tt.want != got {
				t.Errorf("IsValid() = %t, expected %t\npr: %v", got, tt.want, tt.pr)
			}
		})
	}
}

func TestHasFeedbackComment(t *testing.T) {
	tests := map[string]struct {
		pr   *pb.PullRequest
		want bool
	}{
		"Comment":    {pr: &pb.PullRequest{ScmCommentID: 1}, want: true},
		"No comment": {pr: &pb.PullRequest{ScmCommentID: 0}, want: false},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.pr.HasFeedbackComment()
			if tt.want != got {
				t.Errorf("HasFeedbackComment() = %t, expected %t\npr: %v", got, tt.want, tt.pr)
			}
		})
	}
}
