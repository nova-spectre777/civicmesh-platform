#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
echo "[1/6] C++20 edge relay"
cmake -S "$ROOT/edge-cpp" -B "$ROOT/edge-cpp/build" >/dev/null
cmake --build "$ROOT/edge-cpp/build" -j2 >/dev/null
ctest --test-dir "$ROOT/edge-cpp/build" --output-on-failure
echo "[2/6] Go gateway"
(cd "$ROOT/gateway-go" && go test ./...)
echo "[3/6] Kotlin offline client"
"$ROOT/scripts/test_kotlin.sh"
echo "[4/6] TypeScript PWA/domain"
(cd "$ROOT/web-ts" && rm -rf dist && npm run build && npm test)
echo "[5/6] Python disaster simulator"
(cd "$ROOT/simulator-python" && PYTHONPATH=. python3 -m unittest discover -s tests -p 'test_*.py' && PYTHONPATH=. python3 -m civicmesh_sim.cli)
echo "[6/6] Go -> TypeScript signed-alert interoperability"
FIXTURE="$(mktemp)";trap 'rm -f "$FIXTURE"' EXIT
(cd "$ROOT/gateway-go" && go run ./cmd/civicmesh-fixture) > "$FIXTURE"
node "$ROOT/scripts/verify_fixture.mjs" "$FIXTURE"
echo "CIVICMESH full verification passed"
