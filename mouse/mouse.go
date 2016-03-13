// Package mouse provides support for the mouse.
package mouse

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../internal/internal.h"
import "C"

import (
	"log"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// Handler receives the mouse events. 
var Handler interface {
	MouseMotion(rel geom.IVec2, pos geom.IVec2, b ButtonState, time uint32)
	MouseButtonDown(b Button, clicks int, pos geom.IVec2, time uint32)
	MouseButtonUp(b Button, clicks int, pos geom.IVec2, time uint32)
	MouseWheel(w geom.IVec2, time uint32)
}

//------------------------------------------------------------------------------

// Position returns the current mouse position, relative to the game window.
// Updated at the start of each game loop iteration.
func Position() geom.IVec2 {
	return geom.IVec2{X: int32(C.mouseX), Y: int32(C.mouseY)}
}

// Delta returns the mouse position relative to the last call of Delta.
func Delta() geom.IVec2 {
	result := internal.MouseDelta
	internal.MouseDelta.X, internal.MouseDelta.Y = 0, 0
	return result
}

// SetRelativeMode enables or disables the relative mode, where the mouse is
// hidden and mouse motions are continuously reported.
func SetRelativeMode(enabled bool) error {
	var err error
	var m C.SDL_bool
	if enabled {
		m = 1
	}
	if C.SDL_SetRelativeMouseMode(m) != 0 {
		err = internal.GetSDLError()
		log.Print(err)
	}
	return err
}

// GetRelativeMode returns true if the relative mode is enabled.
func GetRelativeMode() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

//------------------------------------------------------------------------------
