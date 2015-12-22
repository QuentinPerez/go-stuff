## How to simulate an exception in Golang

- Using `defer` + `channel`

```go
defer func(exc chan<- interface{}) {
	if err := recover(); err != nil {
		exc <- err
	}
}(exception)
```


## Example

```console
Exception caught 0 throw
No exception \o/
No exception \o/
Exception caught 1 throw
No exception \o/
Exception caught 2 throw
Exception caught 3 throw
Exception caught 4 throw
Too many panics
```
