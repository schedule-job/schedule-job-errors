package schedule_errors

import (
	"fmt"
	"log"
	"time"
)

type QueryError struct {
	Err error
}

func (e *QueryError) Error() string {
	var now = time.Now().Local().String()
	log.Fatalln(e.Err.Error() + " " + now)
	return fmt.Sprintf("Internal Server Error : Query Error (%s)", now)
}
