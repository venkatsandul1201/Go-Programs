package main

import (
	"fmt"
	"strings"
)

func main() {
	P := fmt.Println
	P("Contains:", strings.Contains("test", "es"))
	P("Count:", strings.Count("test", "t"))
	P("HasPrefix:", strings.HasPrefix("test", "te"))
	P("HasSuffix:", strings.HasSuffix("test", "st"))
	P("Index:", strings.Index("test", "e"))
	P("Join:", strings.Join([]string{"a", "b"}, "-"))
	P("Repeat:", strings.Repeat("a", 5))
	P("Split:", strings.Split("a-b-c-d-e", "-"))
	P("ToLower:", strings.ToLower("TEST"))
	P("ToUpper:", strings.ToUpper("test"))
	P("Len:", len("Hello World"))
	P("Char:", "Hello"[1])
}
