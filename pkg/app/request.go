package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/kim118000/game2023/pkg/logging"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
