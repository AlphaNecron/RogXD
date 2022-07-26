package rogxd

import (
	"fmt"
	"rogxd/rogxd/bus/controllers"
)

func StartDaemon() {
	rogConn := controllers.NewRogXConn()
	ctrl := rogConn.LedController()
	c2 := rogConn.ChargeController()
	c2.WatchChargeLimit(func(limit byte) {
		fmt.Println(limit)
	})
	modes := ctrl.LedModes()
	for key, mode := range modes {
		fmt.Printf("%s: %v\n", key, mode)
	}
	select {}
}
