package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type config struct {
	DBDriver      string `mapstructure:"BD_DRIVER"`
	DBHost        string `mapstructure:"BD_HOST"`
	DBPort        string `mapstructure:"BD_PORT"`
	DBUser        string `mapstructure:"BD_USER"`
	DBPassword    string `mapstructure:"BD_PASSWORD"`
	DBName        string `mapstructure:"BD_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"MJWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

// executada antes do método main
// func init()

func LoadConfig(path string) (*config, error) {
	var cfg *config
	//..
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
