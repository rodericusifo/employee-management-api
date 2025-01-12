package constant

import (
	"time"

	"github.com/rodericusifo/employee-management-api/internal/pkg/types"
)

type Default any
type DefaultEnv any
type DefaultSelects []types.SelectQuerySQLOperation
type DefaultSelectsJoin []types.SelectJoinQuerySQLOperation

var (
	DEFAULT_TIME_LAYOUT = Default(time.DateTime)
)
var (
	DEFAULT_ENV_SERVER_PORT = DefaultEnv(8080)
)
var (
	DEFAULT_SELECTS_COLUMNS      = DefaultSelects([]types.SelectQuerySQLOperation{})
	DEFAULT_SELECTS_JOIN_COLUMNS = DefaultSelectsJoin([]types.SelectJoinQuerySQLOperation{})
)
