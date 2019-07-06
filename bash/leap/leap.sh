#!/usr/bin/env bash
set -o errexit
set -o nounset

usage() {
  echo "Usage: $(basename "$0") <year>"
}

re='^[0-9]+$'
if [ $# -ne 1 ] || ! [[ "$1" =~ $re ]]; then
  usage
  exit 1
fi


if [ $(( $1 % 4 )) -eq 0 ] && [ $(( $1 % 100 )) -ne 0 ] || [ $(( $1 % 400 )) -eq 0 ]; then
  echo "true"
else
  echo "false"
fi
