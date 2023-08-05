package gopher

type (
	Entity interface {
		Declare() *Schema
	}

	Config struct {
		TableName string
		Comment   string
	}

	Schema struct {
		Config *Config
		Fields []*Field
	}
)
