package schedule_errors

type ForbiddenError struct{}

func (e *ForbiddenError) Error() string {
	return "Forbidden Error"
}
