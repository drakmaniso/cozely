// Package mouse provides support for the mouse.
package mouse

// #cgo windows LDFLAGS: -lSDL2
// #cgo linux freebsd darwin pkg-config: sdl2
// #include "../internal/sdl.h"
import "C"

import (
	"fmt"
	"time"

	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/internal"
)

//------------------------------------------------------------------------------

// Handler receives the mouse events.
var Handler interface {
	MouseMotion(motion geom.IVec2, position geom.IVec2, timestamp time.Duration)
	MouseButtonDown(b Button, clicks int, timestamp time.Duration)
	MouseButtonUp(b Button, clicks int, timestamp time.Duration)
	MouseWheel(motion geom.IVec2, timestamp time.Duration)
} = DefaultHandler{}

// DefaultHandler implements default behavior for all mouse events.
type DefaultHandler struct{}

func (dh DefaultHandler) MouseMotion(rel geom.IVec2, pos geom.IVec2, timestamp time.Duration) {}
func (dh DefaultHandler) MouseButtonDown(b Button, clicks int, timestamp time.Duration)       {}
func (dh DefaultHandler) MouseButtonUp(b Button, clicks int, timestamp time.Duration)         {}
func (dh DefaultHandler) MouseWheel(w geom.IVec2, timestamp time.Duration)                    {}

//------------------------------------------------------------------------------

// Position returns the current mouse position, relative to the game window.
// Updated at the start of each game loop iteration.
func Position() geom.IVec2 {
	return internal.MousePosition
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
	var m C.SDL_bool
	if enabled {
		m = 1
	}
	if C.SDL_SetRelativeMouseMode(m) != 0 {
		return fmt.Errorf("impossible to set relative mouse mode: %s", internal.GetSDLError())
	}
	return nil
}

// GetRelativeMode returns true if the relative mode is enabled.
func GetRelativeMode() bool {
	return C.SDL_GetRelativeMouseMode() == C.SDL_TRUE
}

//------------------------------------------------------------------------------
