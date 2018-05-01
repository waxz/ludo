package main

import (
	"fmt"
	"libretro"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var keyBinds = map[glfw.Key]uint32{
	glfw.KeyX:         libretro.DeviceIDJoypadA,
	glfw.KeyZ:         libretro.DeviceIDJoypadB,
	glfw.KeyA:         libretro.DeviceIDJoypadY,
	glfw.KeyS:         libretro.DeviceIDJoypadX,
	glfw.KeyUp:        libretro.DeviceIDJoypadUp,
	glfw.KeyDown:      libretro.DeviceIDJoypadDown,
	glfw.KeyLeft:      libretro.DeviceIDJoypadLeft,
	glfw.KeyRight:     libretro.DeviceIDJoypadRight,
	glfw.KeyEnter:     libretro.DeviceIDJoypadStart,
	glfw.KeyBackspace: libretro.DeviceIDJoypadSelect,
}

var joy [libretro.DeviceIDJoypadR3 + 1]bool

func joystickCallback(joy int, event int) {
	switch event {
	case 262145:
		fmt.Printf("[Input]: Joystick #%d plugged.\n", joy)
		break
	case 262146:
		fmt.Printf("[Input]: Joystick #%d unplugged.\n", joy)
		break
	default:
		fmt.Printf("[Input]: Joystick #%d unhandled event: %d.\n", joy, event)
	}
}

func inputInit() {
	glfw.SetJoystickCallback(joystickCallback)
}

func inputPoll() {
	for k, v := range keyBinds {
		joy[v] = (window.GetKey(k) == glfw.Press)
	}

	// Close the window when the user hits the Escape key.
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

func inputState(port uint, device uint32, index uint, id uint) int16 {
	if port > 0 || index > 0 || device != libretro.DeviceJoypad {
		return 0
	}

	if id < 255 && joy[id] {
		return 1
	}
	return 0
}
