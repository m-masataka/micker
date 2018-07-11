package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	//"github.com/opencontainers/runc/libcontainer/mount"
)

func main() {
	rootfs := "/root/test/device"
	flag := syscall.MS_SLAVE | syscall.MS_REC
	if err := syscall.Mount("", "/", "", uintptr(flag), ""); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	syscall.Mount("", rootfs, "", syscall.MS_PRIVATE, "")
	fmt.Println("OK4")
	syscall.Mount( rootfs, rootfs, "bind", syscall.MS_BIND|syscall.MS_REC, "")

	newroot, err := syscall.Open(rootfs, syscall.O_DIRECTORY|syscall.O_RDONLY, 0)
        if err != nil {
		fmt.Println(err)
		os.Exit(-1)
        }

	if err := syscall.Fchdir(newroot); err != nil {
		fmt.Println(err)
		os.Exit(-1)
        }

	if err := syscall.PivotRoot(".", "."); err != nil {
                fmt.Errorf("pivot_root %s", err)
		os.Exit(-1)
        }

	out, err := exec.Command("/bin/ls", "-la", "../").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(out))
}
