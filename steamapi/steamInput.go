package steamapi

import "github.com/hajimehoshi/go-steamworks"

type SteamController uint

func (controller SteamController) handle() steamworks.InputHandle_t {
	return inputHandles[controller]
}

func (controller SteamController) GetDigitalState() DigitalState {
	return getDigitalState(controller.handle())
}
