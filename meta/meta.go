package meta

type Meta struct {
	Project Project `toml:"project"`
	Source  Source  `toml:"source"`
	Env     Env     `toml:"env"`
	Run     Run     `toml:"run"`
}

type Project struct {
	Name     string `toml:"name"`
	Language string `toml:"language"`
	Enabled  bool   `toml:"enabled"`
}

type Source struct {
	Type  string `toml:"type"`
	Owner string `toml:"owner"`
	Repo  string `toml:"repo"`
}

type Env struct {
	Type string `toml:"type"`
	Path string `toml:"path"`
}

type Run struct {
	Cmd string `toml:"cmd"`
}

