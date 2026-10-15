Find indices of 2 elements in a slice that add up to `target`:
![01-svg](images/01.svg)

Naive solution -- O(n^2):

```go
for i := range nums {
    for j := range nums {
        if nums[i] + nums[j] == target && i != j {
            return []int{i, j}
        }
    }
}
```

Optimized solution -- O(n):

1. Create a hashmap: num : idx
2. For each `num`, calculate `candidate = target - num`
   2.1. Check if `candidate` is in map -- ~O(1)
   2.2. If so, return its index
