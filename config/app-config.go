package config

type AppConfig struct {
	Database struct {
		Host     string `json:"host" env:"DB_HOST" env-description:"Database host" env-default:"localhost"`
		Port     string `json:"port" env:"DB_PORT" env-description:"Database port" env-default:"5432"`
		Username string `json:"username" env:"DB_USER" env-description:"Database user name" env-default:"postgres"`
		Password string `json:"password" env:"DB_PASSWORD" env-description:"Database user password" env-default:"Postgres"`
		Name     string `json:"name" env:"DB_NAME" env-description:"Database name" env-default:"parking_pos"`
	} `json:"database"`
	Server struct {
		Host string `json:"host" env:"SRV_HOST,HOST" env-description:"Server host" env-default:"localhost"`
		Port string `json:"port" env:"SRV_PORT,PORT" env-description:"Server port" env-default:"3000"`
	} `json:"server"`
}
