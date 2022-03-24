package main

import (
	"github.com/lyft/clutch/backend/gateway"

	"github.com/colbylwilliams/clutch-teamcloud-gateway/backend/cmd/assets"
	"github.com/colbylwilliams/clutch-teamcloud-gateway/backend/module/echo"
	teamcloudmod "github.com/colbylwilliams/clutch-teamcloud-gateway/backend/module/teamcloud"
	teamcloudservice "github.com/colbylwilliams/clutch-teamcloud-gateway/backend/service/teamcloud"
)

func main() {
	flags := gateway.ParseFlags()

	components := gateway.CoreComponentFactory

	// Add custom components.
	components.Modules[echo.Name] = echo.New
	components.Modules[teamcloudmod.Name] = teamcloudmod.New

	components.Services[teamcloudservice.Name] = teamcloudservice.New

	gateway.Run(flags, components, assets.VirtualFS)
}
