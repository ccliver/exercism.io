#!/usr/bin/env bash

main() {
	strand1_chars=( $(echo -n "$1" | grep -o .) )
	strand2_chars=( $(echo -n "$2" | grep -o .) )

	index=0
	hammingDistance=0
	for char in "${strand1_chars[@]}"
	do
		if [[ "$char" != "${strand2_chars[$index]}" ]]; then
			hammingDistance=$(( hammingDistance + 1 ))
		fi

		index=$(( index + 1 ))
	done

	echo $hammingDistance
}

if [[ $# -ne 2 ]]; then
	echo "Usage: $(basename "$0") <string1> <string2>"
	exit 1
fi

if [[ -z "$1" ]] && [[ -z "$2" ]]; then
	echo 0
	exit 0
fi

if [[ -z "$1" ]]; then
	echo "left strand must not be empty"
	exit 1
fi

if [[ -z "$2" ]]; then
	echo "right strand must not be empty"
	exit 1
fi

if [[ $(echo -n "$1" | wc -c ) -ne $(echo -n "$2" | wc -c) ]]; then
	echo "left and right strands must be of equal length"
	exit 1
fi

main "$1" "$2"
