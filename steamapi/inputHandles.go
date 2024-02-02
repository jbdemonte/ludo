package steamapi

import (
	"github.com/hajimehoshi/go-steamworks"
	"log"
)

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
