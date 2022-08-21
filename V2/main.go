package main

import (
	"CRUDGEN/V2/parser"
	_go "CRUDGEN/V2/writer/go"
	"CRUDGEN/V2/writer/java"
)

func main() {
	tables := parser.ReadFile()
	java.GenerateSpringProject(tables, "SpringProject")
	_go.GenerateGinProject(tables, "GinProject")
}
