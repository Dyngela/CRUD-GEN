package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
)

func generateJavaController(table parser.Table, path string) {
	var str string
	str = str + generateJavaControllerImport(table)
	str = str + fmt.Sprintf(
		`public class %sController {

	%s
	
	%s

	%s

	%s

	%s

	%s

}`, strcase.ToCamel(table.TableName),
		generateJavaControllerDependencyInjection(table),
		generateJavaGetAllController(table),
		generateJavaGetByIdController(table),
		generateJavaDeleteController(table),
		generateJavaCreateController(table),
		generateJavaUpdateController(table))
	path = path + "/" + strcase.ToCamel(table.TableName) + "Controller.java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(str)
}

func generateJavaControllerImport(table parser.Table) string {
	return fmt.Sprintf(`package com.ne.%s;

import com.ne.%s.%sDTO;
import com.ne.%s.%sService;

import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.List;

@RestController
@RequestMapping("%s")
@Slf4j
@AllArgsConstructor
`, strcase.ToLowerCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName))
}

func generateJavaControllerDependencyInjection(table parser.Table) string {
	return fmt.Sprintf("private final %sService %sService;",
		strcase.ToCamel(table.TableName), strcase.ToLowerCamel(table.TableName))
}

func generateJavaGetAllController(table parser.Table) string {
	return fmt.Sprintf(`@GetMapping("")
    public ResponseEntity<List<%sDTO>> getAll%s() {
        return ResponseEntity.ok(%sService.getAll%s());
    }`, strcase.ToCamel(table.TableName), strcase.ToCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToCamel(table.TableName))
}

func generateJavaGetByIdController(table parser.Table) string {
	column := getTablePrimaryKeyColumn(table)
	return fmt.Sprintf(`@GetMapping("{id}")
    public ResponseEntity<%sDTO> get%sById(@PathVariable("id") %s id) {
        return ResponseEntity.ok(%sService.get%sById(id));
    }`, strcase.ToCamel(table.TableName), strcase.ToCamel(table.TableName), column.DataType,
		strcase.ToLowerCamel(table.TableName), strcase.ToCamel(table.TableName))
}

func generateJavaDeleteController(table parser.Table) string {
	column := getTablePrimaryKeyColumn(table)
	return fmt.Sprintf(`@DeleteMapping("{id}")
    public ResponseEntity<String> delete%s(@PathVariable("id") %s id) {
        %sService.delete(id);
        return ResponseEntity.status(HttpStatus.GONE).body("%s deleted successfully");
    }`, strcase.ToCamel(table.TableName), column.DataType,
		strcase.ToLowerCamel(table.TableName),
		strcase.ToCamel(table.TableName))
}

func generateJavaCreateController(table parser.Table) string {
	return fmt.Sprintf(`@PostMapping("")
    public ResponseEntity<String> create%s(@RequestBody %sDTO %s) {
        %sService.save(%s);
        return ResponseEntity.status(HttpStatus.CREATED).body("%s created successfully");
    }`, strcase.ToCamel(table.TableName), strcase.ToCamel(table.TableName), strcase.ToLowerCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName),
		strcase.ToCamel(table.TableName))
}

func generateJavaUpdateController(table parser.Table) string {
	return fmt.Sprintf(`@PutMapping("")
    public ResponseEntity<String> update%s(@RequestBody %sDTO %s) {
        %sService.save(%s);
        return ResponseEntity.status(HttpStatus.ACCEPTED).body("%s updated successfully");
    }`, strcase.ToCamel(table.TableName), strcase.ToCamel(table.TableName), strcase.ToLowerCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName),
		strcase.ToCamel(table.TableName))
}
