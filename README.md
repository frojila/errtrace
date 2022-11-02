# Errtrace

Errtrace is an error propagation package that wraps multiple errors and providing stacktrace info for every wrapped errors.

## How to install
```bash
go get github.com/frojila/errtrace
```

## Example usage
```go
func main() {
	// create new error
	err1 := errtrace.New("test error one")

	// wrap error without message
	err2 := errtrace.Wrap(err1)

	// wrap error with message
	err3 := errtrace.Message("test error three").Wrap(err2)

	log.Print(err3)

	// check if err3 contain err1
	if errors.Is(err3, err1) {
		fmt.Println("err3 is contain err1")
	}

	// check if err is a valid errtrace
	ok := errtrace.Valid(err3)
	fmt.Println(ok) // should print true

	ok = errtrace.Valid(errors.New("any-errror"))
	fmt.Println(ok) // should print false
}
```
the output log of example above is should be like this:
```bash
2022/11/02 23:12:40     error in "main.main": test error three 
                                at /home/mario/project/errtrace/dummy/main.go:19
                        caused by "main.main"
                                at /home/mario/project/errtrace/dummy/main.go:16
                        caused by "main.main": test error one 
                                at /home/mario/project/errtrace/dummy/main.go:13
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
