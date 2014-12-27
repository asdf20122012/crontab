package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"syscall"
)

var d *bool = flag.Bool("d", true, "daemon")

func fork() (int, syscall.Errno) {
	r1, r2, err := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		return 0, err
	}

	if runtime.GOOS == "darwin" && r2 == 1 {
		r1 = 0
	}

	return int(r1), 0

}

func daemon() {
	pid, err := fork()
	if err != 0 {
		log.Fatalln("Daemon Error!")
	}
	if pid != 0 {
		os.Exit(0)
	}
	syscall.Umask(0)
	syscall.Setsid()
	os.Chdir("/")

	f, _ := os.Open("/dev/null")
	devnull := f.Fd()

	syscall.Dup2(int(devnull), int(os.Stdin.Fd()))
	syscall.Dup2(int(devnull), int(os.Stdout.Fd()))
	syscall.Dup2(int(devnull), int(os.Stderr.Fd()))
}