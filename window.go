// Copyright (c) 2013-2017 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package carol

//------------------------------------------------------------------------------

import (
	"github.com/drakmaniso/carol/internal"
)

//------------------------------------------------------------------------------

// HasFocus returns true if the game windows has focus.
func HasFocus() bool {
	return internal.HasFocus
}

// HasMouseFocus returns true if the mouse is currently inside the game window.
func HasMouseFocus() bool {
	return internal.HasMouseFocus
}

// WindowSize returns the size of the window in (screen) pixels.
func WindowSize() (width, height int32) {
	return internal.Window.Width, internal.Window.Height
}

//TODO
// ScreenSize returns the size of the framebuffer in (framebuffer) pixels.
// func ScreenSize() pixel.Coord {
// 	return gpu.Framebuffer.Size
// }

//TODO
// PixelSize returns the size of framebuffer pixels, in screen pixels.
// func PixelSize() pixel.Coord {
// 	return gpu.Framebuffer.PixelSize
// }

//------------------------------------------------------------------------------
