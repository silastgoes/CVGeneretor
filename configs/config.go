package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	Env       *env
	TokenAuth *jwtauth.JWTAuth
}

type env struct {
	Environment   string `mapstructure:"Environment"`
	DBDriver      string `mapstructure:"DB_Driver"`
	DBHost        string `mapstructure:"DB_Host"`
	DBPort        string `mapstructure:"DB_Port"`
	DBUser        string `mapstructure:"DB_User"`
	DBPassword    string `mapstructure:"DB_Password"`
	DBNames       string `mapstructure:"DB_Names"`
	WebServerPort string `mapstructure:"Web_Server_Port"`
	JWTSecret     string `mapstructure:"JWT_Secret"`
	JWTExperesIn  int    `mapstructure:"JWT_ExperesIn"`
}

func Load(path string) *conf {
	cfg = &conf{}
	env := &env{}
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(env); err != nil {
		panic(err)
	}
	cfg.Env = env
	cfg.TokenAuth = jwtauth.New("HS256", []byte(env.JWTSecret), nil)
	return cfg
}
