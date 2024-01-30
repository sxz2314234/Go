/*
// incline error interface

	type error interface{
		Error() string
	}
*/
package errormechanism

import "fmt"

type DivideError struct {
	Divider int
	Dividee int
}

func (d *DivideError) Error() string {
	strFormat := `
	Can not proceed, the divider is zero.
	dividee:%d
	divider:0`

	return fmt.Sprintf(strFormat, d.Dividee)
}

func Divide(vardividee int, vardivider int) (result int, errormessage string) {
	if vardivider == 0 {
		dData := DivideError{
			Dividee: vardividee,
			Divider: vardivider,
		}
		errormessage := dData.Error()

		return -1, errormessage
	} else {
		return vardividee / vardivider, ""
	}
}
