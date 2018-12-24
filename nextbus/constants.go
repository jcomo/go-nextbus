package nextbus

import (
	"errors"
)

var errNotImplemented error

func init() {
	errNotImplemented = errors.New("not implemented")
}
