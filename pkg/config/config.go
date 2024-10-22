package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var config Config

type Config struct {
  DBHost     string `mapstructure:"DB_HOST"`
  DBName     string `mapstructure:"DB_NAME"`
  DBUser     string `mapstructure:"DB_USER"`
  DBPassword string `mapstructure:"DB_PASSWORD"`
  DBPort     string `mapstructure:"DB_PORT"`

  ADMINPASSWORD string `mapstructure:"ADMIN_PASSWORD"`
  ADMINEMAIL string `mapstructure:"ADMIN_EMAIL"`
  ADMINNAME string `mapstructure:"ADMIN_NAME"`
  ADMINSECRET   string `mapstructure:"ADMIN_SECRET"`
  USERSECRET    string `mapstructure:"USER_SECRET"`

  DBACCOUNTSID string `mapstructure:"DB_ACCOUNTSID"`
  DBAUTHTOKEN  string `mapstructure:"DB_AUTHTOKEN"`
  DBSERVICESID string `mapstructure:"DB_SERVICESID"`

  AWS_REGION            string `mapstructure:"AWS_REGION"`
  AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
  AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`


  REDIS_HOST string `mapstructure:"REDIS_HOST"`
  REDIS_PORT string `mapstructure:"REDIS_PORT"`

  SMTP_USERNAME string `mapstructure:"SMTP_USERNAME"`
  SMTP_PASSWORD string `mapstructure:"SMTP_PASSWORD"`
  SMTP_HOST     string `mapstructure:"SMTP_HOST"`
  SMTP_PORT     string `mapstructure:"SMTP_PORT"`

  RAZORPAY_KEY_ID     string `mapstructure:"RAZORPAY_KEY_ID"`
  RAZORPAY_KEY_SECRET string `mapstructure:"RAZORPAY_KEY_SECRET"`

  TEST_DBHost     string `mapstructure:"TEST_DB_HOST"`
  TEST_DBName     string `mapstructure:"TEST_DB_NAME"`
  TEST_DBUser     string `mapstructure:"TEST_DB_USER"`
  TEST_DBPassword string `mapstructure:"TEST_DB_PASSWORD"`
  TEST_DBPort     string `mapstructure:"TEST_DB_PORT"`
}

var envs = []string{
  "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "DB_AUTHTOKEN", "DB_ACCOUNTSID", "DB_SERVICESID", "DB_ACCOUNTSID", "DB_AUTHTOKEN","DB_SERVICESID", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SMTP_USERNAME", "SMTP_PASSWORD", "SMTP_HOST", "SMTP_PORT", "RAZORPAY_KEY_ID", "RAZORPAY_KEY_SECRET","ADMIN_SECRET", "USER_SECRET", "TEST_DB_HOST", "TEST_DB_NAME", "TEST_DB_USER", "TEST_DB_PASSWORD", "TEST_DB_PORT","REDIS_HOST","REDIS_PORT","ADMIN_PASSWORD","ADMIN_EMAIL","ADMIN_NAME"}

func LoadConfig() (Config, error) {

  viper.AddConfigPath(".")
  viper.SetConfigFile(".env")

  if err := viper.ReadInConfig(); err != nil{
    return config, fmt.Errorf("error reading config file: %v", err)
  }

  fmt.Println("env in LOad Config")

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
