# GHRC

GHRC is a tool for managing **G**it**H**ub **R**epository **C**onfigurations in a declarative way.

Repository configuration includes

- Labels
- Protected Branches
- Merging Strategies
- Visibility(private or public)
- Topics
- Repository Metadata (Description, Homepage, etc.)

## Installation

```sh
$ go get github.com/shotat/ghrc
```

## Environment Variables

- `GHRC_GITHUB_TOKEN` (**Required**)
  - Your GitHub API Token
  - `repo` scope is required
- `GHRC_GITHUB_API`
  - GitHub API BaseURL
    - Default: `https://api.github.com/`
    - Enterprise: `https://<domain>/api/v3/`

## Usage

### Import an existing repository state

```sh
# ghrc import --owner <repository owner name> --name <repository name>
$ ghrc import --owner shotat --name ghrc > .ghrc.yaml
```

### Check expected changes without changing the actual state

```sh
$ ghrc plan
```

This command reads `.ghrc.yaml` file implicitly.
You can specify a config file explicitly with `-f`.

```sh
$ ghrc plan -f .ghrc.yaml
```

### Example

```
$ ghrc plan

~ Repo: shotat/ghrc
{*state.Repo}.Description:
	-: "GHRC is a tool for managing GitHub Repository Configurations in a declarative way."
	+: "GHRC is an awesome tool for managing GitHub Repository Configurations in a declarative way."
{*state.Repo}.AllowRebaseMerge:
	-: false
	+: true
Sort({*state.Repo}.Topics)[1->?]:
	-: "git"
	+: <non-existent>

+ Label: awesome
{*state.Label}:
	-: (*state.Label)(nil)
	+: &state.Label{Name: "awesome", Description: "Awesome issue", Color: "d73a4b"}

~ Label: bug
{*state.Label}.Color:
	-: "d73a4a"
	+: "d73a4b"

~ Label: documentation

~ Label: review

- Label: question
{*state.Label}:
	-: &state.Label{Name: "question", Description: "Further information is requested", Color: "d876e3"}
	+: (*state.Label)(nil)

~ Protection: master
{*state.Protection}.RequiredStatusChecks.Strict:
	-: false
	+: true
{*state.Protection}.RequiredStatusChecks.Contexts[0->?]:
	-: "ci/dockercloud"
	+: <non-existent>
{*state.Protection}.RequiredPullRequestReviews.RequiredApprovingReviewCount:
	-: 1
	+: 2
```


### Apply specs to the actual state


```sh
$ ghrc apply
```

This command reads `.ghrc.yaml` file implicitly.
You can specify a config file explicitly with `-f`.

```sh
$ ghrc apply -f .ghrc.yaml
```

## Schema

**Root**

| Field | Type | Description |
| ---------- | ---- | ----------- |
| metadata | Metadata Object | **Required**. Provides metadata about the repository. |
| spec | Spec Object | **Required**. A desired state for the repository. |

**Spec Object**

| Field | Type | Description |
| ---------- | ---- | ----------- |
| owner | string | **Required**. A repository owner name. If the repository is owned by a user, specify user login name. Or if owned by a organization, specify organization login name. |
| name | string | **Required**. A repository name. |

**Metadata Object**

| Field | Type | Description |
| ---------- | ---- | ----------- |
| repo | Repo Object | General configurations for the repository. |
| labels | [Label Object] | List of label configurations. |
| protections | [Protection Object] | List ob protection configurations. |

**Repo Object**

Check details in [GitHub API docs: Repositories](https://developer.github.com/v3/repos/)

| Field | Type | Description |
| ---------- | ---- | ----------- |
| description | string | A short description of the repository. |
| homepage | string | A URL with more information about the repository. |
| private | boolean | Either true to make a repository private or false to make public. |
| allowMergeCommit | boolean | Either true to allow merging pull requests with a merge commit, or false to prevent merging pull requests with merge commits. |
| allowSquashMerge | boolean | Either true to allow squash-merging pull requests, or false to prevent squash-merging. |
| allowRebaseMerge | boolean | Either true to allow rebase-merging pull requests, or false to prevent rebase-merging. |
| topics | [string] | An array of topics. **Note: Topic names cannot contain uppercase letters.** |

**Label Object**

Check details in [GitHub API docs: Labels](https://developer.github.com/v3/issues/labels/)

| Field | Type | Description |
| ---------- | ---- | ----------- |
| name | string | The name of the label. |
| description | string | A short description of the label. |
| color | string | The [hexadecimal color code](https://www.color-hex.com/) for the label, without the leading `#`. |

**Protection Object**

Check details in [GitHub API docs: Branches](https://developer.github.com/v3/repos/branches/)

| Field | Type | Description |
| ---------- | ---- | ----------- |
| branch | string | A branch name to be protected. |
| enforceAdmins | boolean | Enforce all configured restrictions for administrators. |
| requiredStatusChecks | RequiredStatusChecks Object | Require status checks to pass before merging. |
| requiredPullRequestReviews | RequiredPullRequestReviews Object | Require at least one approving review on a pull request, before merging. |

**RequiredStatusChecks Object**

| Field | Type | Description |
| ---------- | ---- | ----------- |
| strict | boolean | Require branches to be up to date before merging. |
| contexts | [string] | The list of status checks to require in order to merge into this branch. |

**RequiredPullRequestReviews Object**

| Field | Type | Description |
| ---------- | ---- | ----------- |
| dismissStaleReviews | boolean | Set to true if you want to automatically dismiss approving reviews when someone pushes a new commit. |
| requireCodeOwnerReviews | boolean | Blocks merging pull requests until code owners review them. |
| requiredApprovingReviewCount | int | Specify the number of reviewers required to approve pull requests. Use a number between 1 and 6. |

### Example

```yaml
metadata:
  owner: shotat
  name: ghrc
spec:
  repo:
    description: GHRC is a tool for managing GitHub Repository Configurations in a declarative way.
    homepage: "https://github.com/shotat/ghrc"
    private: false
    allowMergeCommit: false
    allowSquashMerge: true
    allowRebaseMerge: false
    topics:
    - cli
    - git
    - github
    - go
  labels:
  - name: bug
    description: Something isn't working
    color: d73a4a
  - name: documentation
    description: Improvements or additions to documentation
    color: 0075ca
  - name: question
    description: Further information is requested
    color: d876e3
  - name: review
    description: "ready for review"
    color: f0f0f0
  protections:
  - branch: master
    requiredStatusChecks:
      strict: false
      contexts:
      - ci/dockercloud
    enforceAdmins: false
    requiredPullRequestReviews:
      dismissStaleReviews: false
      requireCodeOwnerReviews: false
      requiredApprovingReviewCount: 1
```
