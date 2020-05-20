package exceptions

import "errors"

var (
	ErrTaskAbort = errors.New("loading task abort")
)
