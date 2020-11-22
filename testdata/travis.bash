#!/bin/bash
set -ex

# set GOPATH if empty (travis sets it, but useful for humans)
if [ -z "$GOPATH" ]; then
GOPATH=$(go env GOPATH)
export GOPATH
fi

# workaround for go1.10's no module support, we copy this run's source code
# into $GOPATH/src/github.com/namecoin/splicesign to avoid downloading our master branch.
#
# this only affects forks running travis runs. ( Pull requests and autobuilds
# will clone into: PWD=/c/Users/travis/gopath/src/github.com/namecoin/splicesign )
if [ "$TRAVIS_GO_VERSION" = "1.10.x" ] && [ ! -d "$GOPATH"/src/github.com/namecoin/splicesign ]; then
  mkdir "$GOPATH"/src/github.com/namecoin && \
  cp -av "$PWD" "$GOPATH"/src/github.com/namecoin/splicesign && \
  cd "$GOPATH"/src/github.com/namecoin/splicesign
fi

echo Fetching dependencies
go get -v -t ./...

echo Building splicesign
go get ./...

echo Running go test suite
go test -v ./...
