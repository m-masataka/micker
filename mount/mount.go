package mount

import (
	"os"
	"path/filepath"
	"syscall"
)

func PrepareRoot(rootfs string) error {
	oldmp := "/.oldroot"
	oldroot := filepath.Join(rootfs, oldmp)
	//flag := syscall.MS_SLAVE | syscall.MS_REC
	if err := syscall.Mount(rootfs, rootfs , "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}

	if err := os.MkdirAll(oldroot, 0700); err != nil {
		return err
	}

	if err := syscall.PivotRoot(rootfs, oldroot); err != nil {
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	if err := syscall.Unmount(oldmp, syscall.MNT_DETACH); err != nil{
		return err
	}

	if err := os.RemoveAll(oldmp); err != nil {
		return err
	}

	return nil
}
