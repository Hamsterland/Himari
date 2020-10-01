package commands

import "github.com/Necroforger/dgrouter/exrouter"

func Truncate(args exrouter.Args) exrouter.Args { return append(args[1:]) }
