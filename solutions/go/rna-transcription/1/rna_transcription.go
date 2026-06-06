package rnatranscription

func ToRNA(dna string) string {

    rna := ""

    for _, ch := range dna {

        if ch == 'G' {

            rna += "C"

        } else if ch == 'C' {

            rna += "G"

        } else if ch == 'T' {

            rna += "A"

        } else if ch == 'A' {

            rna += "U"
        }
    }

    return rna
}

