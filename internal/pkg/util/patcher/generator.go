package patcher

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/util/generator"
)

var (
	GenerateHashFromPassword   = generator.GenerateHashFromPassword
	GenerateJWTTokenFromClaims = generator.GenerateJWTTokenFromClaims
)
