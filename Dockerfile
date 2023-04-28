# =========================================
# Build stage
# =========================================
FROM --platform=$BUILDPLATFORM golang:1.20-alpine3.17 as build

ARG TARGETOS
ARG TARGETARCH
ARG BUILD_DATE
ARG COMMIT_SHA
ARG VERSION

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify && go mod tidy

RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH \
    go build \
    -ldflags="-s -w \
    -X /internal/version.BuildDate=${BUILD_DATE} \
    -X github.com/brpaz/copier-run/internal/version.Version=${VERSION} \
    -X github.com/brpaz/copier-run/internal/version.GitCommit=${COMMIT_SHA} \
    -extldflags '-static'" -a \
    -o /out/copier-run ./main.go

# ====================================
# Production stage
# ====================================
FROM alpine:3.17

COPY --from=build /out/copier-run /bin

ENTRYPOINT ["/bin/copier-run"]
