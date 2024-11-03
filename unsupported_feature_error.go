package schedule_errors

import "fmt"

type UnsupportedFeatureError struct {
	Feature string
}

func (e *UnsupportedFeatureError) Error() string {
	return fmt.Sprintf("'%s' is not supported", e.Feature)
}
