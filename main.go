package main

import (
	"github.com/sskmy1024/PartnerAssistant/api"
	"github.com/sskmy1024/PartnerAssistant/infrastructures"
)

func main() {
	infrastructures.InitEnvironment()

	s := infrastructures.NewServer()
	api.Router(s)

	s.Logger.Fatal(s.Start(":8080"))
}
