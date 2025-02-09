package core

type ICodeGenerator interface {
	Generate(config *Config) error
}
