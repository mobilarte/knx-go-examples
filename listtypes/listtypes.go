// Copyright 2022 Martin Müller.
// Licensed under the MIT license which can be found in the LICENSE file.

// Prints a list of supported DPTs, with default value and unit.
package main

import (
	"fmt"
	"sort"

	"github.com/mobilarte/knx-exp/knx/dpt"
)

// Custom sort for DPT strings having shape "x.y"
type byDPT []string

func (s byDPT) Len() int {
	return len(s)
}

func (s byDPT) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byDPT) Less(i, j int) bool {
	var il, ir int
	var jl, jr int

	fmt.Sscanf(s[i], "%d.%d", &il, &ir)
	fmt.Sscanf(s[j], "%d.%d", &jl, &jr)

	if il < jl {
		return true
	} else if il == jl {
		return ir < jr
	} else {
		return false
	}
}

func main() {

	fmt.Printf("List of implemented DPTs in knx-exp http://github.com/mobilarte/knx-exp\n")
	fmt.Printf("%10s %20s %20s\n", "DPT", "Default Value", "Unit")

	keys := dpt.ListSupportedTypes()
	sort.Sort(byDPT(keys))

	for _, key := range keys {
		d, _ := dpt.Produce(key)

		buf := d.Pack()
		err := d.Unpack(buf)

		if err != nil {
			fmt.Println("Error unpacking")
		}

		fmt.Printf("%10s %20s %20s\n", key, d, d.(dpt.DatapointMeta).Unit())
	}

	fmt.Printf("Total number of DPTs: %d\n", len(keys))
}
