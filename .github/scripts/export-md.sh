#!/usr/bin/env bash
set -euo pipefail

# This script is used to export the markdown files from the docs site
declare -a files=(
  README.md
  complexity.md
  ./array/README.md
  ./strings/README.md
  ./linkedlist/README.md
  ./stack/README.md
  ./queue/README.md
  ./hashtable/README.md
  ./tree/README.md
  ./heap/README.md
  ./recursion/README.md
  ./dnc//README.md
  ./bit//README.md
  ./backtracking/README.md
  ./graph/README.md
  ./greedy/README.md
  ./dp/README.md
)

for file in "${files[@]}"; do
   cat "$file"
   printf "\n"
done