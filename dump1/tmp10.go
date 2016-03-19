package main

import "fmt"

type sampleType struct {
	a int
}

func changeSample(t **sampleType) {
	*t = new(sampleType)
	(**t).a = 444

}

func main() {

	t := new(sampleType)

	fmt.Println(t)

	changeSample(&t)

	fmt.Println(t)
}
