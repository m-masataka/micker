package main

import (
	"fmt"

	"github.com/m-masataka/micker/cmd"
	"github.com/m-masataka/micker/cgroups"
	"github.com/m-masataka/micker/utils"
)

var pid = 3957

func main(){
	cmd.Execute()
	cid := utils.CreateContainerID()
	cgroups.AddCIDtoCgroups(cid)
	cgroups.AddProcstoCgroups(cid, pid)
	cgroups.RemoveCIDfromCgroups(cid)
	fmt.Println(cid)
}
