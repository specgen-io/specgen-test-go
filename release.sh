#!/bin/bash
set -e

if [ -n "$1" ]; then
    VERSION=$1
else
    echo 'Version not set'
    exit 1
fi

echo "Releasing $VERSION"

if [ -n "$2" ]; then
    GITHUB_TOKEN=$2
else
    echo 'Github token not set'
    exit 1
fi

go get github.com/aktau/github-release

GH_ACCOUNT=specgen-io
GH_REPO=specgen-test-service-go
PROGRAM_NAME=service-go
RELEASE_NAME=$VERSION

echo "Creating release in Github: $RELEASE_NAME"
$GOPATH/bin/github-release release --security-token $GITHUB_TOKEN --user $GH_ACCOUNT --repo $GH_REPO --tag $RELEASE_NAME

echo "Releasing zips/${PROGRAM_NAME}_darwin_amd64.zip"
$GOPATH/bin/github-release upload  --security-token $GITHUB_TOKEN --user $GH_ACCOUNT --repo $GH_REPO --tag $RELEASE_NAME --name ${PROGRAM_NAME}_darwin_amd64.zip  --file zips/${PROGRAM_NAME}_darwin_amd64.zip
echo "Releasing zips/${PROGRAM_NAME}_linux_amd64.zip"
$GOPATH/bin/github-release upload  --security-token $GITHUB_TOKEN --user $GH_ACCOUNT --repo $GH_REPO --tag $RELEASE_NAME --name ${PROGRAM_NAME}_linux_amd64.zip   --file zips/${PROGRAM_NAME}_linux_amd64.zip
echo "Releasing zips/${PROGRAM_NAME}_windows_amd64.zip"
$GOPATH/bin/github-release upload  --security-token $GITHUB_TOKEN --user $GH_ACCOUNT --repo $GH_REPO --tag $RELEASE_NAME --name ${PROGRAM_NAME}_windows_amd64.zip --file zips/${PROGRAM_NAME}_windows_amd64.zip

echo "Done releasing $VERSION"
