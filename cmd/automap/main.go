package main

import (
	"bytes"
	"github.com/skhoroshavin/automap/internal"
	"log"
	"os"
)

func main() {
	buf := bytes.Buffer{}
	err := internal.AutoMap(&buf, "")
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("automap_gen.go")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = file.Write(buf.Bytes())
	if err != nil {
		log.Fatalln(err)
	}

	log.Print("Successfully generated automap_gen.go")
}
