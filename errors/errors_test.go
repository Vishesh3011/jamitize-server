package errors

import (
	"example/internal/types"
	"fmt"
	"testing"
)

func TestAppError(t *testing.T) {
	t.Run("TestAppError001", func(t *testing.T) {
		err := NewAppError(nil, types.InternalServerError, "Test message", types.Application)
		fmt.Println(err.Error())
		fmt.Println(err.ErrorStr())
		fmt.Println(err.Json())
		fmt.Println(err.Status())
		fmt.Println(err.Layer())
		fmt.Println(err.Message())
	})
}
