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

	steamworks.SteamInput().Init(false)

	go runCallbacks()

	go getHandles()

	go listenControllers()
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

	f.WriteString(ln + "\n")
}

// when no controllers are connected on startup GetActionSetHandle returns 0 (a bug?)
// so, we need to add a workaround to support hot plugged controllers
func getHandles() {
	for !isActionSetDefined() {
		time.Sleep(50 * time.Millisecond)
		if shouldUnload {
			return
		}
		getActionSetHandle()
	}

	activateActionSetForAll()
	getDigitalBinding()
}

func listenControllers() {

	for !shouldUnload {
		if !isActionSetDefined() {
			time.Sleep(10 * time.Millisecond)
			continue
		}

		updateControllerList()

		time.Sleep(100 * time.Millisecond)
	}
}
