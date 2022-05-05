package scm

import (
	"context"
	"errors"

	pb "github.com/autograde/quickfeed/ag"
	"go.uber.org/zap"
)

// SCM is a common interface for different source code management solutions,
// i.e., GitHub and GitLab.
type SCM interface {
	// Creates a new organization.
	CreateOrganization(context.Context, *OrganizationOptions) (*pb.Organization, error)
	// Updates an organization
	UpdateOrganization(context.Context, *OrganizationOptions) error
	// Gets an organization.
	GetOrganization(context.Context, *GetOrgOptions) (*pb.Organization, error)
	// Create a new repository.
	CreateRepository(context.Context, *CreateRepositoryOptions) (*Repository, error)
	// Get repository by ID or name
	GetRepository(context.Context, *RepositoryOptions) (*Repository, error)
	// Get repositories within organization.
	GetRepositories(context.Context, *pb.Organization) ([]*Repository, error)
	// Delete repository.
	DeleteRepository(context.Context, *RepositoryOptions) error
	// Add user as repository collaborator with provided permissions
	UpdateRepoAccess(context.Context, *Repository, string, string) error
	// Returns true if there are no commits in the given repository
	RepositoryIsEmpty(context.Context, *RepositoryOptions) bool
	// List the webhooks associated with the provided repository or organization.
	ListHooks(context.Context, *Repository, string) ([]*Hook, error)
	// Creates a new webhook for organization if the name of organization
	// is provided. Otherwise creates a hook for the given repo.
	CreateHook(context.Context, *CreateHookOptions) error
	// Create team.
	CreateTeam(context.Context, *NewTeamOptions) (*Team, error)
	// Delete team.
	DeleteTeam(context.Context, *TeamOptions) error
	// Get a single team by ID or name.
	GetTeam(context.Context, *TeamOptions) (*Team, error)
	// Fetch all teams for organization.
	GetTeams(context.Context, *pb.Organization) ([]*Team, error)
	// Add repo to team.
	AddTeamRepo(context.Context, *AddTeamRepoOptions) error
	// AddTeamMember adds a member to a team.
	AddTeamMember(context.Context, *TeamMembershipOptions) error
	// RemoveTeamMember removes team member.
	RemoveTeamMember(context.Context, *TeamMembershipOptions) error
	// UpdateTeamMembers adds or removes members of an existing team based on list of users in TeamOptions.
	UpdateTeamMembers(context.Context, *UpdateTeamOptions) error
	// GetUserName returns the currently logged in user's login name.
	GetUserName(context.Context) (string, error)
	// GetUserNameByID returns the login name of user with the given remoteID.
	GetUserNameByID(context.Context, uint64) (string, error)
	// Returns a provider specific clone path.
	CreateCloneURL(*URLPathOptions) string
	// Promotes or demotes organization member, based on Role field in OrgMembership.
	UpdateOrgMembership(context.Context, *OrgMembershipOptions) error
	// RevokeOrgMembership removes user from the organization.
	RemoveMember(context.Context, *OrgMembershipOptions) error
	// Lists all authorizations for authenticated user.
	GetUserScopes(context.Context) *Authorization

	// CreateIssue creates an issue.
	CreateIssue(context.Context, *CreateIssueOptions) (*Issue, error)
	// EditIssue edits an issue .
	EditIssue(ctx context.Context, issueNumber int, opt *CreateIssueOptions) (*Issue, error)
	// GetIssue fetches a specific issue.
	GetIssue(ctx context.Context, issueNumber int, opt *RepositoryOptions) (*Issue, error)
	// GetIssues fetches all issues in a repository.
	GetIssues(ctx context.Context, opt *RepositoryOptions) ([]*Issue, error)

	// CreateIssueComment creates a comment on a GitHub issue.
	CreateIssueComment(ctx context.Context, number int, opt *IssueCommentOptions) (uint64, error)
	// EditIssueComment edits a comment on a GitHub issue.
	EditIssueComment(ctx context.Context, commentID int64, opt *IssueCommentOptions) error

	// RequestReviewers requests reviewers for a pull request.
	RequestReviewers(ctx context.Context, opt *RequestReviewersOptions) error

	// Accepts repository invite.
	AcceptRepositoryInvites(context.Context, *RepositoryInvitationOptions) error
}

// NewSCMClient returns a new provider client implementing the SCM interface.
func NewSCMClient(logger *zap.SugaredLogger, provider, token string) (SCM, error) {
	switch provider {
	case "github":
		return NewGithubSCMClient(logger, token), nil
	case "gitlab":
		return NewGitlabSCMClient(token), nil
	case "fake":
		return NewFakeSCMClient(), nil
	}
	return nil, errors.New("invalid provider: " + provider)
}

