package github

import (
	"fmt"
	"github.com/octokit/go-octokit/octokit"
)

const (
	GitHubHost  string = "github.com"
	OAuthAppURL string = "http://owenou.com/gh"
)

type GitHub struct {
	Project *Project
	Config  *Config
}

func (gh *GitHub) PullRequest(id string) (pr *octokit.PullRequest, err error) {
	client := gh.octokit()
	prService, err := client.PullRequests(&octokit.PullRequestsURL, octokit.M{"owner": gh.Project.Owner, "repo": gh.Project.Name, "number": id})
	if err != nil {
		return nil, err
	}

	pr, result := prService.Get()
	if result.HasError() {
		err = result.Err
	}

	return pr, nil
}

func (gh *GitHub) CreatePullRequest(base, head, title, body string) (pr *octokit.PullRequest, err error) {
	client := gh.octokit()
	prService, err := client.PullRequests(&octokit.PullRequestsURL, octokit.M{"owner": gh.Project.Owner, "repo": gh.Project.Name})
	if err != nil {
		return nil, err
	}

	params := octokit.PullRequestParams{Base: base, Head: head, Title: title, Body: body}
	pr, result := prService.Create(params)
	if result.HasError() {
		err = result.Err
	}

	return pr, nil
}

func (gh *GitHub) CreatePullRequestForIssue(base, head, issue string) (pr *octokit.PullRequest, err error) {
	client := gh.octokit()
	prService, err := client.PullRequests(&octokit.PullRequestsURL, octokit.M{"owner": gh.Project.Owner, "repo": gh.Project.Name})
	if err != nil {
		return nil, err
	}

	params := octokit.PullRequestForIssueParams{Base: base, Head: head, Issue: issue}
	pr, result := prService.Create(params)
	if result.HasError() {
		err = result.Err
	}

	return pr, nil
}

func (gh *GitHub) Repository(project Project) (repo *octokit.Repository, err error) {
	client := gh.octokit()
	repoService, err := client.Repositories(&octokit.RepositoryURL, octokit.M{"owner": project.Owner, "repo": project.Name})
	if err != nil {
		return nil, err
	}

	repo, result := repoService.Get()
	if result.HasError() {
		err = result.Err
	}

	return repo, nil
}

// TODO: detach GitHub from Project
func (gh *GitHub) IsRepositoryExist(project Project) bool {
	repo, err := gh.Repository(project)

	return err == nil && repo != nil
}

func (gh *GitHub) CreateRepository(project Project, description, homepage string, isPrivate bool) (repo *octokit.Repository, err error) {
	var repoURL octokit.Hyperlink
	if project.Owner != gh.Config.FetchUser() {
		repoURL = octokit.OrgRepositoriesURL
	} else {
		repoURL = octokit.UserRepositoriesURL
	}

	client := gh.octokit()
	repoService, err := client.Repositories(&repoURL, octokit.M{"org": project.Owner})
	if err != nil {
		return nil, err
	}

	params := octokit.Repository{Name: project.Name, Description: description, Homepage: homepage, Private: isPrivate}
	repo, result := repoService.Create(params)
	if result.HasError() {
		err = result.Err
	}

	return repo, nil
}

func (gh *GitHub) Releases() (releases []octokit.Release, err error) {
	client := gh.octokit()
	releasesService, err := client.Releases(nil, octokit.M{"owner": gh.Project.Owner, "repo": gh.Project.Name})
	if err != nil {
		return nil, err
	}

	releases, result := releasesService.GetAll()
	if result.HasError() {
		err = result.Err
		return nil, err
	}

	return releases, nil
}

func (gh *GitHub) CIStatus(sha string) (status *octokit.Status, err error) {
	client := gh.octokit()
	statusesService, err := client.Statuses(nil, octokit.M{"owner": gh.Project.Owner, "repo": gh.Project.Name, "ref": sha})

	statuses, result := statusesService.GetAll()
	if result.HasError() {
		err = result.Err
		return nil, err
	}

	if len(statuses) > 0 {
		status = &statuses[0]
	}

	return status, nil
}

func (gh *GitHub) ForkRepository(name, owner string, noRemote bool) (repo *octokit.Repository, err error) {
	config := gh.Config
	project := Project{Name: name, Owner: config.User}
	r, err := gh.Repository(project)
	if err == nil && r != nil {
		err = fmt.Errorf("Error creating fork: %s exists on %s", r.FullName, GitHubHost)
		return nil, err
	}

	client := gh.octokit()
	repoService, err := client.Repositories(&octokit.ForksURL, octokit.M{"owner": owner, "repo": name})
	repo, result := repoService.Create(nil)
	if result.HasError() {
		err = result.Err
		return nil, err
	}

	return repo, nil
}

func (gh *GitHub) Issues() (issues []octokit.Issue, err error) {
	client := gh.octokit()
	issuesService, err := client.Issues(&octokit.RepoIssuesURL, octokit.M{"owner": gh.Project.Owner, "repo": gh.Project.Name})
	if err != nil {
		return nil, err
	}

	issues, result := issuesService.GetAll()
	if result.HasError() {
		err = result.Err
		return nil, err
	}

	return issues, nil
}

func (gh *GitHub) ExpandRemoteUrl(owner, name string, isSSH bool) (url string) {
	project := gh.Project
	if owner == "origin" {
		config := gh.Config
		owner = config.FetchUser()
	}

	return project.GitURL(name, owner, isSSH)
}

func findOrCreateToken(user, password, twoFactorCode string) (token string, err error) {
	basicAuth := octokit.BasicAuth{Login: user, Password: password, OneTimePassword: twoFactorCode}
	client := octokit.NewClient(basicAuth)
	authsService, err := client.Authorizations(nil, nil)
	if err != nil {
		return "", err
	}

	auths, result := authsService.GetAll()
	if result.HasError() {
		err = result.Err
		return "", err
	}

	for _, auth := range auths {
		if auth.NoteURL == OAuthAppURL {
			token = auth.Token
			break
		}
	}

	if token == "" {
		authParam := octokit.AuthorizationParams{}
		authParam.Scopes = append(authParam.Scopes, "repo")
		authParam.Note = "gh"
		authParam.NoteURL = OAuthAppURL

		auth, result := authsService.Create(authParam)
		if result.HasError() {
			err = result.Err
			return "", err
		}

		token = auth.Token
	}

	return token, nil
}

func (gh *GitHub) octokit() *octokit.Client {
	config := gh.Config
	config.FetchCredentials()
	tokenAuth := octokit.TokenAuth{AccessToken: config.Token}

	return octokit.NewClient(tokenAuth)
}

func New() *GitHub {
	project := CurrentProject()
	c := CurrentConfig()
	c.FetchUser()

	return &GitHub{project, c}
}

// TODO: detach project from GitHub
func NewWithoutProject() *GitHub {
	c := CurrentConfig()
	c.FetchUser()

	return &GitHub{nil, c}
}
