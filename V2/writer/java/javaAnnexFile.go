package java

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
)

func generateJavaApplicationProperties(path string) {
	applicationProperties := fmt.Sprintf(`spring:
  datasource:
    username: #your database username
    password: #your database password
    url: #your database url
  jpa:
    hibernate:
      ddl-auto: #which way you handle data option : update or create-drop
    properties:
      hibernate:
        dialect: #your sql dialect
        format_sql: true
      show-sql: true`)
	path = path + "/application.yml"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(applicationProperties)

}

func generateJavaMainClass(path string, projectName string) {
	mainClass := fmt.Sprintf(`package com.ne;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class %sApplication {
    public static void main(String[] args) {
        SpringApplication.run(%sApplication.class, args);
    }
}`, strcase.ToCamel(projectName), strcase.ToCamel(projectName))

	path = path + "/" + projectName + "Application.java"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(mainClass)
}

func generateJavaGitignore(path string) {
	gitignore := fmt.Sprintf(`HELP.md
target/
!.mvn/wrapper/maven-wrapper.jar
!**/src/main/**/target/
!**/src/test/**/target/

### STS ###
.apt_generated
.classpath
.factorypath
.project
.settings
.springBeans
.sts4-cache

### IntelliJ IDEA ###
.idea
*.iws
*.iml
*.ipr

### NetBeans ###
/nbproject/private/
/nbbuild/
/dist/
/nbdist/
/.nb-gradle/
build/
!**/src/main/**/build/
!**/src/test/**/build/

### VS Code ###
.vscode/
`)
	path = path + "/.gitignore"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(gitignore)
}

func generateJavaPomXML(path string, projectName string) {
	pomXML := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.7.2</version>
        <relativePath/> <!-- lookup parent from repository -->
    </parent>
    <groupId>com</groupId>
    <artifactId>ne</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>%s</name>
    <description>No description yet</description>
    <properties>
        <java.version>17</java.version>
    </properties>
    <dependencies>
<!--        <dependency>-->
<!--            <groupId>org.springframework.boot</groupId>-->
<!--            <artifactId>spring-boot-starter-data-jdbc</artifactId>-->
<!--        </dependency>-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>

		<dependency>
			<groupId>mysql</groupId>
			<artifactId>mysql-connector-java</artifactId>
			<version>8.0.30</version>
		</dependency>

        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>

        <dependency>
            <groupId>org.modelmapper</groupId>
            <artifactId>modelmapper</artifactId>
            <version>3.1.0</version>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-jpa</artifactId>
        </dependency>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
                <configuration>
                    <excludes>
                        <exclude>
                            <groupId>org.projectlombok</groupId>
                            <artifactId>lombok</artifactId>
                        </exclude>
                    </excludes>
                </configuration>
            </plugin>
        </plugins>
    </build>

</project>
`, projectName)
	path = path + "/pom.xml"
	fe, _ := os.Create(path)
	_, _ = fe.WriteString(pomXML)
}
