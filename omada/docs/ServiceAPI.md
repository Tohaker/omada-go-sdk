# ServiceAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchExportDhcpReservationListToFile**](ServiceAPI.md#batchexportdhcpreservationlisttofile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/batch-export | Batch Export DHCP reservation list to file
[**ClearDnsCacheList**](ServiceAPI.md#cleardnscachelist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/dns-cache-data/clear | Clear DNS cache list
[**CreateDdns**](ServiceAPI.md#createddns) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns | Create a new Dynamic DNS entry
[**CreateDhcpReservation**](ServiceAPI.md#createdhcpreservation) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp | Create a new DHCP reservation
[**CreateMdns**](ServiceAPI.md#createmdns) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns | Create a new mDNS rule
[**DeleteDdns**](ServiceAPI.md#deleteddns) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns/{ddnsId} | Delete an exist Dynamic DNS entry
[**DeleteDhcpReservation**](ServiceAPI.md#deletedhcpreservation) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/{mac} | Delete an exist DHCP reservation
[**DeleteMdns**](ServiceAPI.md#deletemdns) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns/{mdnsId} | Delete an exist mDNS rule
[**DownloadDhcpImportResult**](ServiceAPI.md#downloaddhcpimportresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/import-result | Download DHCP import result
[**ExportDhcpReservation**](ServiceAPI.md#exportdhcpreservation) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/{mac}/export | Export DHCP reservation to IP_MAC Binding
[**ExportDhcpReservationListToFile**](ServiceAPI.md#exportdhcpreservationlisttofile) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/export | Export DHCP reservation list to file
[**GetDdnsGrid**](ServiceAPI.md#getddnsgrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns | Get Dynamic DNS list
[**GetDhcpReservationGrid**](ServiceAPI.md#getdhcpreservationgrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp | Get DHCP reservation list
[**GetDnsCacheList**](ServiceAPI.md#getdnscachelist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/dns-cache-data | Get DNS cache list
[**GetDnsCacheSetting**](ServiceAPI.md#getdnscachesetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-cache | Get DNS cache
[**GetDnsProxy**](ServiceAPI.md#getdnsproxy) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-proxy | Get DNS proxy setting
[**GetGridAllDhcpUserList**](ServiceAPI.md#getgridalldhcpuserlist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/user-list | Get the dhcp user list of all servers
[**GetIgmp**](ServiceAPI.md#getigmp) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/igmp | Get IGMP setting
[**GetIptv**](ServiceAPI.md#getiptv) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/iptv | Get IPTV setting
[**GetMdnsGrid**](ServiceAPI.md#getmdnsgrid) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns | Get mDNS rule list
[**GetSnmpSetting**](ServiceAPI.md#getsnmpsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/snmp | Get SNMP setting
[**GetSshSetting**](ServiceAPI.md#getsshsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/ssh | Get SSH setting
[**GetUpnpSetting**](ServiceAPI.md#getupnpsetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/upnp | Get UPnP setting
[**ImportDhcpReservationListFromFile**](ServiceAPI.md#importdhcpreservationlistfromfile) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/files/dhcp/import | Import DHCP reservation list from file
[**ModifyDdns**](ServiceAPI.md#modifyddns) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/ddns/{ddnsId} | Modify an exist Dynamic DNS entry
[**ModifyDhcpReservation**](ServiceAPI.md#modifydhcpreservation) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/{mac} | Modify an exist DHCP reservation
[**ModifyDnsCacheSetting**](ServiceAPI.md#modifydnscachesetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-cache | Modify DNS cache setting
[**ModifyDnsProxy**](ServiceAPI.md#modifydnsproxy) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dns-proxy | Modify DNS proxy setting
[**ModifyIgmp**](ServiceAPI.md#modifyigmp) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/igmp | Modify IGMP setting
[**ModifyIptv**](ServiceAPI.md#modifyiptv) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/iptv | Modify IPTV setting
[**ModifyMdns**](ServiceAPI.md#modifymdns) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/mdns/{mdnsId} | Modify an exist mDNS rule
[**ModifySnmpSetting**](ServiceAPI.md#modifysnmpsetting) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/snmp | Modify SNMP setting
[**ReservationFromDhcpUserList**](ServiceAPI.md#reservationfromdhcpuserlist) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/setting/service/dhcp/user-list/reservation | Reservation From Dhcp User List
[**UpdateSshSetting**](ServiceAPI.md#updatesshsetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/ssh | Modify SSH setting
[**UpdateUpnpSetting**](ServiceAPI.md#updateupnpsetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/upnp | Modify UPnP setting



## BatchExportDhcpReservationListToFile

> OperationResponseWithoutResult BatchExportDhcpReservationListToFile(ctx, omadacId, siteId).DhcpReservationFilterVO(dhcpReservationFilterVO).Execute()

Batch Export DHCP reservation list to file



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
	dhcpReservationFilterVO := *openapiclient.NewDhcpReservationFilterVO(*openapiclient.NewBatchSelectMacsVO("SelectType_example")) // DhcpReservationFilterVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.BatchExportDhcpReservationListToFile(context.Background(), omadacId, siteId).DhcpReservationFilterVO(dhcpReservationFilterVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.BatchExportDhcpReservationListToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchExportDhcpReservationListToFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.BatchExportDhcpReservationListToFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchExportDhcpReservationListToFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpReservationFilterVO** | [**DhcpReservationFilterVO**](DhcpReservationFilterVO.md) |  | 

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


## ClearDnsCacheList

> OperationResponseWithoutResult ClearDnsCacheList(ctx, omadacId, siteId).Execute()

Clear DNS cache list



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
	resp, r, err := apiClient.ServiceAPI.ClearDnsCacheList(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ClearDnsCacheList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearDnsCacheList`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ClearDnsCacheList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearDnsCacheListRequest struct via the builder pattern


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


## CreateDdns

> OperationResponseResIdOpenApiVO CreateDdns(ctx, omadacId, siteId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()

Create a new Dynamic DNS entry



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
	createDdnsOpenApiVO := *openapiclient.NewCreateDdnsOpenApiVO("InterfacePortId_example", "Password_example", int32(123), false, "Username_example") // CreateDdnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.CreateDdns(context.Background(), omadacId, siteId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.CreateDdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDdns`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.CreateDdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDdnsOpenApiVO** | [**CreateDdnsOpenApiVO**](CreateDdnsOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDhcpReservation

> OperationResponseResIdOpenApiVO CreateDhcpReservation(ctx, omadacId, siteId).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()

Create a new DHCP reservation



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
	createDhcpReservationOpenApiVO := *openapiclient.NewCreateDhcpReservationOpenApiVO("Mac_example", "NetId_example", false) // CreateDhcpReservationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.CreateDhcpReservation(context.Background(), omadacId, siteId).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.CreateDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDhcpReservation`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.CreateDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDhcpReservationOpenApiVO** | [**CreateDhcpReservationOpenApiVO**](CreateDhcpReservationOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMdns

> OperationResponseResIdOpenApiVO CreateMdns(ctx, omadacId, siteId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()

Create a new mDNS rule



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
	createMdnsRuleOpenApiVO := *openapiclient.NewCreateMdnsRuleOpenApiVO("Name_example", []string{"ProfileIds_example"}, false) // CreateMdnsRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.CreateMdns(context.Background(), omadacId, siteId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.CreateMdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMdns`: OperationResponseResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.CreateMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMdnsRuleOpenApiVO** | [**CreateMdnsRuleOpenApiVO**](CreateMdnsRuleOpenApiVO.md) |  | 

### Return type

[**OperationResponseResIdOpenApiVO**](OperationResponseResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDdns

> OperationResponseWithoutResult DeleteDdns(ctx, omadacId, siteId, ddnsId).Execute()

Delete an exist Dynamic DNS entry



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
	ddnsId := "ddnsId_example" // string | Dynamic DNS entry ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.DeleteDdns(context.Background(), omadacId, siteId, ddnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.DeleteDdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDdns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.DeleteDdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ddnsId** | **string** | Dynamic DNS entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDdnsRequest struct via the builder pattern


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


## DeleteDhcpReservation

> OperationResponseWithoutResult DeleteDhcpReservation(ctx, omadacId, siteId, mac).Execute()

Delete an exist DHCP reservation



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
	mac := "mac_example" // string | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.DeleteDhcpReservation(context.Background(), omadacId, siteId, mac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.DeleteDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDhcpReservation`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.DeleteDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDhcpReservationRequest struct via the builder pattern


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


## DeleteMdns

> OperationResponseWithoutResult DeleteMdns(ctx, omadacId, siteId, mdnsId).Execute()

Delete an exist mDNS rule



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
	mdnsId := "mdnsId_example" // string | mDNS rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.DeleteMdns(context.Background(), omadacId, siteId, mdnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.DeleteMdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMdns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.DeleteMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mdnsId** | **string** | mDNS rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMdnsRequest struct via the builder pattern


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


## DownloadDhcpImportResult

> OperationResponse DownloadDhcpImportResult(ctx, omadacId, siteId).Execute()

Download DHCP import result



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
	resp, r, err := apiClient.ServiceAPI.DownloadDhcpImportResult(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.DownloadDhcpImportResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDhcpImportResult`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.DownloadDhcpImportResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDhcpImportResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportDhcpReservation

> OperationResponseWithoutResult ExportDhcpReservation(ctx, omadacId, siteId, mac).ExportDhcpReservationOpenApiVO(exportDhcpReservationOpenApiVO).Execute()

Export DHCP reservation to IP_MAC Binding



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
	mac := "mac_example" // string | mac
	exportDhcpReservationOpenApiVO := *openapiclient.NewExportDhcpReservationOpenApiVO("InterfaceId_example") // ExportDhcpReservationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ExportDhcpReservation(context.Background(), omadacId, siteId, mac).ExportDhcpReservationOpenApiVO(exportDhcpReservationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ExportDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDhcpReservation`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ExportDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | mac | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **exportDhcpReservationOpenApiVO** | [**ExportDhcpReservationOpenApiVO**](ExportDhcpReservationOpenApiVO.md) |  | 

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


## ExportDhcpReservationListToFile

> OperationResponseWithoutResult ExportDhcpReservationListToFile(ctx, omadacId, siteId).Execute()

Export DHCP reservation list to file



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
	resp, r, err := apiClient.ServiceAPI.ExportDhcpReservationListToFile(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ExportDhcpReservationListToFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportDhcpReservationListToFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ExportDhcpReservationListToFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDhcpReservationListToFileRequest struct via the builder pattern


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


## GetDdnsGrid

> OperationResponseDdnsOpenApiGridVODdnsOpenApiVO GetDdnsGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsService(sortsService).SortsUpdateInterval(sortsUpdateInterval).SortsStatus(sortsStatus).Execute()

Get Dynamic DNS list



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
	sortsService := "sortsService_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsUpdateInterval := "sortsUpdateInterval_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.GetDdnsGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsService(sortsService).SortsUpdateInterval(sortsUpdateInterval).SortsStatus(sortsStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetDdnsGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDdnsGrid`: OperationResponseDdnsOpenApiGridVODdnsOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetDdnsGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDdnsGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsService** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsUpdateInterval** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseDdnsOpenApiGridVODdnsOpenApiVO**](OperationResponseDdnsOpenApiGridVODdnsOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDhcpReservationGrid

> OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO GetDhcpReservationGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SortsIp(sortsIp).SortsNetName(sortsNetName).SortsDescription(sortsDescription).SortsStatus(sortsStatus).SearchKey(searchKey).Execute()

Get DHCP reservation list



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
	sortsMac := "sortsMac_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsIp := "sortsIp_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsNetName := "sortsNetName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsDescription := "sortsDescription_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsStatus := "sortsStatus_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field : mac, ip, description. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.GetDhcpReservationGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SortsMac(sortsMac).SortsIp(sortsIp).SortsNetName(sortsNetName).SortsDescription(sortsDescription).SortsStatus(sortsStatus).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetDhcpReservationGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDhcpReservationGrid`: OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetDhcpReservationGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDhcpReservationGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsMac** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsIp** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsNetName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsDescription** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsStatus** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field : mac, ip, description. | 

### Return type

[**OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO**](OperationResponseDhcpReservationOpenApiGridVODhcpReservationOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsCacheList

> OperationResponseDnsCacheOpenApiVO GetDnsCacheList(ctx, omadacId, siteId).DnsCacheQueryVO(dnsCacheQueryVO).Execute()

Get DNS cache list



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
	dnsCacheQueryVO := *openapiclient.NewDnsCacheQueryOpenApiV2VO() // DnsCacheQueryOpenApiV2VO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.GetDnsCacheList(context.Background(), omadacId, siteId).DnsCacheQueryVO(dnsCacheQueryVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetDnsCacheList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsCacheList`: OperationResponseDnsCacheOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetDnsCacheList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsCacheListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsCacheQueryVO** | [**DnsCacheQueryOpenApiV2VO**](DnsCacheQueryOpenApiV2VO.md) |  | 

### Return type

[**OperationResponseDnsCacheOpenApiVO**](OperationResponseDnsCacheOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsCacheSetting

> OperationResponseDnsCacheOpenApiVO GetDnsCacheSetting(ctx, omadacId, siteId).Execute()

Get DNS cache



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
	resp, r, err := apiClient.ServiceAPI.GetDnsCacheSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetDnsCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsCacheSetting`: OperationResponseDnsCacheOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetDnsCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDnsCacheOpenApiVO**](OperationResponseDnsCacheOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsProxy

> OperationResponseDnsProxySettingQueryOpenApiVO GetDnsProxy(ctx, omadacId, siteId).Execute()

Get DNS proxy setting



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
	resp, r, err := apiClient.ServiceAPI.GetDnsProxy(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetDnsProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsProxy`: OperationResponseDnsProxySettingQueryOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetDnsProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDnsProxySettingQueryOpenApiVO**](OperationResponseDnsProxySettingQueryOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridAllDhcpUserList

> OperationResponseDhcpUserGridVODhcpUserVO GetGridAllDhcpUserList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get the dhcp user list of all servers



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
	resp, r, err := apiClient.ServiceAPI.GetGridAllDhcpUserList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetGridAllDhcpUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAllDhcpUserList`: OperationResponseDhcpUserGridVODhcpUserVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetGridAllDhcpUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAllDhcpUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseDhcpUserGridVODhcpUserVO**](OperationResponseDhcpUserGridVODhcpUserVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIgmp

> OperationResponseIgmpOpenApiVO GetIgmp(ctx, omadacId, siteId).Execute()

Get IGMP setting



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
	resp, r, err := apiClient.ServiceAPI.GetIgmp(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetIgmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIgmp`: OperationResponseIgmpOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetIgmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIgmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIgmpOpenApiVO**](OperationResponseIgmpOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIptv

> OperationResponseIptvOpenApiVO GetIptv(ctx, omadacId, siteId).Execute()

Get IPTV setting



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
	resp, r, err := apiClient.ServiceAPI.GetIptv(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetIptv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIptv`: OperationResponseIptvOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetIptv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIptvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseIptvOpenApiVO**](OperationResponseIptvOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMdnsGrid

> OperationResponseGridVOMdnsRuleOpenApiVO GetMdnsGrid(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get mDNS rule list



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
	resp, r, err := apiClient.ServiceAPI.GetMdnsGrid(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetMdnsGrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMdnsGrid`: OperationResponseGridVOMdnsRuleOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetMdnsGrid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMdnsGridRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOMdnsRuleOpenApiVO**](OperationResponseGridVOMdnsRuleOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnmpSetting

> OperationResponseSnmpSettingOpenApiVO GetSnmpSetting(ctx, omadacId, siteId).Execute()

Get SNMP setting



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
	resp, r, err := apiClient.ServiceAPI.GetSnmpSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetSnmpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnmpSetting`: OperationResponseSnmpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetSnmpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnmpSettingRequest struct via the builder pattern


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


## GetSshSetting

> OperationResponseSSHSetting GetSshSetting(ctx, omadacId, siteId).Execute()

Get SSH setting



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
	resp, r, err := apiClient.ServiceAPI.GetSshSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetSshSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshSetting`: OperationResponseSSHSetting
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetSshSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshSettingRequest struct via the builder pattern


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


## GetUpnpSetting

> OperationResponseUpnpSettingOpenApiVO GetUpnpSetting(ctx, omadacId, siteId).Execute()

Get UPnP setting



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
	resp, r, err := apiClient.ServiceAPI.GetUpnpSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetUpnpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpnpSetting`: OperationResponseUpnpSettingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetUpnpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpnpSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseUpnpSettingOpenApiVO**](OperationResponseUpnpSettingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportDhcpReservationListFromFile

> OperationResponseWithoutResult ImportDhcpReservationListFromFile(ctx, omadacId, siteId).ImportDhcpReservationListFromFileRequest(importDhcpReservationListFromFileRequest).Execute()

Import DHCP reservation list from file



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
	importDhcpReservationListFromFileRequest := *openapiclient.NewImportDhcpReservationListFromFileRequest() // ImportDhcpReservationListFromFileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ImportDhcpReservationListFromFile(context.Background(), omadacId, siteId).ImportDhcpReservationListFromFileRequest(importDhcpReservationListFromFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ImportDhcpReservationListFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportDhcpReservationListFromFile`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ImportDhcpReservationListFromFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportDhcpReservationListFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **importDhcpReservationListFromFileRequest** | [**ImportDhcpReservationListFromFileRequest**](ImportDhcpReservationListFromFileRequest.md) |  | 

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


## ModifyDdns

> OperationResponseWithoutResult ModifyDdns(ctx, omadacId, siteId, ddnsId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()

Modify an exist Dynamic DNS entry



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
	ddnsId := "ddnsId_example" // string | Dynamic DNS entry ID
	createDdnsOpenApiVO := *openapiclient.NewCreateDdnsOpenApiVO("InterfacePortId_example", "Password_example", int32(123), false, "Username_example") // CreateDdnsOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifyDdns(context.Background(), omadacId, siteId, ddnsId).CreateDdnsOpenApiVO(createDdnsOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifyDdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDdns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifyDdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ddnsId** | **string** | Dynamic DNS entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createDdnsOpenApiVO** | [**CreateDdnsOpenApiVO**](CreateDdnsOpenApiVO.md) |  | 

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


## ModifyDhcpReservation

> OperationResponseWithoutResult ModifyDhcpReservation(ctx, omadacId, siteId, mac).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()

Modify an exist DHCP reservation



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
	mac := "mac_example" // string | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33.
	createDhcpReservationOpenApiVO := *openapiclient.NewCreateDhcpReservationOpenApiVO("Mac_example", "NetId_example", false) // CreateDhcpReservationOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifyDhcpReservation(context.Background(), omadacId, siteId, mac).CreateDhcpReservationOpenApiVO(createDhcpReservationOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifyDhcpReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDhcpReservation`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifyDhcpReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mac** | **string** | MAC address of the DHCP reservation, format: AA-BB-CC-11-22-33. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDhcpReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createDhcpReservationOpenApiVO** | [**CreateDhcpReservationOpenApiVO**](CreateDhcpReservationOpenApiVO.md) |  | 

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


## ModifyDnsCacheSetting

> OperationResponseWithoutResult ModifyDnsCacheSetting(ctx, omadacId, siteId).DnsCacheOpenApiVO(dnsCacheOpenApiVO).Execute()

Modify DNS cache setting



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
	dnsCacheOpenApiVO := *openapiclient.NewDnsCacheOpenApiVO() // DnsCacheOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifyDnsCacheSetting(context.Background(), omadacId, siteId).DnsCacheOpenApiVO(dnsCacheOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifyDnsCacheSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsCacheSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifyDnsCacheSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDnsCacheSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsCacheOpenApiVO** | [**DnsCacheOpenApiVO**](DnsCacheOpenApiVO.md) |  | 

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


## ModifyDnsProxy

> OperationResponseWithoutResult ModifyDnsProxy(ctx, omadacId, siteId).DnsProxySettingOpenApiVO(dnsProxySettingOpenApiVO).Execute()

Modify DNS proxy setting



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
	dnsProxySettingOpenApiVO := *openapiclient.NewDnsProxySettingOpenApiVO(false) // DnsProxySettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifyDnsProxy(context.Background(), omadacId, siteId).DnsProxySettingOpenApiVO(dnsProxySettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifyDnsProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyDnsProxy`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifyDnsProxy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDnsProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsProxySettingOpenApiVO** | [**DnsProxySettingOpenApiVO**](DnsProxySettingOpenApiVO.md) |  | 

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


## ModifyIgmp

> OperationResponseWithoutResult ModifyIgmp(ctx, omadacId, siteId).IgmpOpenApiVO(igmpOpenApiVO).Execute()

Modify IGMP setting



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
	igmpOpenApiVO := *openapiclient.NewIgmpOpenApiVO(false, int32(123)) // IgmpOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifyIgmp(context.Background(), omadacId, siteId).IgmpOpenApiVO(igmpOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifyIgmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIgmp`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifyIgmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIgmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **igmpOpenApiVO** | [**IgmpOpenApiVO**](IgmpOpenApiVO.md) |  | 

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


## ModifyIptv

> OperationResponseWithoutResult ModifyIptv(ctx, omadacId, siteId).IptvConfigOpenApiVO(iptvConfigOpenApiVO).Execute()

Modify IPTV setting



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
	iptvConfigOpenApiVO := *openapiclient.NewIptvConfigOpenApiVO(false, int32(123), []openapiclient.IptvPortConfigOpenApiVO{*openapiclient.NewIptvPortConfigOpenApiVO("PortId_example", int32(123))}, "WanPortId_example") // IptvConfigOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifyIptv(context.Background(), omadacId, siteId).IptvConfigOpenApiVO(iptvConfigOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifyIptv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyIptv`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifyIptv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIptvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iptvConfigOpenApiVO** | [**IptvConfigOpenApiVO**](IptvConfigOpenApiVO.md) |  | 

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


## ModifyMdns

> OperationResponseWithoutResult ModifyMdns(ctx, omadacId, siteId, mdnsId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()

Modify an exist mDNS rule



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
	mdnsId := "mdnsId_example" // string | mDNS rule ID
	createMdnsRuleOpenApiVO := *openapiclient.NewCreateMdnsRuleOpenApiVO("Name_example", []string{"ProfileIds_example"}, false) // CreateMdnsRuleOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifyMdns(context.Background(), omadacId, siteId, mdnsId).CreateMdnsRuleOpenApiVO(createMdnsRuleOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifyMdns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMdns`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifyMdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**mdnsId** | **string** | mDNS rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createMdnsRuleOpenApiVO** | [**CreateMdnsRuleOpenApiVO**](CreateMdnsRuleOpenApiVO.md) |  | 

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


## ModifySnmpSetting

> OperationResponseWithoutResult ModifySnmpSetting(ctx, omadacId, siteId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()

Modify SNMP setting



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
	snmpSettingOpenApiVO := *openapiclient.NewSnmpSettingOpenApiVO(false, false) // SnmpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ModifySnmpSetting(context.Background(), omadacId, siteId).SnmpSettingOpenApiVO(snmpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ModifySnmpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifySnmpSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ModifySnmpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySnmpSettingRequest struct via the builder pattern


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


## ReservationFromDhcpUserList

> OperationResponseWithoutResult ReservationFromDhcpUserList(ctx, omadacId, siteId).DhcpUserFilterVO(dhcpUserFilterVO).Execute()

Reservation From Dhcp User List



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
	dhcpUserFilterVO := *openapiclient.NewDhcpUserFilterVO(*openapiclient.NewBatchSelectMacsVO("SelectType_example")) // DhcpUserFilterVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ReservationFromDhcpUserList(context.Background(), omadacId, siteId).DhcpUserFilterVO(dhcpUserFilterVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ReservationFromDhcpUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReservationFromDhcpUserList`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ReservationFromDhcpUserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReservationFromDhcpUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dhcpUserFilterVO** | [**DhcpUserFilterVO**](DhcpUserFilterVO.md) |  | 

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


## UpdateSshSetting

> OperationResponseWithoutResult UpdateSshSetting(ctx, omadacId, siteId).SSHSetting(sSHSetting).Execute()

Modify SSH setting



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
	sSHSetting := *openapiclient.NewSSHSetting(false, int32(123)) // SSHSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.UpdateSshSetting(context.Background(), omadacId, siteId).SSHSetting(sSHSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.UpdateSshSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSshSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.UpdateSshSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSshSettingRequest struct via the builder pattern


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


## UpdateUpnpSetting

> OperationResponseWithoutResult UpdateUpnpSetting(ctx, omadacId, siteId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()

Modify UPnP setting



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
	upnpSettingOpenApiVO := *openapiclient.NewUpnpSettingOpenApiVO(false) // UpnpSettingOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.UpdateUpnpSetting(context.Background(), omadacId, siteId).UpnpSettingOpenApiVO(upnpSettingOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.UpdateUpnpSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUpnpSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.UpdateUpnpSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUpnpSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upnpSettingOpenApiVO** | [**UpnpSettingOpenApiVO**](UpnpSettingOpenApiVO.md) |  | 

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

