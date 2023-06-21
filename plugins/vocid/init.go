package vocid

import (
	"github.com/infinitybotlist/sysmanage-web/plugins/frontend"
	"github.com/infinitybotlist/sysmanage-web/types"
)

const ID = "vocid"

func InitPlugin(c *types.PluginConfig) error {
	frontend.AddLink(c, frontend.Link{
		Title:       "Vocid",
		Description: "Open the Vocid panel",
		LinkText:    "Open!",
		Href:        "@root",
	})

	return nil
}
