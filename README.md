# Learn Go with Tests
> The test speaks to us more clearly, as if it were an assertion of truth, *not a sequence of operations*

A practice repository following [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) to learn TDD and Go fundamentals.

## Progress (Day 1)
- [x] Hello, World
- [x] Integers
- [x] Iteration

## Progress (Day 2)
- [x] Arrays and slices
- [x] Structs, methods & interfaces

## Learning Reference

### Arrays vs Slices
Arrays are fixed-length sequences with size as part of their type (like `[5]int`), while slices are flexible-length views into arrays (like `[]int`) that can grow and shrink dynamically.

The easiest way to remember: use arrays when you need a fixed size that never changes (rare), use slices for everything else (common).

### Underscore (_) in Range
The underscore is a blank identifier that tells Go "I know there's a value here, but I want to ignore it." With `range`, it's commonly used to skip the index when you only need the values.

Example with `range`:
```go
// When you need both index and value
for index, value := range numbers {
    fmt.Println(index, value)    // Using both
}

// When you only need values
for _, value := range numbers {
    fmt.Println(value)           // Ignoring index
}
```

### Variadic Arguments
Variadic functions can take a variable number of arguments using `...`. Inside the function, the parameters become a slice.

Example:
```go
// Variadic function definition
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Different ways to call it
sum(1, 2)           // Pass individual numbers
sum(1, 2, 3, 4, 5)  // Pass more numbers
nums := []int{1, 2, 3}
sum(nums...)        // Spread a slice (like Python's *)
```

## Todo
- [ ] Pointers & errors
- [ ] Maps
- [ ] Dependency Injection
- [ ] Mocking
- [ ] Concurrency
- [ ] Select
- [ ] Reflection
- [ ] Sync
- [ ] Context
- [ ] Property based tests
- [ ] Math
- [ ] Reading files
- [ ] Templating
- [ ] Generics
- [ ] Revisit HTTP server