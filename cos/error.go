package cos

import "fmt"

type Error struct {
	Code      string
	Message   string
	RequestID string
	Internal  error
}

func (err Error) Error() string {
	return fmt.Sprintf("[%s]%s <- %s", err.Code, err.Message, err.RequestID)
}
