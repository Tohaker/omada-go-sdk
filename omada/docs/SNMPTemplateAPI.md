# SNMPTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSnmpSettingTemplate**](SNMPTemplateAPI.md#getsnmpsettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/snmp | Get SNMP template setting
[**ModifySnmpSettingTemplate**](SNMPTemplateAPI.md#modifysnmpsettingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/setting/service/snmp | Modify SNMP template setting



## GetSnmpSettingTemplate

> OperationResponseSnmpSettingOpenApiVO GetSnmpSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get SNMP template setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SNMPTemplateAPI.GetSnmpSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SNMPTemplateAPI.GetSnmpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnmpSettingTemplate`: OperationResponseSnmpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `SNMPTemplateAPI.GetSnmpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnmpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSnmpSettingOpenApiVO**](OperationResponseSnmpSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySnmpSettingTemplate

> OperationResponseWithoutResult ModifySnmpSettingTemplate(ctx, omadacId, siteTemplateId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()

Modify SNMP template setting



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Tohaker/omada-go-sdk/omada"
)

func main() {
	omadacId := "omadacId_example" // string | Omada ID
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	snmpSettingOpenApiVO := *openapiclient.NewSnmpSettingOpenApiVO(false, false) // SnmpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SNMPTemplateAPI.ModifySnmpSettingTemplate(context.Background(), omadacId, siteTemplateId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SNMPTemplateAPI.ModifySnmpSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySnmpSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SNMPTemplateAPI.ModifySnmpSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySnmpSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **snmpSettingOpenApiVO** | [**SnmpSettingOpenApiVO**](SnmpSettingOpenApiVO.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

