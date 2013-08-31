package nanairo

import (
    "fmt"
    "math"
    "strconv"
)


var CUBE_STEPS = [6]int{0x00, 0x5F, 0x87, 0xAF, 0xD7, 0xFF}
var BASIC16 = [][]int{
        {0, 0, 0}, {205, 0, 0}, {0, 205, 0}, {205, 205, 0},
        {0, 0, 238}, {205, 0, 205}, {0, 205, 205}, {229, 229, 229},
        {127, 127, 127}, {255, 0, 0}, {0, 255, 0}, {255, 255, 0},
        {92, 92, 255}, {255, 0, 255}, {0, 255, 255}, {255, 255, 255}}
var COLOR_TABLE = make([][]int, 256, 256)

func init() {
    for i := 0; i < 256; i++ {
        COLOR_TABLE[i] = make([]int, 3, 3)
        r, g, b := xterm2rgb(i)
        COLOR_TABLE[i][0] = r
        COLOR_TABLE[i][1] = g
        COLOR_TABLE[i][2] = b
    }
}

func xterm2rgb(xcolor int) (int, int, int) {
    if xcolor < 16 {
        return BASIC16[xcolor][0], BASIC16[xcolor][1], BASIC16[xcolor][2]
    } else if 16 <= xcolor && xcolor <= 231 {
        xcolor -= 16
        return CUBE_STEPS[(xcolor / 36) % 6], CUBE_STEPS[(xcolor / 6) % 6], CUBE_STEPS[xcolor % 6]
    } else if 232 <= xcolor && xcolor <= 255 {
        c := 8 + (xcolor - 232) * 0x0A
        return c, c, c
    }

    return BASIC16[xcolor][0], BASIC16[xcolor][1], BASIC16[xcolor][2]
}

func rgb2xterm(r int, g int, b int) int {
    if r < 5 && g < 5 && b < 5 {
        return 16
    }
    best_match := 0
    smallest_distance := 10000000000
    for c := 16; c < 256; c++ {
        d := int(math.Pow(float64(COLOR_TABLE[c][0] - r), 2) +
                 math.Pow(float64(COLOR_TABLE[c][1] - g), 2) +
                 math.Pow(float64(COLOR_TABLE[c][2] - b), 2))
        if d < smallest_distance {
            smallest_distance = d
            best_match = c
        }
    }
    return best_match
}

func parseColor(color string) (int, int, int) {
    r, g, b := int64(0), int64(0), int64(0)
    if color[0] == '#' && len(color) == 4 {
        tr := fmt.Sprintf("%c%c", color[1], color[1])
        r, _ = strconv.ParseInt(tr, 16, 0)
        tr = fmt.Sprintf("%c%c", color[2], color[2])
        g, _ = strconv.ParseInt(tr, 16, 0)
        tr = fmt.Sprintf("%c%c", color[3], color[3])
        b, _ = strconv.ParseInt(tr, 16, 0)
    }
    return int(r), int(g), int(b)
}

func esc(vs ...interface{}) (string) {
    r := ""
    for _, v := range vs {
        r += fmt.Sprintf(";%d", v)
    }
    return fmt.Sprintf("\x1b[%sm", r)
}

func FgColor(color string, text string) (coloredText string) {
    r, g, b := parseColor(color)
    c := rgb2xterm(r, g, b)
    coloredText = esc(38, 5, c) + text + esc(39)
    return coloredText
}

func BgColor(color string, text string) (coloredText string) {
    r, g, b := parseColor(color)
    c := rgb2xterm(r, g, b)
    coloredText = esc(48, 5, c) + text + esc(49)
    return coloredText
}

func Highlight(color string, text string) (coloredText string) {
    r, g, b := parseColor(color)
    c := rgb2xterm(r, g, b)
    coloredText = esc(1, 38, 5, c, 48, 5, c) + text + esc(49, 38, 22)
    return coloredText
}
