package main

import (
	"encoding/json"
	"log"
	"os"
)

func getTranslation(text string, from string, to string) map[string]map[string]string {
	key := os.Getenv("AZURE_TRANSLATOR_API_KEY")
	region := os.Getenv("AZURE_REGION")
	url := "https://api.cognitive.microsofttranslator.com/translate?api-version=3.0"

	query := map[string]string{
		"from": from,
		"to":   to,
	}

	body := []map[string]string{
		{
			"Text": text,
		},
	}

	headers := map[string]string{
		"Ocp-Apim-Subscription-Key":    key,
		"Ocp-Apim-Subscription-Region": region,
		"Content-Type":                 "application/json",
	}

	res, _ := MakeRequest("POST", url, query, body, headers)

	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	prettyJSON, _ := json.MarshalIndent(result, "", "  ")

	m := []map[string][]map[string]string{}
	json.Unmarshal([]byte(prettyJSON), &m)

	translatedText := m[0]["translations"][0]["text"]

	fromMap := map[string]string{
		"lang": from,
		"text": text,
	}

	toMap := map[string]string{
		"lang": to,
		"text": translatedText,
	}

	var translation = map[string]map[string]string{}

	translation["from"] = fromMap
	translation["to"] = toMap

	return translation
}
