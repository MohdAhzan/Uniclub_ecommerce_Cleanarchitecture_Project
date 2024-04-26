package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`

	DBACCOUNTSID string `mapstructure:"DB_ACCOUNTSID"`
	DBAUTHTOKEN  string `mapstructure:"DB_AUTHTOKEN"`
	DBSERVICESID string `mapstructure:"DB_SERVICESID"`

	AWS_REGION            string `mapstructure:"AWS_REGION"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`

	SMTP_USERNAME string `mapstructure:"SMTP_USERNAME"`
	SMTP_PASSWORD string `mapstructure:"SMTP_PASSWORD"`
	SMTP_HOST     string `mapstructure:"SMTP_HOST"`
	SMTP_PORT     string `mapstructure:"SMTP_PORT"`

	RAZORPAY_KEY_ID     string `mapstructure:"RAZORPAY_KEY_ID"`
	RAZORPAY_KEY_SECRET string `mapstructure:"RAZORPAY_KEY_SECRET"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "DB_AUTHTOKEN", "DB_ACCOUNTSID", "DB_SERVICESID", "DB_ACCOUNTSID", "DB_AUTHTOKEN",
	"DB_SERVICESID", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SMTP_USERNAME", "SMTP_PASSWORD", "SMTP_HOST", "SMTP_PORT", "RAZORPAY_KEY_ID", "RAZORPAY_KEY_SECRET",
}

func LoadConfig() (Config, error) {

	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {

		return config, err
	}
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil

}
