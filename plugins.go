package main

import (
	"sysmanage/plugins/vocid"

	"github.com/infinitybotlist/sysmanage-web/plugins/actions"
	"github.com/infinitybotlist/sysmanage-web/plugins/frontend"
	"github.com/infinitybotlist/sysmanage-web/plugins/nginx"
	"github.com/infinitybotlist/sysmanage-web/plugins/persist"
	"github.com/infinitybotlist/sysmanage-web/plugins/systemd"
	"github.com/infinitybotlist/sysmanage-web/types"
)

var meta = types.ServerMeta{
	ConfigVersion: 1,
	Plugins: []types.Plugin{
		{
			ID: nginx.ID,
			Init: nginx.InitPlugin,
			Frontend: types.Provider{
				Provider: "@core",
			},
		},
		{
			ID: systemd.ID,
			Init: systemd.InitPlugin,
			Frontend: types.Provider{
				Provider: "@core",
			},
		},
		// Persist has no frotend, it is a backend plugin
		{
			ID: persist.ID,
			Init: persist.InitPlugin,
		},
		{
			ID: actions.ID,
			Init: actions.InitPlugin,
			Frontend: types.Provider{
				Provider: "@core",
			},
		},
		// Frontend has no frontend, it is a backend plugin
		{
			ID: frontend.ID,
			Init: frontend.InitPlugin,
		},
		// Vocid
		{
			ID: vocid.ID,
			Init: vocid.InitPlugin,
			Frontend: types.Provider{
				Provider: "frontend/extplugins/vocid", // This is the path to the plugin's frontend
			},
		},
	},
	Frontend: types.FrontendConfig{
		FrontendProvider: types.Provider{
			Provider: "frontend",
		},
		ComponentProvider: types.Provider{
			Provider: "@core",
		},
		CorelibProvider: types.Provider{
			Provider: "@core",
		}, // Use a custom corelib provider if you want to modify the corelib
		ExtraFiles: map[string]string{
			"frontend/src/lib/images":          "$lib/images",
			"frontend/src/lib/images/logo.png": "static/favicon.png",
		},
	},
}
