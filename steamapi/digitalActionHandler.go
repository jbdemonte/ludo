package steamapi

import (
	"github.com/hajimehoshi/go-steamworks"
)

type Action int

const (
	ButtonA         Action = 0
	ButtonB         Action = 1
	ButtonDpadUp    Action = 2
	ButtonDpadRight Action = 3
	ButtonDpadDown  Action = 4
	ButtonDpadLeft  Action = 5
)

type DigitalState map[Action]steamworks.InputDigitalActionData_t

var inputBinds = map[Action]steamworks.InputDigitalActionHandle_t{}

func getDigitalBinding() {
	steamInput := steamworks.SteamInput()

	inputBinds[ButtonA] = steamInput.GetDigitalActionHandle("Action_A")
	inputBinds[ButtonB] = steamInput.GetDigitalActionHandle("Action_B")
	inputBinds[ButtonDpadUp] = steamInput.GetDigitalActionHandle("Action_Up")
	inputBinds[ButtonDpadRight] = steamInput.GetDigitalActionHandle("Action_Right")
	inputBinds[ButtonDpadDown] = steamInput.GetDigitalActionHandle("Action_Down")
	inputBinds[ButtonDpadLeft] = steamInput.GetDigitalActionHandle("Action_Left")
}

func getDigitalState(handle steamworks.InputHandle_t) DigitalState {
	steamInput := steamworks.SteamInput()
	return DigitalState{
		ButtonA:         steamInput.GetDigitalActionData(handle, inputBinds[ButtonA]),
		ButtonB:         steamInput.GetDigitalActionData(handle, inputBinds[ButtonB]),
		ButtonDpadUp:    steamInput.GetDigitalActionData(handle, inputBinds[ButtonDpadUp]),
		ButtonDpadDown:  steamInput.GetDigitalActionData(handle, inputBinds[ButtonDpadDown]),
		ButtonDpadLeft:  steamInput.GetDigitalActionData(handle, inputBinds[ButtonDpadLeft]),
		ButtonDpadRight: steamInput.GetDigitalActionData(handle, inputBinds[ButtonDpadRight]),
	}
}
