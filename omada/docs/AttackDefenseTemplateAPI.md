# AttackDefenseTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttackDefenseSettingTemplate**](AttackDefenseTemplateAPI.md#getattackdefensesettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/attack-defense | Get attack and defense template setting
[**ModifyAttackDefenseSettingTemplate**](AttackDefenseTemplateAPI.md#modifyattackdefensesettingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/attack-defense | Modify attack and defense template setting
[**ResetAttackDefenseSettingTemplate**](AttackDefenseTemplateAPI.md#resetattackdefensesettingtemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/attack-defense/reset | Reset attack and defense template setting



## GetAttackDefenseSettingTemplate

> OperationResponseAttackDefenseSetting GetAttackDefenseSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get attack and defense template setting



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDefenseTemplateAPI.GetAttackDefenseSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDefenseTemplateAPI.GetAttackDefenseSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttackDefenseSettingTemplate`: OperationResponseAttackDefenseSetting
	fmt.Fprintf(os.Stdout, "Response from `AttackDefenseTemplateAPI.GetAttackDefenseSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttackDefenseSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseAttackDefenseSetting**](OperationResponseAttackDefenseSetting.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAttackDefenseSettingTemplate

> OperationResponseWithoutResult ModifyAttackDefenseSettingTemplate(ctx, omadacId, siteTemplateId).AttackDefenseSetting(attackDefenseSetting).Execute()

Modify attack and defense template setting



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID
	attackDefenseSetting := *openapiclient.NewAttackDefenseSetting(false, false, false, false, false, false, false, false, false, false, false, false, false, false) // AttackDefenseSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDefenseTemplateAPI.ModifyAttackDefenseSettingTemplate(context.Background(), omadacId, siteTemplateId).AttackDefenseSetting(attackDefenseSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDefenseTemplateAPI.ModifyAttackDefenseSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAttackDefenseSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AttackDefenseTemplateAPI.ModifyAttackDefenseSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAttackDefenseSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attackDefenseSetting** | [**AttackDefenseSetting**](AttackDefenseSetting.md) |  | 

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


## ResetAttackDefenseSettingTemplate

> OperationResponseWithoutResult ResetAttackDefenseSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Reset attack and defense template setting



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDefenseTemplateAPI.ResetAttackDefenseSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDefenseTemplateAPI.ResetAttackDefenseSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAttackDefenseSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AttackDefenseTemplateAPI.ResetAttackDefenseSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetAttackDefenseSettingTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

