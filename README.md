# same

same lists Git branches which have an empty diff against the `master` branch.

This includes branches which may have at one point been different to master, but have since had the `master` branch merged in or have otherwise "caught up".

## Installation

```shell script
go install github.com/jameshiew/same
```

## Usage

Useful for piping into `xargs` for executing further Git commands e.g.

### Show branches with no diff

```shell script
same | xargs git show-branch
```

### Delete branches with no diff

```shell script
same | xargs git branch -D
```
