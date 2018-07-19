package errors

import "strings"

type PublicError interface {
	error
	Public() string
}

// model error implements the error interface type, because it has the Error()
// method built in
type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}

// privateError implements the error interface type, because it has the Error()
// method build in
type privateError string

func (e privateError) Error() string {
	return string(e)
}
