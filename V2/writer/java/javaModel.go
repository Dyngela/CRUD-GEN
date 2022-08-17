package java

import (
	"CRUDGEN/V2/parser"
	"fmt"
	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

func generateJavaModel(table parser.Table, path string) {
	var str string
	str = str + generateJavaModelImport(table)
	str = str + fmt.Sprintf(
		`public class %s implements Serializable {
	%s
	%s
}`, cases.Title(language.Und).String(table.TableName), generateJavaModelClass(table), generateJavaModelRelation(table))
	path = path + "/" + table.TableName + ".java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(str)
}

func generateJavaModelImport(table parser.Table) string {
	return fmt.Sprintf(`package com.ne.%s;

import lombok.*;

import javax.persistence.*;
import java.time.LocalDateTime;

@Getter
@Builder
@Setter
@Entity
@AllArgsConstructor
@NoArgsConstructor
`,
		strcase.ToLowerCamel(table.TableName))
}

func generateJavaModelClass(table parser.Table) string {
	var fieldsWriter = ""
	for _, f := range table.Columns {
		if f.IsPrimaryKey == true {
			fieldsWriter = fieldsWriter + "\n\t@Id\n\t"
			fieldsWriter = fieldsWriter + fmt.Sprintf(`@SequenceGenerator(name = "%s_sequence", sequenceName = "%s_sequence")`, strcase.ToSnake(table.TableName), strcase.ToSnake(table.TableName))
			fieldsWriter = fieldsWriter + "\n\t"
			fieldsWriter = fieldsWriter + fmt.Sprintf(`@GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "%s_sequence")`, strcase.ToSnake(table.TableName))
			fieldsWriter = fieldsWriter + "\n\t"
			if f.Length == 0 {
				fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(name = "%s")`, strcase.ToSnake(f.ColumnName))
				fieldsWriter = fieldsWriter + "\n\t"
			} else {
				fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(length = %d, name = "%s")`, f.Length, strcase.ToSnake(f.ColumnName))
				fieldsWriter = fieldsWriter + "\n\t"
			}
			fieldsWriter = fieldsWriter + fmt.Sprintf("private %s %s; \n\n\t", strcase.ToCamel(f.DataType), strcase.ToLowerCamel(f.ColumnName))
			continue
		}

		if f.Length == 0 {
			fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(unique = %v, nullable = %v, name = "%s")`,
				f.IsUnique, f.IsNullable, strcase.ToSnake(f.ColumnName))
			fieldsWriter = fieldsWriter + "\n\t"
		} else {
			fieldsWriter = fieldsWriter + fmt.Sprintf(`@Column(length = %d, precision = %d, unique = %v, nullable = %v, name = "%s")`,
				f.Length, f.Precision, f.IsUnique, f.IsNullable, strcase.ToSnake(f.ColumnName))
			fieldsWriter = fieldsWriter + "\n\t"
		}
		fieldsWriter = fieldsWriter + fmt.Sprintf("private %s %s; \n\n\t", strcase.ToCamel(f.DataType),
			strcase.ToLowerCamel(f.ColumnName))
	}
	return fieldsWriter
}

func generateJavaModelRelation(table parser.Table) string {
	var relation string
	for i := 0; i < len(table.Columns); i++ {
		reference := table.Columns[i].Reference
		if len(reference) > 0 {
			for r := 0; r < len(reference); r++ {
				if reference[r].MappingType == "OneToMany" {
					relation = relation + fmt.Sprintf(`@OneToMany(mappedBy = "%s"%s, fetch = FetchType.EAGER)`,
						strcase.ToLowerCamel(table.TableName), findCascadeType(reference[r]))
					relation = relation + "\n\t"
					relation = relation + fmt.Sprintf("private List<%s> %s\n\n\t",
						strcase.ToCamel(reference[r].ReferenceTable), strcase.ToLowerCamel(reference[r].ReferenceTable))
				}
				if reference[r].MappingType == "ManyToOne" {
					relation = relation + fmt.Sprintf("@ManyToOne(fetch = FetchType.LAZY)\n\t")
					relation = relation + fmt.Sprintf("@JoinColumn(name = \"%s\")\n\t",
						strcase.ToLowerCamel(findPrimaryKeyAccordingToATableName(reference[r].ReferenceTable)))
					relation = relation + fmt.Sprintf("private %s %s\n\n\t",
						strcase.ToCamel(reference[r].ReferenceTable), strcase.ToLowerCamel(reference[r].ReferenceTable))
				}
			}
		}
	}
	return relation
}

func findPrimaryKeyAccordingToATableName(tableName string) string {
	for i := 0; i < len(parser.Tables); i++ {
		if parser.Tables[i].TableName == tableName {
			for col := 0; col < len(parser.Tables[i].Columns); col++ {
				if parser.Tables[i].Columns[col].IsPrimaryKey {
					return parser.Tables[i].Columns[col].ColumnName
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
