package nanairo

import (
    "fmt"
    "testing"
)

func TestFgColor(t *testing.T) {
    fmt.Println(FgColor("#f93", "Hello World"))
}

func TestFgColorHtml6(t *testing.T) {
    fmt.Println(FgColor("#f19e30", "Hello World"))
}

func TestFgColorHtmlName(t *testing.T) {
    fmt.Println(FgColor("limegreen", "Hello World"))
}

func TestBgColor(t *testing.T) {
    fmt.Println(BgColor("#f93", "Hello World"))
}

func TestBold(t *testing.T) {
    fmt.Println(Bold(FgColor("#fff", "Hello World")))
}

func TestItalic(t *testing.T) {
    fmt.Println(Italic("Hello World"))
}

func TestUnderline(t *testing.T) {
    fmt.Println(Underline("Hello World"))
}

func TestHighlight(t *testing.T) {
    fmt.Println(Highlight("#39e", "Hello World"))
}
