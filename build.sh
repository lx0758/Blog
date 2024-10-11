#!/usr/bin/env bash

pwd=$PWD

error() {
  echo "Build $1 failed!!!"
  exit 1
}

buildAdmin() {
  echo "Building Admin..."
  cd "$pwd"/admin
  npm install
  if [[ $? != 0 ]]; then error "Admin"; fi
  npm run build
  if [[ $? != 0 ]]; then error "Admin"; fi
  rm -rf "$pwd"/blog/res/static/admin/*
  mv "$pwd"/admin/dist/* "$pwd"/blog/res/static/admin/
  rm -rf "$pwd"/admin/dist/
  echo "Build Admin succeed!"
}

buildBlog() {
  echo "Building Blog..."
  cd "$pwd"/blog
  platforms=("linux/amd64" "windows/amd64")
  for platform in "${platforms[@]}"; do
    IFS='/' read -r os arch <<< "$platform"
    filename="blog-$os"
    if [[ "$os" == "windows" ]]; then
      filename=${filename}".exe"
    fi
    GOOS=$os GOARCH=$arch go build -o "${filename}" -v
    if [[ $? != 0 ]]; then error "Blog"; fi
  done
  rm -rf "$pwd"/blog/res/static/admin/*
  touch "$pwd"/blog/res/static/admin/.gitkeep
  echo "Build Blog succeed!"
}

build() {
  buildAdmin
  buildBlog
  echo "Build succeed!"
}

build