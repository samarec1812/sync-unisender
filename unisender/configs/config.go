package configs

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath("app/unisender/configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	//path, err := os.Getwd()
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(path)

	//err := filepath.Walk(".",
	//	func(path string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			return err
	//		}
	//		fmt.Println(path, info.Size())
	//		return nil
	//	})
	//if err != nil {
	//	log.Println(err)
	//}
	return viper.ReadInConfig()
}
