# gopass

gopass is a cross-platform library to read passwords from the terminal
without echoing input.  At the moment, gopass is known to work on OS X,
Linux, FreeBSD, and Windows.

## Getting gopass

The usual `go get github.com/bpollack/gopass` is all you need to install
gopass.  Note that, to build gopass on Windows, you will need an appropriate
MinGW installation available in your path to build the cgo module.

## Usage

`gopass` currently only exports a single function, `GetPass`, which
takes a prompt and returns the *trimmed* input--i.e., with leading and
trailing whitespace removed.  The correct usage is the rather
straight-forward

    name, err := gopass.GetPass("Enter your password: ")
    if err == nil {
        fmt.Printf("Your password is %v\n", name)
    }

The authors generally recommend that you do not actually echo passwords
back to the user.
