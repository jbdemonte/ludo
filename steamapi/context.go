package steamapi

import "github.com/hajimehoshi/go-steamworks"

var shouldUnload = false
var ready = false

type InputEvent int

const (
	Connected              InputEvent = 1
	Disconnected           InputEvent = 0
	MaxSteamInputConnected int        = 4
)

type InputCallback func(handle uint64, event InputEvent)

var fInputCallback InputCallback

var controllerForGamepadIndex = make(map[int]steamworks.InputHandle_t)
