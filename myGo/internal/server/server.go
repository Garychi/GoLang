package server

import "Lv/internal/bootstrap"

func Run() {

	Conf := bootstrap.LoadConfig()

	r := SetupRouter()

	server := CreateServer(r, Conf.App.Port, Conf.App.Host)
	ctx := SignalListenAndServe(server)

	<-ctx.Done()
}
