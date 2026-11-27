# Style guide

This guide describes how code, tests, and docs in this repository are written. New code should look like the code around it. When this guide and a file disagree, follow this guide in new code and leave the existing file as it is.

## Files and packages

- Each problem has a solution in `topic/name.go` and its tests in `topic/name_test.go`, in the same package. For example [array/reverse_inplace.go](./array/reverse_inplace.go) and [array/reverse_inplace_test.go](./array/reverse_inplace_test.go).
- File names are `snake_case`. Function and type names are `CamelCase` and describe the problem, such as `NetworkDelayTime` or `IntToRoman`.
- A solution can be unexported when the tests are its only caller, such as `put` and `get` in [linkedlist/lru_cache.go](./linkedlist/lru_cache.go).
- Use the standard library only.

## Comments

- The function that solves a problem has one doc comment line in this form:

  ```Go
  // NetworkDelayTime solves the problem in O(n*Log n) time and O(n) space.
  ```

  Write complexity as `O(1)`, `O(n)`, `O(Log n)`, `O(n*Log n)`, `O(n^2)`, `O(2^n)`, or `O(n!)`, and use other letters such as `k`, `V`, or `E` when they describe the input better.
- Every other exported type, function, method, variable, and struct field has a doc comment of one short sentence that starts with its name.
- Errors are package variables with a doc comment that says when they occur:

  ```Go
  // ErrQueueEmpty occurs when dequeuing an empty queue.
  var ErrQueueEmpty = errors.New("empty queue")
  ```

- Unexported helpers usually have no doc comment. Add one only when the name does not say what the helper does. It starts with the helper's name.
- Do not add comments inside function bodies. Add one only when a careful reader would get the code wrong without it, as in [dnc/rate_limit.go](./dnc/rate_limit.go).
- Do not leave commented out code, TODOs, or notes about the change you made.

## Code

- Keep functions short. Return early for empty or trivial input, then handle the general case.
- Prefer what the standard library and the language already provide: the builtin `min` and `max`, `slices`, `maps`, `cmp`, `strings.Builder`, `strings.SplitSeq`, and `for i := range n`. Write `any`, not `interface{}`.
- Reuse the types, helpers, and constants that already exist in the package before adding new ones.
- Give variables names that say what they hold, such as `maxValue`, `revenues`, or `queue`. Use `i`, `j`, and `k` for indexes and loop counters only.
- Use `var name T` for a zero value that a loop fills in, and `:=` everywhere else.
- Group related constants and unexported types in `const (...)` and `type (...)` blocks. Use `iota` for enumerations.
- Use `make` with a size or capacity when it is known.
- Return errors instead of panicking. Wrap errors with `fmt.Errorf("...: %w", err)`.
- Keep receiver names the same within a type. `container/heap` types use one letter receivers and one line methods:

  ```Go
  func (m minHeap) Len() int           { return len(m) }
  func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
  func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
  func (m *minHeap) Push(x any)        { *m = append(*m, x.(int)) }
  ```

- Use blank lines only to separate steps of an algorithm.
- Formatting and import order are enforced by `gofumpt` and `gci`. Run `make fmt` and do not format code by hand.

## Tests

- Each test function starts with a block comment that states the signature and the problem:

  ```Go
  /*
  TestReverseInPlace tests solution(s) with the following signature and problem description:

  	ReverseInPlace(list []int, start, end int)

  Given a slice of integers, a start index, and an end index, reverse the integers in the in-place
  without using any extra memory.

  For example given {1,2,3,4,5,6} and start of 0 and end of 4, return {5,4,3,2,1,6} because:
  ...
  */
  ```

  Indent the signature and any ASCII diagram with a tab. Write the problem as a short paragraph followed by one example that starts with "For example given".
