# immutable-go

Straightforward, functional immutable data collections for Go.

## Installation

```bash
go get -u github.com/dowlandaiello/immutable-go
```

## Lists

### Initialization

```go
import "github.com/dowlandaiello/immutable-go"

list := immutable.NewList(1, 2, 3, 4) // Initializes a new list of integers
```

### Getting Values

```go
import "github.com/dowlandaiello/immutable-go"

list := immutable.NewList(1, 2, 3, 4) // Initializes a new list of integers

firstElement := list(0) // Get the element at index 0
secondElement := list(1) // Get the element at index 1
...
```

### Setting Values

```go
import "github.com/dowlandaiello/immutable-go"

list := immutable.NewList(1, 2, 3, 4) // Initialize a new list of integers
newList := list.Set(0, 37) // Set the element at index 0 to 37
```

Or, push a value:

```go
import "github.com/dowlandaiello/immutable-go"

list := immutable.NewList(1, 2, 3, 4) // Initialize a new list of integers
newList := list.Push(102) // Push an integer with the value 102 to the list
```

Finally, pop a value:

```go
import "github.com/dowlandaiello/immutable-go"

list := immutable.NewList(1, 2, 3, 4) // Initialize a new list of integers
newList := list.Pop() // Remove an integer from the list
```

### Getting a List's Used Size

```go
import "github.com/dowlandaiello/immutable-go"

list := immutable.NewList(1, 2, 3, 4) // Initialize a new list of integers
length := list.Size() // Get the size of the list
```
