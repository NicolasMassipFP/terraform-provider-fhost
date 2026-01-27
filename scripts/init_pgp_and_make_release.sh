#!/bin/bash

echo "In current version, private key is expected in variable GPG_PRIVATE_KEY_B64"
gpg --list-secret-keys
echo "${GPG_PRIVATE_KEY_B64}" | base64 -d | gpg --batch --import
gpg --list-secret-keys
echo "GPG key registered for release preparation"

echo "generating documentation"
make docs

echo "Starting release"
goreleaser release --clean
