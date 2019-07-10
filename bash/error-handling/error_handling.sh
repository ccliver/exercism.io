#!/usr/bin/env bash

usage() {
	echo "Usage: ./error_handling <greetee>"
}

main() {
	echo "Hello, $1"
}

if [ $# -ne 1 ]; then
	usage
	exit 1
fi

main "$1"
