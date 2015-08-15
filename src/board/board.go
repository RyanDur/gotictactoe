package board

type Board interface {
	Set(piece string, x, y int)
	Get(x, y int) string
	Contents() [][]string
	Entries() int
	Size() int
	IsFull() bool
}

type gameGrid struct {
	size   int
	spaces []string
}

func NewBoard(size int) Board {
	board := new(gameGrid)
	board.size = size
	board.spaces = make([]string, size*size)
	return board
}

func (b *gameGrid) Set(piece string, x, y int) {
	b.spaces[b.index(x, y)] = piece
}

func (b *gameGrid) Get(x, y int) string {
	return b.spaces[b.index(x, y)]
}

func (b *gameGrid) Contents() [][]string {
	size := b.size
	spaces := make([]string, size*size)
	copy(spaces, b.spaces)
	result := make([][]string, size)
	for i := 0; i < size; i++ {
		result[i] = spaces[(i * size):((i * size) + (size))]
	}
	return result
}

func (b *gameGrid) Entries() int {
	result := 0
	for _, entry := range b.spaces {
		if entry != "" {
			result++
		}
	}
	return result
}

func (b *gameGrid) Size() int {
	return b.size
}

func (b *gameGrid) IsFull() bool {
	return b.Entries() == b.Size()*b.Size()
}

func (b *gameGrid) index(x, y int) int {
	return (x * b.size) + y
}
