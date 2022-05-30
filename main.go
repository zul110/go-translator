package main

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	text := "Hej!"
	from := "sv"
	to := "en"

	godotenv.Load()

	translation := getTranslation(text, from, to)
	b, _ := json.Marshal(translation)

	fmt.Printf("%+v", string(b))
}
