package config

var (
	Version = ""
)

type Cfg struct {
	TplPath    string
	CreateType string
	Args       []string
}
