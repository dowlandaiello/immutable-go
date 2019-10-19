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
