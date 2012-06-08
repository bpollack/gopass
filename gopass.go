// Author: johnsiilver@gmail.com (John Doak)

/*
gopass is a library for getting hidden input from a terminal.

This library's main use is to allow a user to enter a password at the
command line without having it echoed to the screen.

The libary currently supports unix systems by manipulating stty.

This code is based upon code by RogerV in the golang-nuts thread:
https://groups.google.com/group/golang-nuts/browse_thread/thread/40cc41e9d9fc9247
*/
package gopass

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

const (
	sttyArg0   = "/bin/stty"
	exec_cwdir = ""
	//ws syscall.WaitStatus = 0
)

// Tells the terminal to turn echo off.
var sttyArgvEOff []string = []string{"stty", "-echo"}

// Tells the terminal to turn echo on.
var sttyArgvEOn []string = []string{"stty", "echo"}

var ws syscall.WaitStatus = 0

// GetPass gets input hidden from the terminal from a user.
// This is accomplished by turning off terminal echo,
// reading input from the user and finally turning on terminal echo. 
// prompt is a string to display before the user's input.
func GetPass(prompt string) (passwd string, err error) {
	// Display the prompt.
	fmt.Print(prompt)

	// File descriptors for stdin, stdout, and stderr.
	fd := []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}

	// Turn off the terminal echo.
	pid, err := syscall.ForkExec(sttyArg0, sttyArgvEOff, &syscall.ProcAttr{Dir: exec_cwdir, Files: fd})
	if err != nil {
		return passwd, fmt.Errorf("failed turning off console echo for password entry:\n\t%s", err)
	}
	rd := bufio.NewReader(os.Stdin)
	syscall.Wait4(pid, &ws, 0, nil)

	line, err := rd.ReadString('\n')
	if err == nil {
		passwd = strings.TrimSpace(line)
	} else {
		err = fmt.Errorf("failed during password entry: %s", err)
	}

	// Turn on the terminal echo.
	pid, e := syscall.ForkExec(sttyArg0, sttyArgvEOn, &syscall.ProcAttr{Dir: exec_cwdir, Files: fd})
	if e == nil {
		syscall.Wait4(pid, &ws, 0, nil)
	} else if err == nil {
		err = fmt.Errorf("failed turning on console echo post password entry:\n\t%s", e)
	}

	// Carraige return after the user input.
	fmt.Println("")

	return passwd, err
}
