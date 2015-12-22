## How to kwon if a channel if full

- Using `select` statement

```go
select {
case c <- false:
default:
	fmt.Println("Channel is full")
}

```

- Using `len()` and `cap()` functions

```go
if len(c) == cap(c) {
	fmt.Println("Channel is full")
}
```

## Example

```console
$> go run main.go
Channel is full
Channel is full
Channel is not full
Channel is not full
```
