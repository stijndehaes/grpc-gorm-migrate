package house

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
		house, err := connection.HouseClient.CreateHouse(context.Background(), &pb.CreateHouseRequest{
			Location: location,
			OwnerId:  owner,
		})
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			return
		}
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.SetHeader([]string{"Id", "Location", "OwnerId"})
		table.Append([]string{house.Id, house.Location, house.OwnerId})
		table.Render()
	},
}

var location string
var owner string

func init() {
	create.Flags().StringVarP(&location, "location", "l", "", "The location of the house")
	_ = create.MarkFlagRequired("location")
	create.Flags().StringVarP(&owner, "owner", "o", "", "The owner of the house")
	_ = create.MarkFlagRequired("owner")
}
