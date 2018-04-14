// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/x/gl"
)

// Declarations ////////////////////////////////////////////////////////////////

type loop01 struct {
	// OpenGL Object
	pipeline *gl.Pipeline
}

// Initialization //////////////////////////////////////////////////////////////

func Example_01FirstTriangle() {
	cozely.Events.Resize = func() {
		s := cozely.WindowSize()
		gl.Viewport(0, 0, int32(s.C), int32(s.R))
	}
	l := loop01{}
	err := cozely.Run(&l)
	if err != nil {
		cozely.ShowError(err)
		return
	}
	//Output:
}

func (l *loop01) Enter() error {
	// Create and configure the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader01.vert"),
		gl.Shader(cozely.Path()+"shader01.frag"),
		gl.Topology(gl.Triangles),
	)

	return cozely.Error("gfx", gl.Err())
}

func (loop01) Leave() error {
	return nil
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop01) React() error {
	return nil
}

func (loop01) Update() error {
	return nil
}

func (l *loop01) Render() error {
	l.pipeline.Bind()
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	gl.Draw(0, 3)
	l.pipeline.Unbind()

	return gl.Err()
}
