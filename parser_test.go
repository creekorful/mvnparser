// MIT License
//
// Copyright (c) 2019 Aloïs Micard
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
//	copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package mvnparser

import (
	"encoding/xml"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	pomStr := `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.example</groupId>
    <artifactId>my-app</artifactId>
    <version>1.0.0-SNAPSHOT</version>

    <packaging>jar</packaging>
    <name>My App</name>

    <parent>
        <groupId>fr.creekorful</groupId>
        <artifactId>parent-project</artifactId>
        <version>1.0.0-SNAPSHOT</version>
    </parent>

    <modules>
        <module>core</module>
        <module>dao</module>
    </modules>

    <repositories>
        <repository>
            <id>private-repository</id>
            <name>My private repository</name>
            <url>http://localhost:8081/repository/maven-private/</url>
        </repository>
    </repositories>

    <pluginRepositories>
        <pluginRepository>
            <id>private-plugin-repository</id>
            <name>My private plugin repository</name>
            <url>http://localhost:8081/repository/maven-private/</url>
        </pluginRepository>
    </pluginRepositories>

    <properties>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <build.number>SNAPSHOT</build.number>
    </properties>

    <dependencyManagement>
        <dependencies>
            <!-- Import wildfly JEE8 BOM. -->
            <dependency>
                <groupId>org.wildfly.bom</groupId>
                <artifactId>wildfly-javaee8-with-tools</artifactId>
                <version>${version.wildfly.bom}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
        </dependencies>
    </dependencyManagement>

    <dependencies>
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.slf4j</groupId>
            <artifactId>slf4j-api</artifactId>
            <version>1.7.22</version>
            <scope>provided</scope>
        </dependency>
        <dependency>
            <groupId>javax.enterprise</groupId>
            <artifactId>cdi-api</artifactId>
            <scope>provided</scope>
        </dependency>
        <dependency>
            <groupId>javax.persistence</groupId>
            <artifactId>javax.persistence-api</artifactId>
            <scope>provided</scope>
        </dependency>
        <dependency>
            <groupId>org.hibernate</groupId>
            <artifactId>hibernate-core</artifactId>
            <scope>provided</scope>
        </dependency>
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>8.0.15</version>
            <scope>provided</scope>
        </dependency>
        <dependency>
            <groupId>io.swagger.core.v3</groupId>
            <artifactId>swagger-jaxrs2</artifactId>
            <version>2.0.8</version>
            <exclusions>
                <exclusion>
                    <groupId>com.fasterxml.jackson.core</groupId>
                    <artifactId>jackson-databind</artifactId>
                </exclusion>
                <exclusion>
                    <groupId>com.fasterxml.jackson.jaxrs</groupId>
                    <artifactId>jackson-jaxrs-json-provider</artifactId>
                </exclusion>
                <exclusion>
                    <groupId>com.fasterxml.jackson.core</groupId>
                    <artifactId>jackson-annotations</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
        <dependency>
            <groupId>com.example</groupId>
            <artifactId>test-framework</artifactId>
            <version>1.0.0</version>
            <classifier>jee8</classifier>
            <type>test-jar</type>
            <scope>test</scope>
        </dependency>
    </dependencies>
	
    <profiles>
        <profile>
            <id>dev</id>
            <build>
                <plugins>
                    <plugin>
                        <groupId>org.wildfly.plugins</groupId>
                        <artifactId>wildfly-maven-plugin</artifactId>
                        <executions>
                            <execution>
                                <id>deploy_jdbc_driver</id>
                                <phase>install</phase>
                                <goals>
                                    <goal>deploy-artifact</goal>
                                </goals>
                                <configuration>
                                    <groupId>mysql</groupId>
                                    <artifactId>mysql-connector-java</artifactId>
                                    <name>mysql</name>
                                    <runtimeName>mysql</runtimeName>
                                </configuration>
                            </execution>
                        </executions>
                    </plugin>
                </plugins>
            </build>
        </profile>
    </profiles>

    <build>
        <finalName>${project.artifactId}-${build.number}</finalName>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.8.0</version>
                <configuration>
                    <release>11</release>
                </configuration>
            </plugin>
            <plugin>
                <artifactId>maven-war-plugin</artifactId>
                <version>3.2.2</version>
                <configuration>
                    <failOnMissingWebXml>false</failOnMissingWebXml>
                </configuration>
            </plugin>
            <plugin>
                <groupId>org.wildfly.plugins</groupId>
                <artifactId>wildfly-maven-plugin</artifactId>
                <version>2.0.1.Final</version>
                <configuration>
                    <runtimeName>${project.artifactId}.war</runtimeName>
                </configuration>
            </plugin>
            <plugin>
                <groupId>org.jacoco</groupId>
                <artifactId>jacoco-maven-plugin</artifactId>
                <version>0.8.2</version>
                <executions>
                    <execution>
                        <goals>
                            <goal>prepare-agent</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
        </plugins>
    </build>
</project>`

	var project MavenProject
	if err := xml.Unmarshal([]byte(pomStr), &project); err != nil {
		t.Errorf("unable to unmarshal pom file. Reason: %s", err)
	}

	// Validate fields
	if project.GroupId != "com.example" {
		t.Errorf("groupId does not match (expected: com.example, found: %s)", project.GroupId)
	}
	if project.ArtifactId != "my-app" {
		t.Errorf("artifactId does not match (expected: my-app, found: %s)", project.ArtifactId)
	}
	if project.Version != "1.0.0-SNAPSHOT" {
		t.Errorf("version does not match (expected: 1.0.0-SNAPSHOT, found: %s)", project.Version)
	}
	if project.Name != "My App" {
		t.Errorf("artifactId does not match (expected: My App, found: %s)", project.Name)
	}
	if len(project.Repositories) != 1 {
		t.Errorf("expecting 1 repository found %d", len(project.Repositories))
	}
	if project.Repositories[0].Id != "private-repository" {
		t.Errorf("repository[0] id does not match (expected: private-repository, found: %s)", project.Repositories[0].Id)
	}
	if project.Repositories[0].Name != "My private repository" {
		t.Errorf("repository[0] name does not match (expected: My private repository, found: %s)", project.Repositories[0].Name)
	}
	if project.Repositories[0].Url != "http://localhost:8081/repository/maven-private/" {
		t.Errorf("repository[0] url does not match (expected: http://localhost:8081/repository/maven-private/, found: %s)", project.Repositories[0].Url)
	}
	// todo test properties
	if len(project.DependencyManagement.Dependencies) != 1 {
		t.Errorf("expecting 1 dependencies in management found %d", len(project.DependencyManagement.Dependencies))
	}
	if project.DependencyManagement.Dependencies[0].Type != "pom" {
		t.Errorf("artifactId does not match (expected: My App, found: %s)", project.Name)
	}

	if len(project.Dependencies) != 8 {
		t.Errorf("expecting 8 dependencies found %d", len(project.Dependencies))
	}

	if project.PluginRepositories[0].Id != "private-plugin-repository" {
		t.Errorf("pluginRepository[0] id does not match (expected: private-plugin-repository, found: %s)", project.PluginRepositories[0].Id)
	}

	if project.PluginRepositories[0].Name != "My private plugin repository" {
		t.Errorf("pluginRepository[0] name does not match (expected: My private plugin repository, found: %s)", project.PluginRepositories[0].Name)
	}

	if project.PluginRepositories[0].Url != "http://localhost:8081/repository/maven-private/" {
		t.Errorf("pluginRepository[0] url does not match (expected: http://localhost:8081/repository/maven-private/, found: %s)", project.PluginRepositories[0].Url)
	}

	if len(project.Modules) != 2 {
		t.Error("Number of module doesn't match")
	}
	if !contains(project.Modules, "core") {
		t.Error("Module core absent")
	}
	if !contains(project.Modules, "dao") {
		t.Error("Module dao absent")
	}

	if project.Parent.ArtifactId != "parent-project" {
		t.Error("Wrong parent artifactId")
	}
	if project.Parent.Version != "1.0.0-SNAPSHOT" {
		t.Error("Wrong parent version")
	}
	if project.Parent.GroupId != "fr.creekorful" {
		t.Error("Wrong parent groupId")
	}
}

func contains(haystack []string, needle string) bool {
	for _, wheat := range haystack {
		if wheat == needle {
			return true
		}
	}

	return false
}
