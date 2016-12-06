Running dd to initialize a new volume which was created from a snapshot in parallel to achieve faster init time.

```bash
$ time /tmp/main -block 1 -parallel 1 -total 300100 -offset 300000
...
offset: 300100 total: 300100 MB
real	0m23.577s

$ time /tmp/main -block 1 -parallel 10 -total 400100 -offset 400000
...
offset: 400100 total: 400100 MB
real	0m5.551s

$ time /tmp/main -block 1 -parallel 50 -total 500100 -offset 500000
offset: 500100 total: 500100 MB
real	0m5.036s
```