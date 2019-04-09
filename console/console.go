package console

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"log"
	"time"
)

const (
	MaxRenderInterval = 1000 * time.Millisecond
	MinRenderInterval = 100 * time.Millisecond
	AppTitle          = "sampler"
	AppVersion        = "0.9.0"
)

const (
	BellCharacter = "\a"
)

type AsciiFont string

const (
	AsciiFontFlat AsciiFont = "flat"
	AsciiFont3D   AsciiFont = "3d"
)

func Init() {

	fmt.Printf("\033]0;%s\007", AppTitle)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
}

func Close() {
	ui.Close()
}
