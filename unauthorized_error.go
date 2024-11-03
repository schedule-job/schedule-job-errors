package schedule_errors

import "fmt"

type UnauthorizedError struct {
	Reason string // 실패 이유
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized access: %s", e.Reason)
}
