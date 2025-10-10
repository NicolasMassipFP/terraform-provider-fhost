# Dockerfile for Go CI with golangci-lint
FROM golang:1.24

# Install golangci-lint v2.4.0 using official script
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /usr/local/bin v2.4.0

RUN go install -v golang.org/x/tools/cmd/godoc@latest
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

WORKDIR /app

# Entrypoint for interactive use
CMD ["bash"]
