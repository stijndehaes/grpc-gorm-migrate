package client

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/client/connection"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/client/house"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/client/user"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "Example client for connection",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		connection.InitClients()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		connection.CloseConnection()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(house.Cmd)
	rootCmd.AddCommand(user.Cmd)
}
