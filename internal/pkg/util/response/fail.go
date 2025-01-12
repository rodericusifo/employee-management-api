package response

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

func ResponseFail(message string, err any) types.Response[any] {
	return types.Response[any]{
		Success: false,
		Message: message,
		Error:   err,
	}
}
