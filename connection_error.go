package schedule_errors

import "fmt"

type ConnectionError struct {
	Address string
	Reason  string
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("failed to connect to server at %s: %s", e.Address, e.Reason)
}
