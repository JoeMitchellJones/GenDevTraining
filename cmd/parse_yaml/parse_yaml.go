package main

import (
	// "fmt"
	"flag"
	"fmt"
	"os"

	yamlparser "github.com/CircleCI-Public/circleci-yaml-language-server/pkg/parser"
	languageservice "github.com/CircleCI-Public/circleci-yaml-language-server/pkg/services"
	"github.com/CircleCI-Public/circleci-yaml-language-server/pkg/utils"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

func main() {
	filepath := ".circleci/config.yml"
	// filepath := "examples/config1.yml"
	// filepath := "/home/adib/circleci/circle/.circleci/config.yml"

	schemaRef := flag.String("schema", "", "Location of the schema")

	flag.Parse()

	schema := *schemaRef
	if schema == "" {
		schema = os.Getenv("SCHEMA_LOCATION")

		if schema == "" {
			fmt.Print("No schema defined")
			return
		}
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Unable to read file \"%s\"", filepath)
		panic(err)
	}
	context := &utils.LsContext{
		Api: utils.ApiContext{
			Token:   "XXXXXXXXXXXX",
			HostUrl: "https://circleci.com",
		},
	}

	yamlparser.ParseFile(content, context)

	cache := utils.CreateCache()
	cache.FileCache.SetFile(utils.CachedFile{
		TextDocument: protocol.TextDocumentItem{
			URI:  uri.File(filepath),
			Text: string(content),
		},
		Project:      utils.Project{},
		EnvVariables: make([]string, 0),
	})

	fileURI := uri.File(filepath)
	languageservice.DiagnosticFile(fileURI, cache, context, schema)

	// fmt.Printf("S-expression:\n%v\n\n", node.RootNode)
}
