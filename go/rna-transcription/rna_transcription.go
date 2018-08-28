package strand

import "strings"

func ToRNA(dna string) string {
  rna := ""
  rnaMap := map[string]string{
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U",
  }

  for _, nucleotide := range strings.Split(dna, "") {
    rna += rnaMap[nucleotide]
  }

  return rna
}
