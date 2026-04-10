# ACLAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteOsgCustomAcls**](ACLAPI.md#batchdeleteosgcustomacls) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/gateway-acls/batch-delete | batch Delete Osg Custom Acls
[**BatchEditOsgCustomAclStatus**](ACLAPI.md#batcheditosgcustomaclstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/gateway-acls/batch-edit | batch Edit Osg Custom Acls
[**ClearOsgAclHitCounts**](ACLAPI.md#clearosgaclhitcounts) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-acls/hit-counts | Clear gateway ACL hit counts
[**ClearOsgHitCounts**](ACLAPI.md#clearosghitcounts) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-custom-acls/hit-counts | Clear gateway custom ACL hit counts
[**CreateEapAcl**](ACLAPI.md#createeapacl) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/eap-acls | Create new EAP ACL
[**CreateOsgAcl**](ACLAPI.md#createosgacl) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-acls | Create new gateway ACL
[**CreateOswAcl**](ACLAPI.md#createoswacl) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/osw-acls | Create new switch ACL
[**DeleteAcl**](ACLAPI.md#deleteacl) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/acls/{aclId} | Delete ACL
[**ExportOsgCustomAcl**](ACLAPI.md#exportosgcustomacl) | **Get** /openapi/v1/{omadacId}/files/sites/{siteId}/acls/export | Export gateway custom ACL rules
[**GetAclConfigTypeSetting**](ACLAPI.md#getaclconfigtypesetting) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-config-mode | Get gateway ACL config mode
[**GetEapAclList**](ACLAPI.md#geteapacllist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/acls/eap-acls | Get EAP ACL list
[**GetOsgAclList**](ACLAPI.md#getosgacllist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-acls | Get gateway ACL list
[**GetOsgCustomAclList**](ACLAPI.md#getosgcustomacllist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-custom-acls | Get osg custom ACL
[**GetOswAclList**](ACLAPI.md#getoswacllist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/acls/osw-acls | Get switch ACL list
[**ModifyAclConfigTypeSetting**](ACLAPI.md#modifyaclconfigtypesetting) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-config-mode | Modify gateway ACL config mode
[**ModifyAclIndex**](ACLAPI.md#modifyaclindex) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/modifyIndex | Modify ACL Index
[**ModifyEapAcl**](ACLAPI.md#modifyeapacl) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/acls/eap-acls/{aclId} | Modify an  EAP ACL
[**ModifyOsgAcl**](ACLAPI.md#modifyosgacl) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-acls/{aclId} | Modify a gateway ACL
[**ModifyOsgCustomAclIndex**](ACLAPI.md#modifyosgcustomaclindex) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-custom-acls/modifyIndex | Modify gateway custom ACL index
[**ModifyOswAcl**](ACLAPI.md#modifyoswacl) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/acls/osw-acls/{aclId} | Modify a switch ACL
[**UpdateOsgCustomAcls**](ACLAPI.md#updateosgcustomacls) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/acls/osg-custom-acls | Update osg custom ACLs



## BatchDeleteOsgCustomAcls

> OperationResponseWithoutResult BatchDeleteOsgCustomAcls(ctx, omadacId, siteId).BatchIds(batchIds).Execute()

batch Delete Osg Custom Acls



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
	batchIds := *openapiclient.NewBatchIds([]string{"Ids_example"}) // BatchIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.BatchDeleteOsgCustomAcls(context.Background(), omadacId, siteId).BatchIds(batchIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.BatchDeleteOsgCustomAcls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchDeleteOsgCustomAcls`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.BatchDeleteOsgCustomAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteOsgCustomAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchIds** | [**BatchIds**](BatchIds.md) |  | 

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


## BatchEditOsgCustomAclStatus

> OperationResponseWithoutResult BatchEditOsgCustomAclStatus(ctx, omadacId, siteId).BatchEditCustomAclOpenApiVO(batchEditCustomAclOpenApiVO).Execute()

batch Edit Osg Custom Acls



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
	batchEditCustomAclOpenApiVO := *openapiclient.NewBatchEditCustomAclOpenApiVO([]string{"Ids_example"}, int32(123), int32(123)) // BatchEditCustomAclOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.BatchEditOsgCustomAclStatus(context.Background(), omadacId, siteId).BatchEditCustomAclOpenApiVO(batchEditCustomAclOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.BatchEditOsgCustomAclStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchEditOsgCustomAclStatus`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.BatchEditOsgCustomAclStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchEditOsgCustomAclStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchEditCustomAclOpenApiVO** | [**BatchEditCustomAclOpenApiVO**](BatchEditCustomAclOpenApiVO.md) |  | 

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


## ClearOsgAclHitCounts

> OperationResponseWithoutResult ClearOsgAclHitCounts(ctx, omadacId, siteId).Execute()

Clear gateway ACL hit counts



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
	resp, r, err := apiClient.ACLAPI.ClearOsgAclHitCounts(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ClearOsgAclHitCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearOsgAclHitCounts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ClearOsgAclHitCounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearOsgAclHitCountsRequest struct via the builder pattern


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


## ClearOsgHitCounts

> OperationResponseWithoutResult ClearOsgHitCounts(ctx, omadacId, siteId).Execute()

Clear gateway custom ACL hit counts



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
	resp, r, err := apiClient.ACLAPI.ClearOsgHitCounts(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ClearOsgHitCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearOsgHitCounts`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ClearOsgHitCounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearOsgHitCountsRequest struct via the builder pattern


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


## CreateEapAcl

> OperationResponseWithoutResult CreateEapAcl(ctx, omadacId, siteId).EapACLConfig(eapACLConfig).Execute()

Create new EAP ACL



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
	eapACLConfig := *openapiclient.NewEapACLConfig("Description_example", int32(123), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // EapACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.CreateEapAcl(context.Background(), omadacId, siteId).EapACLConfig(eapACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.CreateEapAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEapAcl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.CreateEapAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEapAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eapACLConfig** | [**EapACLConfig**](EapACLConfig.md) |  | 

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


## CreateOsgAcl

> OperationResponseWithoutResult CreateOsgAcl(ctx, omadacId, siteId).GatewayACLConfig(gatewayACLConfig).Execute()

Create new gateway ACL



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
	gatewayACLConfig := *openapiclient.NewGatewayACLConfig("Description_example", int32(123), *openapiclient.NewGatewayDirectionEntity(), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), int32(123), false, false) // GatewayACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.CreateOsgAcl(context.Background(), omadacId, siteId).GatewayACLConfig(gatewayACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.CreateOsgAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOsgAcl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.CreateOsgAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOsgAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayACLConfig** | [**GatewayACLConfig**](GatewayACLConfig.md) |  | 

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


## CreateOswAcl

> OperationResponseWithoutResult CreateOswAcl(ctx, omadacId, siteId).SwitchACLConfig(switchACLConfig).Execute()

Create new switch ACL



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
	switchACLConfig := *openapiclient.NewSwitchACLConfig(int32(123), "Description_example", int32(123), *openapiclient.NewSwitchACLEtherTypeEntity(false), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // SwitchACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.CreateOswAcl(context.Background(), omadacId, siteId).SwitchACLConfig(switchACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.CreateOswAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOswAcl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.CreateOswAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOswAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **switchACLConfig** | [**SwitchACLConfig**](SwitchACLConfig.md) |  | 

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


## DeleteAcl

> OperationResponseWithoutResult DeleteAcl(ctx, omadacId, siteId, aclId).Execute()

Delete ACL



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
	aclId := "aclId_example" // string | ACL ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.DeleteAcl(context.Background(), omadacId, siteId, aclId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.DeleteAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAcl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.DeleteAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**aclId** | **string** | ACL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAclRequest struct via the builder pattern


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


## ExportOsgCustomAcl

> OperationResponse ExportOsgCustomAcl(ctx, omadacId, siteId).Execute()

Export gateway custom ACL rules



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
	resp, r, err := apiClient.ACLAPI.ExportOsgCustomAcl(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ExportOsgCustomAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportOsgCustomAcl`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ExportOsgCustomAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportOsgCustomAclRequest struct via the builder pattern


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


## GetAclConfigTypeSetting

> OperationResponseGatewayACLConfigModeEntity GetAclConfigTypeSetting(ctx, omadacId, siteId).Execute()

Get gateway ACL config mode



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
	resp, r, err := apiClient.ACLAPI.GetAclConfigTypeSetting(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.GetAclConfigTypeSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAclConfigTypeSetting`: OperationResponseGatewayACLConfigModeEntity
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.GetAclConfigTypeSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAclConfigTypeSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseGatewayACLConfigModeEntity**](OperationResponseGatewayACLConfigModeEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEapAclList

> OperationResponseGridVOEapACLInfo GetEapAclList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get EAP ACL list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.GetEapAclList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.GetEapAclList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEapAclList`: OperationResponseGridVOEapACLInfo
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.GetEapAclList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEapAclListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOEapACLInfo**](OperationResponseGridVOEapACLInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsgAclList

> OperationResponseGridVOGatewayACLInfo GetOsgAclList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get gateway ACL list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.GetOsgAclList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.GetOsgAclList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsgAclList`: OperationResponseGridVOGatewayACLInfo
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.GetOsgAclList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsgAclListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOGatewayACLInfo**](OperationResponseGridVOGatewayACLInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOsgCustomAclList

> OperationResponseGridVOGatewayCustomACLInfoEntity GetOsgCustomAclList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get osg custom ACL



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.GetOsgCustomAclList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.GetOsgCustomAclList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsgCustomAclList`: OperationResponseGridVOGatewayCustomACLInfoEntity
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.GetOsgCustomAclList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsgCustomAclListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOGatewayCustomACLInfoEntity**](OperationResponseGridVOGatewayCustomACLInfoEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOswAclList

> OperationResponseGridVOSwitchACLInfo GetOswAclList(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get switch ACL list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	omadacId := "omadacId_example" // string | Omada ID
	siteId := "siteId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.GetOswAclList(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.GetOswAclList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOswAclList`: OperationResponseGridVOSwitchACLInfo
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.GetOswAclList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOswAclListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**OperationResponseGridVOSwitchACLInfo**](OperationResponseGridVOSwitchACLInfo.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAclConfigTypeSetting

> OperationResponseWithoutResult ModifyAclConfigTypeSetting(ctx, omadacId, siteId).GatewayACLConfigModeEntity(gatewayACLConfigModeEntity).Execute()

Modify gateway ACL config mode



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
	gatewayACLConfigModeEntity := *openapiclient.NewGatewayACLConfigModeEntity(int32(123)) // GatewayACLConfigModeEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.ModifyAclConfigTypeSetting(context.Background(), omadacId, siteId).GatewayACLConfigModeEntity(gatewayACLConfigModeEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ModifyAclConfigTypeSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAclConfigTypeSetting`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ModifyAclConfigTypeSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAclConfigTypeSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayACLConfigModeEntity** | [**GatewayACLConfigModeEntity**](GatewayACLConfigModeEntity.md) |  | 

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


## ModifyAclIndex

> OperationResponseWithoutResult ModifyAclIndex(ctx, omadacId, siteId).DragSortIndexOpenapiVO(dragSortIndexOpenapiVO).Execute()

Modify ACL Index



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
	dragSortIndexOpenapiVO := *openapiclient.NewDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}, "Type_example") // DragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.ModifyAclIndex(context.Background(), omadacId, siteId).DragSortIndexOpenapiVO(dragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ModifyAclIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAclIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ModifyAclIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAclIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dragSortIndexOpenapiVO** | [**DragSortIndexOpenapiVO**](DragSortIndexOpenapiVO.md) |  | 

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


## ModifyEapAcl

> OperationResponseWithoutResult ModifyEapAcl(ctx, omadacId, siteId, aclId).EapACLConfig(eapACLConfig).Execute()

Modify an  EAP ACL



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
	aclId := "aclId_example" // string | ACL ID
	eapACLConfig := *openapiclient.NewEapACLConfig("Description_example", int32(123), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // EapACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.ModifyEapAcl(context.Background(), omadacId, siteId, aclId).EapACLConfig(eapACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ModifyEapAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyEapAcl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ModifyEapAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**aclId** | **string** | ACL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyEapAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **eapACLConfig** | [**EapACLConfig**](EapACLConfig.md) |  | 

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


## ModifyOsgAcl

> OperationResponseWithoutResult ModifyOsgAcl(ctx, omadacId, siteId, aclId).GatewayACLConfig(gatewayACLConfig).Execute()

Modify a gateway ACL



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
	aclId := "aclId_example" // string | ACL ID
	gatewayACLConfig := *openapiclient.NewGatewayACLConfig("Description_example", int32(123), *openapiclient.NewGatewayDirectionEntity(), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), int32(123), false, false) // GatewayACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.ModifyOsgAcl(context.Background(), omadacId, siteId, aclId).GatewayACLConfig(gatewayACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ModifyOsgAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOsgAcl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ModifyOsgAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**aclId** | **string** | ACL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOsgAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gatewayACLConfig** | [**GatewayACLConfig**](GatewayACLConfig.md) |  | 

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


## ModifyOsgCustomAclIndex

> OperationResponseWithoutResult ModifyOsgCustomAclIndex(ctx, omadacId, siteId).DragSortIndexOpenapiVO(dragSortIndexOpenapiVO).Execute()

Modify gateway custom ACL index



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
	dragSortIndexOpenapiVO := *openapiclient.NewDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}, "Type_example") // DragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.ModifyOsgCustomAclIndex(context.Background(), omadacId, siteId).DragSortIndexOpenapiVO(dragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ModifyOsgCustomAclIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOsgCustomAclIndex`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ModifyOsgCustomAclIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOsgCustomAclIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dragSortIndexOpenapiVO** | [**DragSortIndexOpenapiVO**](DragSortIndexOpenapiVO.md) |  | 

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


## ModifyOswAcl

> OperationResponseWithoutResult ModifyOswAcl(ctx, omadacId, siteId, aclId).SwitchACLConfig(switchACLConfig).Execute()

Modify a switch ACL



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
	aclId := "aclId_example" // string | ACL ID
	switchACLConfig := *openapiclient.NewSwitchACLConfig(int32(123), "Description_example", int32(123), *openapiclient.NewSwitchACLEtherTypeEntity(false), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // SwitchACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.ModifyOswAcl(context.Background(), omadacId, siteId, aclId).SwitchACLConfig(switchACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.ModifyOswAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOswAcl`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.ModifyOswAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**aclId** | **string** | ACL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOswAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **switchACLConfig** | [**SwitchACLConfig**](SwitchACLConfig.md) |  | 

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


## UpdateOsgCustomAcls

> OperationResponseWithoutResult UpdateOsgCustomAcls(ctx, omadacId, siteId).GatewayCustomACLUpdateEntity(gatewayCustomACLUpdateEntity).Execute()

Update osg custom ACLs



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
	gatewayCustomACLUpdateEntity := *openapiclient.NewGatewayCustomACLUpdateEntity() // GatewayCustomACLUpdateEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLAPI.UpdateOsgCustomAcls(context.Background(), omadacId, siteId).GatewayCustomACLUpdateEntity(gatewayCustomACLUpdateEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLAPI.UpdateOsgCustomAcls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOsgCustomAcls`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLAPI.UpdateOsgCustomAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOsgCustomAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayCustomACLUpdateEntity** | [**GatewayCustomACLUpdateEntity**](GatewayCustomACLUpdateEntity.md) |  | 

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

