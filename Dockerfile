#Stage 1 build and test
#docker.io prefix required by podman
# use podman build . --build-arg BUILD_VERSION="jikjikjik" --build-arg BUILD_HASH="0001100"
FROM docker.io/golang:alpine as builder
ARG BUILD_HEADTAG
ARG BUILD_HASH
ARG BUILD_BRANCH
RUN mkdir /build
WORKDIR /build
COPY go.mod .
COPY go.sum .

# TODO this doesn't work with podman 3.x but does with 4.x
#RUN --mount=type=cache,target=/root/.cache go mod download
RUN go mod download
COPY . .
#RUN --mount=type=cache,target=/root/.cache make build
RUN apk --no-cache add gcc build-base git
RUN make build HEAD_TAG="$BUILD_HEADTAG" VERSION_HASH=$BUILD_HASH BRANCH_NAME=$BUILD_BRANCH

# test that that the build is good and app launches
RUN /build/bin/previewd version

#RUN go test -v

# generate clean, final image for end users
FROM alpine:3.11.3
RUN apk update
RUN apk add git
COPY --from=builder /build/bin/pocketshorten .

# executable
ENTRYPOINT [ "./pocketshorten" ]
CMD ["testserver"]
# arguments that can be override
#CMD [ "3", "300" ]
