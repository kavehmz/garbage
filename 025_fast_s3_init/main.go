package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	total := flag.Int64("total", 0, "Total block size in MB")
	block := flag.Int64("block", 100, "block init size in MB [dd bs]. count will be 1")
	parallel := flag.Int64("parallel", 2, "size")
	device := flag.String("device", "/dev/xvdf", "name")
	offset := flag.Int64("offset", 0, "starting offset that will be used to skip")

	flag.Parse()

	p := make(chan bool, *parallel)

	c := *offset
	for c*(*block) < *total {
		go func(c int64) {
			ifd := fmt.Sprintf("if=%s", *device)
			of := fmt.Sprintf("of=/dev/null")
			bs := fmt.Sprintf("bs=%dM", *block)
			count := fmt.Sprintf("count=1")
			skip := fmt.Sprintf("skip=%d", c)
			out, err := exec.Command("dd", ifd, of, bs, count, skip).CombinedOutput()
			if err != nil {
				log.Println(err, out)
			}
			<-p
		}(c)
		p <- true
		c++
		fmt.Println("offset:", c, "total:", c*(*block), "MB")
	}

	for i := int64(0); i < *parallel; i++ {
		p <- true
	}
}
