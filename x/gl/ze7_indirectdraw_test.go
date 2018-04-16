// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package gl_test

import (
	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/coord"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/space"
	"github.com/cozely/cozely/x/gl"
	"github.com/cozely/cozely/x/math32"
)

// Declarations ////////////////////////////////////////////////////////////////

// Input Bindings
// (same as in FirstCube example)

type loop07 struct {
	// OpenGL objects
	pipeline    *gl.Pipeline
	perFrameUBO gl.UniformBuffer

	// Tranformation matrices
	screenFromView  space.Matrix // projection matrix
	viewFromWorld   space.Matrix // view matrix
	worldFromObject space.Matrix // model matirx

	// Cube state
	position   coord.XYZ
	yaw, pitch float32
}

// Uniform buffer
// (same as in FirstCube example)
// type perObject struct {
// 	screenFromObject space.Matrix
// }

// Indirect Command Buffer
var commands = []gl.DrawIndirectCommand{
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   0,
		BaseInstance:  1,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   6,
		BaseInstance:  1,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   12,
		BaseInstance:  2,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   18,
		BaseInstance:  3,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   24,
		BaseInstance:  4,
	},
	{
		VertexCount:   6,
		InstanceCount: 1,
		FirstVertex:   30,
		BaseInstance:  5,
	},
}

// Instance Buffer
var draws = []struct {
	color color.LRGB `layout:"1" divisor:"1"`
}{
	{color.LRGB{R: 0.2, G: 0, B: 0.6}},
	{color.LRGB{R: 0.2, G: 0, B: 0.6}},
	{color.LRGB{R: 0, G: 0.3, B: 0.1}},
	{color.LRGB{R: 0, G: 0.3, B: 0.1}},
	{color.LRGB{R: 0.8, G: 0.3, B: 0}},
	{color.LRGB{R: 0.8, G: 0.3, B: 0}},
}

// Vertex buffer
type simplemesh []struct {
	position coord.XYZ `layout:"0"`
}

// Initialization //////////////////////////////////////////////////////////////

func Example_indirectDraw() {
	defer cozely.Recover()

	cozely.Configure(cozely.Multisample(8))
	l := loop07{}
	cozely.Events.Resize = func() {
		s := cozely.WindowSize()
		gl.Viewport(0, 0, int32(s.C), int32(s.R))
		r := float32(s.C) / float32(s.R)
		l.screenFromView = space.Perspective(math32.Pi/4, r, 0.001, 1000.0)
	}
	err := cozely.Run(&l)
	if err != nil {
		panic(err)
	}
	//Output:
}

func (l *loop07) Enter() {
	input.Bind(bindings)
	context.Activate(1)

	// Create and configure the pipeline
	l.pipeline = gl.NewPipeline(
		gl.Shader(cozely.Path()+"shader07.vert"),
		gl.Shader(cozely.Path()+"shader07.frag"),
		gl.VertexFormat(0, simplemesh{}),
		gl.VertexFormat(1, draws),
		gl.Topology(gl.Triangles),
		gl.CullFace(false, true),
		gl.DepthTest(true),
	)
	gl.Enable(gl.FramebufferSRGB)
	//TODO: bug related to depth or face culling when run in test sequence

	// Create the uniform buffer
	l.perFrameUBO = gl.NewUniformBuffer(&perObject{}, gl.DynamicStorage)

	// Create the Indirect Command Buffer
	icbo := gl.NewIndirectBuffer(commands, gl.DynamicStorage)
	ibo := gl.NewVertexBuffer(draws, gl.DynamicStorage)

	// Create and fill the vertex buffer
	vbo := gl.NewVertexBuffer(simplecube(), 0)

	// Initialize worldFromObject and viewFromWorld matrices
	l.position = coord.XYZ{0, 0, 0}
	l.yaw = -0.6
	l.pitch = 0.3
	l.computeWorldFromObject()
	l.computeViewFromWorld()

	// Bind the vertex buffer to the pipeline
	l.pipeline.Bind()
	vbo.Bind(0, 0)
	icbo.Bind()
	ibo.Bind(1, 0)
	l.pipeline.Unbind()
}

func (loop07) Leave() {
}

// Game Loop ///////////////////////////////////////////////////////////////////

func (loop07) Update() {
}

func (l *loop07) React() {
	m := input.Cursor.Delta().XY()
	s := input.Cursor.Position().XY()

	if rotate.JustPressed(1) || move.JustPressed(1) || zoom.JustPressed(1) {
		input.Cursor.Hide()
	}
	if rotate.JustReleased(1) || move.JustReleased(1) || zoom.JustReleased(1) {
		input.Cursor.Show()
	}

	if rotate.Pressed(1) {
		l.yaw += 4 * m.X / s.X
		l.pitch += 4 * m.Y / s.Y
		switch {
		case l.pitch < -math32.Pi/2:
			l.pitch = -math32.Pi / 2
		case l.pitch > +math32.Pi/2:
			l.pitch = +math32.Pi / 2
		}
		l.computeWorldFromObject()
	}

	if move.Pressed(1) {
		d := m.Times(2).Slashxy(s)
		l.position.X += d.X
		l.position.Y -= d.Y
		l.computeWorldFromObject()
	}

	if zoom.Pressed(1) {
		d := m.Times(2).Slashxy(s)
		l.position.X += d.X
		l.position.Z += d.Y
		l.computeWorldFromObject()
	}

	if quit.JustPressed(1) {
		cozely.Stop(nil)
	}
}

func (l *loop07) Render() {
	l.pipeline.Bind()
	gl.ClearDepthBuffer(1.0)
	gl.ClearColorBuffer(color.LRGBA{0.9, 0.9, 0.9, 1.0})

	u := perObject{
		screenFromObject: l.screenFromView.
			Times(l.viewFromWorld).
			Times(l.worldFromObject),
	}
	l.perFrameUBO.SubData(&u, 0)
	l.perFrameUBO.Bind(0)

	gl.DrawIndirect(0, 6)

	l.pipeline.Unbind()
}

func (l *loop07) computeWorldFromObject() {
	rot := space.EulerZXY(l.pitch, l.yaw, 0)
	l.worldFromObject = space.Translation(l.position).Times(rot)
}

func (l *loop07) computeViewFromWorld() {
	l.viewFromWorld = space.LookAt(
		coord.XYZ{0, 0, 3},
		coord.XYZ{0, 0, 0},
		coord.XYZ{0, 1, 0},
	)
}

////////////////////////////////////////////////////////////////////////////////

func simplecube() simplemesh {
	return simplemesh{
		// Front Face
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		// Back Face
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, -0.5}},
		// Right Face
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{+0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
		// Left Face
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		// Bottom Face
		{coord.XYZ{-0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		{coord.XYZ{-0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, -0.5}},
		{coord.XYZ{+0.5, -0.5, +0.5}},
		// Top Face
		{coord.XYZ{-0.5, +0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{-0.5, +0.5, -0.5}},
		{coord.XYZ{+0.5, +0.5, +0.5}},
		{coord.XYZ{+0.5, +0.5, -0.5}},
	}
}
