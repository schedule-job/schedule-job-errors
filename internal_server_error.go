package schedule_errors

import "fmt"

type InternalServerError struct {
	Err error
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Internal Server Error : %s", e.Err.Error())
}
