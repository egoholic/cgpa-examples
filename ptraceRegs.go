package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	var r syscall.PtraceRegs
	cmd := exec.CmdCommand(os.Args[1], os.Args[2:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	err := cmd.Start()
	if err != nil {
		fmt.Println("Start:", err)
		return
	}

	err = cmd.Wait()
	fmt.Printf("State: %v\n", err)
	wpid := cmd.Process.Pid

	err = syscall.PtraceGetRegs(wpid, &r)
	if err != nil {
		fmt.Println("PtraceGetRegs:", err)
		return
	}
	fmt.Printf("Registers: %#v\n", r)
	fmt.Printf("R15=%d, Gs=%d\n", r.R15, r.Gs)

}
