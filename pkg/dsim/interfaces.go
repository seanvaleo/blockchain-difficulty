package dsim

type (
	// Blockchain is an interface for blockchains
	Blockchain interface {
		Name() string
		Algorithm() Algorithm
		Length() uint64
		AddBlock(uint64)
		GetBlock(uint64) *Block
	}
	// Algorithm is an interface for difficulty algorithms
	Algorithm interface {
		Name() string
		Window() uint64
		NextDifficulty([]*Block) uint64
	}
)
