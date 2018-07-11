package device

import (
	"
)

func rootfsParentMountPrivate(rootfs string) error {
	}

	return nil
}

func prepareRoot(config *configs.Config) error {
	flag := unix.MS_SLAVE | unix.MS_REC
	if config.RootPropagation != 0 {
		flag = config.RootPropagation
	}
	if err := unix.Mount("", "/", "", uintptr(flag), ""); err != nil {
		return err
	}

	// Make parent mount private to make sure following bind mount does
	// not propagate in other namespaces. Also it will help with kernel
	// check pass in pivot_root. (IS_SHARED(new_mnt->mnt_parent))
	if err := rootfsParentMountPrivate(config.Rootfs); err != nil {
		return err
	}
	unix.Mount("", parentMount, "", unix.MS_PRIVATE, "")

	return unix.Mount(config.Rootfs, config.Rootfs, "bind", unix.MS_BIND|unix.MS_REC, "")
}


