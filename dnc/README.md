# Divide-and-Conquer

## Rehearsal

### Binary Search
Given a sorted set of integers like `{1,2,3,4,6}`, and a target int like `4` find its position in the set like `3`. [Solution](binary_search.go) [Test](binary_search_test.go)

### Square Root with Binary Search
Given a number and percision, return the square root of the number using the binary search algorithm. [Solution](square_root.go) [Test](square_root_test.go)

### Rate Limit
Given a number of allowed calls per second, write an IsAllowed function which returns false if the call should be rate limited and true if the call should be allowed (i.e. is within the range of allowed calls per second). [Solution](rate_limit.go) [Test](rate_limit_test.go)

### Towers of Hanoi
Given n, number of disks, and start and end tower, return the moves it takes to move all disks from start to end tower. The disks are stacked on top of each other with the lightest being on top and heaviest being in the bottom. A heavier disk cannot be placed on a lighter disk. You can move one disk at a time. [Solution](towers_of_hanoi.go) [Test](towers_of_hanoi_test.go)

### Merge Sort
Given a list of integers like `{3,1,2}`, return a sorted set like `{1,2,3}` using Merge Sort. [Solution](merge_sort.go) [Test](merge_sort_test.go)