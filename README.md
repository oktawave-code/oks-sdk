# Go API client for swagger

Oktawave KaaS API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClustersApi* | [**ClustersGet**](docs/ClustersApi.md#clustersget) | **Get** /clusters | 
*ClustersApi* | [**ClustersInstancesNameDelete**](docs/ClustersApi.md#clustersinstancesnamedelete) | **Delete** /clusters/instances/{name} | 
*ClustersApi* | [**ClustersInstancesNameGet**](docs/ClustersApi.md#clustersinstancesnameget) | **Get** /clusters/instances/{name} | 
*ClustersApi* | [**ClustersInstancesNamePost**](docs/ClustersApi.md#clustersinstancesnamepost) | **Post** /clusters/instances/{name} | 
*ClustersApi* | [**ClustersKubeconfigNameGet**](docs/ClustersApi.md#clusterskubeconfignameget) | **Get** /clusters/kubeconfig/{name} | 
*ClustersApi* | [**ClustersNameDelete**](docs/ClustersApi.md#clustersnamedelete) | **Delete** /clusters/{name} | 
*ClustersApi* | [**ClustersNameGet**](docs/ClustersApi.md#clustersnameget) | **Get** /clusters/{name} | 
*ClustersApi* | [**ClustersNamePost**](docs/ClustersApi.md#clustersnamepost) | **Post** /clusters/{name} | 

## Documentation For Models

 - [Date](docs/Date.md)
 - [IaasDictionaryEntry](docs/IaasDictionaryEntry.md)
 - [IaasTicket](docs/IaasTicket.md)
 - [K44SClusterDetailsDto](docs/K44SClusterDetailsDto.md)
 - [K44SNodesList](docs/K44SNodesList.md)
 - [K44SNodesSpecification](docs/K44SNodesSpecification.md)
 - [K44SOperation](docs/K44SOperation.md)
 - [K44sClusterCreateDto](docs/K44sClusterCreateDto.md)
 - [K44sInstance](docs/K44sInstance.md)
 - [K44sInstanceStatus](docs/K44sInstanceStatus.md)
 - [K44sInstanceSubregion](docs/K44sInstanceSubregion.md)
 - [K44sInstanceType](docs/K44sInstanceType.md)
 - [Node](docs/Node.md)

## Documentation For Authorization

## bearer
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author


