def verify(isbn):
	isbn = isbn.replace('-', '')
	
	if len(isbn) != 10:
		return False

	if isbn[9] == 'X':
		isbn = isbn[:8] + '10'

	try:
		isbn = [int(x) for x in isbn]
	except ValueError:
		return False

	total = 0
	return ((isbn[0] * 10) + (isbn[1] * 9) + (isbn[2] * 8) + (isbn[3] * 7) + (isbn[4] * 6) + (isbn[5] * 5) + (isbn[6] * 4) + (isbn[7] * 3) + (isbn[8] * 2) + (isbn[9] * 1)) % 11 == 0
