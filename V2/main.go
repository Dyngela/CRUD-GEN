package main

import (
	"CRUDGEN/V2/parser"
	"CRUDGEN/V2/writer/java"
)

func main() {
	tables := parser.ReadFile()
	java.GenerateSpringProject(tables, "Test")
	//_go.GenerateGinProject(tables, "GinProject")
}
