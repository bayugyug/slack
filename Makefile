VERSION=v0.0.1
TODAY = $(shell date +"%Y%m%d_%H%M%S")
BUILDFLAGS=
NAME=slack


clean:
	rm -f ${NAME}
	rm -f /tmp/test.coverprofile
	find . -type f -name '*.coverprofile' -exec rm -f {} \;
	find . -type f -name '*coverage.xml'  -exec rm -f {} \;
	find . -type f -name '*coverage.txt'  -exec rm -f {} \;

ci-test: clean
	@./scripts/prep.sh ci-test $(NAME)

ci-lint: clean
	@./scripts/prep.sh ci-lint $(NAME)

vet:
	@go vet ${BUILDFLAG} ./...

ci-cover:
	@./scripts/ci-cover.sh

codecov:
	@./scripts/codecov.sh

proto:
	go generate ./...

coverage:
	@./scripts/coverage.sh

cover:
	go test -v ./... -cover -coverprofile=test.coverprofile
	gocov convert test.coverprofile | gocov-xml > coverage.xml
%:
	@:
args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

try:
	@echo $(call args,defaultstring)

simulate:
	@go run cmd/simulator/main.go

.PHONY: clean lint ci-test ci-cover proto cover codecov

init:
	chmod +x .githooks/pre-commit
	git config core.hooksPath .githooks
