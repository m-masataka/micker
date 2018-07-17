package cgroups

import (
	"os"
	"fmt"
	"strconv"
	"io/ioutil"
)

const (
	cPath    = "/sys/fs/cgroup"
	cDevice  = cPath + "/devices/micker"
	cPids    = cPath + "/pids/micker"
	cBlkio   = cPath + "/blkio/micker"
	cFreezer = cPath + "/freezer/micker"
	cNet     = cPath + "/net_cls,net_prio/micker"
	cHuget   = cPath + "/hugetlb/micker"
	cPe      = cPath + "/perf_event/micker"
	cCpu     = cPath + "/cpu,cpuacct/micker"
	cMemory  = cPath + "/memory/micker"
	cCpuset  = cPath + "/cpuset/micker"
	cSystemd = cPath + "/systemd/micker"
)
var pathArray = [...] string{
		cDevice,
		cPids,
		cBlkio,
		cFreezer,
		cNet,
		cHuget, 
		cPe,
		cCpu,
		cMemory,
		cCpuset,
		cSystemd,
	}

func init() {
	for _, p := range pathArray { 
		if _, err := os.Stat(p); err == nil {
			continue
		}
		if err := os.Mkdir(p, 0755); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}

func AddCIDtoCgroups(cid string) error {
	for _, p := range pathArray { 
		if _, err := os.Stat(p + "/" + cid); err == nil {
			continue
		}
		if err := os.Mkdir(p + "/" + cid, 0755); err != nil {
			return err
		}
	}
	return nil
}

func RemoveCIDfromCgroups(cid string) error {
	for _, p := range pathArray {
		if _, err := os.Stat(p + "/" + cid); err != nil {
			fmt.Println(p + "/" +cid +" directory doesn't exist")
			continue
		}
		if err := os.Remove(p + "/" + cid); err != nil {
			return err
		}
	}
	return nil
}

func AddProcstoCgroups(cid string, pid int) error {
	procsPath := "cgroup.procs"
	for _, p := range pathArray { 
		if _, err := os.Stat(p + "/" + cid + "/" +procsPath); err != nil {
			return err
		}
		content := []byte(strconv.Itoa(pid) + "\n")
		ioutil.WriteFile( p + "/" + cid + "/" + procsPath, content, os.ModePerm)
	}
	return nil
}

func MemoryLimit(cid string, ml string) error {
	if _, err := os.Stat(cMemory +"/"+ cid +"/memory.limit_in_bytes"); err != nil {
		return err
	}
	content := []byte(ml + "\n")
	ioutil.WriteFile(cMemory +"/"+ cid + "/memory.limit_in_bytes", content, os.ModePerm)
	return nil
}
