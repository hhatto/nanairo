package nanairo

import (
    "fmt"
    "testing"
)

func TestFgColor(t *testing.T) {
    fmt.Println(FgColor("#f93", "Hello World"))
}

func TestBgColor(t *testing.T) {
    fmt.Println(BgColor("#f93", "Hello World"))
}

func TestHighlight(t *testing.T) {
    fmt.Println(Highlight("#39e", "Hello World"))
}
