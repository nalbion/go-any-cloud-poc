HANDLER ?= index
PACKAGE ?= index
FOLDER ?= lambda/
PROJECT ?= src/github.com/nalbion/go-any-cloud-poc

# For use on developer machines - runs make within a Docker container
dev:
	@docker run --rm \
		-v $(PWD):/go/$(PROJECT) \
		-e "HANDLER=$(HANDLER)" -e "PACKAGE=$(PACKAGE)" \
		-w /go/$(PROJECT) \
		nalbion/go-lambda-build make test lambda/$(PACKAGE).zip

# Useful for debugging build/CI issues
bash:
	@docker run -it \
		-v $(PWD):/go/$(PROJECT) \
		-e "HANDLER=$(HANDLER)" -e "PACKAGE=$(PACKAGE)" \
		-w /go/$(PROJECT) \
		nalbion/go-lambda-build /bin/bash

# Build the Docker image for building for AWS Lambda
docker:
	@mkdir -p lambda/_gopath/src && cp -r /usr/lib/go_appengine/goroot/src/appengine* lambda/_gopath/src/
	@docker build -t nalbion/go-lambda-build lambda
	@rm -rf lambda/_gopath
# Push the Docker image to the hub
docker-push:
	@docker push nalbion/go-lambda-build


test:
	@echo -ne "vet..."\\r
	@govendor vet +local
	@echo -ne "tests..."\\r
	@mv lambda/ .tmp/
	@mkdir -p testresults codecoverage
	@$(eval PKGS := $(shell go list ./lib/... | grep -v /vendor/))
	@$(eval PKGS_DELIM := $(shell echo $(PKGS) | sed -e 's/ /,/g'))
	@go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}}go test -test.v -test.timeout=120s -covermode=count -coverprofile={{.Name}}_{{len .Imports}}_{{len .Deps}}.coverprofile -coverpkg $(PKGS_DELIM) {{.ImportPath}}{{end}}' $(PKGS) | xargs -I {} bash -c {} | go-junit-report > testresults/$(PACKAGE).xml
	@gocovmerge `ls *.coverprofile` > cover.out
	@gocov convert cover.out | gocov-xml > codecoverage/$(PACKAGE).xml
	@rm -f *.coverprofile cover.out
	@mv .tmp lambda


# Build Lambda function packages - includes a Python shim to the Go code
lambda/$(PACKAGE).zip: clean
	@mv vendor/github.com/eawsy/ lambda/unvendor
	@go build -buildmode=plugin -ldflags='-w -s' -o lambda/$(HANDLER).so ./lambda
	@chown $(shell stat -c '%u:%g' .) lambda/$(HANDLER).so
	@cd lambda; zip -q $(PACKAGE).zip $(HANDLER).so
	@cd lambda; mv /shim $(HANDLER); zip -q -r $(PACKAGE).zip $(HANDLER); mv $(HANDLER) /shim
	@mv lambda/unvendor/ vendor/github.com/eawsy
	@chown $(shell stat -c '%u:%g' .) lambda/$(PACKAGE).zip
	@echo DONE!


.PHONY: clean test deploy docker docker-push
clean:
	@rm -rf lambda/$(PACKAGE).zip lambda/$(HANDLER).so
