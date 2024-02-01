package input

import (
	"github.com/libretro/ludo/libretro"
	"github.com/libretro/ludo/steamapi"
)

var steamInputBinds = map[steamapi.Action]uint32{
	steamapi.ButtonDpadUp:    libretro.DeviceIDJoypadUp,
	steamapi.ButtonDpadDown:  libretro.DeviceIDJoypadDown,
	steamapi.ButtonDpadLeft:  libretro.DeviceIDJoypadLeft,
	steamapi.ButtonDpadRight: libretro.DeviceIDJoypadRight,

	steamapi.ButtonA: libretro.DeviceIDJoypadA,
	steamapi.ButtonB: libretro.DeviceIDJoypadB,
	/*
		glfw.ButtonSquare:   libretro.DeviceIDJoypadY,
		glfw.ButtonTriangle: libretro.DeviceIDJoypadX,

		glfw.ButtonLeftBumper:  libretro.DeviceIDJoypadL,
		glfw.ButtonRightBumper: libretro.DeviceIDJoypadR,

		glfw.ButtonLeftThumb:  libretro.DeviceIDJoypadL3,
		glfw.ButtonRightThumb: libretro.DeviceIDJoypadR3,

		glfw.ButtonStart: libretro.DeviceIDJoypadStart,
		glfw.ButtonBack:  libretro.DeviceIDJoypadSelect,
		glfw.ButtonGuide: ActionMenuToggle,
		/**/
}
