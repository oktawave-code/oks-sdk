# Go API client for swagger

Oktawave KaaS API

# Overview


- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

# Usage
```bash
import "github.com/oktawave-code/kaas_sdk"
```

# Documentation for API Endpoints


- ClustarsApi: /docs/ClustarsApi.md


# Documentation For Models

- Date: /docs/Date.md
 
- K44SClusterDetailsDto: /docs/K44SClusterDetailsDto.md
 
- K44SNodesList: /docs/K44SNodesList.md
 
- K44SNodesSpecification: /docs/K44SNodesSpecification.md
 
- K44sClusterCreateDto: /docs/K44sClusterCreateDto.md
 
- K44sInstance: /docs/K44sInstance.md
 
- K44sInstanceStatus: /docs/K44sInstanceStatus.md
 
- K44sInstanceSubregion: /docs/K44sInstanceSubregion.md
 
- K44sInstanceType: /docs/K44sInstanceType.md
 
- K44sTaskDto: /docs/K44sTaskDto.md
 
- Node: /docs/Node.md


# Documentation For Authorization

# Example
```bash
package main

import (
	"context"
	"github.com/oktawave-code/kaas_sdk"
)
func main() {
	auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
	r, err := client.Service.Operation(auth, args)
}
```



