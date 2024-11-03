package schedule_errors

import "fmt"

type LogWriteError struct {
	Reason string // 실패한 이유
}

func (e *LogWriteError) Error() string {
	return fmt.Sprintf("failed to write log: %s", e.Reason)
}
