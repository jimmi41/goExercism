package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return "This is the number " + fmt.Sprintf("%.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {

    return "This is a box containing the number " +
        fmt.Sprintf("%.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {

	fancy, ok := fnb.(FancyNumber)

	if !ok {
		return 0
	}

	num, _ := strconv.Atoi(fancy.Value())

	return num
}



func DescribeFancyNumberBox(fnb FancyNumberBox) string {

    num := ExtractFancyNumber(fnb)

    return "This is a fancy box containing the number " +
        fmt.Sprintf("%.1f", float64(num))
}

func DescribeAnything(i any) string {

    switch v := i.(type) {

    case int:
        return DescribeNumber(float64(v))

    case float64:
        return DescribeNumber(v)

    case NumberBox:
        return DescribeNumberBox(v)

    case FancyNumberBox:
        return DescribeFancyNumberBox(v)

    default:
        return "Return to sender"
    }
}