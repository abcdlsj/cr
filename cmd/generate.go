//go:generate go run generate.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	crFile, err := os.Create("../cr.go")
	if err != nil {
		panic(err)
	}
	defer crFile.Close()
	crFile.WriteString("// DO NOT EDIT THIS FILE\n\n")
	crFile.WriteString("package cr\n\n")

	crTestFile, err := os.Create("../cr_test.go")
	if err != nil {
		panic(err)
	}
	defer crTestFile.Close()

	crTestFile.WriteString("// DO NOT EDIT THIS FILE\n\n")
	crTestFile.WriteString("package cr\n\n")
	crTestFile.WriteString("import (\n\t\"fmt\"\n\t\"testing\"\n)\n\n")
	crTestFile.WriteString("var globalStr = \"Hello, World!\"\n\n")

	colors := []string{"BLACK", "RED", "GREEN", "YELLOW", "BLUE", "MAGENTA", "CYAN", "WHITE"}

	for _, light := range []bool{false, true} {
		for _, style := range []string{"nil", "BOLD", "FAINT", "ITALIC", "UNDERLINE", "BLINK"} {
			for _, fg := range colors {
				for _, bg := range colors {
					var bgf = bg
					if bg == fg {
						bgf = "nil"
					}
					var fgf = fg
					if light {
						fgf = "light " + fg
					}
					comment := fmt.Sprintf("// fg: %s, bg: %s, style: %v\n", strings.ToLower(fgf), strings.ToLower(bgf), strings.ToLower(style))

					crFile.WriteString(comment)
					crFile.WriteString(genCode(fg, bg, light, style) + "\n\n")

					crTestFile.WriteString(comment)
					crTestFile.WriteString(genTest(fg, bg, light, style) + "\n\n")
				}
			}
		}
	}
}

func genFuncName(fg, bg string, light bool, styles ...string) string {
	lowerExceptFirst := func(s string) string {
		return s[:1] + strings.ToLower(s[1:])
	}
	name := "P"
	if light {
		name += "L"
	}
	name += lowerExceptFirst(fg)
	if fg != bg {
		name += "Bg" + lowerExceptFirst(bg)
	}

	for _, style := range styles {
		if style == "nil" {
			return name
		}
		name += lowerExceptFirst(style)
	}

	return name
}

func genCode(fg, bg string, light bool, styles ...string) string {
	name := genFuncName(fg, bg, light, styles...)
	var fgf, bgf string

	if light {
		fgf = fg + ".L()"
	} else {
		fgf = fg
	}

	if bg != fg {
		bgf = bg + ".Pointer()"
	} else {
		bgf = "nil"
	}

	if len(styles) == 1 && styles[0] == "nil" {
		return fmt.Sprintf("func "+name+"(txt string) string {\n%s"+"}", fmt.Sprintf("\treturn P(txt, %s, %s)\n", fgf, bgf))
	}

	return fmt.Sprintf("func "+name+"(txt string) string {\n%s"+"}", fmt.Sprintf("\treturn P(txt, %s, %s, %s)\n", fgf, bgf, strings.Join(styles, ", ")))
}

func genTest(fg, bg string, light bool, styles ...string) string {
	name := genFuncName(fg, bg, light, styles...)

	return fmt.Sprintf("func Test%s(t *testing.T) {\n\tfmt.Println(%s(globalStr))\n}", name, name)
}
