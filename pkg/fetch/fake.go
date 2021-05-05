package fetch

import "go.i3wm.org/i3/v4"

type FromFake struct{}

// Fetch fetches a fake tree generated manually
func (_ FromFake) Fetch() (i3.Tree, error) {
	return fakeTree(), nil
}

func fakeTree() i3.Tree {
	// Horizontal Split
	ws1 := &i3.Node{
		Name:   "1",
		Type:   i3.NodeType(i3.WorkspaceNode),
		Layout: i3.Layout(i3.SplitH),
		Nodes: []*i3.Node{
			{
				Name: "Reddit.com - Mozilla Firefox",
				Type: i3.NodeType(i3.Con),
			},
		},
	}

	// Stacked
	ws2 := &i3.Node{
		Name:   "2",
		Type:   i3.NodeType(i3.WorkspaceNode),
		Layout: i3.Layout(i3.Stacked),
		Nodes: []*i3.Node{
			{
				Name: "Twitter.com - Mozilla Firefox",
				Type: i3.NodeType(i3.Con),
			},
			{
				Name: "Stackoverflow.com - Google Chrome",
				Type: i3.NodeType(i3.Con),
			},
			{
				Name: "duckduckgo.com - Chromium",
				Type: i3.NodeType(i3.Con),
			},
		},
	}

	// Vertical Split
	ws3 := &i3.Node{
		Name:   "3",
		Type:   i3.NodeType(i3.WorkspaceNode),
		Layout: i3.Layout(i3.SplitV),
		Nodes: []*i3.Node{
			{
				Name: "Mozilla Firefox",
				Type: i3.NodeType(i3.Con),
			},
			{
				Name: "VLC media player",
				Type: i3.NodeType(i3.Con),
			},
		},
	}

	// Tabbed
	ws4 := &i3.Node{
		Name:   "4",
		Type:   i3.NodeType(i3.WorkspaceNode),
		Layout: i3.Layout(i3.Tabbed),
		Nodes: []*i3.Node{
			{
				Name: "kubernetes.io - Mozilla Firefox",
				Type: i3.NodeType(i3.Con),
			},
			{
				Name: "VLC media player",
				Type: i3.NodeType(i3.Con),
			},
			{
				Name: "Slack",
				Type: i3.NodeType(i3.Con),
			},
		},
	}

	// Slightly more complex with nested panes
	ws5 := &i3.Node{
		Name:   "5",
		Type:   i3.NodeType(i3.WorkspaceNode),
		Layout: i3.Layout(i3.SplitH),
		Nodes: []*i3.Node{
			{
				Type:   i3.NodeType(i3.Con),
				Layout: i3.Layout(i3.SplitV),
				Nodes: []*i3.Node{
					{
						Name: "/bin/bash",
						Type: i3.NodeType(i3.Con),
					},
					{
						Name: "/bin/bash",
						Type: i3.NodeType(i3.Con),
					},
				},
			},
			{
				Type:   i3.NodeType(i3.Con),
				Layout: i3.Layout(i3.SplitV),
				Nodes: []*i3.Node{
					{
						Name: "/bin/bash",
						Type: i3.NodeType(i3.Con),
					},
					{
						Name: "/bin/bash",
						Type: i3.NodeType(i3.Con),
					},
				},
			},
		},
	}

	ws6 := &i3.Node{
		Name:   "6",
		Type:   i3.NodeType(i3.WorkspaceNode),
		Layout: i3.Layout(i3.SplitH),
		Nodes: []*i3.Node{
			{
				Type:   i3.NodeType(i3.Con),
				Layout: i3.Layout(i3.SplitV),
				Nodes: []*i3.Node{
					{
						Name:    "VLC media player",
						Type:    i3.NodeType(i3.Con),
						Focused: true,
					},
				},
			},
			{
				Type:   i3.NodeType(i3.Con),
				Layout: i3.Layout(i3.SplitV),
				Nodes: []*i3.Node{
					{
						Name: "/bin/bash",
						Type: i3.NodeType(i3.Con),
					},
					{
						Name: "/bin/bash",
						Type: i3.NodeType(i3.Con),
					},
				},
			},
		},
	}

	node := i3.Node{
		Name: "root",
		Type: i3.NodeType(i3.Root),

		Nodes: []*i3.Node{
			{
				Name:   "HDMI-0",
				Type:   i3.NodeType(i3.OutputNode),
				Layout: i3.Layout(i3.OutputLayout),
				Nodes: []*i3.Node{
					ws1,
					ws2,
					ws3,
					ws4,
					ws5,
				},
			},
			{
				Name:   "HDMI-1",
				Type:   i3.NodeType(i3.OutputNode),
				Layout: i3.Layout(i3.OutputLayout),
				Nodes: []*i3.Node{
					ws6,
				},
			},
		},
	}

	return i3.Tree{
		Root: &node,
	}
}
