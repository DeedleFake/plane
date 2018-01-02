package plane

import "sort"

// An Object represents something in the game world.
type Object interface {
	Update(*Engine)
	Draw(*Engine)
}

// An ObjectGroup is a depth-sorted group of objects that can be used
// as an Object itself.
type ObjectGroup struct {
	objects []objectE
}

// Update calls the Update method of every object in the group.
func (og ObjectGroup) Update(e *Engine) {
	for _, obj := range og.objects {
		obj.Update(e)
	}
}

// Draw calls the draw method of every object in the group.
func (og ObjectGroup) Draw(e *Engine) {
	for _, obj := range og.objects {
		obj.Draw(e)
	}
}

type objectE struct {
	Object
	depth int
}

func (og *ObjectGroup) objectI(depth int) int {
	return sort.Search(len(og.objects), func(i int) bool {
		return og.objects[i].depth <= depth
	})
}

// Add adds an object at the given depth. Higher depth values are
// drawn behind lower values. If an object already exists at a given
// depth, the object is replaced.
func (og *ObjectGroup) Add(depth int, obj Object) {
	t := objectE{
		Object: obj,
		depth:  depth,
	}

	i := og.objectI(depth)
	if i >= len(og.objects) {
		og.objects = append(og.objects, t)
		return
	}
	if og.objects[i].depth == depth {
		og.objects[i].Object = obj
		return
	}

	og.objects = append(og.objects[:i], append([]objectE{t}, og.objects[i:]...)...)
}

// Get returns the object at the given depth, or nil if no such object
// exists.
func (og *ObjectGroup) Get(depth int) Object {
	i := og.objectI(depth)
	if (i >= len(og.objects)) || (og.objects[i].depth != depth) {
		return nil
	}

	return og.objects[i].Object
}

// Remove removes the object at the given depth. If no such object
// exists, this does nothing.
func (og *ObjectGroup) Remove(depth int) {
	i := og.objectI(depth)
	if (i >= len(og.objects)) || (og.objects[i].depth != depth) {
		return
	}

	og.objects = append(og.objects[:i], og.objects[i+1:]...)
}
