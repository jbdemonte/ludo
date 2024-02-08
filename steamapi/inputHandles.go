package steamapi

import (
	"fmt"
	"github.com/hajimehoshi/go-steamworks"
)

func updateControllerList() {
	steamInput := steamworks.SteamInput()

	var newHandleList []steamworks.InputHandle_t

	previousHandleList := make([]steamworks.InputHandle_t, 0, len(controllerForGamepadIndex))
	for key := range controllerForGamepadIndex {
		if controllerForGamepadIndex[key] > 0 {
			previousHandleList = append(previousHandleList, controllerForGamepadIndex[key])
		}
	}

	Addlog(fmt.Sprintf("previousHandleList len = %d", len(previousHandleList)))

	steamInput.RunFrame()

	// https://steamcommunity.com/groups/steamworks/discussions/0/4206994023678108167/
	handles := steamInput.GetConnectedControllers()
	Addlog(fmt.Sprintf("handles len = %d", len(handles)))

	for _, handle := range handles {
		Addlog(fmt.Sprintf("> GetGamepadIndexForController %d = %d", handle, steamInput.GetGamepadIndexForController(handle)))
	}

	for i := 0; i < MaxSteamInputConnected; i++ {
		handle := steamInput.GetControllerForGamepadIndex(i)
		Addlog(fmt.Sprintf("GetControllerForGamepadIndex %d = %d", i, handle))
		controllerForGamepadIndex[i] = handle
		if handle > 0 {
			newHandleList = append(newHandleList, handle)
		}
	}

	for _, handle := range previousHandleList {
		if handle > 0 && !isInHandleList(newHandleList, handle) {
			fInputCallback(uint64(handle), Disconnected)
		}
	}
}

func isInHandleList(source []steamworks.InputHandle_t, handle steamworks.InputHandle_t) bool {
	for _, h := range source {
		if h == handle {
			return true
		}
	}
	return false
}
