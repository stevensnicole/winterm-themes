// Taken from https://gitlab.com/hackebrot/palette
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	plist "github.com/DHowett/go-plist"
	"github.com/fatih/structs"
	"github.com/tomnomnom/xtermcolor"
)

// itermpalette representing a decoded .itermcolors file
type itermpalette struct {
	Ansi0        itermcolor `plist:"Ansi 0 Color"`
	Ansi1        itermcolor `plist:"Ansi 1 Color"`
	Ansi2        itermcolor `plist:"Ansi 2 Color"`
	Ansi3        itermcolor `plist:"Ansi 3 Color"`
	Ansi4        itermcolor `plist:"Ansi 4 Color"`
	Ansi5        itermcolor `plist:"Ansi 5 Color"`
	Ansi6        itermcolor `plist:"Ansi 6 Color"`
	Ansi7        itermcolor `plist:"Ansi 7 Color"`
	Ansi8        itermcolor `plist:"Ansi 8 Color"`
	Ansi9        itermcolor `plist:"Ansi 9 Color"`
	Ansi10       itermcolor `plist:"Ansi 10 Color"`
	Ansi11       itermcolor `plist:"Ansi 11 Color"`
	Ansi12       itermcolor `plist:"Ansi 12 Color"`
	Ansi13       itermcolor `plist:"Ansi 13 Color"`
	Ansi14       itermcolor `plist:"Ansi 14 Color"`
	Ansi15       itermcolor `plist:"Ansi 15 Color"`
	Background   itermcolor `plist:"Background Color"`
	Badge        itermcolor `plist:"Badge Color"`
	Bold         itermcolor `plist:"Bold Color"`
	Cursor       itermcolor `plist:"Cursor Color"`
	CursorGuide  itermcolor `plist:"Cursor Guide Color"`
	CursorText   itermcolor `plist:"Cursor Text Color"`
	Foreground   itermcolor `plist:"Foreground Color"`
	Link         itermcolor `plist:"Link Color"`
	SelectedText itermcolor `plist:"Selected Text Color"`
	Selection    itermcolor `plist:"Selection Color"`
	Tab          itermcolor `plist:"Tab Color"`
}

// itermcolor representing a single key value entry
// of a decoded .itermcolors file with its components
type itermcolor struct {
	Red   float64 `plist:"Red Component"`
	Green float64 `plist:"Green Component"`
	Blue  float64 `plist:"Blue Component"`
	Alpha float64 `plist:"Alpha Component"`
}

// HEX translates an itermcolor into a hex string
// representation of the color, such as #98bb99
func (i itermcolor) HEX() string {
	return fmt.Sprintf(
		"#%02x%02x%02x",
		uint8(i.Red*255.0+0.5),
		uint8(i.Green*255.0+0.5),
		uint8(i.Blue*255.0+0.5),
	)
}

// RGB255 is a convenience function to translate
// an itermcolor with values from 0 to 255
func (i itermcolor) RGB255() (r, g, b uint8) {
	r = uint8(i.Red*255.0 + 0.5)
	g = uint8(i.Green*255.0 + 0.5)
	b = uint8(i.Blue*255.0 + 0.5)
	return
}

// RGBA implements the color.Color interface on itermcolor
// so we can use the xtermcolor.FromColor function
func (i itermcolor) RGBA() (r, g, b, a uint32) {
	r = uint32(i.Red * 0xFFFF)
	g = uint32(i.Green * 0xFFFF)
	b = uint32(i.Blue * 0xFFFF)
	a = 0xFFFF
	return
}

// Xterm wraps the xtermcolor function that returns
// the according xterm color code for a itermcolor
func (i itermcolor) Xterm() uint8 {
	return xtermcolor.FromColor(i)
}

func main() {
	colorsfile := flag.String("colors", "", "a itermcolors file to be loaded")
	flag.Parse()

	if *colorsfile == "" {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", "colors file is required")
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(*colorsfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	var itp itermpalette
	decoder := plist.NewDecoder(bytes.NewReader(dat))

	if err := decoder.Decode(&itp); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	s := structs.New(itp)
	for _, f := range s.Fields() {
		t := f.Tag("plist")
		c := f.Value().(itermcolor)
		r, g, b := c.RGB255()

		fmt.Printf("%-24v", t)
		fmt.Printf("%-20v", fmt.Sprintf("Hex RGB %v", c.HEX()))
		fmt.Printf("%-20v", fmt.Sprintf("RGB %3v %3v %3v", r, g, b))
		fmt.Printf("%-20v", fmt.Sprintf("Xterm code %3v", c.Xterm()))
		fmt.Println("")
	}
}
