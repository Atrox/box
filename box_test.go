package box

import "testing"

// TODO: Add tests...
func TestPrint(*testing.T) {
	box := New()

	box.String("First line", "second")
	box.Print("First line", "Second line", "Third line")
	box.Println("First line", "Second line", "Third line")
}
