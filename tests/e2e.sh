#!/usr/bin/env bash

set -eoux pipefail

assert_empty() {
  OUTPUT="$(same)"
  if [[ "${OUTPUT}" ]]; then
    echo 'same listed branches'
    echo "${OUTPUT}"
    exit 1
  fi
}

assert_lists_branch() {
  OUTPUT="$(same)"
  if [[ ! "${OUTPUT}" == $1 ]]; then
    echo "same didn't list exactly '$1'"
    echo "${OUTPUT}"
    exit 1
  fi
}

TMPDIR="$(mktemp -d)"
echo "Tests will run in ${TMPDIR}"
pushd "${TMPDIR}"

# TODO: print better error message when called in a directory without a git repository

git init

# TODO: print better error message when called in a repository with no commits

echo 'a' >a.txt

# TODO: print better error message when called in a repository with no commits

git add a.txt

# TODO: print better error message when called in a repository with no commits

git commit -m "Initial commit"

assert_empty

git checkout -B new-branch

assert_lists_branch 'new-branch'

echo 'b' >b.txt

assert_lists_branch 'new-branch'

git add b.txt

assert_lists_branch 'new-branch'

git commit -m "Further commit"

assert_empty

git checkout master

assert_empty

popd
echo 'Tests finished'
