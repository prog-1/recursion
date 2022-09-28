# Exercises

## 1. Partition

### Requirements

- Time complexity `O(N)`
- Space complexity `O(1)` (no allocations) 

### Prototype

```go
func rearrange(s []int, pivot int) (i int)
```

### Problem

Implement the function `rearrange` which accepts the slice and the index named
`pivot`. The function must rearrange elements in the slice into two segments:

1. Elements with values lower than `s[pivot]`.
2. Everything else.

The function returns the index of the split such as for every element value of
`s[:i]` is lower than `s[pivot]` and `s[i:]` is everything else.

Note: Your slice values may be different depending on the algorithm that you
come up with. But the returned `i` must be the same.

### Example 1 (no repeated elements)

```go
s := []int{8, 4, 1, 0, 5, 9, 3, 7, 2, 6}
fmt.Println(rearrange(s, 4), s) // Output: 5 [2 4 1 0 3 5 9 7 8 6]
```

### Example 2 (repeated elements)

```go
s := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 6}
fmt.Println(rearrange(s, 9), s) // Output: 4 [2 5 1 5 6 9 6 7 6 9]
```

### Example 3 (empty left side)

```go
s := []int{9, 6, 1, 7, 5, 9, 6, 5, 2, 0}
fmt.Println(rearrange(s, 9), s) // Output: 0 [0 6 1 7 5 9 6 5 2 9]
```

## 2. Numbers

1) Output all numbers from 1 to `n` using recursion. It is not allowed to use `for` loops.

2) Output all numbers from `n` to 1. The program must be implemented using the same function as in (1).
The only change must be a place where you output the numbers. The purpose of this task is to understand
how recursion winds and unwinds.

Basically both (1) and (2) can be implemented using the same function puting `fmt.Println` before and
after the recursive call.

## 3. Fibonacci

Output first `n` Fibonacci numbers.

Estimate time / space complexity of your implementation. It is not
allowed to use `for` loops.

## 4. Is power of 2?

Your program must check whether the given number is an exact power of 2.

Estimate time / space complexity of your implementation. It is not allowed to use `for` loops.

## 5. Sum digits.

Your program must sum all digits of the given number.

Estimate time / space complexity of your implementation. It is not allowed to use `for` loops.

## 6. Max number.

1) Your program must find max number in the given slice. 

2) Your program must find the second max in the given slice. 

Estimate time / space complexity of your implementation. It is not allowed to use `for` loops.

## Optional: Non-recursive quick sort

Attempt to implement non-recursive version of quick sort. Hint: use a slice to emulate hardware stack.

Estimate time / space complexity of your implementation.
