package user

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	pb "github.com/stijndehaes/grpc-gorm-migrate/pb"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/client/connection"
)

var list = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := connection.UserClient.GetUsers(context.Background(), &pb.UsersRequest{})
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			return
		}
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.SetHeader([]string{"Id", "Name"})
		for _, user := range response.Users {
			table.Append([]string{user.Id, user.Name})
		}
		table.Render()
	},
}
