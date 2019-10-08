# mvnparser

Go parser for maven Project Object Model (POM) file

# how to use it ?

Let's take the following POM file

```xml
<?xml version="1.0" encoding="UTF-8"?>
    <project xmlns="http://maven.apache.org/POM/4.0.0" 
	         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
        <modelVersion>4.0.0</modelVersion>
    
        <groupId>com.example</groupId>
        <artifactId>my-app</artifactId>
        <version>1.0.0-SNAPSHOT</version>
    
        <dependencies>
            <dependency>
                <groupId>junit</groupId>
                <artifactId>junit</artifactId>
                <scope>test</scope>
            </dependency>
            <dependency>
                <groupId>javax.enterprise</groupId>
                <artifactId>cdi-api</artifactId>
                <scope>provided</scope>
            </dependency>
        </dependencies>
    
        <build>
            <plugins>
                <plugin>
                    <groupId>org.apache.maven.plugins</groupId>
                    <artifactId>maven-compiler-plugin</artifactId>
                    <version>3.8.0</version>
                    <configuration>
                        <release>11</release>
                    </configuration>
                </plugin>
            </plugins>
        </build>
    </project>
```

You can read the pom file using 

```go
package main

import (
	"github.com/creekorful/mvnparser"
	"encoding/xml"
	"log"
)

func main() {
	// filled with previously declared xml
	pomStr := "..."
	
	// Load project from bytes
    var project MavenProject
    if err := xml.Unmarshal([]byte(pomStr), &project); err != nil {
        log.Fatalf("unable to unmarshal pom file. Reason: %s", err)
    }
    
    log.Print(project.GroupId) // -> com.example
    log.Print(project.ArtifactId) // -> my-app
    log.Print(project.Version) // -> 1.0.0-SNAPSHOT
    
    // iterate over dependencies
    for _, dep := range project.Dependencies {
    	log.Print(dep.GroupId)
    	log.Print(dep.ArtifactId)
    	log.Print(dep.Version)
    	
    	// ...
    }
}

```