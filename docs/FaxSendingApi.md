# \FaxSendingApi

All URIs are relative to *https://api.fax.to/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FaxDocumentIdCostsGet**](FaxSendingApi.md#FaxDocumentIdCostsGet) | **Get** /fax/{document_id}/costs | This API is for computing the cost of the fax to be sent
[**FaxHistoryGet**](FaxSendingApi.md#FaxHistoryGet) | **Get** /fax-history | This API gets the history of a fax
[**FaxJobIdStatusGet**](FaxSendingApi.md#FaxJobIdStatusGet) | **Get** /fax/{fax_job_id}/status | This API gets the status of a fax
[**FaxPost**](FaxSendingApi.md#FaxPost) | **Post** /fax | This API is for sending a new fax in any fax numbers anywhere in the world
[**FileCleanGet**](FaxSendingApi.md#FileCleanGet) | **Get** /file-clean | This API is used for cleaning a document
[**FileGeneratePreviewGet**](FaxSendingApi.md#FileGeneratePreviewGet) | **Get** /file-generate-preview | This API generates a preview of a document
[**FilesDocumentIDDelete**](FaxSendingApi.md#FilesDocumentIDDelete) | **Delete** /files/{document_id} | This API deletes a document
[**FilesGet**](FaxSendingApi.md#FilesGet) | **Get** /files | This API gets all the files of a certain user


# **FaxDocumentIdCostsGet**
> InlineResponse2001 FaxDocumentIdCostsGet(ctx, documentId, apiKey, optional)
This API is for computing the cost of the fax to be sent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **documentId** | **int32**| The id of the document to be sent | 
  **apiKey** | **string**|  | 
 **optional** | ***FaxDocumentIdCostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FaxDocumentIdCostsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **faxNumber** | **optional.String**| The fax number of the recipient | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FaxHistoryGet**
> InlineResponse2003 FaxHistoryGet(ctx, apiKey, optional)
This API gets the history of a fax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 
 **optional** | ***FaxHistoryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FaxHistoryGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.String**| The number of record to return | 
 **page** | **optional.String**| The page you want to get | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FaxJobIdStatusGet**
> InlineResponse2002 FaxJobIdStatusGet(ctx, faxJobId, apiKey)
This API gets the status of a fax

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **faxJobId** | **int32**| The id of the fax job | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FaxPost**
> InlineResponse200 FaxPost(ctx, faxNumber, documentId, apiKey)
This API is for sending a new fax in any fax numbers anywhere in the world

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **faxNumber** | [****](.md)| The fax number of the recipient | 
  **documentId** | [****](.md)| The id of the document to be sent | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileCleanGet**
> InlineResponse2009 FileCleanGet(ctx, documentId, apiKey)
This API is used for cleaning a document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **documentId** | **int32**| The id of the document | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FileGeneratePreviewGet**
> InlineResponse20010 FileGeneratePreviewGet(ctx, documentId, apiKey)
This API generates a preview of a document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **documentId** | **int32**| The id of the document | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesDocumentIDDelete**
> InlineResponse20011 FilesDocumentIDDelete(ctx, documentId, apiKey)
This API deletes a document

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **documentId** | **int32**| The id of the document | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesGet**
> InlineResponse2008 FilesGet(ctx, apiKey, optional)
This API gets all the files of a certain user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 
 **optional** | ***FilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.String**| The number of record to return | 
 **page** | **optional.String**| The page you want to get | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

