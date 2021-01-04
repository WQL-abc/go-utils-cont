# What is this?
If you want to know if "Hello" is included in "Hello, World!", Use `string.Contains("Hello, World!","Hello")`.  
If you want to know whether "1" is included in []string{"1", "2", "3"}, you have to implement it yourself.  
At that time, I made something that was put together and can be used even a little.

# Installation
```commandline
go get -u github.com/go-utils/cont
```

# Usage
```go
import (
    "fmt"
    "time"

    "github.com/go-utils/cont"
)

func main() {
    if cont.Contains([]string{"1", "2", "3"}, "1") {
        fmt.Println("ex1: OK!")
    }

    if cont.Contains([]int{1, 2, 3}, 1) {
        fmt.Println("ex2: OK!")
    }

    if cont.Contains("123", "1") {
        fmt.Println("ex3: OK!")
    }

    now := time.Now()
    mp := map[string]interface{}{"a": "1", "b": 2, "c": "3", "d": now}
    if cont.Contains(mp, now) {
        fmt.Println("ex4: OK!")
    }
    if cont.Contains(mp, 2) {
        fmt.Println("ex5: OK!")
    }
    if cont.Contains(mp, "c") {
        fmt.Println("ex6: OK!")
    }
}

```

### Result
```
ex1: OK!
ex2: OK!
ex3: OK!
ex4: OK!
ex5: OK!
ex6: OK!
```

Check [Go Playground](https://play.golang.org/p/XW1jxv-PNPl)

# Support
`string`, `slice`, `array`, `map`

# License
[MIT](./LICENSE)