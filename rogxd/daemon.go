package rogxd

import (
	"rogxd/rogxd/bus/controllers"
	"rogxd/rogxd/ota"
)

func StartDaemon() {
	rogConn := controllers.NewRogXConn()
	server := ota.New(rogConn)
	server.Start()
}
