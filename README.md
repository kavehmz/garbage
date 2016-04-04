# garbage
tests and garbage

```go
func callPointertoFunc() {
	x := new(func(int) bool)
	*x = func(x int) bool { return false }

	fmt.Println((*x)(4))
}

```


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/kavehmz/garbage/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

