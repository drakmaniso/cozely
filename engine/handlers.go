// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package engine

import (
	"github.com/drakmaniso/glam/geom"
	"github.com/drakmaniso/glam/key"
	"github.com/drakmaniso/glam/mouse"
)

//------------------------------------------------------------------------------

// Handler implements the game loop.
var Handler interface {
	Update()
	Draw()
	Quit()
}

//------------------------------------------------------------------------------

var defaultHandler DefaultHandler

// DefaultHandler provides a default implementation for all handlers.
type DefaultHandler struct{}

func (dh DefaultHandler) Update() {}
func (dh DefaultHandler) Draw()   {}
func (dh DefaultHandler) Quit()   {}

func (dh DefaultHandler) KeyDown(l key.Label, p key.Position, time uint32) {}
func (dh DefaultHandler) KeyUp(l key.Label, p key.Position, time uint32)   {}

func (dh DefaultHandler) MouseMotion(rel geom.IVec2, pos geom.IVec2, time uint32) {}
func (dh DefaultHandler) MouseButtonDown(b mouse.Button, clicks int, time uint32) {}
func (dh DefaultHandler) MouseButtonUp(b mouse.Button, clicks int, time uint32)   {}
func (dh DefaultHandler) MouseWheel(w geom.IVec2, time uint32)                    {}

//------------------------------------------------------------------------------