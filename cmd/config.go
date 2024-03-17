package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Origin string `mapstructure:"origin" yaml:"origin"`
	DB     struct {
		Host     string `mapstructure:"host" yaml:"host"`
		Port     int    `mapstructure:"port" yaml:"port"`
		Username string `mapstructure:"username" yaml:"username"`
		Password string `mapstructure:"password" yaml:"password"`
		Database string `mapstructure:"database" yaml:"database"`
	} `mapstructure:"db" yaml:"db"`
}

func setDefaults() {
	viper.SetDefault("db.host", "http://localhost:8000")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.username", "root")
	viper.SetDefault("db.password", "password")
	viper.SetDefault("db.database", "bookstack")
	viper.SetDefault("origin", "http://localhost:3000")
}

func (c Config) getDatabase() (*gorm.DB, error) {
	engine, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true",
			c.DB.Username,
			c.DB.Password,
			c.DB.Host,
			c.DB.Port,
			c.DB.Database,
		),
	}), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Truncate(time.Microsecond)
		},
	})
	if err != nil {
		return nil, err
	}

	return engine.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").Session(&gorm.Session{}), nil
}