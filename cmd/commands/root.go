package commands

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/config"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/core"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configPath string
	VexoraApp  *core.Vexora

	rootCmd = &cobra.Command{
		Use:   "vexora_api",
		Short: "Backend application for vexora mobile app",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	cobra.OnInitialize(initApp)

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "../config.yaml", "config file (default is $HOME/.cobra.yaml)")

	serveCmd := newServeCmd()

	rootCmd.AddCommand(serveCmd)
}

func initApp() {
	viper := config.NewConfig()

	app := core.NewFiber(viper)

	db, err := core.NewDB(viper)
	if err != nil {
		logrus.Fatal("unable to initialize db: %s", err.Error())
	}

	rds, err := core.NewRedis(viper)
	if err != nil {
		logrus.Fatal("unable to initialize redis: %s", err.Error())
	}

	jwt := services.NewJWTService(viper.GetString("app.secret"))

	imageKit := services.NewImageKitService(viper.GetString("imagekit.private_key"), viper.GetString("imagekit.public_key"), viper.GetString("imagekit.url_endpoint"))

	spotify := services.NewSpotifyService(viper.GetString("spotify.client_id"), viper.GetString("spotify.client_secret"))
	if err != nil {
		logrus.Fatal("unable to initialize spotify service: %s", err.Error())
	}

	mail := services.NewMailService(viper.GetString("mail.host"), viper.GetInt("mail.port"), viper.GetString("mail.email"), viper.GetString("mail.password"))

	VexoraApp = &core.Vexora{
		Viper:    viper,
		App:      app,
		DB:       db,
		Redis:    rds,
		JWT:      jwt,
		ImageKit: imageKit,
		Spotify:  spotify,
		Mail:     mail,
	}
	core.Init(VexoraApp)
}
