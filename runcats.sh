#!/bin/bash
set -xeu


CF_GOPATH=$HOME/go/src/github.com/cloudfoundry/

echo "Setting integration config..."
export CONFIG="/Users/amu36/Documents/git/cats-integration-config/m25-test-01.json"

echo "Moving cf-acceptance-tests onto the gopath..."
mkdir -p $CF_GOPATH
cp -R ../cf-acceptance-tests $CF_GOPATH
cd $HOME/go/src/github.com/cloudfoundry/cf-acceptance-tests

export CF_DIAL_TIMEOUT=11

./bin/test \
-keepGoing \
-randomizeAllSpecs \
-skipPackage=helpers \
-slowSpecThreshold=120 \
-nodes=3 \