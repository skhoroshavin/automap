package main

import (
	"automap/internal"
	"log"
	"os"
)

func main() {
	file, err := os.Create("automap_gen.go")
	if err != nil {
		log.Fatalln(err)
	}

	err = internal.AutoMap(file, "")
	if err != nil {
		_ = os.Remove("automap_gen.go")
		log.Fatalln(err)
	}

	log.Print("Successfully generated automap_gen.go")
}
