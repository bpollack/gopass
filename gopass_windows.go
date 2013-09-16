// Author: benjamin@bitquabit.com (Benjamin Pollack)
package gopass

/*
#include <conio.h>
*/
import "C"

import (
	"errors"
	"fmt"
)

func GetPass(prompt string) (string, error) {
	fmt.Print("Password: ")
	pass := make([]rune, 0)
	var ch rune
	for {
		ch = rune(C._getwch())
		if ch == '\n' || ch == '\r' {
			break
		} else if ch == 3 {
			return "", errors.New("break")
		} else if ch == '\b' && len(pass) > 0 {
			pass = pass[0 : len(pass)-1]
		} else {
			pass = append(pass, ch)
		}
	}
	C._putch('\r')
	C._putch('\n')
	return string(pass), nil
}
