package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	Charset    string

	EtcdHost string
	EtcdPort string

	UserServiceAddress string
	TaskServiceAddress string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("config.ini配置文件错误，请仔细检查文件", err)
	}
	LoadMysqlData(file)
	LoadEtcd(file)
	LoadRabbitMq(file)
	LoadServer(file)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadRabbitMq(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}

func LoadServer(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
}
