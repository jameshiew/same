# same
same lists Git branches which have an empty diff against `master` branch.

This includes branches which may have at one point been different to master, but have since had the `master` branch merged in or have otherwise "caught up".

Useful for piping into `xargs` for executing further Git commands e.g.

```shell script
./same | xargs git show-branch
```
