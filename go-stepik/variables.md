# Variables in Go

```golang
a := 5

var a int = 5

var a = 5

var a int
a = 5
```

## Constants

```go
const pi float64 = 3.1415

const (
    a int = 45
    b float32 = 3.3
)
```

## Iota

```go
const (
    Sunday    = 0
    Monday    = 1
    Tuesday   = 2
    Wednesday = 3
    Thursday  = 4
    Friday    = 5
    Saturday  = 6
)
```

Can be replaced with

```go
const (
    Sunday    = iota
    Monday    
    Tuesday   
    Wednesday 
    Thursday  
    Friday    
    Saturday  
)
```
