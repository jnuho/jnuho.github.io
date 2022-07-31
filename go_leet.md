
# Leetcode in Go

- [Reference](https://github.com/austingebauer/go-leetcode)
- [1-two_sum](https://raw.githubusercontent.com/jnuho/jnuho.github.io/master/hellogo/algorithm/1-two_sum/solution.go)

```go
 
 // nil value if not initialized
 var m map[string] int

// Map types are reference types, like pointers or slices
// and so the value of m above is nil; it doesn’t point to an initialized map.
// A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic;
// don’t do that. To initialize a map, use the built in make function:
m = make(map[string]int)

//The make function allocates and initializes a hash map data structure and returns a map value
// that points to it. The specifics of that data structure are an implementation
// detail of the runtime and are not specified by the language itself.
// In this article we will focus on the use of maps, not their implementation.

```