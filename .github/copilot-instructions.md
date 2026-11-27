# Copilot instructions

The full guide is in [AGENTS.md](../AGENTS.md), and all style rules are in [STYLE.md](../STYLE.md). Read both before you write code. The most important rules:

- Work on one existing issue per pull request, and start the description with `# Fixes issue #N`.
- Change only the lines needed to fix the issue. Do not reformat, rename, refactor, or edit unrelated files.
- Every bug fix adds a test case that fails before the fix.
- Do not change existing expected test values to make code pass.
- Run `make fmt`, `make vet`, `make test`, and `golangci-lint run --config=.github/linters/.golangci.yml ./... ./.github/cmd/...` before pushing.
- Fill in the pull request template and say whether AI was used.
