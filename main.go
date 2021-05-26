package main

import (
	"github.com/Spores-Labs/spores-nft-backend/app"
	"github.com/Spores-Labs/spores-nft-backend/conf"
)

//APIv1 ...
func APIv1() app.IHandlerBuildable {
	hPool := app.HandlerPool{}
	hPool.Push(&app.UserHandler{})
	return &hPool
}
func main() {
	confOption := conf.GetConfigOption()
	hbuilder := APIv1()
	s := app.NewServiceServer(confOption.Addr, hbuilder)
	s.Start()
}
