package house

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
		response, err := connection.HouseClient.GetHouses(context.Background(), &pb.HousesRequest{})
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			return
		}
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.SetHeader([]string{"Id", "Location", "OwnerId"})
		for _, house := range response.Houses {
			table.Append([]string{house.Id, house.Location, house.OwnerId})
		}
		table.Render()
	},
}
