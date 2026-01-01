package configs

type DatatabaseType string

const DatabaseTypeMySql DatatabaseType = "mysql"

type Database struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
