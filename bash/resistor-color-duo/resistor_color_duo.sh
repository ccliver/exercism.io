#!/usr/bin/env bash

RESISTORS=(
	black
	brown
	red
	orange
	yellow
	green
	blue
	violet
	grey
	white
)

main() {
	RESISTOR_CODE=
	INDEX=0
	for resistor in "${!RESISTORS[@]}"
	do
		if [ "${RESISTORS[$resistor]}" == "$1" ]; then
			RESISTOR_CODE="$INDEX"
			break
		fi
		INDEX=$((INDEX + 1))
	done
	
	if [ -z "$RESISTOR_CODE" ]; then
		echo "invalid color"
		exit 1
	fi

	INDEX=0
	for resistor in "${!RESISTORS[@]}"
	do
		if [ "${RESISTORS[$resistor]}" == "$2" ]; then
			RESISTOR_CODE="$RESISTOR_CODE${INDEX}"
			break
		fi
		INDEX=$((INDEX + 1))
	done

	echo "$RESISTOR_CODE"
}

main "$1" "$2"
