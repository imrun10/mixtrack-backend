package git

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Initialise(cmd *cobra.Command, args []string) error {
	fmt.Println("Initialising the project...")
	if len(args) < 1 {
		return fmt.Errorf("main project is required")
	}

	return nil
}
