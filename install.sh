#!/bin/bash
#
# installation Script for the Open Source project fetch (http://github.com/KushagraIndurkhya/fetch)
#
#

GIT_USER=KushagraIndurkhya
GIT_PROJECT=fetch
BASE_URL=https://github.com/$GIT_USER/$GIT_PROJECT/releases/download
RELEASE=v1.0.0
BINARY=fetch

if [[ -e $BINARY ]]
then
  rm fetch
fi


URL="$BASE_URL/$RELEASE/$BINARY"

set -e
echo "Fetching from: $URL"
wget -q -O $BINARY "$URL"
file $BINARY
chmod a+x $BINARY
mkdir -p /usr/local/bin
mv $BINARY /usr/local/bin
echo "Done"