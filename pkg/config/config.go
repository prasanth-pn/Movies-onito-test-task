package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct{
	Port string `mapstructure:"PORT"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

var envs=[]string{"PORT","DB_SOURCE"}


func LoadConfig()(Config,error){

	var config Config

	viper.AddConfigPath("./pkg/config/env")

	viper.SetConfigFile("./pkg/config/env/dev.env")
	viper.ReadInConfig()
	for _,env:=range envs{
		if err:=viper.BindEnv(env);err!=nil{
			return config,err
		}
	}
	if err:=viper.Unmarshal(&config);err!=nil{
		return config,err
	}
if err:=validator.New().Struct(&config);err!=nil{
	return config,err
}
return config,nil


}