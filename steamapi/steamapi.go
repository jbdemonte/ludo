package steamapi

import "C"
import (
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/go-steamworks"
)

const appID = 2792270

func Init() {
	if steamworks.RestartAppIfNecessary(appID) {
		os.Exit(1)
	}
	if !steamworks.Init() {
		panic("steamworks.Init failed")
	}

	go listenControllers()
}

var shouldUnload = false

func Unload() {
	shouldUnload = true
	time.Sleep(200)
	steamworks.SteamInput().Shutdown()
}

var ready = false

type InputEvent int

const (
	Connected    InputEvent = 1
	Disconnected InputEvent = 0
)

type InputCallback func(handle uint64, event InputEvent)

var fInputCallback InputCallback
var inputHandles []steamworks.InputHandle_t

func SetInputCallback(cb InputCallback) {
	fInputCallback = cb
	ready = true
}

func isInHandleList(source []steamworks.InputHandle_t, handle steamworks.InputHandle_t) bool {
	for _, h := range source {
		if h == handle {
			return true
		}
	}
	return false
}

func isNewHandle(handle steamworks.InputHandle_t) bool {
	return !isInHandleList(inputHandles, handle)
}

func addHandle(handle steamworks.InputHandle_t) {
	log.Println("[Steam] Adding input", handle)
	inputHandles = append(inputHandles, handle)
	fInputCallback(uint64(handle), Connected)
}

func removeHandle(handle steamworks.InputHandle_t) {
	log.Println("[Steam] Removing input", handle)
	var handles []steamworks.InputHandle_t
	for _, h := range inputHandles {
		if h != handle {
			handles = append(handles, h)
		}
	}
	inputHandles = handles
	fInputCallback(uint64(handle), Disconnected)
}

func steamRunCallbacks() {
	for !shouldUnload {
		steamworks.RunCallbacks()
		time.Sleep(25 * time.Millisecond)
	}
}

func readControllers() {
	steamInput := steamworks.SteamInput()
	actionSetHandle := steamInput.GetActionSetHandle("GameControls")

	getDigitalBinding()

	for !shouldUnload {
		for _, inputHandle := range inputHandles {
			steamInput.ActivateActionSet(inputHandle, actionSetHandle)

		}
		time.Sleep(1 * time.Millisecond)
	}
}

// debug
func Addlog(ln string) {
	f, err := os.OpenFile("steamwork.log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(ln)
	f.WriteString("\n")
}

func listenControllers() {

	steamInput := steamworks.SteamInput()

	steamInput.Init(false)

	go steamRunCallbacks()

	go readControllers()

	for !shouldUnload {
		if !ready {
			continue
		}

		var handles = steamInput.GetConnectedControllers()

		//ntf.DisplayAndLog(ntf.Info, "Input", "handles len %d", len(handles))

		for _, handle := range handles {
			if handle > 0 && isNewHandle(handle) {
				addHandle(handle)
			}
		}

		for _, handle := range inputHandles {
			if !isInHandleList(handles, handle) {
				removeHandle(handle)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
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
