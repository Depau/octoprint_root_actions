package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func reboot() error {
	cmd := exec.Command("/usr/bin/reboot")
	return cmd.Run()
}

func systemctl_restart(unit string) error {
	cmd := exec.Command("/usr/bin/systemctl", "restart", unit)
	return cmd.Run()
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s [reboot|restart_octoprint|restart_klippy|restart_ustreamer]\n", os.Args[0])
	}

	var err error
	switch os.Args[1] {
	case "reboot":
		err = reboot()
	case "restart_octoprint":
		err = systemctl_restart("octoprint.service")
	case "restart_klippy":
		err = systemctl_restart("klippy.service")
	case "restart_ustreamer":
		err = systemctl_restart("ustreamer.service")
	default:
		err = errors.New("invalid command")
	}

	if err != nil {
		log.Fatalln("error executing action: ", err)
	}
}
