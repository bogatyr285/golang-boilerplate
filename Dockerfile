# image for compiling binary
ARG BUILDER_IMAGE=golang:1.18.3-alpine3.16
# here we'll run binary app
ARG RUNNER_IMAGE=$BUILDER_IMAGE

FROM $BUILDER_IMAGE AS builder
### variables
ARG PROJECT_PATH="/go/src/github.com/bogatyr285/golang-boilerplate"
# disable golang package proxying for such modules
ARG GOPRIVATE="github.com/bogatyr285"
# key for accessing private repos
ARG GITHUB_TOKEN

RUN apk update && apk upgrade && \
    apk add --no-cache git make build-base curl jq

# install swagger
RUN curl -o /usr/local/bin/swagger -L'#' "https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64" && \
    chmod +x /usr/local/bin/swagger

# configure git to work with private repos
RUN git config --global url."https://$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"

ENV GO111MODULE on
ENV GOPRIVATE ${GOPRIVATE}

### copying project files
WORKDIR ${PROJECT_PATH}
# copy gomod 
COPY go.mod go.sum ./
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# creates build/main files
RUN make build
RUN make generate
RUN make proto

FROM ${RUNNER_IMAGE}
ARG PROJECT_PATH="/go/src/github.com/bogatyr285/golang-boilerplate"

COPY --from=builder ${PROJECT_PATH}/build/doer-api .
COPY --from=builder ${PROJECT_PATH}/config.yaml ./config.yaml

CMD ["/go/doer-api"]