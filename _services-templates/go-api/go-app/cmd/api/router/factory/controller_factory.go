package factory

import "go-app/cmd/api/controllers/ping"

type ControllerBuilder interface {
	BuildPingController() ping.PingController
}

type controllerBuildImpl struct {}

func NewCtrlFactory() ControllerBuilder {
	return &controllerBuildImpl{}
}

func (ctrlFactory *controllerBuildImpl) BuildPingController() ping.PingController {
	return ping.NewPingController()
}
