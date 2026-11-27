# Guide for AI agents

This file is for AI coding agents that contribute to this repository. Read [CONTRIBUTING.md](./CONTRIBUTING.md) as well. The rules below are stricter and take precedence.

Read [STYLE.md](./STYLE.md) before you write code, tests, or docs. All style rules are there.

## Scope

- Work on an existing GitHub issue only. One issue per pull request.
- Do not open pull requests for refactors, renames, style changes, new problems, or dependency updates that no issue asks for.
- Start the pull request description with `# Fixes issue #N`.

## Keep the diff small

- Change only the lines needed to fix the issue. Do not reformat, reorder, rename, or tidy code around it, and do not edit unrelated files.
- Keep the existing approach of a solution unless the issue is about the approach. Fix it, do not rewrite it.
- Update the complexity in the doc comment only when the complexity changes.

## Tests

- Every bugfix adds a test case that fails before the fix and passes after it. Say in the pull request how it failed.
- Do not change an expected value in an existing test to make the code pass. If the value is wrong, explain why in the pull request.
- Keep test coverage at 100%. Check the package you changed with `go test -cover ./topic/`, which must report `coverage: 100.0% of statements`. If it does not, `make test` writes `profile.cov`, and `go tool cover -func=profile.cov | grep -v '100.0%'` lists the functions that are not fully covered.

## Before you push

Run these and fix every failure:

```bash
make fmt
make vet
make test
golangci-lint run --config=.github/linters/.golangci.yml ./... ./.github/cmd/...
```

- Do not commit `profile.cov` or files in `bin/`.
- Do not add `//nolint` unless it names the linter and gives a reason.

## Pull request description

- Fill in every section of the [pull request template](./.github/pull-request-template.md).
- Write short, plain sentences. Describe what was wrong, what changed, and which test covers it. Pull requests #219 and #220 are good examples.
- Answer the AI question in the template, and name the tool.
- Be honest under downsides. "None that I know of" is fine when it is true.

## New problems

Add a problem only when a maintainer has agreed to it in an issue. Link it in two places, because the `.github/cmd` exporter reads them:

- the `## Rehearsal` section of the topic `README.md`, as `- [Title](./name_test.go), [Solution](./name.go)`
- the table of contents in the root [README.md](./README.md)
