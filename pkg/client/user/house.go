package user

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/stijndehaes/grpc-gorm-migrate/pb"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/client/connection"
)

var house = &cobra.Command{
	Use: "houses",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := connection.UserClient.UserWithHouses(context.Background(), &pb.UserHousesRequest{
			Id: id,
		})
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			return
		}
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.SetHeader([]string{"Id", "Name", "House Id", "Location"})
		for _, uh := range response.UserHouses {
			table.Append([]string{uh.UserId, uh.UserName, uh.HouseId, uh.HouseLocation})
		}
		table.Render()
	},
}

var id string

func init() {
	house.Flags().StringVarP(&id, "id", "", "", "The id of the user to list houses for")
	_ = house.MarkFlagRequired("id")
}
