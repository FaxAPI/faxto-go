# \UserManagementApi

All URIs are relative to *https://api.fax.to/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubUserPost**](UserManagementApi.md#SubUserPost) | **Post** /subuser | This API creates a subuser
[**UserLoginPost**](UserManagementApi.md#UserLoginPost) | **Post** /user/login | This API is used for logging in on an existing user account
[**UserPost**](UserManagementApi.md#UserPost) | **Post** /user | This API registers a new user account


# **SubUserPost**
> InlineResponse2006 SubUserPost(ctx, email, password, apiKey)
This API creates a subuser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | [****](.md)| The unique email of the user | 
  **password** | [****](.md)| The password of the subuser | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLoginPost**
> InlineResponse2005 UserLoginPost(ctx, email, password)
This API is used for logging in on an existing user account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | [****](.md)| The unique email of the user | 
  **password** | [****](.md)| The password of the user | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPost**
> InlineResponse2005 UserPost(ctx, email, password)
This API registers a new user account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | [****](.md)| The unique email of the user | 
  **password** | [****](.md)| The password of the user | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

