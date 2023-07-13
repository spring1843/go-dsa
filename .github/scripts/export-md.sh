#!/usr/bin/env bash
# Exports the markdown files from the docs site
set -euo pipefail

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
source "${SCRIPT_DIR}/sections.sh"

declare -a FILES=(
  README.md
  complexity.md
)

FILES+=("${SECTIONS[@]/%//README.md}")

for file in "${FILES[@]}"; do
   cat "$file"
   printf "\n"
done
