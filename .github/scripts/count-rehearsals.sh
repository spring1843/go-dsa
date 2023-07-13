#!/usr/bin/env bash
# Counts the number of rehearsal files in each section
set -euo pipefail

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
source "${SCRIPT_DIR}/sections.sh"

ALL_TESTS=0
for section in "${sections[@]}"; do
   SECTION_TESTS=$(ls -l "$section" | grep "_test.go" | wc -l)
   ALL_TESTS=$((ALL_TESTS + SECTION_TESTS))
done
echo $ALL_TESTS