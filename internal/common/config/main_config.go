package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type MainConfig struct {
	// app
	Port      uint   `koanf:"port" validate:"gte=1"`
	SecretKey string `koanf:"secret_key" validate:"required"`

	// database
	DatabaseName     string `koanf:"db_name" validate:"required"`
	DatabaseHost     string `koanf:"db_host" validate:"required"`
	DatabasePort     uint   `koanf:"db_port" validate:"gte=1"`
	DatabasePassword string `koanf:"db_password" validate:"required"`
	DatabaseUsername string `koanf:"db_username" validate:"required"`

	// mail
	MailPort     uint   `koanf:"mail_port" validate:"gte=1"`
	MailHost     string `koanf:"mail_host" validate:"required"`
	MailUser     string `koanf:"mail_user" validate:"required"`
	MailPassword string `koanf:"mail_password" validate:"required"`

	// zarinpal
	ZarinpalRequestURL  string `koanf:"zarinpal_request_url" validate:"required"`
	ZarinpalMerchantID  string `koanf:"zarinpal_merchant_id" validate:"required"`
	ZarinpalVerifyURL   string `koanf:"zarinpal_verify_url" validate:"required"`
	ZarinpalPayURL      string `koanf:"zarinpal_pay_url" validate:"required"`
	ZarinpalCallbackURL string `koanf:"zarinpal_callback_url" validate:"required"`

	// upload directory
	UploadDirectory string `koanf:"upload_dir"`
}

func NewMainConfig() MainConfig {
	return MainConfig{}
}

func (mainConfig *MainConfig) LoadConfigs() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("unable to get environment variable")
	}

	k := koanf.New(".")

	if err := k.Load(env.Provider("ECOMMERCE_", "", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "ECOMMERCE_"))
	}), nil); err != nil {
		log.Fatalln("unable to load environment variable")
	}

	if err := k.UnmarshalWithConf("", &mainConfig, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true}); err != nil {
		log.Fatalln("unable to parse value of environment value")
	}

	if err := validator.New().Struct(mainConfig); err != nil {
		log.Fatalln("environment variable validation failed")
	}
}

func (mainConfig *MainConfig) GetAppPort() string {
	return fmt.Sprintf(":%d", mainConfig.Port)
}
