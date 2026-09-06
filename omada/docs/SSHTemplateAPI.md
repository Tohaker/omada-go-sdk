# SSHTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSshSettingTemplate**](SSHTemplateAPI.md#getsshsettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ssh | Get SSH template setting
[**UpdateSshSettingTemplate**](SSHTemplateAPI.md#updatesshsettingtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/ssh | Modify SSH template setting



## GetSshSettingTemplate

> OperationResponseSSHSetting GetSshSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get SSH template setting



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
	resp, r, err := apiClient.SSHTemplateAPI.GetSshSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHTemplateAPI.GetSshSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshSettingTemplate`: OperationResponseSSHSetting
	fmt.Fprintf(os.Stdout, "Response from `SSHTemplateAPI.GetSshSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseSSHSetting**](OperationResponseSSHSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSshSettingTemplate

> OperationResponseWithoutResult UpdateSshSettingTemplate(ctx, omadacId, siteTemplateId).SSHSetting(sSHSetting).Execute()

Modify SSH template setting



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
	sSHSetting := *openapiclient.NewSSHSetting(false) // SSHSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHTemplateAPI.UpdateSshSettingTemplate(context.Background(), omadacId, siteTemplateId).SSHSetting(sSHSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHTemplateAPI.UpdateSshSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSshSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `SSHTemplateAPI.UpdateSshSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSshSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sSHSetting** | [**SSHSetting**](SSHSetting.md) |  | 

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

