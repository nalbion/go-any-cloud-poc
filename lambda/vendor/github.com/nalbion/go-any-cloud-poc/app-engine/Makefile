PROJECT ?= src/github.com/nalbion/go-any-cloud-poc

# For use on developer machines - runs make within a Docker container
dev:
	@docker run --rm \
		-v $(PWD):/go/$(PROJECT) \
		-w /go/$(PROJECT) \
		nalbion/go-app-engine-build make test app-engine/app.zip

bash:
	@docker run -it \
		-v $(PWD):/go/$(PROJECT) \
		-w /go/$(PROJECT) \
		nalbion/go-app-engine-build /bin/bash

# Build the Docker image for building for AWS Lambda
docker:
	@docker build -t nalbion/go-app-engine-build app-engine
# Push the Docker image to the hub
docker-push:
	@docker push nalbion/go-app-engine-build

test:
	@echo -ne "vet..."\\r
	#@govendor update +local
	@govendor vet +local
	@echo -ne "tests..."\\r
	@mv app-engine/ .tmp/
	@mkdir -p testresults codecoverage
	@$(eval PKGS := $(shell go list ./lib/... | grep -v /vendor/))
	@$(eval PKGS_DELIM := $(shell echo $(PKGS) | sed -e 's/ /,/g'))
	@go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}}go test -test.v -test.timeout=120s -covermode=count -coverprofile={{.Name}}_{{len .Imports}}_{{len .Deps}}.coverprofile -coverpkg $(PKGS_DELIM) {{.ImportPath}}{{end}}' $(PKGS) | xargs -I {} bash -c {} | go-junit-report > testresults/app.xml
	@gocovmerge `ls *.coverprofile` > cover.out
    #@govendor test +local -v -coverprofile=	codecoverage/cover.out -coverpkg=lib lib | go-junit-report > testresults/app.xml
#	@go test $(PKGS) | go-junit-report > testresults/app.xml
	@gocov convert cover.out | gocov-xml > codecoverage/app.xml
	@rm -f *.coverprofile cover.out
	@mv .tmp app-engine


# Build Lambda function packages - includes a Python shim to the Go code
app-engine/app.zip: clean
	@mv lambda/ .tmp/
	@go build -o app-engine/app.so ./app-engine
	@chown $(shell stat -c '%u:%g' .) app-engine/app.so
	@cd app-engine; zip -q app.zip app.so
	@chown $(shell stat -c '%u:%g' .) app-engine/app.zip
	@mv .tmp/ lambda
	@echo DONE!


.PHONY: clean test deploy docker docker-push
clean:
	@rm -rf app-engine/app.zip app-engine/app.so
