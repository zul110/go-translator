# GO Translator

More of an experimentation with Go, a simple **Swedish (sv) to English (en)** translator that uses [Azure Translator APIs](https://azure.microsoft.com/en-us/services/cognitive-services/translator/) for translation.

#### Motivation(s)

The project has 2 main motivations:

1. Learning Go - A new language I'm trying to learn after JS/TS, Python, and C#
2. Translation from Swedish to English - Something to help me in Sweden, where I now reside

#### Configuration

Although the project is pretty straightforward, and does not require much configuration, one can set the `text` to be translated, the `from` language code, and the `to` language code in `main.go`. Language codes are publicly available [here](https://api.cognitive.microsofttranslator.com/languages?api-version=3.0).
Furthermore, the following `Environment variables` need to be added to a `.env` file:

- `AZURE_TRANSLATOR_API_KEY`
- `AZURE_REGION`

#### Roadmap

As I've mentioned previously, one of the main goals of this project is to learn the Go language. However, this is just the beginning, and as of now, the project shall evolve to something more useful. One of the main features shall be switching the `to` and `from` languages, which shall follow changing the `text` to be translated. However, again, since this is experimental, I may not continue it myself. Who knows ¯\\_(ツ)_/¯

##### Dependencies

- [`joho/dotenv`](github.com/joho/godotenv)
