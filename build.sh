#!/usr/bin/env bash

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"

error() {
  echo "Build $1 failed!!!"
  exit 1
}

buildAdmin() {
  echo "Building Admin..."
  cd "$SCRIPT_DIR"/admin
  npm install
  if [[ $? != 0 ]]; then error "Admin"; fi
  npm run build
  if [[ $? != 0 ]]; then error "Admin"; fi
  rm -rf "$SCRIPT_DIR"/blog/res/static/admin/*
  mv "$SCRIPT_DIR"/admin/dist/* "$SCRIPT_DIR"/blog/res/static/admin/
  rm -rf "$SCRIPT_DIR"/admin/dist/
  echo "Build Admin succeed!"
}

buildBlog() {
  echo "Building Blog..."
  cd "$SCRIPT_DIR"/blog
  platforms=("linux/amd64" "windows/amd64")
  for platform in "${platforms[@]}"; do
    IFS='/' read -r os arch <<< "$platform"
    filename="blog-$os"
    if [[ "$os" == "windows" ]]; then
      filename=${filename}".exe"
    fi
    GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o "${filename}" -v
    if [[ $? != 0 ]]; then error "Blog"; fi
  done
  rm -rf "$SCRIPT_DIR"/blog/res/static/admin/*
  touch "$SCRIPT_DIR"/blog/res/static/admin/.gitkeep
  echo "Build Blog succeed!"
}

build() {
  buildAdmin
  buildBlog
  echo "Build succeed!"
}

build