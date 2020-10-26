package main

func main() {
	var balance uint32 = 15_000_000_00 // 15 миллионов в копейках
	var transfer uint32 = 10_000_000_00 // 10 миллионов в копейках
	total := balance + transfer // int32 + int32 будет int32
	println(total)
}
