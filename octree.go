package gorgb

const (
	treeSize  = 19173961
	numColors = 1 << 24
)

var lookup = [][]int{
	{0, 1, 4, 5, 2, 3, 6, 7},
	{1, 0, 5, 4, 3, 2, 7, 6},
	{2, 3, 6, 7, 0, 1, 4, 5},
	{3, 2, 7, 6, 1, 0, 5, 4},
	{4, 5, 0, 1, 6, 7, 2, 3},
	{5, 4, 1, 0, 7, 6, 3, 2},
	{6, 7, 2, 3, 4, 5, 0, 1},
	{7, 6, 3, 2, 5, 4, 1, 0},
}

type Octree struct {
	data []int
}

func NewOctree() *Octree {
	t := Octree{}
	t.data = make([]int, treeSize)
	t.initialize(0, numColors)
	return &t
}

func (t *Octree) initialize(index, value int) {
	t.data[index] = value
	if value == 1 {
		return
	}
	value /= 8
	for i := 0; i < 8; i++ {
		t.initialize(8*index+1+i, value)
	}
}

func (t *Octree) Pop(r, g, b int) (int, int, int) {
	var path [8]int
	for i := 7; i >= 0; i-- {
		br := r & 1
		bg := g & 1
		bb := b & 1
		path[i] = (br << 2) | (bg << 1) | (bb << 0)
		r >>= 1
		g >>= 1
		b >>= 1
	}
	var index int
	for i := 0; i < 8; i++ {
		base := 8*index + 1
		for j := 0; j < 8; j++ {
			child := lookup[path[i]][j]
			new_index := base + child
			if t.data[new_index] > 0 {
				path[i] = child
				t.data[new_index]--
				index = new_index
				break
			}
		}
	}
	r = 0
	g = 0
	b = 0
	for i := 0; i < 8; i++ {
		r <<= 1
		g <<= 1
		b <<= 1
		r |= (path[i] >> 2) & 1
		g |= (path[i] >> 1) & 1
		b |= (path[i] >> 0) & 1
	}
	return r, g, b
}
