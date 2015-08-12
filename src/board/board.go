package board

type Board struct {
	size   int
	spaces []string
}

func NewBoard(size int) *Board {
	board := new(Board)
	board.size = size
	board.spaces = make([]string, size*size)
	return board
}

func (b *Board) Set(piece string, x, y int) {
	b.spaces[b.index(x, y)] = piece
}

func (b *Board) Get(x, y int) string {
	return b.spaces[b.index(x, y)]
}

func (b *Board) Contents() [][]string {
	size := b.size
	spaces := make([]string, size*size)
	copy(spaces, b.spaces)
	result := make([][]string, size)
	for i := 0; i < size; i++ {
		result[i] = spaces[(i * size):((i * size) + (size))]
	}
	return result
}

func (b *Board) index(x, y int) int {
	return (x * b.size) + y
}
