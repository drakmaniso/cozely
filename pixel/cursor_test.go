// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package pixel_test

import (
	"testing"

	"github.com/drakmaniso/glam"
	"github.com/drakmaniso/glam/palette"
	"github.com/drakmaniso/glam/pixel"
)

//------------------------------------------------------------------------------

var curScreen = pixel.NewCanvas(pixel.Zoom(2))

var cursor = pixel.NewCursor()

func init() {
	cursor.Canvas(curScreen)
}

var curBg, curFg palette.Index

//------------------------------------------------------------------------------

func TestCursor_print(t *testing.T) {
	do(func() {
		err := glam.Run(curLoop{})
		if err != nil {
			t.Error(err)
		}
	})
}

//------------------------------------------------------------------------------

type curLoop struct {
	glam.Handlers
}

//------------------------------------------------------------------------------

func (curLoop) Enter() error {
	palette.Load("C64")
	curBg = palette.Find("white")
	curFg = palette.Find("black")
	cursor.ColorShift(curFg - 1)
	return nil
}

//------------------------------------------------------------------------------

func (curLoop) Update() error {
	return nil
}

//------------------------------------------------------------------------------

func (curLoop) Draw() error {
	curScreen.Clear(curBg)

	cursor.Locate(16, 8, 0)

	cursor.Font(pixel.Font(0))
	cursor.Println("a quick brown fox \"jumps\" over the (lazy) dog.")
	cursor.Println("A QUICK BROWN FOX \"JUMPS\" OVER THE (LAZY) DOG.")
	cursor.Println("0123456789!@#$^&*()-+=_~[]{}|\\;:'\",.<>/?%")
	cursor.Println("12+34 56-7.8 90*13 24/35 -5 +2 3*(2+5) 4<5 6>2 2=1+1 *f := &x;")
	cursor.Println()

	cursor.Font(pixel.Font(0))
	cursor.Locate(16, 100, 0)
	cursor.Write([]byte("Foo"))
	cursor.Move(1, 3, 0)
	cursor.WriteRune('B')
	cursor.Move(2, 2, 0)
	cursor.WriteRune('a')
	cursor.Move(3, 1, 0)
	cursor.WriteRune('r')
	cursor.MoveTo(32, 132, 0)
	cursor.Flush()
	cursor.Write([]byte("Boo\n"))
	cursor.Write([]byte("Choo"))
	cursor.Flush()

	cursor.Locate(16, 200, 0)
	cursor.Font(tinela9)
	cursor.Print("Tinela")
	cursor.Font(simpela10)
	cursor.Print("Simpela10")
	cursor.Font(simpela12)
	cursor.Print("Simpela12")
	cursor.Font(cozela10)
	cursor.Print("Cozela10")
	cursor.Font(cozela12)
	cursor.Print("Cozela12")
	cursor.Font(chaotela12)
	cursor.Print("Chaotela12")

	cursor.Locate(curScreen.Size().X-200, 9, 0)
	cursor.Font(pixel.Font(0))
	cursor.Printf("Position x=%d, y=%d\n", curScreen.Mouse().X, curScreen.Mouse().Y)

	curScreen.Display()
	return nil
}

//------------------------------------------------------------------------------