package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

type State struct {
	Args    []string
	Command string
	Path    string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	currentUser := getUser()
	hostname := getHostName()
	currentPath := getPath()

	for {
		_, _ = fmt.Fprintf(os.Stdin, "%s@%s:%s$", hostname, currentUser, currentPath)

		args := getInput(reader, currentPath)

		output, err := execShellCommand(&args)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		currentPath = args.Path

		_, _ = fmt.Fprintf(os.Stdin, "%s", output)
	}
}

func getUser() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return currentUser.Username
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

func getPath() string {
	fullPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return fullPath
}

func getInput(reader *bufio.Reader, path string) State {
	input, err := reader.ReadString('\n')
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	input = removeRedundantSpaces(input)

	var args []string
	for _, arg := range strings.Split(input, " ") {
		args = append(args, arg)
	}

	return State{
		Args:    args[1:],
		Command: args[0],
		Path:    path,
	}
}

func execShellCommand(currentState *State) (string, error) {
	switch (*currentState).Command {
	case "cd":
		return cd(currentState)
	case "pwd":
		return pwd(currentState)
	case "echo":
		return echo(currentState)
	case "kill":
		return kill(currentState)
	case "ps":
		return ps(currentState)
	case "fork":
		return fork(currentState)
	case "exec":
		return doExec(currentState)
	case "quit":
		return quit()
	case "":
		return "", nil
	default:
		return "", fmt.Errorf("%s: command not found", currentState.Command)
	}
}

func removeRedundantSpaces(input string) string {
	return strings.Join(strings.Fields(input), " ")
}

func cd(state *State) (string, error) {
	err := os.Chdir(state.Args[0])
	if err != nil {
		return "", err
	}

	newLocation, err := os.Getwd()
	if err != nil {
		return "", err
	}

	state.Path = newLocation
	return "", nil
}

func pwd(state *State) (string, error) {
	return state.Path + "\n", nil
}

func echo(state *State) (string, error) {
	return strings.Join(state.Args, " ") + "\n", nil
}

func kill(state *State) (string, error) {
	pid, err := strconv.Atoi(state.Args[0])
	if err != nil {
		return "", err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return "", err
	}

	err = process.Kill()
	if err != nil {
		return "", err
	}

	return "", nil
}

func ps(state *State) (string, error) {
	output, err := exec.Command(state.Command).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func fork(state *State) (string, error) {
	cmd := exec.Command(state.Command, state.Args...)
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(cmd.Process.Pid), nil
}

func doExec(state *State) (string, error) {
	cmd := exec.Command(state.Command, state.Args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func quit() (string, error) {
	os.Exit(0)
	return "", nil
}
