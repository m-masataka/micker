package command

import (
	"github.com/spf13/cobra"
	"github.com/m-masataka/micker/utils"
	"github.com/m-masataka/micker/container"
)

type runOption struct {
	rootpath string
	memorylimit string
}

var (
	rop = &runOption{}
)

var runCmd = &cobra.Command{
	Use: "run",
	Short: "Run micker container",
	Long: `Run micker container`,
	Run: func(cmd *cobra.Command, args []string){
		cid := utils.CreateUUID()
		container.ExecCommand(rop.rootpath, rop.memorylimit, cid)
	},
}
