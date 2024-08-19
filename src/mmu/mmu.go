package mmu

type MMU struct {
	Memory [0xFFFF]byte
}

func (mmu *MMU) Init() {}

func (mmu *MMU) Tick() {}

func NewMMU() *MMU {
	return &MMU{}
}