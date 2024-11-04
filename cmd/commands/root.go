package commands

import (
	"github.com/SyahrulBhudiF/Vexora-Api/internal/config"
	"github.com/SyahrulBhudiF/Vexora-Api/internal/core"
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

	VexoraApp = &core.Vexora{
		Config: viper,
		App:    app,
		DB:     db,
		Redis:  rds,
	}
	core.Init(VexoraApp)
}
