# \AttackDefenseAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttackDefenseSetting**](AttackDefenseAPI.md#GetAttackDefenseSetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/attack-defense | Get attack and defense setting
[**GetDefaultAttackDefense**](AttackDefenseAPI.md#GetDefaultAttackDefense) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/attack-defense/default | Get attack and defense default setting
[**ModifyAttackDefenseSetting**](AttackDefenseAPI.md#ModifyAttackDefenseSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/attack-defense | Modify attack and defense setting
[**ResetAttackDefenseSetting**](AttackDefenseAPI.md#ResetAttackDefenseSetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/attack-defense/reset | Reset attack and defense setting



## GetAttackDefenseSetting

> OperationResponseAttackDefenseSettingForQuery GetAttackDefenseSetting(ctx, omadacId, siteId).Execute()

Get attack and defense setting



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
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDefenseAPI.GetAttackDefenseSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDefenseAPI.GetAttackDefenseSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttackDefenseSetting`: OperationResponseAttackDefenseSettingForQuery
	fmt.Fprintf(os.Stdout, "Response from `AttackDefenseAPI.GetAttackDefenseSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttackDefenseSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseAttackDefenseSettingForQuery**](OperationResponseAttackDefenseSettingForQuery.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultAttackDefense

> OperationResponseAttackDefenseSetting GetDefaultAttackDefense(ctx, omadacId, siteId).Execute()

Get attack and defense default setting



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
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDefenseAPI.GetDefaultAttackDefense(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDefenseAPI.GetDefaultAttackDefense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultAttackDefense`: OperationResponseAttackDefenseSetting
	fmt.Fprintf(os.Stdout, "Response from `AttackDefenseAPI.GetDefaultAttackDefense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultAttackDefenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseAttackDefenseSetting**](OperationResponseAttackDefenseSetting.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAttackDefenseSetting

> OperationResponseWithoutResult ModifyAttackDefenseSetting(ctx, omadacId, siteId).AttackDefenseSetting(attackDefenseSetting).Execute()

Modify attack and defense setting



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
	siteId := "siteId_example" // string | Site ID
	attackDefenseSetting := *openapiclient.NewAttackDefenseSetting(false, false, false, false, false, false, false, false, false, false, false, false, false, false) // AttackDefenseSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDefenseAPI.ModifyAttackDefenseSetting(context.Background(), omadacId, siteId).AttackDefenseSetting(attackDefenseSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDefenseAPI.ModifyAttackDefenseSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAttackDefenseSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AttackDefenseAPI.ModifyAttackDefenseSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAttackDefenseSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attackDefenseSetting** | [**AttackDefenseSetting**](AttackDefenseSetting.md) |  | 

### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAttackDefenseSetting

> OperationResponseWithoutResult ResetAttackDefenseSetting(ctx, omadacId, siteId).Execute()

Reset attack and defense setting



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
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDefenseAPI.ResetAttackDefenseSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDefenseAPI.ResetAttackDefenseSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAttackDefenseSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `AttackDefenseAPI.ResetAttackDefenseSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetAttackDefenseSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

