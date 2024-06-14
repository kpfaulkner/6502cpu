package pkg

type RAM struct {
	ram [64 * 1024]uint8
}

func NewRAM() *RAM {
	return &RAM{}
}
