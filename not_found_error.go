package schedule_errors

import "fmt"

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	if e.Name != "" {
		return fmt.Sprintf("name '%s' not found", e.Name)
	}
	return "Not Found"
}
