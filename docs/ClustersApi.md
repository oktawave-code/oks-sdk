# \ClustersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | 
[**ClustersInstancesNameDelete**](ClustersApi.md#ClustersInstancesNameDelete) | **Delete** /clusters/instances/{name} | 
[**ClustersInstancesNameGet**](ClustersApi.md#ClustersInstancesNameGet) | **Get** /clusters/instances/{name} | 
[**ClustersInstancesNamePost**](ClustersApi.md#ClustersInstancesNamePost) | **Post** /clusters/instances/{name} | 
[**ClustersInstancesNameTasksGet**](ClustersApi.md#ClustersInstancesNameTasksGet) | **Get** /clusters/instances/{name}/tasks | 
[**ClustersInstancesNameTasksTaskIdGet**](ClustersApi.md#ClustersInstancesNameTasksTaskIdGet) | **Get** /clusters/instances/{name}/tasks/{taskId} | 
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersInstancesNameDelete**
> []K44sTaskDto ClustersInstancesNameDelete(ctx, k44SNodesList, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k44SNodesList** | [**K44SNodesList**](K44SNodesList.md)|  | 
  **name** | **string**|  | 

### Return type

[**[]K44sTaskDto**](K44sTaskDto.md)

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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersInstancesNamePost**
> []K44sTaskDto ClustersInstancesNamePost(ctx, k44SNodesSpecification, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k44SNodesSpecification** | [**K44SNodesSpecification**](K44SNodesSpecification.md)|  | 
  **name** | **string**|  | 

### Return type

[**[]K44sTaskDto**](K44sTaskDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersInstancesNameTasksGet**
> []K44sTaskDto ClustersInstancesNameTasksGet(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**[]K44sTaskDto**](K44sTaskDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersInstancesNameTasksTaskIdGet**
> K44sTaskDto ClustersInstancesNameTasksTaskIdGet(ctx, taskId, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**K44sTaskDto**](K44sTaskDto.md)

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

 - **Content-Type**: application/json
 - **Accept**: application/json

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

 - **Content-Type**: application/json
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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersNamePost**
> K44SClusterDetailsDto ClustersNamePost(ctx, k44sClusterCreateDto, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **k44sClusterCreateDto** | [**K44sClusterCreateDto**](K44sClusterCreateDto.md)|  | 
  **name** | **string**|  | 

### Return type

[**K44SClusterDetailsDto**](K44SClusterDetailsDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

