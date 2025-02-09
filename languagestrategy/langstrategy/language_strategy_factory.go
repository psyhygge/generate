package langstrategy

import (
	"fmt"
	"generate_dao/languagestrategy/ilangface"
)

type LanguageStrategyFactory struct {
}

func (lsf *LanguageStrategyFactory) CreateStrategy(lang string) (ilangface.ILanguageStrategy, error) {
	switch lang {
	case "go":
		return NewGoStrategy(), nil
	case "java":
		return NewJavaStrategy(), nil
	default:
		return nil, fmt.Errorf("unsupported languagestrategy type: %s", lang)
	}
}
