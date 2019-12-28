import re

def word_count(phrase):
    words = {}
    for word in re.split(r'[ \t_,\n]+', phrase.lower()):
        word = word.rstrip(".:!&@$%^'")
        word = word.lstrip("'")
        if len(word) == 0:
            continue

        if word in words:
            words[word] += 1
        else:
            words[word] = 1
    return words
