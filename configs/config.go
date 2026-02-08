package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPath        string `mapstructure:"DB_PATH"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExperesIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig() (*conf, error) {
	var cfg conf

	viper.SetConfigFile(".env") // arquivo exato
	viper.SetConfigType("env")  // tipo correto
	viper.AddConfigPath(".")    // raiz do projeto
	viper.AutomaticEnv()        // permite sobrescrever por env var

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return &cfg, nil
}
