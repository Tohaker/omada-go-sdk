# \OSPFAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteOspfInterface**](OSPFAPI.md#BatchDeleteOspfInterface) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ospf/interface/batch-delete | Batch Delete Ospf Interface
[**BatchDeleteOspfProcess**](OSPFAPI.md#BatchDeleteOspfProcess) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ospf/process/batch-delete | Batch Delete Ospf Process
[**CreateOspfInterface**](OSPFAPI.md#CreateOspfInterface) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ospf/interface | Create Ospf Interface
[**CreateOspfProcess**](OSPFAPI.md#CreateOspfProcess) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/ospf/process | Create Ospf Process
[**DeleteOspfInterface**](OSPFAPI.md#DeleteOspfInterface) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/interface/{ospfInterfaceId} | Delete Ospf Interface
[**DeleteOspfProcess**](OSPFAPI.md#DeleteOspfProcess) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/switches/{switchMac}/process/{ospfProcessId} | Delete Ospf Process
[**GetGridOspfInterface**](OSPFAPI.md#GetGridOspfInterface) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ospf/interface | Get Ospf Interface list
[**GetGridOspfProcess**](OSPFAPI.md#GetGridOspfProcess) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ospf/process | Get Ospf Process list
[**GetOspfDevice**](OSPFAPI.md#GetOspfDevice) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ospf/device | Get Ospf Device list
[**GetOswVlanIf**](OSPFAPI.md#GetOswVlanIf) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/vlan-interface/switches/{switchMac} | Get Switch Vlan Interface list
[**ModifyOspfInterface**](OSPFAPI.md#ModifyOspfInterface) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/ospf/interface/{ospfInterfaceId} | Modify Ospf Interface
[**ModifyOspfProcess**](OSPFAPI.md#ModifyOspfProcess) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/ospf/process/{ospfProcessId} | Modify Ospf Process



## BatchDeleteOspfInterface

> OperationResponseWithoutResult BatchDeleteOspfInterface(ctx, omadacId, siteId).BatchOspfInterfaceOpenApiVO(batchOspfInterfaceOpenApiVO).Execute()

Batch Delete Ospf Interface



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
	batchOspfInterfaceOpenApiVO := *openapiclient.NewBatchOspfInterfaceOpenApiVO([]string{"InterfaceIdList_example"}, int32(123)) // BatchOspfInterfaceOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.BatchDeleteOspfInterface(context.Background(), omadacId, siteId).BatchOspfInterfaceOpenApiVO(batchOspfInterfaceOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.BatchDeleteOspfInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteOspfInterface`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.BatchDeleteOspfInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteOspfInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchOspfInterfaceOpenApiVO** | [**BatchOspfInterfaceOpenApiVO**](BatchOspfInterfaceOpenApiVO.md) |  | 

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


## BatchDeleteOspfProcess

> OperationResponseWithoutResult BatchDeleteOspfProcess(ctx, omadacId, siteId).BatchOspfProcessOpenApiVO(batchOspfProcessOpenApiVO).Execute()

Batch Delete Ospf Process



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
	batchOspfProcessOpenApiVO := *openapiclient.NewBatchOspfProcessOpenApiVO([]string{"ProcessIdList_example"}, int32(123)) // BatchOspfProcessOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.BatchDeleteOspfProcess(context.Background(), omadacId, siteId).BatchOspfProcessOpenApiVO(batchOspfProcessOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.BatchDeleteOspfProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteOspfProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.BatchDeleteOspfProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteOspfProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchOspfProcessOpenApiVO** | [**BatchOspfProcessOpenApiVO**](BatchOspfProcessOpenApiVO.md) |  | 

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


## CreateOspfInterface

> OperationResponseWithoutResult CreateOspfInterface(ctx, omadacId, siteId).CreateOspfInterfaceRequest(createOspfInterfaceRequest).Execute()

Create Ospf Interface



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
	createOspfInterfaceRequest := *openapiclient.NewCreateOspfInterfaceRequest() // CreateOspfInterfaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.CreateOspfInterface(context.Background(), omadacId, siteId).CreateOspfInterfaceRequest(createOspfInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.CreateOspfInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOspfInterface`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.CreateOspfInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOspfInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOspfInterfaceRequest** | [**CreateOspfInterfaceRequest**](CreateOspfInterfaceRequest.md) |  | 

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


## CreateOspfProcess

> OperationResponseWithoutResult CreateOspfProcess(ctx, omadacId, siteId).CreateOspfProcessRequest(createOspfProcessRequest).Execute()

Create Ospf Process



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
	createOspfProcessRequest := *openapiclient.NewCreateOspfProcessRequest() // CreateOspfProcessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.CreateOspfProcess(context.Background(), omadacId, siteId).CreateOspfProcessRequest(createOspfProcessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.CreateOspfProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOspfProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.CreateOspfProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOspfProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOspfProcessRequest** | [**CreateOspfProcessRequest**](CreateOspfProcessRequest.md) |  | 

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


## DeleteOspfInterface

> OperationResponseWithoutResult DeleteOspfInterface(ctx, omadacId, siteId, switchMac, ospfInterfaceId).UserInfoBriefDTO(userInfoBriefDTO).Execute()

Delete Ospf Interface



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
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF
	ospfInterfaceId := "ospfInterfaceId_example" // string | Ospf Interface ID
	userInfoBriefDTO := *openapiclient.NewUserInfoBriefDTO() // UserInfoBriefDTO |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.DeleteOspfInterface(context.Background(), omadacId, siteId, switchMac, ospfInterfaceId).UserInfoBriefDTO(userInfoBriefDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.DeleteOspfInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOspfInterface`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.DeleteOspfInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**ospfInterfaceId** | **string** | Ospf Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOspfInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userInfoBriefDTO** | [**UserInfoBriefDTO**](UserInfoBriefDTO.md) |  | 

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


## DeleteOspfProcess

> OperationResponseWithoutResult DeleteOspfProcess(ctx, omadacId, siteId, switchMac, ospfProcessId).UserInfoBriefDTO(userInfoBriefDTO).Execute()

Delete Ospf Process



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
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF
	ospfProcessId := "ospfProcessId_example" // string | Ospf Process ID
	userInfoBriefDTO := *openapiclient.NewUserInfoBriefDTO() // UserInfoBriefDTO |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.DeleteOspfProcess(context.Background(), omadacId, siteId, switchMac, ospfProcessId).UserInfoBriefDTO(userInfoBriefDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.DeleteOspfProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOspfProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.DeleteOspfProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 
**ospfProcessId** | **string** | Ospf Process ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOspfProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userInfoBriefDTO** | [**UserInfoBriefDTO**](UserInfoBriefDTO.md) |  | 

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


## GetGridOspfInterface

> OperationResponseGridVOOspfInterfaceOpenApiVO GetGridOspfInterface(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Ospf Interface list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.GetGridOspfInterface(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.GetGridOspfInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridOspfInterface`: OperationResponseGridVOOspfInterfaceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.GetGridOspfInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridOspfInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOspfInterfaceOpenApiVO**](OperationResponseGridVOOspfInterfaceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridOspfProcess

> OperationResponseGridVOOspfProcessOpenApiVO GetGridOspfProcess(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get Ospf Process list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.GetGridOspfProcess(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.GetGridOspfProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridOspfProcess`: OperationResponseGridVOOspfProcessOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.GetGridOspfProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridOspfProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOOspfProcessOpenApiVO**](OperationResponseGridVOOspfProcessOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOspfDevice

> OperationResponseListOspfDeviceOpenApiVO GetOspfDevice(ctx, omadacId, siteId).Execute()

Get Ospf Device list



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
	resp, r, err := apiClient.OSPFAPI.GetOspfDevice(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.GetOspfDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOspfDevice`: OperationResponseListOspfDeviceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.GetOspfDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOspfDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListOspfDeviceOpenApiVO**](OperationResponseListOspfDeviceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswVlanIf

> OperationResponseListVlanInterfaceOpenApiVO GetOswVlanIf(ctx, omadacId, siteId, switchMac).Execute()

Get Switch Vlan Interface list



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
	switchMac := "switchMac_example" // string | Switch MAC address, like AA-BB-CC-DD-EE-FF

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.GetOswVlanIf(context.Background(), omadacId, siteId, switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.GetOswVlanIf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswVlanIf`: OperationResponseListVlanInterfaceOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.GetOswVlanIf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**switchMac** | **string** | Switch MAC address, like AA-BB-CC-DD-EE-FF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswVlanIfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListVlanInterfaceOpenApiVO**](OperationResponseListVlanInterfaceOpenApiVO.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOspfInterface

> OperationResponseWithoutResult ModifyOspfInterface(ctx, omadacId, siteId, ospfInterfaceId).ModifyOspfInterfaceRequest(modifyOspfInterfaceRequest).Execute()

Modify Ospf Interface



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
	ospfInterfaceId := "ospfInterfaceId_example" // string | Ospf Interface ID
	modifyOspfInterfaceRequest := *openapiclient.NewModifyOspfInterfaceRequest() // ModifyOspfInterfaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.ModifyOspfInterface(context.Background(), omadacId, siteId, ospfInterfaceId).ModifyOspfInterfaceRequest(modifyOspfInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.ModifyOspfInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOspfInterface`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.ModifyOspfInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ospfInterfaceId** | **string** | Ospf Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOspfInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyOspfInterfaceRequest** | [**ModifyOspfInterfaceRequest**](ModifyOspfInterfaceRequest.md) |  | 

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


## ModifyOspfProcess

> OperationResponseWithoutResult ModifyOspfProcess(ctx, omadacId, siteId, ospfProcessId).ModifyOspfProcessRequest(modifyOspfProcessRequest).Execute()

Modify Ospf Process



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
	ospfProcessId := "ospfProcessId_example" // string | Ospf Process ID
	modifyOspfProcessRequest := *openapiclient.NewModifyOspfProcessRequest() // ModifyOspfProcessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSPFAPI.ModifyOspfProcess(context.Background(), omadacId, siteId, ospfProcessId).ModifyOspfProcessRequest(modifyOspfProcessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSPFAPI.ModifyOspfProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOspfProcess`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OSPFAPI.ModifyOspfProcess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ospfProcessId** | **string** | Ospf Process ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOspfProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modifyOspfProcessRequest** | [**ModifyOspfProcessRequest**](ModifyOspfProcessRequest.md) |  | 

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

