package configs

type Auth struct {
	Hash Hash `yaml:"hash"`
}

type Hash struct {
	Cost int `yaml:"cost"`
}
