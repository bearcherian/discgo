package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/v37/github"
	"github.com/joho/godotenv"
)

const (
	discordRepoOwner = "discord"
	discordRepoName  = "discord-api-docs"
)

var arrayOf = "array of"

var objectPattern = regexp.MustCompile(`\[(?P<object>.+)\]`)

// markers pattern matches legend markers used with field names and type names
var markers = regexp.MustCompile(`[?*\\]+`)

var structFormat = "type %s struct {\n%s}"

type parsedContent struct {
	FileName string
	Structs  []string
}

var parsedStructs map[string]bool = make(map[string]bool)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env load failed: ", err)
	}

	var client *github.Client

	githubUser := os.Getenv("GITHUB_USER")
	gitHubPassword := os.Getenv("GITHUB_PASSWORD")

	if githubUser != "" {
		fmt.Println("using authenticated client")
		tp := github.BasicAuthTransport{
			Username: githubUser,
			Password: gitHubPassword,
		}
		client = github.NewClient(tp.Client())
	} else {
		fmt.Println("usin unauthenticated client")
		client = github.NewClient(nil)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("unable to get current working directory: %s\n", err)
		return
	}
	outputPath := flag.String("o", fmt.Sprintf("%s/structs", workingDir), "output directory path for generated Go files")

	// get the resources folder
	_, repoContents, _, err := client.Repositories.GetContents(
		context.Background(),
		discordRepoOwner,
		discordRepoName,
		"/docs",
		nil,
	)
	if err != nil {
		panic(fmt.Sprintf("cannot get repo contents %v", err))
	}

	parsedStructs := parseDirContents(context.Background(), client, "", repoContents)

	fmt.Printf("parsed %d files to write\n", len(parsedStructs))

	if _, err := os.Stat(*outputPath); err != nil {
		if err := os.Mkdir(*outputPath, 0755); err != nil {
			fmt.Println("unable to create output directory")
			return
		}
	}

	for _, p := range parsedStructs {
		if len(p.Structs) == 0 {
			continue
		}
		fmt.Println("writing ", fmt.Sprintf("%s/%s", *outputPath, p.FileName))

		fileContent := "package structs\n\n"
		fileContent = fileContent + strings.Join(p.Structs, "\n\n")
		if err := os.WriteFile(fmt.Sprintf("%s/%s", *outputPath, p.FileName), []byte(fileContent), 0644); err != nil {
			fmt.Printf("unable to write to file %s: %s\n", p.FileName, err)
		}
	}

}

// parseDirContents recursivley traverses directory contents and reads structs from markdown files into a map of structs
func parseDirContents(ctx context.Context, client *github.Client, dirName string, contents []*github.RepositoryContent) []parsedContent {

	currentDir := dirName
	var p []parsedContent
	for _, item := range contents {
		if *item.Type == "dir" {
			_, items, _, err := client.Repositories.GetContents(ctx, discordRepoOwner, discordRepoName, item.GetPath(), nil)
			if err != nil {
				fmt.Printf("unable to get repository item %s. %s", *item.Path, err)
			}

			p = append(p, parseDirContents(ctx, client, item.GetName(), items)...)

		} else if *item.Type == "file" {

			fileContnts, _, _, err := client.Repositories.GetContents(ctx, discordRepoOwner, discordRepoName, item.GetPath(), nil)
			if err != nil {
				fmt.Printf("unable to get file contents (%s): %s", item.GetName(), err)
			}
			contents, err := fileContnts.GetContent()
			if err != nil {
				fmt.Printf("unable to get contents: %s\n", err)
			}
			structs := parseFileContents(contents)

			fileName := strings.ReplaceAll(strings.ToLower(*item.Name), " ", "_")
			fileName = strings.ReplaceAll(fileName, ".md", ".go")
			if currentDir != "" {
				fileName = fmt.Sprintf("%s_%s", currentDir, fileName)
			}
			p = append(p, parsedContent{
				FileName: fileName,
				Structs:  structs,
			})
		}
	}

	return p
}

