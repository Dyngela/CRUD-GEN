package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
	"strings"
)

func generateJavaModel(table parser.Table, path string) {
	var str string
	relations, tablesToBeImported, columnExcluded := generateJavaModelRelation(table)
	str = str + generateJavaModelImport(table, tablesToBeImported)
	str = str + fmt.Sprintf(
		`@Table(name = "%s")
public class %s implements Serializable {
	%s
	%s
}`, table.TableName, strcase.ToCamel(table.TableName), generateJavaModelClass(table, columnExcluded), relations)
	path = path + "/" + strcase.ToCamel(table.TableName) + ".java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(str)
}

func generateJavaModelImport(table parser.Table, tablesToImport []string) string {
	var importTable string
	for i := 0; i < len(tablesToImport); i++ {
		importTable = importTable + fmt.Sprintf("import com.ne.%s.%s;\n",
			strcase.ToLowerCamel(tablesToImport[i]), strcase.ToCamel(tablesToImport[i]))
	}
	return fmt.Sprintf(`package com.ne.%s;

%s
import lombok.*;

import java.io.Serializable;
import javax.persistence.*;
import java.time.LocalDateTime;	
import java.util.List;

@Getter
@Builder
@Setter
@Entity
@AllArgsConstructor
@NoArgsConstructor
`,
		strcase.ToLowerCamel(table.TableName), importTable)
}

func generateJavaModelClass(table parser.Table, columnExcluded []string) string {
	var fieldsWriter = ""
	for _, f := range table.Columns {
		if findIfColumnIsAManyToOneRelation(f, columnExcluded) {
			continue
		}
		if f.IsPrimaryKey == true {
			fieldsWriter = fieldsWriter + "\n\t@Id\n\t"
			//fieldsWriter = fieldsWriter + fmt.Sprintf(`@SequenceGenerator(name = "%s_sequence", sequenceName = "%s_sequence")`, strcase.ToSnake(table.TableName), strcase.ToSnake(table.TableName))
			//fieldsWriter = fieldsWriter + "\n\t"
			//fieldsWriter = fieldsWriter + fmt.Sprintf(`@GeneratedValue(strategy = GenerationType.AUTO, generator = "%s_sequence")`, strcase.ToSnake(table.TableName))
			fieldsWriter = fieldsWriter + fmt.Sprintf(`@GeneratedValue(strategy = GenerationType.AUTO)`)
			fieldsWriter = fieldsWriter + "\n\t"
			if f.Length == 0 {
				fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(name = "%s")`, strings.ToLower(f.ColumnName))
				fieldsWriter = fieldsWriter + "\n\t"
			} else {
				fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(length = %d, name = "%s")`, f.Length, strings.ToLower(f.ColumnName))
				fieldsWriter = fieldsWriter + "\n\t"
			}
			fieldsWriter = fieldsWriter + fmt.Sprintf("private %s %s; \n\n\t", strcase.ToCamel(f.DataType), strcase.ToLowerCamel(f.ColumnName))
			continue
		}

		if f.Length == 0 {
			fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(unique = %v, nullable = %v, name = "%s")`,
				f.IsUnique, f.IsNullable, strings.ToLower(f.ColumnName))
			fieldsWriter = fieldsWriter + "\n\t"
		} else {
			fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(length = %d, precision = %d, unique = %v, nullable = %v, name = "%s")`,
				f.Length, f.Precision, f.IsUnique, f.IsNullable, strings.ToLower(f.ColumnName))
			fieldsWriter = fieldsWriter + "\n\t"
		}
		fieldsWriter = fieldsWriter + fmt.Sprintf("private %s %s; \n\t", strcase.ToCamel(f.DataType),
			strcase.ToLowerCamel(f.ColumnName))
	}
	return fieldsWriter
}

func generateJavaModelRelation(table parser.Table) (string, []string, []string) {
	var relation string
	var referencedTableToBeImported []string
	var columnExcluded []string
	for i := 0; i < len(table.Columns); i++ {
		reference := table.Columns[i].Reference
		if len(reference) > 0 {
			for r := 0; r < len(reference); r++ {
				if reference[r].MappingType == "OneToMany" {
					relation = relation + fmt.Sprintf(`@OneToMany(mappedBy = "%s"%s, fetch = FetchType.LAZY)`,
						strcase.ToLowerCamel(table.TableName), findCascadeType(reference[r]))
					relation = relation + "\n\t"
					relation = relation + fmt.Sprintf("private List<%s> %s;\n\t",
						strcase.ToCamel(reference[r].ReferenceTable), getTableAndFieldConcat(reference[r].ReferenceTable, reference[r].ForeignKeyName))
					referencedTableToBeImported = append(referencedTableToBeImported, strcase.ToCamel(reference[r].ReferenceTable))
				}
				if reference[r].MappingType == "ManyToOne" {
					relation = relation + fmt.Sprintf("@ManyToOne(fetch = FetchType.LAZY)\n\t")
					relation = relation + fmt.Sprintf("@JoinColumn(name = \"%s\")\n\t",
						findPrimaryKeyAccordingToATableName(reference[r].ReferenceTable))
					//findPrimaryKeyAccordingToATableName(table.TableName))
					relation = relation + fmt.Sprintf("private %s %s;\n\t",
						strcase.ToCamel(reference[r].ReferenceTable), strcase.ToLowerCamel(reference[r].ReferenceTable))
					referencedTableToBeImported = append(referencedTableToBeImported, strcase.ToCamel(reference[r].ReferenceTable))
					columnExcluded = append(columnExcluded, reference[r].FieldName)
				}
			}
		}
	}
	return relation, referencedTableToBeImported, columnExcluded
}

func findPrimaryKeyAccordingToATableName(tableName string) string {
	for i := 0; i < len(parser.Tables); i++ {
		if parser.Tables[i].TableName == tableName {
			for col := 0; col < len(parser.Tables[i].Columns); col++ {
				if parser.Tables[i].Columns[col].IsPrimaryKey {
					return strings.ToLower(parser.Tables[i].Columns[col].ColumnName)
				}
			}
		}
	}
	return ""
}

func findCascadeType(reference parser.Reference) string {
	var cascadeType string
	if reference.OnUpdate == "CASCADE" && reference.OnDelete == "CASCADE" {
		cascadeType = ", cascade = CascadeType.ALL"
	}
	if reference.OnUpdate == "CASCADE" && reference.OnDelete == "" {
		cascadeType = ", cascade = CascadeType.PERSIST"
	}
	if reference.OnUpdate == "" && reference.OnDelete == "CASCADE" {
		cascadeType = ", cascade = CascadeType.REMOVE"
	}
	if reference.OnUpdate == "" && reference.OnDelete == "" {
		cascadeType = ""
	}

	return cascadeType
}

func findIfColumnIsAManyToOneRelation(column parser.Column, str []string) bool {
	for i := 0; i < len(str); i++ {
		if column.ColumnName == str[i] {
			return true
		}
	}
	return false
}

func getTableAndFieldConcat(table string, field string) string {
	return strcase.ToLowerCamel(table + field)
}
