#!/usr/bin/env bash

usage() {
	echo "Usage: ./$(basename "$0") <number>"
}

main() {
	output=''
	(( $1 % 3 == 0 )) && output="Pling"
	(( $1 % 5 == 0 )) && output+="Plang"
	(( $1 % 7 == 0 )) && output+="Plong"

	if [[ -z $output ]]; then
		echo "$1"
	else
		echo "$output"
	fi
}

main "$1"
