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
var NAMED_COLOR = map[string] string {
    "aliceblue":            "#f0f8ff",
    "antiquewhite":         "#faebd7",
    "aqua":                 "#00ffff",
    "aquamarine":           "#7fffd4",
    "azure":                "#f0ffff",
    "beige":                "#f5f5dc",
    "bisque":               "#ffe4c4",
    "black":                "#000000",
    "blanchedalmond":       "#ffebcd",
    "blue":                 "#0000ff",
    "blueviolet":           "#8a2be2",
    "brown":                "#a52a2a",
    "burlywood":            "#deb887",
    "cadetblue":            "#5f9ea0",
    "chartreuse":           "#7fff00",
    "chocolate":            "#d2691e",
    "coral":                "#ff7f50",
    "cornflowerblue":       "#6495ed",
    "cornsilk":             "#fff8dc",
    "crimson":              "#dc143c",
    "cyan":                 "#00ffff",
    "darkblue":             "#00008b",
    "darkcyan":             "#008b8b",
    "darkgoldenrod":        "#b8860b",
    "darkgray":             "#a9a9a9",
    "darkgrey":             "#a9a9a9",
    "darkgreen":            "#006400",
    "darkkhaki":            "#bdb76b",
    "darkmagenta":          "#8b008b",
    "darkolivegreen":       "#556b2f",
    "darkorange":           "#ff8c00",
    "darkorchid":           "#9932cc",
    "darkred":              "#8b0000",
    "darksalmon":           "#e9967a",
    "darkseagreen":         "#8fbc8f",
    "darkslateblue":        "#483d8b",
    "darkslategray":        "#2f4f4f",
    "darkslategrey":        "#2f4f4f",
    "darkturquoise":        "#00ced1",
    "darkviolet":           "#9400d3",
    "deeppink":             "#ff1493",
    "deepskyblue":          "#00bfff",
    "dimgray":              "#696969",
    "dimgrey":              "#696969",
    "dodgerblue":           "#1e90ff",
    "firebrick":            "#b22222",
    "floralwhite":          "#fffaf0",
    "forestgreen":          "#228b22",
    "fuchsia":              "#ff00ff",
    "gainsboro":            "#dcdcdc",
    "ghostwhite":           "#f8f8ff",
    "gold":                 "#ffd700",
    "goldenrod":            "#daa520",
    "gray":                 "#808080",
    "grey":                 "#808080",
    "green":                "#008000",
    "greenyellow":          "#adff2f",
    "honeydew":             "#f0fff0",
    "hotpink":              "#ff69b4",
    "indianred":            "#cd5c5c",
    "indigo":               "#4b0082",
    "ivory":                "#fffff0",
    "khaki":                "#f0e68c",
    "lavender":             "#e6e6fa",
    "lavenderblush":        "#fff0f5",
    "lawngreen":            "#7cfc00",
    "lemonchiffon":         "#fffacd",
    "lightblue":            "#add8e6",
    "lightcoral":           "#f08080",
    "lightcyan":            "#e0ffff",
    "lightgoldenrodyellow": "#fafad2",
    "lightgreen":           "#90ee90",
    "lightgray":            "#d3d3d3",
    "lightgrey":            "#d3d3d3",
    "lightpink":            "#ffb6c1",
    "lightsalmon":          "#ffa07a",
    "lightseagreen":        "#20b2aa",
    "lightskyblue":         "#87cefa",
    "lightslategray":       "#778899",
    "lightslategrey":       "#778899",
    "lightsteelblue":       "#b0c4de",
    "lightyellow":          "#ffffe0",
    "lime":                 "#00ff00",
    "limegreen":            "#32cd32",
    "linen":                "#faf0e6",
    "magenta":              "#ff00ff",
    "maroon":               "#800000",
    "mediumaquamarine":     "#66cdaa",
    "mediumblue":           "#0000cd",
    "mediumorchid":         "#ba55d3",
    "mediumpurple":         "#9370db",
    "mediumseagreen":       "#3cb371",
    "mediumslateblue":      "#7b68ee",
    "mediumspringgreen":    "#00fa9a",
    "mediumturquoise":      "#48d1cc",
    "mediumvioletred":      "#c71585",
    "midnightblue":         "#191970",
    "mintcream":            "#f5fffa",
    "mistyrose":            "#ffe4e1",
    "moccasin":             "#ffe4b5",
    "navajowhite":          "#ffdead",
    "navy":                 "#000080",
    "oldlace":              "#fdf5e6",
    "olive":                "#808000",
    "olivedrab":            "#6b8e23",
    "orange":               "#ffa500",
    "orangered":            "#ff4500",
    "orchid":               "#da70d6",
    "palegoldenrod":        "#eee8aa",
    "palegreen":            "#98fb98",
    "paleturquoise":        "#afeeee",
    "palevioletred":        "#db7093",
    "papayawhip":           "#ffefd5",
    "peachpuff":            "#ffdab9",
    "peru":                 "#cd853f",
    "pink":                 "#ffc0cb",
    "plum":                 "#dda0dd",
    "powderblue":           "#b0e0e6",
    "purple":               "#800080",
    "red":                  "#ff0000",
    "rosybrown":            "#bc8f8f",
    "royalblue":            "#4169e1",
    "saddlebrown":          "#8b4513",
    "salmon":               "#fa8072",
    "sandybrown":           "#f4a460",
    "seagreen":             "#2e8b57",
    "seashell":             "#fff5ee",
    "sienna":               "#a0522d",
    "silver":               "#c0c0c0",
    "skyblue":              "#87ceeb",
    "slateblue":            "#6a5acd",
    "slategray":            "#708090",
    "slategrey":            "#708090",
    "snow":                 "#fffafa",
    "springgreen":          "#00ff7f",
    "steelblue":            "#4682b4",
    "tan":                  "#d2b48c",
    "teal":                 "#008080",
    "thistle":              "#d8bfd8",
    "tomato":               "#ff6347",
    "turquoise":            "#40e0d0",
    "violet":               "#ee82ee",
    "wheat":                "#f5deb3",
    "white":                "#ffffff",
    "whitesmoke":           "#f5f5f5",
    "yellow":               "#ffff00",
    "yellowgreen":          "#9acd32",
}


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
    var smallest_distance int64 = 10000000000
    for c := 16; c < 256; c++ {
        d := int(math.Pow(float64(COLOR_TABLE[c][0] - r), 2) +
                 math.Pow(float64(COLOR_TABLE[c][1] - g), 2) +
                 math.Pow(float64(COLOR_TABLE[c][2] - b), 2))
        if int64(d) < smallest_distance {
            smallest_distance = int64(d)
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
    } else if color[0] == '#' && len(color) == 7 {
        tr := fmt.Sprintf("%c%c", color[1], color[2])
        r, _ = strconv.ParseInt(tr, 16, 0)
        tr = fmt.Sprintf("%c%c", color[3], color[4])
        g, _ = strconv.ParseInt(tr, 16, 0)
        tr = fmt.Sprintf("%c%c", color[5], color[6])
        b, _ = strconv.ParseInt(tr, 16, 0)
    } else if "" != NAMED_COLOR[color] {
        htmlColor := NAMED_COLOR[color]
        tr := fmt.Sprintf("%c%c", htmlColor[1], htmlColor[2])
        r, _ = strconv.ParseInt(tr, 16, 0)
        tr = fmt.Sprintf("%c%c", htmlColor[3], htmlColor[4])
        g, _ = strconv.ParseInt(tr, 16, 0)
        tr = fmt.Sprintf("%c%c", htmlColor[5], htmlColor[6])
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

func Bold(text string) (boldText string) {
    boldText = esc(1) + text + esc(22)
    return boldText
}

func Italic(text string) (italicText string) {
    italicText = esc(3) + text + esc(23)
    return italicText
}

func Underline(text string) (underlineText string) {
    underlineText = esc(4) + text + esc(24)
    return underlineText
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
