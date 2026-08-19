#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
mkdir -p "$ROOT/mobile-kotlin/build"
mapfile -t SOURCES < <(find "$ROOT/mobile-kotlin/src/main/kotlin" "$ROOT/mobile-kotlin/src/test/kotlin" -name '*.kt' -print)
kotlinc "${SOURCES[@]}" -include-runtime -d "$ROOT/mobile-kotlin/build/tests.jar"
java -jar "$ROOT/mobile-kotlin/build/tests.jar"
