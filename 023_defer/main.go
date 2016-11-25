package main

func a() {
	defer func() {
		println("a")
	}()
	defer func() {
		recover()
	}()
	defer func() {
		println("b")
	}()

	i := 1
	b := 1 / i
	println(b)
}

func main() {
	a()
}
