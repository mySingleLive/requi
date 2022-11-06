package layout

import (
	"strings"
)

type Block struct {
	title         string
	width         int
	height        int
	paddingTop    int
	paddingRight  int
	paddingBottom int
	paddingLeft   int
	top           string
	right         string
	bottom        string
	left          string
	topLeft       string
	topRight      string
	bottomLeft    string
	bottomRight   string
}

func NewBlock() Block {
	return Block{
		title:         "",
		width:         0,
		height:        0,
		paddingTop:    0,
		paddingRight:  0,
		paddingBottom: 0,
		paddingLeft:   0,
		top:           "-",
		bottom:        "-",
		left:          "│",
		right:         "│",
		topLeft:       "┌",
		topRight:      "┐",
		bottomLeft:    "└",
		bottomRight:   "┘",
	}
}

func (b Block) Width(w int) Block {
	b.width = w
	return b
}

func (b Block) Height(h int) Block {
	b.height = h
	return b
}

func length(content string) int {
	size := len(content)
	esc := false
	opt := false
	ret := 0
	for i := 0; i < size; i++ {
		ch := content[i]
		if ch == '' {
			esc = true
		} else if esc && ch == '[' {
			esc = false
			opt = true
		} else if opt && ch == 'm' {
			esc = false
			opt = false
		} else if !esc && !opt {
			ret++
		}
	}
	return ret
}

func multiline(content string) ([]string, int) {
	var lines []string
	size := len(content)
	line := strings.Builder{}
	cw := 0
	for i := 0; i < size; i++ {
		ch := content[i]
		if ch == '\n' {
			str := line.String()
			w := length(str)
			if w > cw {
				cw = w
			}
			lines = append(lines, str)
			line.Reset()
		} else {
			line.WriteByte(ch)
		}
	}
	return lines, cw
}

func (b Block) Render(content string) string {
	lines, cw := multiline(content)
	ch := len(lines)
	str := strings.Builder{}
	rw := b.width - 2
	if rw < cw {
		rw = cw
	}
	rh := b.height - 2
	if rh < ch {
		rh = ch
	}
	str.WriteString(b.topLeft)
	for i := 0; i < rw; i++ {
		str.WriteString(b.top)
	}
	str.WriteString(b.topRight)
	str.WriteByte('\n')
	for i := 0; i < rh; i++ {
		str.WriteString(b.left)
		if i < ch {
			line := lines[i]
			lw := length(line)
			rest := rw - lw
			str.WriteString(line)
			if rest > 0 {
				for j := 0; j < rest; j++ {
					str.WriteByte(' ')
				}
			}
		}
		str.WriteString(b.right)
		str.WriteByte('\n')
	}
	str.WriteString(b.bottomLeft)
	for i := 0; i < rw; i++ {
		str.WriteString(b.bottom)
	}
	str.WriteString(b.bottomRight)
	return str.String()
}
