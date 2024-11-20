package main

func main() {
	func() {
		for i := 0; i < 100; i++ {
			println(i)
		}
	}()
}
