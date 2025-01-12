package getter

import (
	"github.com/rodericusifo/employee-management-api/internal/pkg/config"
)

func GetEnvConfig() config.EnvConfig {
	return config.Env
}

func GetMysqlDatabaseSQLConnection() config.MysqlDatabaseSQLConnection {
	return config.MysqlDBSQL
}
func GetPostgresDatabaseSQLConnection() config.PostgresDatabaseSQLConnection {
	return config.PostgresDBSQL
}

func GetRedisDatabaseCacheConnection() config.RedisDatabaseCacheConnection {
	return config.RedisDBCache
}

func GetJWTAuthConfig() config.JWTAuthConfig {
	return config.JWTAuth
}
