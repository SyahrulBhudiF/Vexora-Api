package core

import (
	"fmt"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Vexora struct {
	Config *viper.Viper
	DB     *gorm.DB
	App    *fiber.App
	Redis  *redis.Client
}

func Init(vexora *Vexora) {
	vexora.Config = config.NewConfig()
}

func (a *Vexora) Start() {
	err := a.App.Listen(fmt.Sprintf("%s:%s", a.Config.GetString("app.host"), a.Config.GetString("app.port")))
	if err != nil {
		logrus.Fatal(err)
	}
}
