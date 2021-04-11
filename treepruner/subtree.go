package treepruner

import "go.i3wm.org/i3/v4"

// condition to prune the subtree
type pruneCond func(*i3.Node) bool

// Prunes a whole subtree based on a criteria
// it explores subtrees recursively until
// * a valid subtree is found
//   which the whole subtree is returned immediately
// * a leaf is found
func pruneSubtree(tree *i3.Tree, checkFn pruneCond) *i3.Tree {
	// we will call it recursively
	// therefore it needs to be declared first
	var helper func(src *i3.Node) *i3.Node

	helper = func(src *i3.Node) *i3.Node {
		if checkFn(src) {
			return src
		}

		subtrees := make([]*i3.Node, 0)
		for _, n := range src.Nodes {
			r := helper(n)
			if r != nil {
				subtrees = append(subtrees, r)
			}
		}

		if len(subtrees) > 0 {
			src.Nodes = subtrees
			return src
		}

		return nil
	}

	t := tree.Root
	if tree.Root != nil {
		t = helper(tree.Root)
	}

	return &i3.Tree{
		Root: t,
	}
}
