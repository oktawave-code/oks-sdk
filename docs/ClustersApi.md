# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | 
[**ClustersInstancesNameDelete**](ClustersApi.md#ClustersInstancesNameDelete) | **Delete** /clusters/instances/{name} | 
[**ClustersInstancesNameGet**](ClustersApi.md#ClustersInstancesNameGet) | **Get** /clusters/instances/{name} | 
[**ClustersInstancesNamePost**](ClustersApi.md#ClustersInstancesNamePost) | **Post** /clusters/instances/{name} | 
[**ClustersKubeconfigNameGet**](ClustersApi.md#ClustersKubeconfigNameGet) | **Get** /clusters/kubeconfig/{name} | 
[**ClustersNameDelete**](ClustersApi.md#ClustersNameDelete) | **Delete** /clusters/{name} | 
[**ClustersNameGet**](ClustersApi.md#ClustersNameGet) | **Get** /clusters/{name} | 
[**ClustersNamePost**](ClustersApi.md#ClustersNamePost) | **Post** /clusters/{name} | 

# **ClustersGet**
> []K44SClusterDetailsDto ClustersGet(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]K44SClusterDetailsDto**](K44SClusterDetailsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersInstancesNameDelete**
> []K44SOperation ClustersInstancesNameDelete(ctx, body, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**K44SNodesList**](K44SNodesList.md)|  | 
  **name** | **string**|  | 

### Return type

[**[]K44SOperation**](K44SOperation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersInstancesNameGet**
> []K44sInstance ClustersInstancesNameGet(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**[]K44sInstance**](K44sInstance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersInstancesNamePost**
> []K44SOperation ClustersInstancesNamePost(ctx, body, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**K44SNodesSpecification**](K44SNodesSpecification.md)|  | 
  **name** | **string**|  | 

### Return type

[**[]K44SOperation**](K44SOperation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersKubeconfigNameGet**
> ClustersKubeconfigNameGet(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersNameDelete**
> K44SClusterDetailsDto ClustersNameDelete(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**K44SClusterDetailsDto**](K44SClusterDetailsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersNameGet**
> K44SClusterDetailsDto ClustersNameGet(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**K44SClusterDetailsDto**](K44SClusterDetailsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersNamePost**
> K44SClusterDetailsDto ClustersNamePost(ctx, body, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**K44sClusterCreateDto**](K44sClusterCreateDto.md)|  | 
  **name** | **string**|  | 

### Return type

[**K44SClusterDetailsDto**](K44SClusterDetailsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

