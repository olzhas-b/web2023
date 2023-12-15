package main

import (
	"context"
	"runtime"

	"github.com/olzhas-b/social-media/modules/daylight"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ctx := context.Background()
	daylight.Start(ctx)
}
