package main

const (
	apiKey            = "EK-qp2KD-pesR75o-9QwAm"
	apiMainnet        = "https://api.ethplorer.io"
	apiKovan          = "https://kovan-api.ethplorer.io"
	apiLimit          = 5
	apiShowZeroValues = false
)

func main() {
	bootloader()

	var dbpointer = initDB()
	defer dbpointer.Close()
	separator()

	listenOnPort(dbpointer)
	separator()
}
