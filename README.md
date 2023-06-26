# repeat
[![Tests](https://github.com/gruzovator/repeat/actions/workflows/go.yml/badge.svg)](https://github.com/gruzovator/repeat/actions/workflows/go.yml)

Set of golang utils to make periodic calls.

## Example
```go
func main() {
    stopFn := repeat.Start(time.Second, func(ctx context.Context) {
        fmt.Println("hello")
    })
    defer stopFn()
    
    // do smth
    time.Sleep(3 * time.Second)
}
```