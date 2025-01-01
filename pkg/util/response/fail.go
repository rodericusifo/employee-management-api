package response

import (
	pkg_types "github.com/rodericusifo/employee-management-api/pkg/types"
)

func ResponseFail(message string, err any) pkg_types.Response[any] {
	return pkg_types.Response[any]{
		Success: false,
		Message: message,
		Error:   err,
	}
}
