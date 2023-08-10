package cr

import (
	"fmt"
	"testing"
)

func TestCR(t *testing.T) {
	text := "Hello, world!"

	colors := []COLOR{BLACK, RED, GREEN, YELLOW, BLUE, MAGENTA, CYAN, WHITE}
	// styles := []STYLE{NORMAL, BOLD, FAINT, ITALIC, UNDERLINE, BLINK}

	fmt.Println("Single style with foreground color:")
	for _, fg := range colors {
		fmt.Println(P(text, fg, nil, BOLD))
	}

	fmt.Println("\nMulti style with background color(foreground color is GREEN):")
	for _, bg := range colors {
		fmt.Println(P(text, GREEN, &bg, ITALIC, BOLD, UNDERLINE))
	}

	var testFunc = map[string]func(s string) string{
		"PBlack":    PBlack,
		"PRed":      PRed,
		"PGreen":    PGreen,
		"PYellow":   PYellow,
		"PBlue":     PBlue,
		"PMagenta":  PMagenta,
		"PWhite":    PWhite,
		"PLBlack":   PLBalck,
		"PLRed":     PLRed,
		"PLGreen":   PLGreen,
		"PLYellow":  PYellow,
		"PLBlue":    PLBlue,
		"PLMagenta": PLMagenta,
		"PLWhite":   PLWhite,
	}

	funcKeys := []string{"PBlack", "PRed", "PGreen", "PYellow", "PBlue", "PMagenta", "PWhite", "PLBlack", "PLRed", "PLGreen", "PLYellow", "PLBlue", "PLMagenta", "PLWhite"}

	fmt.Println("\nSimple functions:")
	for _, key := range funcKeys {
		fmt.Println(key, testFunc[key](text))
	}
}
