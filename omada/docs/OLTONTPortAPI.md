# OLTONTPortAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditOntEthPort**](OLTONTPortAPI.md#editontethport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/{serviceId}/eth-ports/edit | Modify an existing ont eth port
[**EditOntPotsPort**](OLTONTPortAPI.md#editontpotsport) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/{serviceId}/pots-ports/edit | Modify an existing ont pots port
[**GetOntEthPortList**](OLTONTPortAPI.md#getontethportlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/{serviceId}/eth-ports/list | Get ont eth port list
[**GetOntPotsPortList**](OLTONTPortAPI.md#getontpotsportlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/profile/service/{serviceId}/pots-ports/list | Get ont pots port list



## EditOntEthPort

> OperationResponseDeviceResponseBodyVoid EditOntEthPort(ctx, omadacId, siteId, deviceMac, serviceId).OntEthPortModifyDTO(ontEthPortModifyDTO).Execute()

Modify an existing ont eth port



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
	serviceId := "serviceId_example" // string | Service Profile ID should be a number between 0-512
	ontEthPortModifyDTO := *openapiclient.NewOntEthPortModifyDTO(int32(123)) // OntEthPortModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONTPortAPI.EditOntEthPort(context.Background(), omadacId, siteId, deviceMac, serviceId).OntEthPortModifyDTO(ontEthPortModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONTPortAPI.EditOntEthPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOntEthPort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONTPortAPI.EditOntEthPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**serviceId** | **string** | Service Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOntEthPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ontEthPortModifyDTO** | [**OntEthPortModifyDTO**](OntEthPortModifyDTO.md) |  | 

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


## EditOntPotsPort

> OperationResponseDeviceResponseBodyVoid EditOntPotsPort(ctx, omadacId, siteId, deviceMac, serviceId).OntPotsPortModifyDTO(ontPotsPortModifyDTO).Execute()

Modify an existing ont pots port



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
	serviceId := "serviceId_example" // string | Service Profile ID should be a number between 0-512
	ontPotsPortModifyDTO := *openapiclient.NewOntPotsPortModifyDTO("VlanConfigMode_example") // OntPotsPortModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONTPortAPI.EditOntPotsPort(context.Background(), omadacId, siteId, deviceMac, serviceId).OntPotsPortModifyDTO(ontPotsPortModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONTPortAPI.EditOntPotsPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOntPotsPort`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONTPortAPI.EditOntPotsPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**serviceId** | **string** | Service Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditOntPotsPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ontPotsPortModifyDTO** | [**OntPotsPortModifyDTO**](OntPotsPortModifyDTO.md) |  | 

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


## GetOntEthPortList

> OperationResponseListOntEthPortDTO GetOntEthPortList(ctx, omadacId, siteId, deviceMac, serviceId).QueryParam(queryParam).Execute()

Get ont eth port list



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
	serviceId := "serviceId_example" // string | Service Profile ID should be a number between 0-512
	queryParam := *openapiclient.NewOntEthPortListQueryDTO() // OntEthPortListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONTPortAPI.GetOntEthPortList(context.Background(), omadacId, siteId, deviceMac, serviceId).QueryParam(queryParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONTPortAPI.GetOntEthPortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOntEthPortList`: OperationResponseListOntEthPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONTPortAPI.GetOntEthPortList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**serviceId** | **string** | Service Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOntEthPortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queryParam** | [**OntEthPortListQueryDTO**](OntEthPortListQueryDTO.md) |  | 

### Return type

[**OperationResponseListOntEthPortDTO**](OperationResponseListOntEthPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOntPotsPortList

> OperationResponseListOntPotsPortDTO GetOntPotsPortList(ctx, omadacId, siteId, deviceMac, serviceId).QueryParam(queryParam).Execute()

Get ont pots port list



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
	serviceId := "serviceId_example" // string | Service Profile ID should be a number between 0-512
	queryParam := *openapiclient.NewOntPotsPortListQueryDTO() // OntPotsPortListQueryDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONTPortAPI.GetOntPotsPortList(context.Background(), omadacId, siteId, deviceMac, serviceId).QueryParam(queryParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONTPortAPI.GetOntPotsPortList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOntPotsPortList`: OperationResponseListOntPotsPortDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONTPortAPI.GetOntPotsPortList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**serviceId** | **string** | Service Profile ID should be a number between 0-512 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOntPotsPortListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **queryParam** | [**OntPotsPortListQueryDTO**](OntPotsPortListQueryDTO.md) |  | 

### Return type

[**OperationResponseListOntPotsPortDTO**](OperationResponseListOntPotsPortDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

