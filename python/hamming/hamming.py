def distance(strand_a, strand_b):
    hamming_distance = 0

    if len(strand_a) != len(strand_b):
        raise ValueError("Strands are different lengths")

    for idx, val in enumerate(strand_a):
        if strand_a[idx] != strand_b[idx]:
            hamming_distance += 1

    return hamming_distance
