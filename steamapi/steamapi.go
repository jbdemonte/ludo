package steamapi

import "C"
import (
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

	go run()
}

func Unload() {
	shouldUnload = true
	time.Sleep(200)
	steamworks.SteamInput().Shutdown()
}

func SetInputCallback(cb InputCallback) {
	fInputCallback = cb
	ready = true
}

func runCallbacks() {
	for !shouldUnload {
		steamworks.RunCallbacks()
		time.Sleep(10 * time.Millisecond)
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

// when no controllers are connected on startup GetActionSetHandle returns 0 (a bug?)
// so, we need to add a workaround to support hot plugged controllers
func ensureActionSetIsDefined() {
	for !isActionSetDefined() {
		time.Sleep(50 * time.Millisecond)
		getActionSetHandle()
	}

	activateActionSetForAll()
	getDigitalBinding()
}

func run() {
	steamInput := steamworks.SteamInput()

	steamInput.Init(false)

	go runCallbacks()

	go ensureActionSetIsDefined()

	for !shouldUnload {
		if !ready {
			continue
		}

		handles := steamInput.GetConnectedControllers()

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
