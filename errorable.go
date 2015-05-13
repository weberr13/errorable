package errorable

type IntErr struct {
	I int
	E error
}
type IntErrs []IntErr

// Create a errable int
func NewIntErr(f func(string) (int, error), s string) *IntErr {
	this := &IntErr{}
	this.I, this.E = f(s)
	return this
}

// create an errable slice of ints by applying function f on a slice of strings
func NewIntErrs(f func(string) (int, error), s ...string) *IntErrs {
	interr := make(IntErrs, len(s), len(s))
	for i, v := range s {
		interr[i] = *NewIntErr(f, v)
	}
	return (&interr)
}

// get the value at i, ignore errors
func (this *IntErrs) Get(i int) int {
	return (*this)[i].I
}

// length of the slice of errable ints
func (this *IntErrs) Len() int {
	return (len(*this))
}

// get the first error in the slice of errable ints
func (this *IntErrs) GetFirstErr() (err error) {
	for _, v := range *this {
		if nil != v.E {
			return v.E
		}
	}
	return nil
}
