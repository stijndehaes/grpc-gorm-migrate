package user

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/stijndehaes/grpc-gorm-migrate/pb"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/client/connection"
)

var create = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := connection.UserClient.CreatUser(context.Background(), &pb.CreateUserRequest{
			Name: name,
		})
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			return
		}
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.SetHeader([]string{"Id", "Name"})
		table.Append([]string{response.Id, response.Name})
		table.Render()
	},
}

var name string

func init() {
	create.Flags().StringVarP(&name, "name", "n", "", "The name of the user to create")
	_ = create.MarkFlagRequired("name")
}
