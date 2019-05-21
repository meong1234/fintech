.PHONY: all test clean
all: build fmt vet lint test

APP=fintech_service
DB_USER="postgres"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell  go list ./... | grep -v "vendor")

DB_NAME=$(APP)
TEST_DB_NAME="$(APP)_test"

APP_EXECUTABLE="./out/$(APP)"

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/lint/golint
	go get -u github.com/axw/gocov/gocov
	go get -u gopkg.in/matm/v1/gocov-html

db.setup: db.create db.migrate

db.create:
	createdb -O$(DB_USER) -Eutf8 $(DB_NAME)

db.migrate:
	cp -r resources out
	ENVIRONMENT=development $(APP_EXECUTABLE) migrate

db.drop:
	dropdb --if-exists -U$(DB_USER) $(DB_NAME)

db.reset: db.drop db.create db.migrate

testdb.migrate:
	cp -r resources out
	ENVIRONMENT=test $(APP_EXECUTABLE) migrate

testdb.create: testdb.drop
	createdb -O$(DB_USER) -Eutf8 $(TEST_DB_NAME)

testdb.drop:
	dropdb --if-exists -U$(DB_USER) $(TEST_DB_NAME)

build-deps:
	dep ensure -v

update-deps:
	dep ensure

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

build: build-deps compile

install:
	go install ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test:
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -p=1

copy-config:
	cp application.yml.sample application.yml

test-cov: testdb.migrate
	gocov test ${ALL_PACKAGES} > docs/cov.json

test-cov-html:
	@echo "\nEXPORTING RESULTS TO COVERAGE.HTML..."
	gocov-html docs/cov.json > docs/coverage.html
	@echo 'TEST RESULTS EXPORTED TO DOCS/COVERAGE.HTML'

test-cov-report:
	@echo "\nGENERATING TEST REPORT."
	gocov report docs/cov.json

remove-cov-json:
	rm docs/cov.json

clean:
	rm -rf ./out/

