package router

import "go-app/cmd/api/router/factory"

func (r routerImpl) routes() {
	factoryCtrl := factory.NewCtrlFactory()

	r.router.GET("/ping", factoryCtrl.BuildPingController().Ping)
}
