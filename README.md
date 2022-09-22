# Recursion

## Quicksort

https://en.wikipedia.org/wiki/Quicksort

Quicksort is an in-place sorting algorithm. Developed by British computer
scientist Tony Hoare in 1959 and published in 1961, it is still a commonly used
algorithm for sorting.

Quicksort is a divide-and-conquer algorithm. It works by selecting a 'pivot'
element from the array and partitioning the other elements into two sub-arrays,
according to whether they are less than or greater than the pivot. For this
reason, it is sometimes called partition-exchange sort. The sub-arrays are
then sorted recursively. This can be done in-place, requiring small additional
amounts of memory to perform the sorting.

Worst complexity O(N^2), average and best O(NLogN).

One of the sort algorithms used in Go is [pdqsort] (pattern-defeating quicksort).

[pdqsort]: https://cs.opensource.google/go/go/+/refs/tags/go1.19.1:src/sort/zsortfunc.go;l=55;bpv=1;bpt=1
