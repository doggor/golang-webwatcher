# Build the app into Docker image.
#
# Build command:
# > docker build -t doggor/webwatcher:{version} .
#
# Launch command:
# > docker run --rm doggor/webwatcher:{version}
#
# Ref: https://blog.wu-boy.com/2017/04/build-minimal-docker-container-using-multi-stage-for-go-app/

# Build stage
FROM golang:alpine AS build-stage
COPY . /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/webwatcher/main.go

# Pack stage
FROM centurylink/ca-certs
COPY --from=build-stage /src/app /
ENTRYPOINT ["/app"]
