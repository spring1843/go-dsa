#!/usr/bin/env bash
# Exports the markdown files from the docs site
set -euo pipefail

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
# shellcheck source=/dev/null
source "${SCRIPT_DIR}/sections.sh"

declare -a files=(
  README.md
  complexity.md
)
files+=("${sections[@]/%//README.md}")

for file in "${files[@]}"; do
   cat "$file"
   printf "\n"
done
