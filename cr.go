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

func PBlack(txt string) string {
	return P(txt, BLACK, nil)
}

func PRed(txt string) string {
	return P(txt, RED, nil)
}

func PGreen(txt string) string {
	return P(txt, GREEN, nil)
}

func PYellow(txt string) string {
	return P(txt, YELLOW, nil)
}

func PBlue(txt string) string {
	return P(txt, BLUE, nil)
}

func PMagenta(txt string) string {
	return P(txt, MAGENTA, nil)
}

func PCyan(txt string) string {
	return P(txt, CYAN, nil)
}

func PWhite(txt string) string {
	return P(txt, WHITE, nil)
}

func PLBalck(txt string) string {
	return P(txt, BLACK.L(), nil)
}

func PLRed(txt string) string {
	return P(txt, RED.L(), nil)
}

func PLGreen(txt string) string {
	return P(txt, GREEN.L(), nil)
}

func PLYellow(txt string) string {
	return P(txt, YELLOW.L(), nil)
}

func PLBlue(txt string) string {
	return P(txt, BLUE.L(), nil)
}

func PLMagenta(txt string) string {
	return P(txt, MAGENTA.L(), nil)
}

func PLCyan(txt string) string {
	return P(txt, CYAN.L(), nil)
}

func PLWhite(txt string) string {
	return P(txt, WHITE.L(), nil)
}
