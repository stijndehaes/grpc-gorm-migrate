package house

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use: "house",
}

func init() {
	Cmd.AddCommand(list)
	Cmd.AddCommand(create)
	Cmd.AddCommand(stream)
}
