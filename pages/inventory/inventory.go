package inventory

import (
	"net/http"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Get inventory page.
func Get(ctx aero.Context) error {
	user := utils.GetUser(ctx)
	viewUser := user

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	inventory, err := arn.GetInventory(viewUser.ID)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching inventory data", err)
	}

	return ctx.HTML(components.Inventory(inventory, viewUser, user))
}
