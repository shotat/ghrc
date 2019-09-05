# GHRC

GHRC is a tool for managing **G**it**H**ub **R**epository **C**onfigurations in a declarative way.

Repository configuration includes

- Visibility
- Labels
- Topics
- Protected Branches
- Description

## Environment Variables

- `GHRC_GITHUB_TOKEN`
- `GHRC_GITHUB_API` (default: "https://api.github.com")

## Usage

### Import an existing repository state

```sh
# ghrc import --owner <repository owner name> --name <repository name>
$ ghrc import --owner shotat --name ghrc
```

### Check expected changes without changing the actual state.

```sh
$ ghrc plan
```

### Apply specs to the actual state.

```sh
$ ghrc apply
```

