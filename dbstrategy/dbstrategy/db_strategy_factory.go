package dbstrategy

import (
	"fmt"
	"generate_dao/dbstrategy/idbface"
)

type DatabaseStrategyFactory struct{}

func (dsf *DatabaseStrategyFactory) CreateStrategy(dbType, dsn string) (idbface.IDatabaseStrategy, error) {
	switch dbType {
	case "mysql":
		return NewMySQLStrategy(dsn), nil
	case "postgres":
		return NewPostgresStrategy(), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
