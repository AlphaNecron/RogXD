package ota

import (
	"rogxd/rogxd/ota/routes"
)

type Party struct {
	Path   string
	Routes []routes.Route
}

var Parties = []Party{
	{
		Path:   "/led",
		Routes: routes.LedRoutes,
	},
	{
		Path:   "/system",
		Routes: routes.SystemRoutes,
	},
}
