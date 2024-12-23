#!/bin/bash

for dir in ./[0-9]* ; do
    if [ -d "$dir" ]; then

        # build execs 
        (cd "$dir" && go build -o "$(basename "$dir")")

        hyperfine --warmup 3  --min-runs 10  "$dir/$(basename "$dir")"

        # cleaup
        rm "$dir/$dir"
    fi
done