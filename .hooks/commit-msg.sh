#!/bin/sh

PREFIX="(build|chore|ci|feat|fix|refactor|revert|test)"
CONVENTION="^${PREFIX}:\s{1}\.*$"
ARG=$1

if ! [[ $ARG =~ $CONVENTION ]]
then
  echo "Commit message is not matching the convention"
  echo "See example: build|chore|ci|feat|fix|refactor|revert|test: message"
  return 0
fi
