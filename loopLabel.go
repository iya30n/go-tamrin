package main

func main() {
	TestLoop:
	// the main loop
	for {
		for {
			for {
				// NOTE: breaks the main loop
				break TestLoop

				// NOTE: continues the main loop
				// continue TestLoop

				// NOTE: restarts the main loop
				// goto TestLoop
			}
		}
	}
}