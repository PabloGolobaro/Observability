package mongo

import (
	"fmt"
)

//Error mongo error
type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	DevMessage string `json:"dev_message"`
}

//New return ner err with devmsg as err
func (e Error) New(err error) error {
	e.DevMessage = err.Error()
	return &e
}

//Error implement error
func (e Error) Error() string {
	return fmt.Sprintf("Msg: %s   DevMsg: %s\n", e.Message, e.DevMessage)
}
