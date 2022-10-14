.ONESHELL:
.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)
VERSION = "0.2.9"
LD_FLAGS = -X beryju.io/gravity/pkg/extconfig.Version=${VERSION}
GO_FLAGS = -ldflags "${LD_FLAGS}" -v
SCHEMA_FILE = schema.yml

ci--env:
	echo "##[set-output name=sha]${GITHUB_SHA}"
	echo "##[set-output name=build]${GITHUB_RUN_ID}"
	echo "##[set-output name=timestamp]$(shell date +%s)"
	echo "##[set-output name=version]${VERSION}"

docker-build:
	go build \
		-ldflags "${LD_FLAGS} -X beryju.io/gravity/pkg/extconfig.BuildHash=${GIT_BUILD_HASH}" \
		-v -a -o gravity .

run:
	export INSTANCE_LISTEN=0.0.0.0
	export DEBUG=true
	export LISTEN_ONLY=true
	go run ${GO_FLAGS} . server

web-build:
	cd web
	npm ci
	npm version ${VERSION} || true
	npm run build

gen-build:
	DEBUG=true go run ${GO_FLAGS} . generateSchema ${SCHEMA_FILE}
	git add ${SCHEMA_FILE}

gen-clean:
	rm -rf gen-ts-api/

gen-client-ts:
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		openapitools/openapi-generator-cli:v6.0.0 generate \
		-i /local/${SCHEMA_FILE} \
		-g typescript-fetch \
		-o /local/gen-ts-api \
		--additional-properties=typescriptThreePlus=true,supportsES6=true,npmName=gravity-api,npmVersion=${VERSION} \
		--git-repo-id BeryJu \
		--git-user-id gravity
	cd gen-ts-api && npm i

gen-client-ts-update: gen-client-ts
	cd gen-ts-api
	npm publish
	cd ../web
	npm i gravity-api@${VERSION}
	git add package*.json
	npm version ${VERSION} || true

gen: gen-build gen-clean gen-client-ts-update

test-env-start:
	cd hack/tests/
	docker compose --project-name gravity-test-env up -d

test-env-stop:
	cd hack/tests/
	docker compose --project-name gravity-test-env down -v

test:
	export BOOTSTRAP_ROLES="dns;dhcp;api;discovery;backup;debug;tsdb"
	export ETCD_ENDPOINT="localhost:2379"
	export DEBUG="true"
	etcdctl del --prefix / || true
	go test -p 1 -race -coverprofile=coverage.txt -covermode=atomic -v ./...
	go tool cover -html coverage.txt -o coverage.html
