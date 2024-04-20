package errors

import "fmt"

type WrongRawCredentialsFormatError struct {
}

func (err *WrongRawCredentialsFormatError) Error() string {
	return fmt.Sprintf("user:pass credentials format is required")
}
