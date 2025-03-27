package main

const a = "Hello World!"

type ID int

var (
	b    = true
	c    = 10
	d    = "Williams"
	e    = 1.2
	f ID = 1
)

func main() {
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
