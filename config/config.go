package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config 定义了应用程序的配置结构
type Config struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
	JWT      JWTConfig      `toml:"jwt"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	Port     string `toml:"port"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	URL  string `toml:"url"`
	Port string `toml:"port"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string `toml:"secret"`
}

// AppConfig 全局配置实例
var AppConfig *Config

// LoadConfig 加载配置文件
func LoadConfig(configPath string) error {
	if configPath == "" {
		configPath = "config.toml"
	}

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return err
	}

	// 解析TOML配置文件
	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return err
	}

	AppConfig = &config
	log.Println("配置文件加载成功:", configPath)
	return nil
}

// GetDatabaseDSN 获取数据库连接字符串
func GetDatabaseDSN() string {
	if AppConfig == nil {
		log.Fatal("配置未加载，请先调用 LoadConfig()")
	}
	return AppConfig.Database.Host + ":" + AppConfig.Database.Port
}

// GetJWTSecret 获取JWT密钥
func GetJWTSecret() string {
	if AppConfig == nil {
		log.Fatal("配置未加载，请先调用 LoadConfig()")
	}
	return AppConfig.JWT.Secret
}

// GetServerAddress 获取服务器地址
func GetServerAddress() string {
	if AppConfig == nil {
		log.Fatal("配置未加载，请先调用 LoadConfig()")
	}
	return AppConfig.Server.URL + ":" + AppConfig.Server.Port
}