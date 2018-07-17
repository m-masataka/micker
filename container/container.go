package container

import (
        "fmt"
        "os"
        "syscall"
        "os/exec"

        "github.com/m-masataka/micker/cgroups"
        "github.com/m-masataka/micker/mount"
        "github.com/m-masataka/micker/network"
        "github.com/docker/docker/pkg/reexec"
)

func ExecCommand(rootfs string, ml string, cid string) error {
        cmd := reexec.Command("nsInitialisation", rootfs, cid)

        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        cmd.SysProcAttr = &syscall.SysProcAttr{
                Cloneflags: syscall.CLONE_NEWNS |
                        syscall.CLONE_NEWUTS |
                        syscall.CLONE_NEWIPC |
                        syscall.CLONE_NEWPID |
                        syscall.CLONE_NEWNET |
                        syscall.CLONE_NEWUSER,
                UidMappings: []syscall.SysProcIDMap{
                        {
                                ContainerID: 0,
                                HostID:      os.Getuid(),
                                Size:        1,
                        },
                },
                GidMappings: []syscall.SysProcIDMap{
                        {
                                ContainerID: 0,
                                HostID:      os.Getgid(),
                                Size:        1,
                        },
                },
        }

        if err := cmd.Start(); err != nil {
                return err
        }

        //Set Network
        ne := &network.NTconfig {
                Bridge: "docker0",
                Vehost: cid[:4] + "-host",
                Vepeer: cid[:4] + "-peer",
                Mtu   : 1500,
                Ipaddr: "172.17.0.9/24",
        }

        if err := ne.Create(cmd.Process.Pid); err != nil {
                fmt.Println(err)
        }

        //Set Cgroups
        if err := cgroups.AddCIDtoCgroups(cid); err != nil {
                fmt.Println(err)
        }
        if err := cgroups.AddProcstoCgroups(cid, cmd.Process.Pid); err != nil {
                fmt.Println(err)
        }
        if err := cgroups.MemoryLimit(cid, ml); err != nil {
                fmt.Println(err)
        }

        if err := cmd.Wait(); err != nil {
                return err
        }
        fmt.Println("OK")
        if err := cgroups.RemoveCIDfromCgroups(cid); err != nil {
                return err
        }
        return nil
}

func init() {
        reexec.Register("nsInitialisation", nsInitialisation)
        if reexec.Init() {
                os.Exit(0)
        }
}

func nsInitialisation() {
        newrootPath := os.Args[1]

        cid := os.Args[2]
        fmt.Println(os.Getuid())
        if err := mount.PrepareRoot(newrootPath); err != nil {
                fmt.Printf("Error running pivot_root - %s\n", err)
                os.Exit(1)
        }

        if err := syscall.Sethostname([]byte("ns-process")); err != nil {
                fmt.Printf("Error setting hostname - %s\n", err)
                os.Exit(1)
        }

        defer func() {
                ne := &network.NTconfig {Vehost: cid[:4] + "-host"}
                if err := ne.Remove(); err != nil {
                        fmt.Println(err)
                }
        }()

        cmd := exec.Command("/bin/sh")

        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        cmd.Env = []string{"PS1=["+ cid[:10] +"]# "}

        if err := cmd.Run(); err != nil {
                fmt.Printf("Error running the /bin/sh command - %s\n", err)
                os.Exit(1)
        }

}
