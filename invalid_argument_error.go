package schedule_errors

import "fmt"

type InvalidArgumentError struct {
	Param   string // 문제가 되는 인자 이름
	Message string // 추가 설명 메시지
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("invalid argument: %s - %s", e.Param, e.Message)
}
