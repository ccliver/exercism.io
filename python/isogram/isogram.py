import re

def is_isogram(string):
    characters_seen = {}

    if string == '':
        return True

    for character in string.lower():
        if re.match(r'[a-zA-z]+', character):
            if character in characters_seen:
                return False
            else:
                characters_seen[character] = True

    return True
