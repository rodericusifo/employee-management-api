package merger

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/checker"
)

func MergeSlices[T comparable](withUnique bool, slices ...[]T) []T {
	resultMerged := make([]T, 0)
	resultMergedAndUnique := make([]T, 0)

	for _, s := range slices {
		resultMerged = append(resultMerged, s...)
	}

	if !withUnique {
		return resultMerged
	}

	for _, v := range resultMerged {
		if !checker.CheckSliceContain(resultMergedAndUnique, v) {
			resultMergedAndUnique = append(resultMergedAndUnique, v)
		}
	}

	return resultMergedAndUnique
}
