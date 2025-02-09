package idbface

type IDatabaseStrategyFactory interface {
	CreateStrategy(dbType, dsn string) (IDatabaseStrategy, error)
}
