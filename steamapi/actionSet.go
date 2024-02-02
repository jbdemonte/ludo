package steamapi

import (
	"github.com/hajimehoshi/go-steamworks"
)

var actionSetHandle steamworks.InputActionSetHandle_t

func getActionSetHandle() {
	actionSetHandle = steamworks.SteamInput().GetActionSetHandle("GameControls")
}

func activateActionSet(inputHandle steamworks.InputHandle_t) {
	steamworks.SteamInput().ActivateActionSet(inputHandle, actionSetHandle)
}

func activateActionSetForAll() {
	steamworks.SteamInput().ActivateActionSet(steamworks.STEAM_INPUT_HANDLE_ALL_CONTROLLERS, actionSetHandle)
}

func isActionSetDefined() bool {
	return actionSetHandle > 0
}
