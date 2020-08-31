package quadtree

import "fmt"

// Bounds represents bounds of the node
type Bounds struct {
	x int
	y int
	w int
	h int
}

// QuadTree is a Quad Tree implementation
type QuadTree struct {
	maxObjects int
	maxLevels  int
	level      int
	bounds     *Bounds
	objects    []*Bounds
	nodes      []*QuadTree
}

// CreateBounds function creates Bounds object
func CreateBounds(x int, y int, w int, h int) *Bounds {
	return &Bounds{x, y, w, h}
}

// Create function creates QuadTree
func Create(bounds *Bounds, maxObjects int, maxLevels int) *QuadTree {
	return &QuadTree{maxObjects, maxLevels, 0, bounds, make([]*Bounds, 0), make([]*QuadTree, 0)}
}

// Split function split the node into 4 subnodes
func (t *QuadTree) Split() {
	nextLevel := t.level + 1
	subWidth := t.bounds.w / 2
	subHeight := t.bounds.h / 2
	x := t.bounds.x
	y := t.bounds.y

	// top right node
	t.nodes[0] = &QuadTree{
		t.maxObjects,
		t.maxLevels,
		nextLevel,
		&Bounds{x + subWidth, y, subWidth, subHeight},
		make([]*Bounds, 0),
		make([]*QuadTree, 0),
	}

	// top left node
	t.nodes[1] = &QuadTree{
		t.maxObjects,
		t.maxLevels,
		nextLevel,
		&Bounds{x, y, subWidth, subHeight},
		make([]*Bounds, 0),
		make([]*QuadTree, 0),
	}

	// bottom left node
	t.nodes[2] = &QuadTree{
		t.maxObjects,
		t.maxLevels,
		nextLevel,
		&Bounds{x, y + subHeight, subWidth, subHeight},
		make([]*Bounds, 0),
		make([]*QuadTree, 0),
	}

	// bottom right node
	t.nodes[2] = &QuadTree{
		t.maxObjects,
		t.maxLevels,
		nextLevel,
		&Bounds{x + subWidth, y + subHeight, subWidth, subHeight},
		make([]*Bounds, 0),
		make([]*QuadTree, 0),
	}
}

// GetIndex function determines which node the object belongs to
// react - bounds of the area to be checked, with x, y, width, height
// returns an array of indexes of the intersecting subnodes (0-3 = top-right, top-left, bottom-left, bottom-right / ne, nw, sw, se)
func (t *QuadTree) GetIndex(rect *Bounds) []int {
	indexes := make([]int, 0)
	verticalMidpoint := t.bounds.x + (t.bounds.w / 2)
	horizontalMidpoint := t.bounds.y + (t.bounds.h / 2)
	startIsNorth := rect.y < horizontalMidpoint
	startIsWest := rect.x < verticalMidpoint
	endIsEast := rect.x+rect.w > verticalMidpoint
	endIsSouth := rect.y+rect.h > horizontalMidpoint

	//top-right quad
	if startIsNorth && endIsEast {
		indexes = append(indexes, 0)
	}
	//top-left quad
	if startIsWest && startIsNorth {
		indexes = append(indexes, 1)
	}
	//bottom-left quad
	if startIsWest && endIsSouth {
		indexes = append(indexes, 2)
	}
	//bottom-right quad
	if endIsEast && endIsSouth {
		indexes = append(indexes, 3)
	}
	return indexes
}

// Insert function inserts the object into the node. If the node
// exceeds the capacity, it will split and add all
// objects to their corresponding subnodes
func (t *QuadTree) Insert(rect *Bounds) {
	// if we have subnodes, call insert on matching subnodes
	if len(t.nodes) != 0 {
		indexes := t.GetIndex(rect)
		for i := 0; i < len(indexes); i++ {
			t.nodes[indexes[i]].Insert(rect)
		}
		return
	}

	//otherwise, store object here
	t.objects = append(t.objects, rect)

	if len(t.objects) <= t.maxObjects || t.level >= t.maxLevels {
		return
	}

	// split if we don't already have subnodes
	if len(t.nodes) == 0 {
		t.Split()
	}

	// add all objects to their corresponding subnode
	for i := 0; i < len(t.objects); i++ {
		indexes := t.GetIndex(t.objects[i])
		for k := 0; k < len(indexes); k++ {
			t.nodes[indexes[k]].Insert(t.objects[i])
		}
	}
	t.objects = make([]*Bounds, 0)
}

// Retrieve function returns all objects that could collide with the given object
func (t *QuadTree) Retrieve(rect *Bounds) []*Bounds {
	indexes := t.GetIndex(rect)
	returnObjects := t.objects

	// if we have subnodes, retrieve their objects
	if len(t.nodes) != 0 {
		for i := 0; i < len(indexes); i++ {
			res := t.nodes[indexes[i]].Retrieve(rect)
			returnObjects = append(returnObjects, res...)
		}
	}

	// remove duplicates
	filterFunc := func(o *Bounds, idx int) bool {
		curr := findIndexOfElement(returnObjects, o)
		return curr >= idx
	}
	return filterRects(returnObjects, filterFunc)
}

func filterRects(arr []*Bounds, test func(el *Bounds, idx int) bool) (ret []*Bounds) {
	for idx, o := range arr {
		if test(o, idx) {
			ret = append(ret, o)
		}
	}
	return
}

func findIndexOfElement(arr []*Bounds, value *Bounds) int {
	for idx, el := range arr {
		if el == value {
			return idx
		}
	}
	return -1
}

// Clear function clears QuadTree
func (t *QuadTree) Clear() {
	t.objects = make([]*Bounds, 0)
	for i := 0; i < len(t.nodes); i++ {
		if len(t.nodes) != 0 {
			t.nodes[i].Clear()
		}
	}
	t.nodes = make([]*QuadTree, 0)
}

func printNode(tr *QuadTree) string {
	res := fmt.Sprintf(
		"level: %v, bounds(x|y|w|h): (%v|%v|%v|%v)\n",
		tr.level,
		tr.bounds.x,
		tr.bounds.y,
		tr.bounds.w,
		tr.bounds.h,
	)
	if len(tr.objects) <= 0 {
		return res
	}
	res += "Objects: [\n"
	for _, o := range tr.objects {
		res += fmt.Sprintf(
			"  bounds(x|y|w|h): (%v|%v|%v|%v),\n",
			o.x,
			o.y,
			o.w,
			o.h,
		)
	}
	res += "]\n"
	return res
}

// Traverse function traverses quad tree
func (t *QuadTree) Traverse() []string {
	res := make([]string, 0)
	res = append(res, printNode(t))
	for _, child := range t.nodes {
		res = append(res, child.Traverse()...)
	}
	return res
}

// PrintTraverse prints traverse
func (t *QuadTree) PrintTraverse() {
	for _, s := range t.Traverse() {
		fmt.Printf(s)
	}
}

// Test function tests quad tree
func Test(q *QuadTree) {
	fmt.Printf("Created QuadTree\n")
	q.PrintTraverse()

	q.Insert(&Bounds{
		200,
		150,
		20,
		20,
	})
	q.Insert(&Bounds{
		20,
		100,
		1,
		30,
	})
	q.Insert(&Bounds{
		11,
		67,
		10,
		30,
	})
	fmt.Printf("Added some elements to the QuadTree\n")
	q.PrintTraverse()
}
