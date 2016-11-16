### Build Hi
```bash
$ cd say_hi
$ go build -buildmode=plugin -o plugin.so hi.go
$ cd ..
```

### Build Bye
```bash
$ cd say_bye
$ go build -buildmode=plugin -o plugin.so bye.go
$ cd .. 
```

### Test with Hi or Bye plugin
```bash
$ cd loader
$ cp ../say_hi/plugin.so .
$ go run main.go
Hello World
$ cp ../say_bye/plugin.so .
$ go run main.go
Bye World
```
