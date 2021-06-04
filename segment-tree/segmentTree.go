package segment_tree


type SegmentTree struct {
	Arr  []int
	Tree []int
}

func (c* SegmentTree)BuildTree(v, tl, tr int) {
	if tl == tr {
		c.Tree[v] = c.Arr[tl];
	} else {
		tm := (tl + tr) / 2
		c.BuildTree(v * 2, tl, tm)
		c.BuildTree(v * 2 + 1, tm + 1, tr)
		c.Tree[v] = c.Tree[v * 2] + c.Tree[v * 2 + 1]
	}
}

func (c *SegmentTree) GetSum(l, r, v, tl, tr int) int {
	if l <= tl && tr <= r {
		return c.Tree[v]
	}
	if tr < l || r < tl {
		return 0
	}
	tm := (tl + tr) / 2
	return c.GetSum(l, r, v * 2, tl, tm) + c.GetSum(l, r, v * 2 + 1, tm + 1, tr)
 }