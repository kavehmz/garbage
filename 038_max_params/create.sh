NUM=4000
VARLIST=$(for i in $(seq 1 $NUM); do echo -n "x$i, ";done)
VALLIST=$(for i in $(seq 1 $NUM); do echo -n "$i, ";done)
RETLIST=$(for i in $(seq 1 $NUM); do echo -n "int64, ";done)

cat > multi_ret <<END
package main

func test($VARLIST x0 int64) ($RETLIST int64) {
	return $VALLIST 0
}

func main() {
    test($VALLIST 0)
}
END

cat > single_ret <<END
package main

func test($VARLIST x0 int64) int64 {
	return 0
}

func main() {
    test($VALLIST 0)
}
END

mkdir -p single
mkdir -p multi
mv single_ret single/main.go
mv multi_ret multi/main.go

echo "building for single return and $NUM params"
time go build -o single/main single/main.go

echo "building for $NUM returns and $NUM params"
time go build -o multi/main multi/main.go