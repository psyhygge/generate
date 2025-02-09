package ilangface

type ILanguageStrategyFactory interface {
	CreateStrategy(lang string) (ILanguageStrategy, error)
}
