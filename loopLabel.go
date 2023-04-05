package main

func main() {
	TestLoop:
	for {
		for {
			for {
				break TestLoop
			}
		}
	}
}