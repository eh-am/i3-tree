package fetch

import "go.i3wm.org/i3/v4"

type FromI3 struct{}

func (i FromI3) Fetch() (i3.Tree, error) {
	return i3.GetTree()
}
