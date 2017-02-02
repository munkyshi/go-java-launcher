package launchlib

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsRunning(t *testing.T) {
	running, err := isRunning(99999)
	require.NoError(t, err)
	assert.False(t, running)

	myPid := os.Getpid()
	running, err = isRunning(myPid)
	require.NoError(t, err)
	assert.True(t, running)
}

func TestIsRunningByPidFile(t *testing.T) {
	running, err := isRunningByPidFile("bogus file")
	require.Error(t, err)
	assert.False(t, running)
	assert.EqualError(t, err, "open bogus file: no such file or directory")

	ioutil.WriteFile("pidfile", []byte("99999"), os.ModePerm)
	running, err = isRunningByPidFile("pidfile")
	require.NoError(t, err)
	assert.False(t, running)

	ioutil.WriteFile("pidfile", []byte(strconv.Itoa(os.Getpid())), os.ModeAppend)
	running, err = isRunningByPidFile("pidfile")
	require.NoError(t, err)
	assert.True(t, running)

	os.Remove("pidfile")
}
