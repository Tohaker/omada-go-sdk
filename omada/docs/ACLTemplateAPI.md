# ACLTemplateAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEapAclTemplate**](ACLTemplateAPI.md#createeapacltemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/eap-acls | Create new eap template ACL
[**CreateOsgAclTemplate**](ACLTemplateAPI.md#createosgacltemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-acls | Create new gateway template ACL
[**DeleteAclTemplate**](ACLTemplateAPI.md#deleteacltemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/{aclId} | Delete Gateway/AP ACL template
[**ExportOsgCustomAclTemplate**](ACLTemplateAPI.md#exportosgcustomacltemplate) | **Get** /openapi/v1/{omadacId}/files/sitetemplates/{siteTemplateId}/acls/export | Export gateway custom ACL template rules
[**GetAclConfigTypeSettingTemplate**](ACLTemplateAPI.md#getaclconfigtypesettingtemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-config-mode | Get gateway ACL config mode template
[**GetEapAclListTemplate**](ACLTemplateAPI.md#geteapacllisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/eap-acls | Get eap ACL template list
[**GetOsgAclListTemplate**](ACLTemplateAPI.md#getosgacllisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-acls | Get gateway ACL template list
[**GetOsgCustomAclListTemplate**](ACLTemplateAPI.md#getosgcustomacllisttemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-custom-acls | Get osg custom ACL template
[**ModifyAclConfigTypeSettingTemplate**](ACLTemplateAPI.md#modifyaclconfigtypesettingtemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-config-mode | Modify gateway ACL config template mode
[**ModifyAclIndexTemplate**](ACLTemplateAPI.md#modifyaclindextemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/modifyIndex | Modify ACL template Index
[**ModifyEapAclTemplate**](ACLTemplateAPI.md#modifyeapacltemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/eap-acls/{aclId} | Modify an EAP template ACL
[**ModifyOsgAclTemplate**](ACLTemplateAPI.md#modifyosgacltemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-acls/{aclId} | Modify a gateway template ACL
[**ModifyOsgCustomAclIndexTemplate**](ACLTemplateAPI.md#modifyosgcustomaclindextemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-custom-acls/modifyIndex | Modify gateway custom ACL template index
[**UpdateOsgCustomAclsTemplate**](ACLTemplateAPI.md#updateosgcustomaclstemplate) | **Patch** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/acls/osg-custom-acls | Update osg custom ACL templates



## CreateEapAclTemplate

> OperationResponseWithoutResult CreateEapAclTemplate(ctx, omadacId, siteTemplateId).EapACLConfig(eapACLConfig).Execute()

Create new eap template ACL



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
	eapACLConfig := *openapiclient.NewEapACLConfig("Description_example", int32(123), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // EapACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.CreateEapAclTemplate(context.Background(), omadacId, siteTemplateId).EapACLConfig(eapACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.CreateEapAclTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEapAclTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.CreateEapAclTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEapAclTemplateRequest struct via the builder pattern


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


## CreateOsgAclTemplate

> OperationResponseWithoutResult CreateOsgAclTemplate(ctx, omadacId, siteTemplateId).GatewayACLConfig(gatewayACLConfig).Execute()

Create new gateway template ACL



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
	gatewayACLConfig := *openapiclient.NewGatewayACLConfig("Description_example", int32(123), *openapiclient.NewGatewayDirectionEntity(), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), int32(123), false, false) // GatewayACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.CreateOsgAclTemplate(context.Background(), omadacId, siteTemplateId).GatewayACLConfig(gatewayACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.CreateOsgAclTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOsgAclTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.CreateOsgAclTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOsgAclTemplateRequest struct via the builder pattern


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


## DeleteAclTemplate

> OperationResponseWithoutResult DeleteAclTemplate(ctx, omadacId, siteTemplateId, aclId).Execute()

Delete Gateway/AP ACL template



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
	aclId := "aclId_example" // string | ACL ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.DeleteAclTemplate(context.Background(), omadacId, siteTemplateId, aclId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.DeleteAclTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAclTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.DeleteAclTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**aclId** | **string** | ACL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAclTemplateRequest struct via the builder pattern


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


## ExportOsgCustomAclTemplate

> OperationResponse ExportOsgCustomAclTemplate(ctx, omadacId, siteTemplateId).Execute()

Export gateway custom ACL template rules



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
	resp, r, err := apiClient.ACLTemplateAPI.ExportOsgCustomAclTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.ExportOsgCustomAclTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportOsgCustomAclTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.ExportOsgCustomAclTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportOsgCustomAclTemplateRequest struct via the builder pattern


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


## GetAclConfigTypeSettingTemplate

> OperationResponseGatewayACLConfigModeEntity GetAclConfigTypeSettingTemplate(ctx, omadacId, siteTemplateId).Execute()

Get gateway ACL config mode template



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
	resp, r, err := apiClient.ACLTemplateAPI.GetAclConfigTypeSettingTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.GetAclConfigTypeSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAclConfigTypeSettingTemplate`: OperationResponseGatewayACLConfigModeEntity
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.GetAclConfigTypeSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAclConfigTypeSettingTemplateRequest struct via the builder pattern


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


## GetEapAclListTemplate

> OperationResponseGridVOEapACLInfo GetEapAclListTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get eap ACL template list



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.GetEapAclListTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.GetEapAclListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEapAclListTemplate`: OperationResponseGridVOEapACLInfo
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.GetEapAclListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEapAclListTemplateRequest struct via the builder pattern


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


## GetOsgAclListTemplate

> OperationResponseGridVOGatewayACLInfo GetOsgAclListTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get gateway ACL template list



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.GetOsgAclListTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.GetOsgAclListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsgAclListTemplate`: OperationResponseGridVOGatewayACLInfo
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.GetOsgAclListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsgAclListTemplateRequest struct via the builder pattern


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


## GetOsgCustomAclListTemplate

> OperationResponseGridVOGatewayCustomACLInfoEntity GetOsgCustomAclListTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get osg custom ACL template



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
	siteTemplateId := "siteTemplateId_example" // string | Site ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.GetOsgCustomAclListTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.GetOsgCustomAclListTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOsgCustomAclListTemplate`: OperationResponseGridVOGatewayCustomACLInfoEntity
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.GetOsgCustomAclListTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOsgCustomAclListTemplateRequest struct via the builder pattern


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


## ModifyAclConfigTypeSettingTemplate

> OperationResponseWithoutResult ModifyAclConfigTypeSettingTemplate(ctx, omadacId, siteTemplateId).GatewayACLConfigModeEntity(gatewayACLConfigModeEntity).Execute()

Modify gateway ACL config template mode



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
	gatewayACLConfigModeEntity := *openapiclient.NewGatewayACLConfigModeEntity(int32(123)) // GatewayACLConfigModeEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.ModifyAclConfigTypeSettingTemplate(context.Background(), omadacId, siteTemplateId).GatewayACLConfigModeEntity(gatewayACLConfigModeEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.ModifyAclConfigTypeSettingTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAclConfigTypeSettingTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.ModifyAclConfigTypeSettingTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAclConfigTypeSettingTemplateRequest struct via the builder pattern


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


## ModifyAclIndexTemplate

> OperationResponseWithoutResult ModifyAclIndexTemplate(ctx, omadacId, siteTemplateId).DragSortIndexTemplateOpenapiVO(dragSortIndexTemplateOpenapiVO).Execute()

Modify ACL template Index



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
	dragSortIndexTemplateOpenapiVO := *openapiclient.NewDragSortIndexTemplateOpenapiVO(map[string]int32{"key": int32(123)}, "Type_example") // DragSortIndexTemplateOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.ModifyAclIndexTemplate(context.Background(), omadacId, siteTemplateId).DragSortIndexTemplateOpenapiVO(dragSortIndexTemplateOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.ModifyAclIndexTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAclIndexTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.ModifyAclIndexTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAclIndexTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dragSortIndexTemplateOpenapiVO** | [**DragSortIndexTemplateOpenapiVO**](DragSortIndexTemplateOpenapiVO.md) |  | 

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


## ModifyEapAclTemplate

> OperationResponseWithoutResult ModifyEapAclTemplate(ctx, omadacId, siteTemplateId, aclId).EapACLConfig(eapACLConfig).Execute()

Modify an EAP template ACL



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
	aclId := "aclId_example" // string | ACL ID
	eapACLConfig := *openapiclient.NewEapACLConfig("Description_example", int32(123), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), false) // EapACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.ModifyEapAclTemplate(context.Background(), omadacId, siteTemplateId, aclId).EapACLConfig(eapACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.ModifyEapAclTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyEapAclTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.ModifyEapAclTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**aclId** | **string** | ACL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyEapAclTemplateRequest struct via the builder pattern


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


## ModifyOsgAclTemplate

> OperationResponseWithoutResult ModifyOsgAclTemplate(ctx, omadacId, siteTemplateId, aclId).GatewayACLConfig(gatewayACLConfig).Execute()

Modify a gateway template ACL



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
	aclId := "aclId_example" // string | ACL ID
	gatewayACLConfig := *openapiclient.NewGatewayACLConfig("Description_example", int32(123), *openapiclient.NewGatewayDirectionEntity(), int32(123), []int32{int32(123)}, []string{"SourceIds_example"}, int32(123), int32(123), false, false) // GatewayACLConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.ModifyOsgAclTemplate(context.Background(), omadacId, siteTemplateId, aclId).GatewayACLConfig(gatewayACLConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.ModifyOsgAclTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOsgAclTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.ModifyOsgAclTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 
**aclId** | **string** | ACL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOsgAclTemplateRequest struct via the builder pattern


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


## ModifyOsgCustomAclIndexTemplate

> OperationResponseWithoutResult ModifyOsgCustomAclIndexTemplate(ctx, omadacId, siteTemplateId).DragSortIndexOpenapiVO(dragSortIndexOpenapiVO).Execute()

Modify gateway custom ACL template index



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
	dragSortIndexOpenapiVO := *openapiclient.NewDragSortIndexOpenapiVO(map[string]int32{"key": int32(123)}, "Type_example") // DragSortIndexOpenapiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.ModifyOsgCustomAclIndexTemplate(context.Background(), omadacId, siteTemplateId).DragSortIndexOpenapiVO(dragSortIndexOpenapiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.ModifyOsgCustomAclIndexTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyOsgCustomAclIndexTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.ModifyOsgCustomAclIndexTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyOsgCustomAclIndexTemplateRequest struct via the builder pattern


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


## UpdateOsgCustomAclsTemplate

> OperationResponseWithoutResult UpdateOsgCustomAclsTemplate(ctx, omadacId, siteTemplateId).GatewayCustomACLUpdateEntity(gatewayCustomACLUpdateEntity).Execute()

Update osg custom ACL templates



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
	gatewayCustomACLUpdateEntity := *openapiclient.NewGatewayCustomACLUpdateEntity() // GatewayCustomACLUpdateEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ACLTemplateAPI.UpdateOsgCustomAclsTemplate(context.Background(), omadacId, siteTemplateId).GatewayCustomACLUpdateEntity(gatewayCustomACLUpdateEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ACLTemplateAPI.UpdateOsgCustomAclsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOsgCustomAclsTemplate`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `ACLTemplateAPI.UpdateOsgCustomAclsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOsgCustomAclsTemplateRequest struct via the builder pattern


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

