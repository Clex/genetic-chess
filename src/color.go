package geneticchess

type Color bool

const (
	White Color = true
	Black Color = false
)

func (c Color) String() string {
	if c == White {
		return "white"
	}

	return "black"
}

func (c *Color) other() Color {
	if *c == White {
		return Black
	}
	return White
}

func (c *Color) swap() {
	*c = !*c
}

func (c *Color) score() float64 {
	if *c == White {
		return 1.0
	}

	return -1.0
}
