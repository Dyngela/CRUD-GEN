package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
)

func generateJavaService(table parser.Table, path string) {
	var str string
	str = str + generateJavaServiceImport(table)
	str = str + fmt.Sprintf(
		`public class %sService {
	%s

	%s

	%s

	%s

	%s
}`, strcase.ToCamel(table.TableName), generateJavaServiceDependencyInjection(table), generateJavaFindAllMethod(table), generateJavaFindByIdMethod(table),
		generateJavaSaveMethod(table), generateJavaDeleteMethod(table))
	path = path + "/" + strcase.ToCamel(table.TableName) + "Service.java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(str)
}

func generateJavaServiceImport(table parser.Table) string {
	return fmt.Sprintf(`package com.ne.%s;

import com.ne.%s.%sDTO;

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
`, strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName), strcase.ToCamel(table.TableName))
}

func generateJavaServiceDependencyInjection(table parser.Table) string {
	return fmt.Sprintf("private final %sRepository %sRepository;",
		strcase.ToCamel(table.TableName), strcase.ToLowerCamel(table.TableName))
}

func generateJavaSaveMethod(table parser.Table) string {
	column := getTablePrimaryKeyColumn(table)
	camelTable := strcase.ToCamel(table.TableName)
	lowerCamelTable := strcase.ToLowerCamel(table.TableName)
	camelColumn := strcase.ToCamel(column.ColumnName)
	return fmt.Sprintf(`public void save(%sDTO %sDTO) {
        if (%sDTO.get%s() != null) {
            ModelMapper mapper = new ModelMapper();
            %s %s = mapper.map(%sDTO, %s.class);
            %s %sToUpdate = %sRepository.findById(%s.get%s()).orElseThrow(()
                    -> new ExceptionHandler("Wrong input data"));
            %s = mapper.map(%sToUpdate, %s.class);
            try {
                %sRepository.save(%s);
            } catch (Exception e) {
                log.error("Error while updating a %s -> " + e);
                throw new ExceptionHandler("A fatal error has occurred while updating your %s");
            }
			return;
        }

        try {
            ModelMapper mapper = new ModelMapper();
            %s %s = mapper.map(%sDTO, %s.class);

            %sRepository.save(%s);
        } catch (Exception e) {
            log.error("Error while updating a %s -> " + e);
            throw new ExceptionHandler("A fatal error has occurred while creating your %s");
        }
    }`, camelTable, lowerCamelTable,
		lowerCamelTable, camelColumn,
		camelTable, lowerCamelTable, lowerCamelTable, camelTable,
		camelTable, lowerCamelTable, lowerCamelTable, lowerCamelTable, camelColumn,
		lowerCamelTable, lowerCamelTable, camelTable,
		lowerCamelTable, lowerCamelTable,
		lowerCamelTable,
		lowerCamelTable,
		camelTable, lowerCamelTable, lowerCamelTable, camelTable,
		lowerCamelTable, lowerCamelTable,
		lowerCamelTable,
		lowerCamelTable,
	)
}

func generateJavaDeleteMethod(table parser.Table) string {
	column := getTablePrimaryKeyColumn(table)
	return fmt.Sprintf(`public void delete(%s id) {
        %s %s = %sRepository.findById(id).orElseThrow(()
                -> new ExceptionHandler("Wrong input data"));
        %sRepository.delete(%s);
    }`, column.DataType, strcase.ToCamel(table.TableName), strcase.ToLowerCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName))
}

func generateJavaFindAllMethod(table parser.Table) string {
	return fmt.Sprintf(`public List<%sDTO> getAll%s() {
        ModelMapper mapper = new ModelMapper();
        List<%s> %s = %sRepository.findAll();
        return mapper.map(%s, new TypeToken<List<%sDTO>>() {}.getType());
    }`, strcase.ToCamel(table.TableName), strcase.ToCamel(table.TableName), strcase.ToCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName),
		strcase.ToCamel(table.TableName))
}

func generateJavaFindByIdMethod(table parser.Table) string {
	column := getTablePrimaryKeyColumn(table)
	return fmt.Sprintf(`public %sDTO get%sById(%s id) {
        ModelMapper mapper = new ModelMapper();
        %s %s = %sRepository.findById(id).orElseThrow(() -> new ExceptionHandler("Cannot find your %s"));
        return mapper.map(%s, %sDTO.class);
    }`, strcase.ToCamel(table.TableName), strcase.ToCamel(table.TableName), column.DataType,
		strcase.ToCamel(table.TableName), strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName), strcase.ToLowerCamel(table.TableName),
		strcase.ToLowerCamel(table.TableName), strcase.ToCamel(table.TableName))
}
