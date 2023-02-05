GO_FILES=cmd
BUILD_DIR=build

build:
	for file in $$(ls ${GO_FILES}); do \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BUILD_DIR}/$$file ${GO_FILES}/$$file/main.go; \
		zip -j ${BUILD_DIR}/$$file.zip ${BUILD_DIR}/$$file; \
	done

.PHONY: clean
clean:
	rm -rf build

.PHONY: fmt
fmt:
	go fmt ./...
	terraform fmt -recursive

.PHONY: deploy
deploy:
	terraform -chdir=terraform/prod get
	terraform -chdir=terraform/prod apply
