package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {

	err := checkForUnsupportedPlatform()
	if err != nil {
		log.Fatal(err)
	}

	err = downloadUpdates()
	if err != nil {
		log.Fatal(err)
	}

	err = installUpdatesAndReboot()
	if err != nil {

		log.Fatal(err)
	}

}

func checkForUnsupportedPlatform() error {
	if runtime.GOOS != "darwin" && runtime.GOOS != "windows" {
		err := errors.New("Unsupported platform")
		return err
	}

	return nil
}

func downloadUpdates() error {
	cmd := exec.Command("/usr/sbin/softwareupdate", "-dla")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("wuauclt /detectnow")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))

	return nil
}

func installUpdatesAndReboot() error {
	cmd := exec.Command("/usr/sbin/softwareupdate", "-dia", "--restart")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("wuauclt /updatenow")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))

	if runtime.GOOS == "windows" {
		cmd = exec.Command("shutdown", "/r")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}

		fmt.Print(string(out))
	}

	return nil
}
