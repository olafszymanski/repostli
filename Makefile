GO_FILES=cmd
BUILD_DIR=build

build:
	for dir in $$(ls ${GO_FILES}); do \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BUILD_DIR}/main ${GO_FILES}/$$dir/main.go; \
		zip -j ${BUILD_DIR}/$$dir.zip ${BUILD_DIR}/main; \
		rm -rf ${BUILD_DIR}/main; \
	done

.PHONY: clean
clean:
	rm -rf build
