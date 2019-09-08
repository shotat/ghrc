# GHRC

*NOTE: work in progress*

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
- `GHRC_GITHUB_API`
  - GitHub API BaseURL
    - Default: `https://api.github.com/`
    - Enterprise: `https://<domain>/api/v3/`

## Usage

### Import an existing repository state.

```sh
# ghrc import --owner <repository owner name> --name <repository name>
$ ghrc import --owner shotat --name ghrc > .ghrc.yaml
```

### Check expected changes without changing the actual state.

```sh
$ ghrc plan
```

### Apply specs to the actual state.

```sh
$ ghrc apply
```

