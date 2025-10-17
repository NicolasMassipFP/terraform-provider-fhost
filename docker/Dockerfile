# Dockerfile for Go CI with golangci-lint
FROM golang:1.24

# Install golangci-lint v2.4.0 using official script
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b /usr/local/bin v2.4.0

# Install goreleaser latest version
RUN go install github.com/goreleaser/goreleaser@latest

RUN go install -v golang.org/x/tools/cmd/godoc@latest
COPY --from=ghcr.io/astral-sh/uv:latest /uv /uvx /bin/

# Create forcepoint user with home directory -- for signature purposes, it need a clean user setup with a used id allows to do the signature
RUN useradd -m -s /bin/bash forcepoint
RUN usermod -aG sudo forcepoint

# gpg for goreleaser need a valid home directory
RUN chown -R forcepoint:forcepoint /home/forcepoint
ENV HOME /home/forcepoint

# Switch to forcepoint user
USER forcepoint

WORKDIR /app

# Entrypoint for interactive use
CMD ["bash"]