- Tests are tables of anonymous structs with positional values. Name the fields after what they hold, such as `edges`, `numberOfNodes`, or `shortestPath`, and not `want` or `expected`. Name an error field `expectedErr`.
- Order cases from the simplest, such as empty input, to the most complex.
- Loop over the table with `for i, test := range tests` and stop at the first failure:

  ```Go
  if got := NetworkDelayTime(test.numberOfNodes, test.messageFrom, test.edges); got != test.shortestPath {
  	t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.shortestPath, got)
  }
  ```

- Compare slices with `slices.Equal` and errors with `errors.Is`.
- Put test helpers at the bottom of the test file as unexported functions.
- Do not use `t.Run`, `t.Parallel`, or third party test libraries.

## READMEs

Each topic directory has a `README.md` that teaches the topic and lists its problems. The exporter in [.github/cmd](./.github/cmd) turns these files into a book, so parts of the structure below are required for it to work.

### Sections

A topic README has a title and four standard sections, in this order:

```md
# Stack

Introduction, with an optional figure.

## Implementation

## Complexity

## Application

## Rehearsal
```

- **Title**: the name of the topic, such as `# Stack` or `# Dynamic Programming`.
- **Introduction**: no heading. Start with a definition of the topic in one or two paragraphs. Then give a real world analogy, such as a backpack for a stack or cars at a toll booth for a queue. Add a figure when a picture helps.
- **Implementation**: how to implement or use the topic in Go, with one or more runnable examples. Prefer the standard library, such as `container/list` or `container/heap`, when it has what the topic needs.
- **Complexity**: the time and space complexity of the main operations, with the reason for each. Compare with other topics when it helps the reader choose between them.
- **Application**: the kinds of problems the topic solves and where it is used in real systems.
- **Rehearsal**: the list of problems. It is always the last section.

A topic with several variations can add `###` subsections under Implementation, such as `### Top-Down` and `### Bottom-Up` in [dp/README.md](./dp/README.md). It can also add its own `##` sections between Implementation and Complexity, such as `## Arithmetic by Shifting` in [bit/README.md](./bit/README.md). A topic that covers several structures, such as [tree/README.md](./tree/README.md), can give each structure its own `##` section with a `### ... Complexity` subsection. Do not add sections after Rehearsal.

### Rehearsal

Each problem is one line in this exact form, because the exporter matches it with a regular expression:

```md
- [Max Stack](./max_stack_test.go), [Solution](./max_stack.go)
```

The title is in Title Case, and the first link is the test file because the test describes the problem. Add new problems to the end of the list, and add the same title to the table of contents in the root [README.md](./README.md).

### Figures

Draw figures as text in a code block marked `ASCII`. The first line is a numbered caption, followed by a blank line:

```ASCII
[Figure 1] Push 1,2,3,4 to a stack and then pop 4 times.

┌───┐┌───┐
│   ││ 1 │
└───┘└───┘
```

Number figures from 1 in each README, and refer to them in the text as _Figure 1_. Tests can refer to them too, as in [graph/topological_sort_test.go](./graph/topological_sort_test.go).

### Code examples

- Mark Go code blocks as `Go`.
- Each example is a complete `package main` program that runs on its own, so readers can paste it into the Go playground.
- Keep examples short and focused on one idea. Leave out error handling that does not help explain the topic.
- Show results with trailing comments: `// Prints [0 1 2 3]` for output, or `// [1,2]` for the state of a data structure after each line.

### Links

- Link to other topics with relative links that start with `../`, such as `[heap](../heap/)`, and to files in the same topic with links that start with `./`. The exporter rewrites these into links to GitHub, and other forms break in the book.
- Link to Go documentation and the Go blog for language features, such as `[slices](https://go.dev/blog/slices)`.

### Writing

- Write short, plain sentences in US English.
- Explain with examples and numbers instead of general claims.
- Use the full name with the abbreviation the first time a term appears, such as Last-In-First-Out (LIFO), then the abbreviation.
- Write complexity in the same notation as the code, such as O(n) and O(Log n).
