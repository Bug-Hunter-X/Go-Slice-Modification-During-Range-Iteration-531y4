# Go Slice Modification During Range Iteration
This repository demonstrates a common error in Go when modifying a slice while iterating over it using a `for range` loop.  The code attempts to remove an element at a specific index, leading to unexpected results due to the changing length of the slice.

The `bug.go` file contains the erroneous code, and `bugSolution.go` provides a corrected version.