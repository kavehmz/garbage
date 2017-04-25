Comparing C/Go callign a function and using a float64 sin

```bash
$ gcc -O3 c/main.c;time ./a.out 
0.0005477222260973496

real	0m0.403s
user	0m0.389s
sys	0m0.005s

$ go build main.go; time ./main
0.0005477221987112738

real	0m0.394s
user	0m0.371s
sys	0m0.008s
```
