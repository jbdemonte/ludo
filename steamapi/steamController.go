package steamapi

import "github.com/hajimehoshi/go-steamworks"

type SteamController uint

func (controller SteamController) handle() steamworks.InputHandle_t {
	return inputHandles[controller]
}

func (controller SteamController) GetDigitalState() DigitalState {
	handle := controller.handle()
	activateActionSet(handle)
	return getDigitalState(handle)
}

func Joystick(index SteamController) SteamController {
	if index > JoystickLast() {
		return 0
	}
	return index
}

// respect the glfw approach, JoystickLast is not usable and is helpful for loops
func JoystickLast() SteamController {
	return SteamController(len(inputHandles))
}
