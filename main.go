package main

import (
	_ "zzw/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"zzw/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
