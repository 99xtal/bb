BINARY_NAME=bb
VERSION=0.1.0
BUILD_DATE=2023-08-23

build:
	go build -o ${BINARY_NAME} \
		-ldflags="-X github.com/99xtal/bb/internal/build.Version=${VERSION} -X github.com/99xtal/bb/internal/build.Date=${BUILD_DATE}" \
		./cmd/bb

