#!/usr/bin/env bash

# Incredible ugly hack
# the idea is to record using asciinema
# and then generate a nice svg for a late frame
# so that it encompasses the whole output
set -euo pipefail

mkdir -p tmp

# the .5s is due to a race condition
asciinema rec tmp/asciinema.cast -c "i3-tree --fetch-strat=fake; sleep .5s"
trap 'rm -f tmp/asciinema.cast' EXIT

mkdir -p docs
# crazy heuristics
cat tmp/asciinema.cast | svg-term --out docs/example.svg --window --at 999 --width 50 --height 30 --padding-x 20

