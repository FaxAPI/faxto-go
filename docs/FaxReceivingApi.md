# \FaxReceivingApi

All URIs are relative to *https://api.fax.to/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AreaCodesCountryCodeStateIdGet**](FaxReceivingApi.md#AreaCodesCountryCodeStateIdGet) | **Get** /areacodes/{COUNTRY_CODE}/{STATE_ID} | This API gets a list of countries with its area code
[**CountriesCountryCodeDidGroupsGet**](FaxReceivingApi.md#CountriesCountryCodeDidGroupsGet) | **Get** /countries/{countryCode}/didgroups | This API gets a list of DID groups
[**CountriesGet**](FaxReceivingApi.md#CountriesGet) | **Get** /countries | This API gets a list of countries available in the Fax.to coverage
[**IncomingFaxesGet**](FaxReceivingApi.md#IncomingFaxesGet) | **Get** /incoming-faxes | This API gets a list of incoming faxes
[**IncomingFaxesRecipientGet**](FaxReceivingApi.md#IncomingFaxesRecipientGet) | **Get** /incoming-faxes/{recipient} | This API gets a list of incoming faxes for a specific recipient
[**NumbersGet**](FaxReceivingApi.md#NumbersGet) | **Get** /numbers | This API gets a list of numbers to get the current configuration of one or multiple number
[**ProvisionNumbersGet**](FaxReceivingApi.md#ProvisionNumbersGet) | **Get** /countries/didgroups/{did_group_id}/provision | This API gets a list of provisioned numbers
[**StatesCountryCodeGet**](FaxReceivingApi.md#StatesCountryCodeGet) | **Get** /states/{COUNTRY_CODE} | This API gets a list of states of a given country available in the Fax.to coverage


# **AreaCodesCountryCodeStateIdGet**
> InlineResponse20014 AreaCodesCountryCodeStateIdGet(ctx, cOUNTRYCODE, sTATEID, apiKey)
This API gets a list of countries with its area code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTRYCODE** | **int32**| Indicates the country code in its ISO 3166-1 alpha-3 format | 
  **sTATEID** | **int32**| The numerical identifier for the state | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountriesCountryCodeDidGroupsGet**
> InlineResponse20015 CountriesCountryCodeDidGroupsGet(ctx, cOUNTRYCODE, areaCode, apiKey, optional)
This API gets a list of DID groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTRYCODE** | **int32**| Indicates the country code of the DID group in its ISO 3166-1 alpha-3 format | 
  **areaCode** | **int32**| The area code of the DID group | 
  **apiKey** | **string**|  | 
 **optional** | ***CountriesCountryCodeDidGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CountriesCountryCodeDidGroupsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **didGroupIds** | **optional.Int32**| Used to display more information about specific DID groups | 
 **stateId** | **optional.Int32**| The numerical identifier for the didGroup&#39;s state | 
 **cityNamePattern** | **optional.Int32**| A string pattern for the beginning of city name | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountriesGet**
> InlineResponse20012 CountriesGet(ctx, apiKey)
This API gets a list of countries available in the Fax.to coverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncomingFaxesGet**
> InlineResponse2004 IncomingFaxesGet(ctx, apiKey, optional)
This API gets a list of incoming faxes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 
 **optional** | ***IncomingFaxesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IncomingFaxesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.String**| The number of record to return | 
 **page** | **optional.String**| The page you want to get | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncomingFaxesRecipientGet**
> InlineResponse2004 IncomingFaxesRecipientGet(ctx, recipient, apiKey, optional)
This API gets a list of incoming faxes for a specific recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int32**| The recipient&#39;s fax number | 
  **apiKey** | **string**|  | 
 **optional** | ***IncomingFaxesRecipientGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IncomingFaxesRecipientGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.String**| The number of record to return | 
 **page** | **optional.String**| The page you want to get | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NumbersGet**
> InlineResponse20017 NumbersGet(ctx, apiKey, optional)
This API gets a list of numbers to get the current configuration of one or multiple number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 
 **optional** | ***NumbersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NumbersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.String**| The number of record to return | 
 **page** | **optional.String**| The page you want to get | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProvisionNumbersGet**
> InlineResponse20016 ProvisionNumbersGet(ctx, didGroupId, apiKey)
This API gets a list of provisioned numbers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **didGroupId** | **int32**| The id of the did group | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatesCountryCodeGet**
> InlineResponse20013 StatesCountryCodeGet(ctx, cOUNTRYCODE, apiKey)
This API gets a list of states of a given country available in the Fax.to coverage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cOUNTRYCODE** | **int32**| Indicates the country code in its ISO 3166-1 alpha-3 format | 
  **apiKey** | **string**|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

