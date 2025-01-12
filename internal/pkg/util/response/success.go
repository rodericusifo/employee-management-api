package response

import "github.com/rodericusifo/employee-management-api/internal/pkg/types"

func ResponseSuccess[T any](message string, data T, meta *types.Meta) types.Response[T] {
	return types.Response[T]{
		Success: true,
		Message: message,
		Meta:    meta,
		Data:    data,
	}
}
