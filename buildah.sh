#!/usr/bin/env bash

ctr1=$(buildah from "${1:-golang:1.15-alpine}")

buildah run "$ctr1" -- apk update
buildah run "$ctr1" -- apk add make

buildah config --workingdir '/build' "$ctr1"

buildah add "$ctr1" "$PWD"
buildah run "$ctr1" -- make build
buildah run "$ctr1" -- mv /build/cloudbot /usr/bin/cloudbot

buildah config --cmd "/usr/bin/cloudbot" "$ctr1"
buildah config --port 3000 "$ctr1"

buildah commit "$ctr1" "${2:-$USER/cloudbot}"

