package routes

import (
	"rogxd/rogxd/ota/models"
	"rogxd/rogxd/sys/dmi"
)

var SystemRoutes = []Route{
	{
		Path:   "/device",
		Method: Get,
		Handler: func(ctx Context) {
			family, boardName := dmi.Prop(dmi.ProductFamily), dmi.Prop(dmi.BoardName)
			if len(family) == 0 {
				ctx.Error(500, "Could not get device family.")
			} else if len(boardName) == 0 {
				ctx.Error(500, "Could not get board name.")
			} else {
				ctx.Json(models.DeviceInfo{
					ProductFamily: family,
					BoardName:     boardName,
				})
			}
		},
	},
}
