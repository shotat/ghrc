# GHRC

GHRC is a tool for managing **G**it**H**ub **R**epository **C**onfigurations in a declarative way.

Repository configuration includes

- Labels
- Protected Branches
- Visibility(private or public)
- Merging Strategies
- Topics
- Description

## Installation

```sh
$ go get github.com/shotat/ghrc
```

## Environment Variables

- `GHRC_GITHUB_TOKEN`
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

| Field Name | Type | Description |
| ---------- | ---- | ----------- |
| metadata | Metadata Object | **Required**. Provides metadata about the repository. |
| spec | Spec Object | **Required**. A desired state for the repository. |

**Spec Object**

**Metadata Object**

