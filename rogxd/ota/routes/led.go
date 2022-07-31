package routes

import (
	"rogxd/rogxd/bus/controllers"
	"rogxd/rogxd/ota/models"
)

var LedRoutes = []Route{
	{
		Method: Get,
		Path:   "/brightness",
		Handler: func(ctx Context) {
			controller := ctx.RogX.LedController
			ctx.Json(models.BrightnessResponse{Brightness: byte(controller.LedBrightness())})
		},
	},
	{
		Method: Patch,
		Path:   "/brightness",
		Handler: func(ctx Context) {
			var body models.BrightnessRequest
			if e := ctx.Body(&body); e != nil {
				ctx.Bad(e.Error())
				return
			}
			controller := ctx.RogX.LedController
			switch body.Operation {
			case models.Prev:
				controller.PrevLedBrightness()
			case models.Next:
				controller.NextLedBrightness()
			default:
				if body.Brightness < 0 || body.Brightness > 3 {
					ctx.Bad("Invalid brightness value.")
					return
				}
				controller.SetLedBrightness(controllers.Brightness(body.Brightness))
			}
			ctx.Json(models.BrightnessResponse{Brightness: byte(controller.LedBrightness())})
		},
	},
	{
		Path:   "/modes",
		Method: Get,
		Handler: func(ctx Context) {
			ctx.Json(ctx.RogX.LedController.LedModes())
		},
	},
}
