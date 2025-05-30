package settings

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		// fmt.Println("Failed to init config: ", err)
		return err
	}
	fmt.Println("Init config success: ", viper.ConfigFileUsed())
	return nil
}
