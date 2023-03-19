package processes

import "os"

func Util() error {
	err := os.Mkdir("util", 0755)
	if err != nil {
		return err
	}

	err = os.Chdir("util")
	if err != nil {
		return err
	}

	file, err := os.Create("config.go")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`package util 

import "github.com/spf13/viper"

` + "type Config struct {" + "\n" + "\tDBDriver string" + "`mapstructure:" + `"DB_DRIVER"` + "`" + "\n" +
		"\tDBSource string" + "`mapstructure:" + `"DB_SOURCE"` + "`" + "\n" +
		"\tHTTPServerAddress string" + "`mapstructure:" + `"HTTP_SERVER_ADDRESS"` + "`" + "\n" +
		"}" + "\n" +
		`
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return config, err
}
`)
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}
