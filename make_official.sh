#!/usr/bin/env bash

set -e

# first build the version string
VERSION_NUM=1.0

# add the git commit id and date
VERSION="$VERSION_NUM (commit $(git rev-parse --short HEAD) @ $(git log -1 --date=short --pretty=format:%cd))"

PROJECT="timestats"

function buildbinary {
    goos=$1
    goarch=$2

    echo "Building official $goos $goarch binary for version '$VERSION'"

    outputfolder="build/$PROJECT-$VERSION_NUM/${goos}_${goarch}"
    echo "Output Folder $outputfolder"
    mkdir -pv $outputfolder

    export GOOS=$goos
    export GOARCH=$goarch

    go build -i -v -o "$outputfolder/$PROJECT" -ldflags "-X \"main.Version=$VERSION\"" "github.com/AstromechZA/$PROJECT"

    echo "Done"
    ls -l "$outputfolder/$PROJECT"
    file "$outputfolder/$PROJECT"
    echo
}

# build for mac
buildbinary darwin amd64

# build for linux
buildbinary linux amd64

# zip up
tar -czvf "$PROJECT-$VERSION_NUM.tar.gz" -C "build" "$PROJECT-$VERSION_NUM"

# output text
echo ""
echo "How to release:"
echo ""
echo "1. Either tag and push the current commit with version $VERSION_NUM or create a new release on github"
echo "   git tag 'v$VERSION_NUM' && git push && git push --tags"
echo "2. Upload the $PROJECT-$VERSION_NUM.tar.gz file as the attached file"
echo "3. Set the release title to '$PROJECT $VERSION_NUM release'"
echo "3. Write a brief description if required"
