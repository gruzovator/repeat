# repeat

Set of golang utils to make periodic calls.

## Example 1
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

## Example 2
```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    repeat.StartWithContext(ctx, time.Second, func(ctx context.Context) {
        fmt.Println("hello")
    })
    
    // do smth
    time.Sleep(3 * time.Second)
}
```



