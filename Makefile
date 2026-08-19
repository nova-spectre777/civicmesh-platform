.PHONY: check cpp go kotlin ts python

check:
	./scripts/check_all.sh

cpp:
	cmake -S edge-cpp -B edge-cpp/build && cmake --build edge-cpp/build && ctest --test-dir edge-cpp/build --output-on-failure

go:
	cd gateway-go && go test ./...

kotlin:
	./scripts/test_kotlin.sh

ts:
	cd web-ts && npm run build && npm test

python:
	cd simulator-python && python3 -m unittest discover -s tests -p 'test_*.py'
