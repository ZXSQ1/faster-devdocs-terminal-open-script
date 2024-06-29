#!/bin/bash

DEPS_ROOT=bin/deps

mkdir -p ./"$DEPS_ROOT"

function get-dedoc {
    VERSION=0.2.5
    URL="https://github.com/toiletbril/dedoc/releases/download/$VERSION"

    for bin in dedoc dedoc.exe; do
        wget -P "$DEPS_ROOT" "$URL"/$bin
    done
}

function get-fzf {
    VERSION=0.53.0
    URL="https://github.com/junegunn/fzf/releases/download/$VERSION/"

    for bin in fzf-$VERSION-linux_amd64.tar.gz fzf-$VERSION-windows_amd64.zip; do
        wget -P "$DEPS_ROOT" "$URL"/$bin

        cd "$DEPS_ROOT"

        tar -xf "$bin"
        unzip "$bin"

        cd -
    done
}

function get-ov {
    VERSION=0.35.0
    URL="https://github.com/noborus/ov/releases/download/$VERSION/"

    for bin in ov_"$VERSION"_linux_amd64.zip ov_"$VERSION"_windows_amd64.zip; do
        wget -P "$DEPS_ROOT" "$URL"/$bin

        cd "$DEPS_ROOT"

        unzip "$bin"

        cd -
    done
}