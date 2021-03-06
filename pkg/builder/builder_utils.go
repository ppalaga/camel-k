/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package builder

import (
	"encoding/xml"
	"os"

	"github.com/apache/camel-k/pkg/util/maven"
	"github.com/apache/camel-k/version"

	"github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
)

// MavenExtraOptions --
func MavenExtraOptions() string {
	if _, err := os.Stat("/tmp/artifacts/m2"); err == nil {
		return "-Dmaven.repo.local=/tmp/artifacts/m2"
	}
	return "-Dcamel.noop=true"
}

// ArtifactIDs --
func ArtifactIDs(artifacts []v1alpha1.Artifact) []string {
	result := make([]string, 0, len(artifacts))

	for _, a := range artifacts {
		result = append(result, a.ID)
	}

	return result
}

// NewProject --
func NewProject(ctx *Context) maven.Project {
	return maven.Project{
		XMLName:           xml.Name{Local: "project"},
		XMLNs:             "http://maven.apache.org/POM/4.0.0",
		XMLNsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: "http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd",
		ModelVersion:      "4.0.0",
		GroupID:           "org.apache.camel.k.integration",
		ArtifactID:        "camel-k-integration",
		Version:           version.Version,
		Properties:        ctx.Request.Platform.Build.Properties,
		DependencyManagement: maven.DependencyManagement{
			Dependencies: maven.Dependencies{
				Dependencies: []maven.Dependency{
					{
						GroupID:    "org.apache.camel",
						ArtifactID: "camel-bom",
						Version:    ctx.Request.Platform.Build.CamelVersion,
						Type:       "pom",
						Scope:      "import",
					},
				},
			},
		},
		Dependencies: maven.Dependencies{
			Dependencies: make([]maven.Dependency, 0),
		},
	}
}
