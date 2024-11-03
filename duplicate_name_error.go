package schedule_errors

import "fmt"

type DuplicateNameError struct {
	Name string
}

func (e *DuplicateNameError) Error() string {
	return fmt.Sprintf("the name '%s' is already in use", e.Name)
}
