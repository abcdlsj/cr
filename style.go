package cr

import "strconv"

type COLOR uint32

const (
	BLACK COLOR = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

func (c COLOR) FG() COLOR {
	return c + 30
}

func (c COLOR) BG() COLOR {
	return c + 40
}

func (c COLOR) L() COLOR {
	return c + 60
}

func (c COLOR) String() string {
	return strconv.FormatUint(uint64(c), 10)
}

func (c COLOR) Pointer() *COLOR {
	return &c
}

type STYLE uint32

const (
	NORMAL STYLE = iota
	BOLD
	FAINT
	ITALIC
	UNDERLINE
	BLINK
)

func (s STYLE) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

func P(txt string, fg COLOR, bg *COLOR, styles ...STYLE) string {
	var s string
	for _, style := range styles {
		s += style.String() + ";"
	}
	if bg == nil {
		return "\033[" + s + fg.FG().String() + "m" + txt + "\033[0m"
	}
	return "\033[" + s + fg.FG().String() + ";" + bg.BG().String() + "m" + txt + "\033[0m"
}
