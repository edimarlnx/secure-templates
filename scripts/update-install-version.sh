#!/usr/bin/env sh

VERSION=${1:-"$(git tag --list --sort=-committerdate | head -n 1)"}
sed -i "s/INSTALL_VERSION=\${INSTALL_VERSION:-\".*\"}/INSTALL_VERSION=\${INSTALL_VERSION:-\"${VERSION}\"}/" scripts/install.sh
git add scripts
git commit -m "Update script version to latest: ${VERSION}"