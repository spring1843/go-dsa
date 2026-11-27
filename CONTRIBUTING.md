# Contributing

Thank you for helping improve go-dsa. This guide covers how to report a problem and how to send a change.

## Start with an issue

Open an issue with the [problem report template](./.github/ISSUE_TEMPLATE/bug_report.md), or find an existing one, before you write code. Each pull request fixes one issue. For a new problem, wait until a maintainer agrees to it in the issue.

## Make the change

1. Fork the repository and create a branch.
2. Keep the change as small as possible, and leave code that the issue does not mention as it is.
3. Follow [STYLE.md](./STYLE.md) for code, tests, and docs.
4. For a bugfix, add a test case that fails before the fix.
5. For a new problem, link it in the `## Rehearsal` section of the topic `README.md` and in the table of contents of the root [README.md](./README.md).

## Run the checks

```bash
make fmt
make vet
make test
golangci-lint run --config=.github/linters/.golangci.yml ./... ./.github/cmd/...
```

The tests must pass and coverage must stay at 100%. To check the coverage of the package you changed:

```bash
go test -cover ./topic/
```

It must report `coverage: 100.0% of statements`. If it does not, find the code that is not covered:

```bash
make test
go tool cover -func=profile.cov | grep -v '100.0%'
go tool cover -html=profile.cov
```

The first `go tool cover` command lists the functions that are not fully covered, and the second opens a page in your browser that highlights the lines that no test runs.

## Open the pull request

Fill in the [pull request template](./.github/pull-request-template.md) with what was wrong, what changed, and which test covers it.

If you used AI tools, follow [AGENTS.md](./AGENTS.md) and say so in the pull request. You are responsible for reviewing and running the code before you submit it.
