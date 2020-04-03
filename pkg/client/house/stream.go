package house

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	pb "github.com/stijndehaes/grpc-gorm-migrate/pb"
	"github.com/stijndehaes/grpc-gorm-migrate/pkg/client/connection"
	"io"
)

var stream = &cobra.Command{
	Use: "stream",
	Run: func(cmd *cobra.Command, args []string) {
		stream, err := connection.HouseClient.StreamHouses(context.Background(), &pb.HousesRequest{})
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
			return
		}
		table := tablewriter.NewWriter(cmd.OutOrStdout())
		table.SetHeader([]string{"Id", "Location", "OwnerId"})

		for {
			house, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Something wrong")
				panic(err)
			}
			table.Append([]string{house.Id, house.Location, house.OwnerId})
		}
		table.Render()
	},
}
