// Copyright (c) 2013-2016 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package main

//------------------------------------------------------------------------------

import . "github.com/drakmaniso/glam/geom"

//------------------------------------------------------------------------------

func cube() []perVertex {
	return []perVertex{
		// Front Face
		{Vec3{-0.5, -0.5, +0.5}, Vec2{0, 1}},
		{Vec3{+0.5, -0.5, +0.5}, Vec2{1, 1}},
		{Vec3{+0.5, +0.5, +0.5}, Vec2{1, 0}},
		{Vec3{-0.5, -0.5, +0.5}, Vec2{0, 1}},
		{Vec3{+0.5, +0.5, +0.5}, Vec2{1, 0}},
		{Vec3{-0.5, +0.5, +0.5}, Vec2{0, 0}},
		// Back Face
		{Vec3{+0.5, -0.5, -0.5}, Vec2{0, 1}},
		{Vec3{-0.5, -0.5, -0.5}, Vec2{1, 1}},
		{Vec3{-0.5, +0.5, -0.5}, Vec2{1, 0}},
		{Vec3{+0.5, -0.5, -0.5}, Vec2{0, 1}},
		{Vec3{-0.5, +0.5, -0.5}, Vec2{1, 0}},
		{Vec3{+0.5, +0.5, -0.5}, Vec2{0, 0}},
		// Right Face
		{Vec3{+0.5, -0.5, +0.5}, Vec2{0, 1}},
		{Vec3{+0.5, -0.5, -0.5}, Vec2{1, 1}},
		{Vec3{+0.5, +0.5, -0.5}, Vec2{1, 0}},
		{Vec3{+0.5, -0.5, +0.5}, Vec2{0, 1}},
		{Vec3{+0.5, +0.5, -0.5}, Vec2{1, 0}},
		{Vec3{+0.5, +0.5, +0.5}, Vec2{0, 0}},
		// Left Face
		{Vec3{-0.5, -0.5, -0.5}, Vec2{0, 1}},
		{Vec3{-0.5, -0.5, +0.5}, Vec2{1, 1}},
		{Vec3{-0.5, +0.5, +0.5}, Vec2{1, 0}},
		{Vec3{-0.5, -0.5, -0.5}, Vec2{0, 1}},
		{Vec3{-0.5, +0.5, +0.5}, Vec2{1, 0}},
		{Vec3{-0.5, +0.5, -0.5}, Vec2{0, 0}},
		// Bottom Face
		{Vec3{-0.5, -0.5, -0.5}, Vec2{0, 1}},
		{Vec3{+0.5, -0.5, -0.5}, Vec2{1, 1}},
		{Vec3{+0.5, -0.5, +0.5}, Vec2{1, 0}},
		{Vec3{-0.5, -0.5, -0.5}, Vec2{0, 1}},
		{Vec3{+0.5, -0.5, +0.5}, Vec2{1, 0}},
		{Vec3{-0.5, -0.5, +0.5}, Vec2{0, 0}},
		// Top Face
		{Vec3{-0.5, +0.5, +0.5}, Vec2{0, 1}},
		{Vec3{+0.5, +0.5, +0.5}, Vec2{1, 1}},
		{Vec3{+0.5, +0.5, -0.5}, Vec2{1, 0}},
		{Vec3{-0.5, +0.5, +0.5}, Vec2{0, 1}},
		{Vec3{+0.5, +0.5, -0.5}, Vec2{1, 0}},
		{Vec3{-0.5, +0.5, -0.5}, Vec2{0, 0}},
	}
}

//------------------------------------------------------------------------------
