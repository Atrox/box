package box

import "testing"

// TODO: Add tests...
func TestPrint(*testing.T) {
	box := New()

	box.Print("First line", "Second line", "Third line")
	box.Println("First line", "Second line", "Third line")
}
