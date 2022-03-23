package main

import (
	"github.com/lyft/clutch/backend/gateway"

	"github.com/colbylwilliams/clutch-teamcloud-gateway/backend/cmd/assets"
	"github.com/colbylwilliams/clutch-teamcloud-gateway/backend/module/echo"
)

func main() {
	flags := gateway.ParseFlags()

	components := gateway.CoreComponentFactory

	// Add custom components.
	components.Modules[echo.Name] = echo.New

	gateway.Run(flags, components, assets.VirtualFS)
}