func parseFileContents(fileContents string) []string {
	lines := strings.Split(fileContents, "\n")
	structure := "Structure"

	var structs []string

	var parsingStructure bool = false
	var typeName string
	var structContents string
	var structLine int
	for _, line := range lines {

		// check for heading 6
		isH4 := strings.Index(line, "###### ") == 0

		if len(line) > 0 && line[0] == '#' {
			if typeName != "" && structContents != "" {
				if _, found := parsedStructs[typeName]; !found {
					structs = append(structs, fmt.Sprintf(structFormat, typeName, structContents))
					parsedStructs[typeName] = true
				}
			}
			parsingStructure = false
			typeName = ""
			structContents = ""
			structLine = 0
		}

		if parsingStructure {

			if strings.Index(line, "|") == 0 {
				structLine++
			}

			// only parse from tables
			line = strings.TrimSpace(line)
			if line == "" || !strings.Contains(line, "|") || structLine <= 2 {
				continue
			}

			// split the table columns
			cols := strings.Split(line, "|")
			if len(cols) < 3 {
				fmt.Printf("parsing failed: %#v\n\tline: '%s'", cols, line)
				continue
			}

			jsonName := markers.ReplaceAllString(cols[1], "")
			jsonName = strings.TrimSpace(jsonName)
			if jsonName == "" || jsonName[0] == '-' {
				continue
			}

			jsonType := strings.TrimSpace(markers.ReplaceAllString(cols[2], ""))
			varDescription := strings.TrimSpace(cols[3])
			varType := typeFromString(jsonType, varDescription)

			varName := strings.Title(strings.Replace(jsonName, "_", " ", -1))
			varName = strings.Replace(varName, " ", "", -1)

			structContents = fmt.Sprintf("%s%s",
				structContents,
				fmt.Sprintf("\t// %s %s\n\t%s %s `json:\"%s\"`\n", varName, varDescription, varName, varType, jsonName),
			)

		}

		isStructure := stringEndsWith(line, structure) || stringEndsWith(line, "Object") || stringEndsWith(line, "Struct")

		if isH4 && isStructure {
			line = strings.Replace(line, "######", "", -1)
			line = strings.Replace(line, "Structure", "", -1)
			line = strings.Replace(line, "Object", "", -1)
			line = strings.Replace(line, "Struct", "", -1)
			line = strings.Replace(line, " ", "", -1)
			typeName = line
			parsingStructure = true
		}
	}

	// append the last one
	if typeName != "" && structContents != "" {
		if _, found := parsedStructs[typeName]; !found {
			structs = append(structs, fmt.Sprintf(structFormat, typeName, structContents))
			parsedStructs[typeName] = true
		}
		parsingStructure = false
		typeName = ""
		structContents = ""
		structLine = 0
	}

	return structs
}

func typeFromString(s, d string) string {

	var typeString string

	if strings.Index(s, arrayOf) == 0 {
		typeString = "[]"
	}

	s = markers.ReplaceAllString(s, "")

	words := strings.Split(s, " ")
	lastWord := strings.ReplaceAll(words[len(words)-1], " ", "")

	switch lastWord {
	case "snowflake":
		fallthrough
	case "snowflakes":
		fallthrough
	case "string":
		fallthrough
	case "strings":
		typeString += "string"
	case "boolean":
		fallthrough
	case "bool":
		typeString += "bool"
	case "float":
		typeString += "float32"
	case "int":
		fallthrough
	case "integer":
		typeString += "int"
	case "Int64":
		typeString += "int64"
	case "Int32":
		typeString += "int32"
	case "UInt8":
		typeString += "uint8"
	case "UInt32":
		typeString += "uint32"
	case "timestamp":
		typeString += "time.Time"
	case "object":
		fallthrough
	case "objects":
		objectType := parseObjectFromString(s)
		if objectType == "" {
			objectType = parseObjectFromString(d)
		}

		typeString += objectType
	default:
		fmt.Printf("unknown type '%s'\n", lastWord)
		return s
	}

	return typeString

}

func parseObjectFromString(s string) string {

	if !strings.Contains(s, "object") && !strings.Contains(s, "objects") {
		return ""
	}

	// find the object name in square brackets
	match := objectPattern.FindStringSubmatchIndex(s)
	var tBytes []byte
	objectType := string(objectPattern.ExpandString(tBytes, "${object}", s, match))
	objectType = strings.ReplaceAll(objectType, " object", "")
	objectType = strings.Title(objectType)
	objectType = strings.Replace(objectType, " ", "", -1)

	return objectType
}

func stringEndsWith(s string, t string) bool {
	return len(s) > len(t) && s[len(s)-len(t):] == t
}
