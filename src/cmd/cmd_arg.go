package cmd

import (
	"fmt"
	"os"
	"strings"
)

type CmdArg struct {
	args map[string]string
}

func NewCmdArg() CmdArg {
	_args := map[string]string{}

	for _, arg := range os.Args {
		ary := strings.Split(arg, "=")
		if len(ary) != 2 {
			continue
		}
		_args[ary[0]] = ary[1]
	}

	return CmdArg{
		args: _args,
	}
}

func (c CmdArg) Validate(keys []string) bool {
	for _, key := range keys {
		if _, existsKey := c.args[key]; !existsKey {
			return false
		}
	}

	return true
}

func (c CmdArg) Get(key string) (*string, error) {
	val, existsKey := c.args[key]
	if !existsKey {
		return nil, fmt.Errorf("not exists key, %s", key)
	}

	return &val, nil
}

func (c CmdArg) IsEmpty() bool {
	return len(c.args) == 0
}
