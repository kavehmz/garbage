This is a draft for using plugins to load app functionality on fly.


"loader" is a http server that uses a plugin function to serve request if one is loaded.

"hi" plugin can be changed and loaded again while loader http server is running to change the functionility.

### Loader
```bash
$ cd loader
$ go run main.go -logtostderr
I1119 23:51:22.397036   26552 main.go:33] Checking time 1479595882
I1119 23:51:23.399146   26552 main.go:33] Checking time 1479595883
I1119 23:51:23.417463   26552 main.go:43] Plugin loaded: 1479595883.so
I1119 23:51:23.417484   26552 main.go:50] Func loaded
...
I1119 23:52:05.396799   26552 main.go:33] Checking time 1479595925
I1119 23:52:05.413332   26552 main.go:43] Plugin loaded: 1479595925.so
I1119 23:52:05.413359   26552 main.go:50] Func loaded
```

### Build Hi, change and build again
```bash
$ cd hi
$ F=$(mktemp)
$ cp hi.go $F.go
$ go build -buildmode=plugin -o plugin.so $F.go
$ cp plugin.so ../loader/$(perl -e 'print time+1').so
$ # Now do some changes and then repeat
$ curl 'http://localhost:8080/test'
$ F=$(mktemp)
$ cp hi.go $F.go
$ go build -buildmode=plugin -o plugin.so $F.go
$ cp plugin.so ../loader/$(perl -e 'print time+1').so
$ curl 'http://localhost:8080/test'
```
First built plugin will be loaded automatically. Th

