package user

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use: "user",
}

func init() {
	Cmd.AddCommand(list)
	Cmd.AddCommand(create)
	Cmd.AddCommand(house)
}
