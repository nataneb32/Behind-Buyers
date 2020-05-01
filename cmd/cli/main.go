package main

import (
	"fmt"

	"../../pkg/acessListener"
	"../../pkg/reportObserver"
)

func main() {

	ro := reportObserver.CreateReportHandler()
	al := acessListener.CreateAcessListener()
	ro.Subscribe(al)

	ro.Report(map[string]interface{}{
		"acess": "test",
	})

	fmt.Println(al.GetData())
}
