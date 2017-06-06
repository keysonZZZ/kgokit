package main

import (
	"github.com/keysonZZZ/kgo/kgoCmd"
	"github.com/keysonZZZ/kgo/kgoGoSource"
)

func buildCmd() {
	kgoGoSource.MustGoInstallWithDefaultEnv("github.com/keysonZZZ/kgo/kgo")
	kgoCmd.ProxyRun("cp -f ./bin/kgo /usr/local/bin/kgo")
}
