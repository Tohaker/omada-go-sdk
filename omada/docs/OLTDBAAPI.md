# OLTDBAAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDBAProfile**](OLTDBAAPI.md#adddbaprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/dba/profiles/add | Create new DBA profile
[**DeleteDBAProfile**](OLTDBAAPI.md#deletedbaprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/dba/profiles/delete | Delete an existing DBA profile
[**EditDBAProfile**](OLTDBAAPI.md#editdbaprofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/dba/profiles/edit | Modify an existing DBA profile
[**GetDBAProfileList**](OLTDBAAPI.md#getdbaprofilelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/dba/profiles/list | Get DBA profile list
[**GetDBAProfilePage**](OLTDBAAPI.md#getdbaprofilepage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/dba/profiles/page | Get DBA profile page



## AddDBAProfile

> OperationResponseDeviceResponseBodyDBAProfileDTO AddDBAProfile(ctx, omadacId, siteId, deviceMac).DBAProfileModifyDTO(dBAProfileModifyDTO).Execute()

Create new DBA profile



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dBAProfileModifyDTO := *openapiclient.NewDBAProfileModifyDTO(int32(123)) // DBAProfileModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTDBAAPI.AddDBAProfile(context.Background(), omadacId, siteId, deviceMac).DBAProfileModifyDTO(dBAProfileModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTDBAAPI.AddDBAProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDBAProfile`: OperationResponseDeviceResponseBodyDBAProfileDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTDBAAPI.AddDBAProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDBAProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dBAProfileModifyDTO** | [**DBAProfileModifyDTO**](DBAProfileModifyDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyDBAProfileDTO**](OperationResponseDeviceResponseBodyDBAProfileDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDBAProfile

> OperationResponseDeviceResponseBodyDBAProfileDeleteResultDTO DeleteDBAProfile(ctx, omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()

Delete an existing DBA profile



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	intIdListRequest := *openapiclient.NewIntIdListRequest([]int32{int32(123)}) // IntIdListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTDBAAPI.DeleteDBAProfile(context.Background(), omadacId, siteId, deviceMac).IntIdListRequest(intIdListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTDBAAPI.DeleteDBAProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDBAProfile`: OperationResponseDeviceResponseBodyDBAProfileDeleteResultDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTDBAAPI.DeleteDBAProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDBAProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **intIdListRequest** | [**IntIdListRequest**](IntIdListRequest.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyDBAProfileDeleteResultDTO**](OperationResponseDeviceResponseBodyDBAProfileDeleteResultDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDBAProfile

> OperationResponseDeviceResponseBodyVoid EditDBAProfile(ctx, omadacId, siteId, deviceMac).DBAProfileModifyDTO(dBAProfileModifyDTO).Execute()

Modify an existing DBA profile



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dBAProfileModifyDTO := *openapiclient.NewDBAProfileModifyDTO(int32(123)) // DBAProfileModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTDBAAPI.EditDBAProfile(context.Background(), omadacId, siteId, deviceMac).DBAProfileModifyDTO(dBAProfileModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTDBAAPI.EditDBAProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDBAProfile`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTDBAAPI.EditDBAProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDBAProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dBAProfileModifyDTO** | [**DBAProfileModifyDTO**](DBAProfileModifyDTO.md) |  | 

### Return type

[**OperationResponseDeviceResponseBodyVoid**](OperationResponseDeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDBAProfileList

> OperationResponseListDBAProfileVO GetDBAProfileList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get DBA profile list



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dto := *openapiclient.NewDBAProfileListQueryDTO() // DBAProfileListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTDBAAPI.GetDBAProfileList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTDBAAPI.GetDBAProfileList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDBAProfileList`: OperationResponseListDBAProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTDBAAPI.GetDBAProfileList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDBAProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**DBAProfileListQueryDTO**](DBAProfileListQueryDTO.md) |  | 

### Return type

[**OperationResponseListDBAProfileVO**](OperationResponseListDBAProfileVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDBAProfilePage

> OperationResponsePageResponseDBAProfileVO GetDBAProfilePage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get DBA profile page



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
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	dto := *openapiclient.NewDBAProfilePageQueryDTO() // DBAProfilePageQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTDBAAPI.GetDBAProfilePage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTDBAAPI.GetDBAProfilePage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDBAProfilePage`: OperationResponsePageResponseDBAProfileVO
	fmt.Fprintf(os.Stdout, "Response from `OLTDBAAPI.GetDBAProfilePage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDBAProfilePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**DBAProfilePageQueryDTO**](DBAProfilePageQueryDTO.md) |  | 

### Return type

[**OperationResponsePageResponseDBAProfileVO**](OperationResponsePageResponseDBAProfileVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

