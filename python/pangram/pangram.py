import re

def is_pangram(sentence):
    letters_seen = ''
    alphabet_re = re.compile('[a-z]')

    for letter in sentence.lower():
        if alphabet_re.search(letter):
            if letter in letters_seen:
                next
            else:
                letters_seen += letter
    if len(letters_seen) == 26:
        return True
    else:
        return False
