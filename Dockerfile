FROM golang:1.24.1-alpine3.20@sha256:3d9132b88a6317b846b55aa8e821821301906fe799932ecbc4f814468c6977a5 as skeleton-builder
WORKDIR /go/src/app
RUN apk update && apk add --no-cache git
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg/mod CGO_ENABLED=0 go build -v ./cmd/skeleton/

FROM cgr.dev/chainguard/static:latest@sha256:95a45fc5fda9aa71dbdc645b20c6fb03f33aec8c1c2581ef7362b1e6e1d09dfb
COPY --from=ghcr.io/grpc-ecosystem/grpc-health-probe:v0.4.28 /ko-app/grpc-health-probe /usr/local/bin/grpc_health_probe
COPY --from=skeleton-builder /go/src/app/skeleton /usr/local/bin/skeleton
ENV PATH="$PATH:/usr/local/bin"
ENTRYPOINT ["skeleton"]
CMD ["serve"]
