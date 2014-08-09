go-bitarray
-----------

A library to compactly store a large amount of integers using bitmasks. Provides constant time access in O(n) space. Therefore it's best used in situations where you are very likely to encounter all integers from 0 to n, but you need a way to keep track of which ones you've seen. Do not use this to store a small amount of numbers. For example, storing the number 8,000,000,000 will allocate 1GB of memory. Storing *all* numbers from 0 to 8,000,000,000 will also use 1GB of memory, and thus 32x less than a standard array or slice of said integers.


**Installation**: `go get github.com/arsham-f/go-bitarray`. It will be imported as `bitarray`

**Usage**

```go
b := bitarray.New(1000)

b.Add(23)
b.Add(25)

b.Contains(23) // true
b.Contains(24) // false
b.Contains(25) // true

b.Remove(25)

b.Contains(23) // true
b.Contains(24) // false
b.Contains(25) // false
```