// OrganizationOptions contains information on how an organization should be
// created.
type OrganizationOptions struct {
	Path              string
	Name              string
	DefaultPermission string
	// prohibit students from creating new repos
	// on the course organization
	RepoPermissions bool
}

// GetOrgOptions contains information on the organization to fetch
type GetOrgOptions struct {
	ID   uint64
	Name string
	// Username field is used to filter organizations
	// where the given user has a certain role.
	Username string
}

// Repository represents a git remote repository.
type Repository struct {
	ID      uint64
	Path    string
	Owner   string // Only used by GitHub.
	WebURL  string // Repository website.
	SSHURL  string // SSH clone URL, used by GitLab.
	HTTPURL string // HTTP(S) clone URL.
	OrgID   uint64
	Size    uint64
}

// RepositoryOptions is used to fetch a single repository by ID or name.
// Either ID or both Path and Owner fields must be set.
type RepositoryOptions struct {
	ID    uint64
	Path  string
	Owner string
}

// Hook contains information about a webhook for a repository.
type Hook struct {
	ID     uint64
	Name   string
	URL    string
	Events []string
}

// CreateRepositoryOptions contains information on how a repository should be created.
type CreateRepositoryOptions struct {
	Organization *pb.Organization
	Path         string
	Private      bool
	Owner        string // The owner of an organization's repo is always the organization itself.
	Permission   string // Default permission level for the given repo. Can be "read", "write", "admin", "none".
}

// CreateHookOptions contains information on how to create a webhook.
// If Organization string is provided, will create a new hook on
// the organization's level. This hook will be triggered on push to any
// of the organization's repositories.
type CreateHookOptions struct {
	URL          string
	Secret       string
	Organization string
	Repository   *Repository
}

// TeamOptions contains information about the team and the organization it belongs to.
// It must include either both IDs or both names for the team and organization/
type TeamOptions struct {
	Organization   string
	OrganizationID uint64
	TeamName       string
	TeamID         uint64
}

// NewTeamOptions used when creating a new team
type NewTeamOptions struct {
	Organization string
	TeamName     string
	Users        []string
}

// UpdateTeamOptions used when updating team members.
type UpdateTeamOptions struct {
	OrganizationID uint64
	TeamID         uint64
	Users          []string
}

// TeamMembershipOptions contain information on organization team and associated user.
// Username and either team ID, or names of both the team and organization must be provided.
type TeamMembershipOptions struct {
	Organization   string
	OrganizationID uint64
	TeamID         uint64
	TeamName       string
	Username       string // GitHub username.
	Role           string // "Member" or "maintainer". A maintainer can add, remove and promote team members.
}

// OrgMembershipOptions represent user's membership in organization
type OrgMembershipOptions struct {
	Organization string
	Username     string // GitHub username.
	Role         string // Role can be "admin" (organization owner) or "member".
}

// URLPathOptions holds elements used when constructing a clone URL string.
type URLPathOptions struct {
	UserToken    string
	Organization string
	Repository   string
}

// AddTeamRepoOptions contains information about the repos to be added to a team.
// All fields must be provided.
type AddTeamRepoOptions struct {
	OrganizationID uint64
	TeamID         uint64
	Repo           string
	Owner          string // Name of the organization. Only used by GitHub.
	Permission     string // Permission level for team members. Can be "push", "pull", "admin".
}

// Team represents a git Team
type Team struct {
	ID           uint64
	Name         string
	Organization string
}

// Authorization stores information about user scopes
type Authorization struct {
	Scopes []string
}

// CreateNewIssueOptions contains information on how to create an Issue.
//type NewIssue struct {
//	Title string
//	Body   *Repository
//	Labels *[]string
//	Assignee *string
//	Assignees *[]string
//}

// Repository represents a git remote repository.
type Issue struct {
	ID         uint64
	Title      string
	Body       string
	Repository string
	Assignee   string
	Status     string
	//	Assignees string
	IssueNumber int
}

// CreateIssueOptions contains information on how to create an Issue.
type CreateIssueOptions struct {
	Organization string
	Repository   string
	Title        string
	Body         string
	State        string
	Labels       *[]string
	Assignee     *string
	Assignees    *[]string
}

// RequestReviewersOptions contains information on how to create or edit a pull request comment.
type IssueCommentOptions struct {
	Organization string
	Repository   string
	Body         string
}

// RequestReviewersOptions contains information on how to assign reviewers to a pull request.
type RequestReviewersOptions struct {
	Organization string
	Repository   string
	Number       int
	Reviewers    []string // Reviewers is a slice of github usernames
}

// RepositoryInvitationOptions contains information on which organization and user to accept invitations for.
type RepositoryInvitationOptions struct {
	Login string // GitHub username.
	Owner string // Name of the organization.
}
