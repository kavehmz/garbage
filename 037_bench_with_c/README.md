Comparing C/Go callign a function and using a float64 sin

```bash
$ gcc -O3 c/main.c;time ./a.out 
0.000548

real	0m0.431s
user	0m0.420s
sys	0m0.004s

$ go build main.go; time ./main
0.0005477221987112738

real	0m0.387s
user	0m0.368s
sys	0m0.007s
```
