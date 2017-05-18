package main

import (
	"github.com/keysonZZZ/kgo/kgoGoSource"
	"github.com/keysonZZZ/kgo/kgoCmd"
)

func buildCmd() {
	kgoGoSource.MustGoInstallWithDefaultEnv("github.com/bronze1man/kmg/kmg")
	kgoCmd.ProxyRun("cp -f ./bin/kmg /usr/local/bin/kmg")
}

