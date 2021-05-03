# i3-tree
like [tree](https://linux.die.net/man/1/tree), but for [i3](https://i3wm.org/)

# install
install with
```bash
go get github.com/eh-am/i3-tree
```

# running
```
> i3-tree

[root] root
├──[output][output] HDMI-0
│  ├──[workspace][splith] 1
│  │  └──[con] Reddit.com - Mozilla Firefox
│  ├──[workspace][stacked] 2
│  │  ├──[con] Twitter.com - Mozilla Firefox
│  │  ├──[con] Stackoverflow.com - Google Chrome
│  │  └──[con] duckduckgo.com - Chromium
│  ├──[workspace][splitv] 3
│  │  ├──[con] Mozilla Firefox
│  │  └──[con] VLC media player
│  ├──[workspace][tabbed] 4
│  │  ├──[con] kubernetes.io - Mozilla Firefox
│  │  ├──[con] VLC media player
│  │  └──[con] Slack
│  └──[workspace][splith] 5
│     ├──[con][splitv] 
│     │  ├──[con] /bin/bash
│     │  └──[con] /bin/bash
│     └──[con][splitv] 
│        ├──[con] /bin/bash
│        └──[con] /bin/bash
└──[output][output] HDMI-1
   └──[workspace][splith] 6
      ├──[con][splitv] 
      │  └──[con] VLC media player
      └──[con][splitv] 
         ├──[con] /bin/bash
         └──[con] /bin/bash
```

<!-- Geneated with aha -->
<!--
<pre>
[root] root
├──[<span style="color:purple;">output</span>][output] HDMI-0
│  ├──[<span style="color:teal;">workspace</span>][<span style="filter: contrast(70%) brightness(190%);color:olive;">splith</span>] 1
│  │  └──[<span style="color:blue;">con</span>] Reddit.com - Mozilla Firefox
│  ├──[<span style="color:teal;">workspace</span>][<span style="filter: contrast(70%) brightness(190%);color:green;">stacked</span>] 2
│  │  ├──[<span style="color:blue;">con</span>] Twitter.com - Mozilla Firefox
│  │  ├──[<span style="color:blue;">con</span>] Stackoverflow.com - Google Chrome
│  │  └──[<span style="color:blue;">con</span>] duckduckgo.com - Chromium
│  ├──[<span style="color:teal;">workspace</span>][<span style="color:olive;">splitv</span>] 3
│  │  ├──[<span style="color:blue;">con</span>] Mozilla Firefox
│  │  └──[<span style="color:blue;">con</span>] VLC media player
│  ├──[<span style="color:teal;">workspace</span>][<span style="color:green;">tabbed</span>] 4
│  │  ├──[<span style="color:blue;">con</span>] kubernetes.io - Mozilla Firefox
│  │  ├──[<span style="color:blue;">con</span>] VLC media player
│  │  └──[<span style="color:blue;">con</span>] Slack
│  └──[<span style="color:teal;">workspace</span>][<span style="filter: contrast(70%) brightness(190%);color:olive;">splith</span>] 5
│     ├──[<span style="color:blue;">con</span>][<span style="color:olive;">splitv</span>] 
│     │  ├──[<span style="color:blue;">con</span>] /bin/bash
│     │  └──[<span style="color:blue;">con</span>] /bin/bash
│     └──[<span style="color:blue;">con</span>][<span style="color:olive;">splitv</span>] 
│        ├──[<span style="color:blue;">con</span>] /bin/bash
│        └──[<span style="color:blue;">con</span>] /bin/bash
└──[<span style="color:purple;">output</span>][output] HDMI-1
   └──[<span style="color:teal;">workspace</span>][<span style="filter: contrast(70%) brightness(190%);color:olive;">splith</span>] 6
      ├──[<span style="color:blue;">con</span>][<span style="color:olive;">splitv</span>] 
      │  └──[<span style="color:blue;">con</span>] VLC media player
      └──[<span style="color:blue;">con</span>][<span style="color:olive;">splitv</span>] 
         ├──[<span style="color:blue;">con</span>] /bin/bash
         └──[<span style="color:blue;">con</span>] /bin/bash
</pre>
-->

![Output example](./docs/example.svg)

# help
```
USAGE
  i3-tree

FLAGS
  -fetch-strat i3            where to fetch the tree from. Available: [i3 fake]
  -prune-strat non-empty-ws  what to prune from the (possible raw) tree i3. Available: [non-empty-ws none]
```
