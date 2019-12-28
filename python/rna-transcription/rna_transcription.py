DNA_to_RNA = {
    'G': 'C',
    'C': 'G',
    'T': 'A',
    'A': 'U'
}

def to_rna(dna_strand):
    return "".join([DNA_to_RNA[nucleotide] for nucleotide in dna_strand])
