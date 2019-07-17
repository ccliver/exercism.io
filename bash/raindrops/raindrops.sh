#!/usr/bin/env bash

usage() {
	echo "Usage: ./$(basename $0) <number>"
}

main() {
	OUTPUT=''

	if (( $1 % 3 == 0 )); then
		OUTPUT="Pling"
	fi
	if (( $1 % 5 == 0 )); then
		OUTPUT="${OUTPUT}Plang"
	fi
	if (( $1 % 7 == 0 )); then
		OUTPUT="${OUTPUT}Plong"
	fi

	if [ -z $OUTPUT ]; then
		echo $1
	else
		echo $OUTPUT
	fi
}

main $1
