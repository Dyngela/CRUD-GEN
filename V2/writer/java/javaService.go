package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

func generateJavaService(table parser.Table, path string) {
	var str string
	str = str + generateJavaServiceImport(table)
	str = str + fmt.Sprintf(
		`public class %s {
	%s

	%s

	%s

	%s
}`, cases.Title(language.Und).String(table.TableName), generateJavaSaveMethod(table), generateJavaDeleteMethod(table),
		generateJavaFindAllMethod(table), generateJavaFindByIdMethod(table))
	path = path + "/" + table.TableName + "Service.java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(str)
}

func generateJavaServiceImport(table parser.Table) string {
	//TODO import les dto correctement
	return fmt.Sprintf(`package com.ne.%s;

import com.ne.orders.DTO.OrderDTO;

import com.ne.exception.ExceptionHandler;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.modelmapper.ModelMapper;
import org.modelmapper.TypeToken;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;

@Service
@AllArgsConstructor
@Slf4j
`, strcase.ToLowerCamel(table.TableName))
}

func generateJavaSaveMethod(table parser.Table) string {
	return ""
}

func generateJavaDeleteMethod(table parser.Table) string {
	return ""
}

func generateJavaFindAllMethod(table parser.Table) string {
	return ""
}

func generateJavaFindByIdMethod(table parser.Table) string {
	return ""
}
