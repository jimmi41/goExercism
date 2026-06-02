package nucleotidecount

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
// Start by uncommenting the following line:
// type Histogram ...

// DNA is a list of nucleotides. Choose a suitable data type.
// Start by uncommenting the following line:
// type DNA ...

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
import "errors"

// Histogram is a mapping from nucleotide to count.
type Histogram map[rune]int

// DNA represents a DNA strand.
type DNA string
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, nucleotide := range d {

		switch nucleotide {

		case 'A', 'C', 'G', 'T':
			h[nucleotide]++

		default:
			return nil,
				errors.New(
					"invalid nucleotide",
				)
		}
	}

	return h, nil
}
