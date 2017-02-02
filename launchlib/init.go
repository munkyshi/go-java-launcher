package launchlib

import (
	"io/ioutil"
	"strconv"

	"github.com/keybase/go-ps"
)

func isRunning(pid int) (bool, error) {
	p, err := ps.FindProcess(pid)
	if err != nil || p == nil {
		return false, err
	} else {
		return true, nil
	}
}

func isRunningByPidFile(pidFile string) (bool, error) {
	bytes, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return false, err
	}

	pid, err := strconv.Atoi(string(bytes[:]))
	if err != nil {
		return false, err
	}

	return isRunning(pid)
}
