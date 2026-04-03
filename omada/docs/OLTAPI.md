# OLTAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditConfig**](OLTAPI.md#editconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/system/system-info/configs/edit | Modify OLT description info
[**EditConfigForMsp**](OLTAPI.md#editconfigformsp) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/olts/{deviceMac}/system/system-info/configs/edit | Modify OLT description info(MSP mode)
[**EditLagsConfigList**](OLTAPI.md#editlagsconfiglist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/l2-feature/eth-port/port/lags/edit | Batch modify ETH LAG(Link Aggregation Group) config
[**EditManagementSystemInterfaces**](OLTAPI.md#editmanagementsysteminterfaces) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/system/management-system-interface/edit | Modify OLT Management System Interface
[**EditPortsConfigList**](OLTAPI.md#editportsconfiglist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/l2-feature/eth-port/port/unit1/edit | Batch Modify ETH Unit1 port of olt
[**GetConfig**](OLTAPI.md#getconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/system/system-info/configs | Get OLT description info
[**GetConfigForMsp**](OLTAPI.md#getconfigformsp) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/olts/{deviceMac}/system/system-info/configs | Get OLT description info(MSP mode)
[**GetDetail**](OLTAPI.md#getdetail) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac} | Get OLT info
[**GetEthAndPonList**](OLTAPI.md#getethandponlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/eth-port/and/pon-port/list | Get pon port and eth port of OLT
[**GetEthAndPonListForMsp**](OLTAPI.md#getethandponlistformsp) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/eth-port/and/pon-port/list | Get pon port and eth port of OLT(MSP mode)
[**GetLagsConfigList**](OLTAPI.md#getlagsconfiglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/l2-feature/eth-port/port/lags/list | Get ETH LAG(Link Aggregation Group) config list
[**GetManagementSystemInterfaces**](OLTAPI.md#getmanagementsysteminterfaces) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/system/management-system-interface/get | Get OLT Management System Interface list
[**GetOltDetailForMsp**](OLTAPI.md#getoltdetailformsp) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/olts/{deviceMac} | Get OLT info(MSP mode)
[**GetOltPortsDDMStatus**](OLTAPI.md#getoltportsddmstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/stat/olt/{deviceMac}/ddm/info | Get DDM(Digital Diagnostic Monitoring) status list of OLT port list
[**GetPortsConfigList**](OLTAPI.md#getportsconfiglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/l2-feature/eth-port/port/unit1/list | Get ETH Unit1 port list of olt
[**GetPortsConfigListForMsp**](OLTAPI.md#getportsconfiglistformsp) | **Get** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/olts/{deviceMac}/l2-feature/eth-port/port/unit1/list | Get ETH Unit1 port list of olt(MSP mode)
[**ModifyOltConfig**](OLTAPI.md#modifyoltconfig) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac} | Modify olt config
[**ModifyOltConfigForMsp**](OLTAPI.md#modifyoltconfigformsp) | **Patch** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/olts/{deviceMac} | Modify olt config(MSP mode)
[**MoveToSite**](OLTAPI.md#movetosite) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/olts/move | Move OLT to another site
[**MspMoveToCustomer1**](OLTAPI.md#mspmovetocustomer1) | **Post** /openapi/v1/msp/{mspId}/customers/{customerId}/sites/{siteId}/cmd/olts/move | Move OLT to another customer(MSP mode)
[**RebootNow**](OLTAPI.md#rebootnow) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/system/system-tools/reboot/now | Reboot OLT



## EditConfig

> DeviceResponseBodyVoid EditConfig(ctx, omadacId, siteId, deviceMac).SystemInfoAppModifyDTO(systemInfoAppModifyDTO).Execute()

Modify OLT description info



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
	systemInfoAppModifyDTO := *openapiclient.NewSystemInfoAppModifyDTO() // SystemInfoAppModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.EditConfig(context.Background(), omadacId, siteId, deviceMac).SystemInfoAppModifyDTO(systemInfoAppModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.EditConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditConfig`: DeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.EditConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **systemInfoAppModifyDTO** | [**SystemInfoAppModifyDTO**](SystemInfoAppModifyDTO.md) |  | 

### Return type

[**DeviceResponseBodyVoid**](DeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditConfigForMsp

> DeviceResponseBodyVoid EditConfigForMsp(ctx, siteId, deviceMac, mspId, customerId).SystemInfoAppModifyDTO(systemInfoAppModifyDTO).Execute()

Modify OLT description info(MSP mode)



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
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId
	systemInfoAppModifyDTO := *openapiclient.NewSystemInfoAppModifyDTO() // SystemInfoAppModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.EditConfigForMsp(context.Background(), siteId, deviceMac, mspId, customerId).SystemInfoAppModifyDTO(systemInfoAppModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.EditConfigForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditConfigForMsp`: DeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.EditConfigForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditConfigForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **systemInfoAppModifyDTO** | [**SystemInfoAppModifyDTO**](SystemInfoAppModifyDTO.md) |  | 

### Return type

[**DeviceResponseBodyVoid**](DeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditLagsConfigList

> DeviceResponseBodyVoid EditLagsConfigList(ctx, omadacId, siteId, deviceMac).EthLagPortAppListDTO(ethLagPortAppListDTO).Execute()

Batch modify ETH LAG(Link Aggregation Group) config



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
	ethLagPortAppListDTO := *openapiclient.NewEthLagPortAppListDTO([]openapiclient.EthLagPortAppDTO{*openapiclient.NewEthLagPortAppDTO("Lag_example")}) // EthLagPortAppListDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.EditLagsConfigList(context.Background(), omadacId, siteId, deviceMac).EthLagPortAppListDTO(ethLagPortAppListDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.EditLagsConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditLagsConfigList`: DeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.EditLagsConfigList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditLagsConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ethLagPortAppListDTO** | [**EthLagPortAppListDTO**](EthLagPortAppListDTO.md) |  | 

### Return type

[**DeviceResponseBodyVoid**](DeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditManagementSystemInterfaces

> OperationResponseWithoutResult EditManagementSystemInterfaces(ctx, omadacId, siteId, deviceMac).ManagementSystemInterfaceDTO(managementSystemInterfaceDTO).Execute()

Modify OLT Management System Interface



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
	managementSystemInterfaceDTO := *openapiclient.NewManagementSystemInterfaceDTO() // ManagementSystemInterfaceDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.EditManagementSystemInterfaces(context.Background(), omadacId, siteId, deviceMac).ManagementSystemInterfaceDTO(managementSystemInterfaceDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.EditManagementSystemInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditManagementSystemInterfaces`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.EditManagementSystemInterfaces`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditManagementSystemInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **managementSystemInterfaceDTO** | [**ManagementSystemInterfaceDTO**](ManagementSystemInterfaceDTO.md) |  | 

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


## EditPortsConfigList

> DeviceResponseBodyVoid EditPortsConfigList(ctx, omadacId, siteId, deviceMac).EthUnit1PortAppListDTO(ethUnit1PortAppListDTO).Execute()

Batch Modify ETH Unit1 port of olt



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
	ethUnit1PortAppListDTO := *openapiclient.NewEthUnit1PortAppListDTO([]openapiclient.EthUnit1PortAppDTO{*openapiclient.NewEthUnit1PortAppDTO("Port_example")}) // EthUnit1PortAppListDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.EditPortsConfigList(context.Background(), omadacId, siteId, deviceMac).EthUnit1PortAppListDTO(ethUnit1PortAppListDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.EditPortsConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPortsConfigList`: DeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.EditPortsConfigList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditPortsConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ethUnit1PortAppListDTO** | [**EthUnit1PortAppListDTO**](EthUnit1PortAppListDTO.md) |  | 

### Return type

[**DeviceResponseBodyVoid**](DeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> OperationResponseSystemInfoAppDTO GetConfig(ctx, omadacId, siteId, deviceMac).Execute()

Get OLT description info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetConfig(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: OperationResponseSystemInfoAppDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseSystemInfoAppDTO**](OperationResponseSystemInfoAppDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigForMsp

> OperationResponseSystemInfoAppDTO GetConfigForMsp(ctx, siteId, deviceMac, mspId, customerId).Execute()

Get OLT description info(MSP mode)



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
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetConfigForMsp(context.Background(), siteId, deviceMac, mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetConfigForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigForMsp`: OperationResponseSystemInfoAppDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetConfigForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseSystemInfoAppDTO**](OperationResponseSystemInfoAppDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetail

> OperationResponseOltDetailVO GetDetail(ctx, omadacId, siteId, deviceMac).Execute()

Get OLT info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetDetail(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetail`: OperationResponseOltDetailVO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetDetail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseOltDetailVO**](OperationResponseOltDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthAndPonList

> OperationResponseEthAndPonListDTO GetEthAndPonList(ctx, omadacId, siteId, deviceMac).Execute()

Get pon port and eth port of OLT



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetEthAndPonList(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetEthAndPonList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthAndPonList`: OperationResponseEthAndPonListDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetEthAndPonList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetEthAndPonListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseEthAndPonListDTO**](OperationResponseEthAndPonListDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEthAndPonListForMsp

> OperationResponseEthAndPonListDTO GetEthAndPonListForMsp(ctx, mspId, customerId, siteId).Execute()

Get pon port and eth port of OLT(MSP mode)



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
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetEthAndPonListForMsp(context.Background(), mspId, customerId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetEthAndPonListForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthAndPonListForMsp`: OperationResponseEthAndPonListDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetEthAndPonListForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEthAndPonListForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseEthAndPonListDTO**](OperationResponseEthAndPonListDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLagsConfigList

> OperationResponseListEthLagPortAppDTO GetLagsConfigList(ctx, omadacId, siteId, deviceMac).Execute()

Get ETH LAG(Link Aggregation Group) config list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetLagsConfigList(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetLagsConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLagsConfigList`: OperationResponseListEthLagPortAppDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetLagsConfigList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetLagsConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListEthLagPortAppDTO**](OperationResponseListEthLagPortAppDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagementSystemInterfaces

> OperationResponseManagementSystemInterfaceDTO GetManagementSystemInterfaces(ctx, omadacId, siteId, deviceMac).Execute()

Get OLT Management System Interface list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetManagementSystemInterfaces(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetManagementSystemInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagementSystemInterfaces`: OperationResponseManagementSystemInterfaceDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetManagementSystemInterfaces`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetManagementSystemInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseManagementSystemInterfaceDTO**](OperationResponseManagementSystemInterfaceDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOltDetailForMsp

> OperationResponseOltDetailVO GetOltDetailForMsp(ctx, siteId, deviceMac, mspId, customerId).Execute()

Get OLT info(MSP mode)



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
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetOltDetailForMsp(context.Background(), siteId, deviceMac, mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetOltDetailForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOltDetailForMsp`: OperationResponseOltDetailVO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetOltDetailForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOltDetailForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseOltDetailVO**](OperationResponseOltDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOltPortsDDMStatus

> OperationResponseListDDMStatusResultVO GetOltPortsDDMStatus(ctx, omadacId, siteId, deviceMac).DDMStatusRequestVO(dDMStatusRequestVO).Execute()

Get DDM(Digital Diagnostic Monitoring) status list of OLT port list



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
	dDMStatusRequestVO := *openapiclient.NewDDMStatusRequestVO() // DDMStatusRequestVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetOltPortsDDMStatus(context.Background(), omadacId, siteId, deviceMac).DDMStatusRequestVO(dDMStatusRequestVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetOltPortsDDMStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOltPortsDDMStatus`: OperationResponseListDDMStatusResultVO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetOltPortsDDMStatus`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOltPortsDDMStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dDMStatusRequestVO** | [**DDMStatusRequestVO**](DDMStatusRequestVO.md) |  | 

### Return type

[**OperationResponseListDDMStatusResultVO**](OperationResponseListDDMStatusResultVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortsConfigList

> OperationResponseListEthUnit1PortAppDTO GetPortsConfigList(ctx, omadacId, siteId, deviceMac).Execute()

Get ETH Unit1 port list of olt



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetPortsConfigList(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetPortsConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortsConfigList`: OperationResponseListEthUnit1PortAppDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetPortsConfigList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPortsConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseListEthUnit1PortAppDTO**](OperationResponseListEthUnit1PortAppDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortsConfigListForMsp

> OperationResponseListEthUnit1PortAppDTO GetPortsConfigListForMsp(ctx, siteId, deviceMac, mspId, customerId).Execute()

Get ETH Unit1 port list of olt(MSP mode)



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
	siteId := "siteId_example" // string | Site ID
	deviceMac := "deviceMac_example" // string | Device MAC address, like AA-BB-CC-DD-EE-FF
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.GetPortsConfigListForMsp(context.Background(), siteId, deviceMac, mspId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.GetPortsConfigListForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortsConfigListForMsp`: OperationResponseListEthUnit1PortAppDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.GetPortsConfigListForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**deviceMac** | **string** | Device MAC address, like AA-BB-CC-DD-EE-FF | 
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortsConfigListForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**OperationResponseListEthUnit1PortAppDTO**](OperationResponseListEthUnit1PortAppDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOltConfig

> OperationResponseOltDetailVO ModifyOltConfig(ctx, omadacId, siteId, deviceMac).OltConfigModifyDTO(oltConfigModifyDTO).Execute()

Modify olt config



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
	oltConfigModifyDTO := *openapiclient.NewOltConfigModifyDTO() // OltConfigModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.ModifyOltConfig(context.Background(), omadacId, siteId, deviceMac).OltConfigModifyDTO(oltConfigModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.ModifyOltConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOltConfig`: OperationResponseOltDetailVO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.ModifyOltConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyOltConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oltConfigModifyDTO** | [**OltConfigModifyDTO**](OltConfigModifyDTO.md) |  | 

### Return type

[**OperationResponseOltDetailVO**](OperationResponseOltDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOltConfigForMsp

> OperationResponseOltDetailVO ModifyOltConfigForMsp(ctx, siteId, mspId, customerId, deviceMac).OltConfigModifyDTO(oltConfigModifyDTO).Execute()

Modify olt config(MSP mode)



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
	siteId := "siteId_example" // string | Site ID
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId
	deviceMac := "deviceMac_example" // string | deviceMac
	oltConfigModifyDTO := *openapiclient.NewOltConfigModifyDTO() // OltConfigModifyDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.ModifyOltConfigForMsp(context.Background(), siteId, mspId, customerId, deviceMac).OltConfigModifyDTO(oltConfigModifyDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.ModifyOltConfigForMsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOltConfigForMsp`: OperationResponseOltDetailVO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.ModifyOltConfigForMsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 
**deviceMac** | **string** | deviceMac | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOltConfigForMspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **oltConfigModifyDTO** | [**OltConfigModifyDTO**](OltConfigModifyDTO.md) |  | 

### Return type

[**OperationResponseOltDetailVO**](OperationResponseOltDetailVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveToSite

> OperationResponseMoveSiteProcessVO MoveToSite(ctx, omadacId, siteId).MoveToSiteVO(moveToSiteVO).Execute()

Move OLT to another site



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
	moveToSiteVO := *openapiclient.NewMoveToSiteVO() // MoveToSiteVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.MoveToSite(context.Background(), omadacId, siteId).MoveToSiteVO(moveToSiteVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.MoveToSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveToSite`: OperationResponseMoveSiteProcessVO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.MoveToSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveToSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **moveToSiteVO** | [**MoveToSiteVO**](MoveToSiteVO.md) |  | 

### Return type

[**OperationResponseMoveSiteProcessVO**](OperationResponseMoveSiteProcessVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MspMoveToCustomer1

> DeviceResponseBodyVoid MspMoveToCustomer1(ctx, siteId, mspId, customerId).MoveToCustomerVO(moveToCustomerVO).Execute()

Move OLT to another customer(MSP mode)



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
	siteId := "siteId_example" // string | Site ID
	mspId := "mspId_example" // string | mspId
	customerId := "customerId_example" // string | customerId
	moveToCustomerVO := *openapiclient.NewMoveToCustomerVO("Customer_example", []string{"DeviceMacs_example"}) // MoveToCustomerVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.MspMoveToCustomer1(context.Background(), siteId, mspId, customerId).MoveToCustomerVO(moveToCustomerVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.MspMoveToCustomer1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MspMoveToCustomer1`: DeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.MspMoveToCustomer1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** | Site ID | 
**mspId** | **string** | mspId | 
**customerId** | **string** | customerId | 

### Other Parameters

Other parameters are passed through a pointer to a apiMspMoveToCustomer1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **moveToCustomerVO** | [**MoveToCustomerVO**](MoveToCustomerVO.md) |  | 

### Return type

[**DeviceResponseBodyVoid**](DeviceResponseBodyVoid.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootNow

> OperationResponseDeviceRebootAppDTO RebootNow(ctx, omadacId, siteId, deviceMac).Execute()

Reboot OLT



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTAPI.RebootNow(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTAPI.RebootNow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootNow`: OperationResponseDeviceRebootAppDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTAPI.RebootNow`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiRebootNowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseDeviceRebootAppDTO**](OperationResponseDeviceRebootAppDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

