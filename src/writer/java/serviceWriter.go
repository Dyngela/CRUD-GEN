package java

import (
	"CRUDGEN/src/api/model"
	"CRUDGEN/src/writer/writerUtils"
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
)

func CreateServices(tables []model.Table) {
	for i := 0; i < len(tables); i++ {
		f, err := os.Create("C:/CRUDGenerator/myproject/nomDeMicroService/src/main/java/com/ne/nomDeMicroService/service/" + strcase.ToCamel(tables[i].Name) + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		className, _ := createService(tables[i])
		methods, _ := createServiceMethod(tables[i])
		_, err2 := f.WriteString(
			className +
				writerUtils.OpeningBracket() +
				methods +
				writerUtils.ClosingBracket())

		if err2 != nil {
			log.Fatal(err2)
		}
	}

	fmt.Println("done creating services")
}

func createServiceMethod(table model.Table) (string, error) {

	return fmt.Sprintf(
		`
	private final %sRepository %sRepository;
	private final ModelMapper modelMapper;
`,
		strcase.ToCamel(table.Name),
		strcase.ToLowerCamel(table.Name)), nil
}

//func findAll(table model.Table) string {
//	fmt.Sprintf(
//		`
//	public %sDTO getAll%s() {
//		%s = %s.findAll();
//		modelMapper.map(
//`)
//}
//
//func findById(table model.Table) string {
//
//}
//
//func save(table model.Table) string {
//
//}
//
//func delete(table model.Table) string {
//
//}

func createService(table model.Table) (string, error) {
	serviceName := strcase.ToCamel(table.Name)
	service := fmt.Sprintf(
		`
package com.ne.%s

@Service
@AllArgsConstructor
public class %sService
`,
		strcase.ToLowerCamel(serviceName),
		serviceName)
	return service, nil
}
