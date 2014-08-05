goroaringbenchmark
==================

Just a little private repo to run benchmarks
on the roaring bitmap format. It is a very 
fast format on current processors, but it is 
in good part thanks to the presence of the POPCNT
instruction. Go has no native support for this
instruction, so assembly shall be required.


Usage
==================

go get  github.com/fzandona/goroar

go get  github.com/tgruben/roaring

go run roaringbenchmark.go
