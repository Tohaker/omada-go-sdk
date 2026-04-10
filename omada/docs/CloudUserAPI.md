# CloudUserAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindMspOwner**](CloudUserAPI.md#bindmspowner) | **Post** /openapi/v1/msp/{mspId}/cmd/cloud/bind | Bind the MSP owner account
[**BindOwner**](CloudUserAPI.md#bindowner) | **Post** /openapi/v1/{omadacId}/cmd/cloud/bind | Bind the owner account
[**GetCloudAccessStatus**](CloudUserAPI.md#getcloudaccessstatus) | **Get** /openapi/v1/{omadacId}/cloud/status | Get cloud access status
[**GetCloudUserInfo**](CloudUserAPI.md#getclouduserinfo) | **Get** /openapi/v1/{omadacId}/cloud/user | Get cloud user information
[**GetMspCloudAccessStatus**](CloudUserAPI.md#getmspcloudaccessstatus) | **Get** /openapi/v1/msp/{mspId}/cloud/status | Get MSP cloud access status
[**GetRemoteBindingStatus**](CloudUserAPI.md#getremotebindingstatus) | **Get** /openapi/v1/{omadacId}/cloud/remote/bind/status | Get remote bind status
[**UnbindMspOwner**](CloudUserAPI.md#unbindmspowner) | **Post** /openapi/v1/msp/{mspId}/cmd/cloud/unbind | Unbind the MSP owner account
[**UnbindOwner**](CloudUserAPI.md#unbindowner) | **Post** /openapi/v1/{omadacId}/cmd/cloud/unbind | Unbind the owner account



## BindMspOwner

> OperationResponseWithoutResult BindMspOwner(ctx, mspId).BindOwnerOpenApiVO(bindOwnerOpenApiVO).Execute()

Bind the MSP owner account



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
	mspId := "mspId_example" // string | MSP ID
	bindOwnerOpenApiVO := *openapiclient.NewBindOwnerOpenApiVO("AccountName_example", "Password_example") // BindOwnerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.BindMspOwner(context.Background(), mspId).BindOwnerOpenApiVO(bindOwnerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.BindMspOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindMspOwner`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.BindMspOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindMspOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindOwnerOpenApiVO** | [**BindOwnerOpenApiVO**](BindOwnerOpenApiVO.md) |  | 

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


## BindOwner

> OperationResponseWithoutResult BindOwner(ctx, omadacId).BindOwnerOpenApiVO(bindOwnerOpenApiVO).Execute()

Bind the owner account



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
	bindOwnerOpenApiVO := *openapiclient.NewBindOwnerOpenApiVO("AccountName_example", "Password_example") // BindOwnerOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.BindOwner(context.Background(), omadacId).BindOwnerOpenApiVO(bindOwnerOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.BindOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindOwner`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.BindOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindOwnerOpenApiVO** | [**BindOwnerOpenApiVO**](BindOwnerOpenApiVO.md) |  | 

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


## GetCloudAccessStatus

> OperationResponseCloudAccessOpenApiVO GetCloudAccessStatus(ctx, omadacId).Execute()

Get cloud access status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.GetCloudAccessStatus(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.GetCloudAccessStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudAccessStatus`: OperationResponseCloudAccessOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.GetCloudAccessStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudAccessStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseCloudAccessOpenApiVO**](OperationResponseCloudAccessOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudUserInfo

> OperationResponse GetCloudUserInfo(ctx, omadacId).Execute()

Get cloud user information



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.GetCloudUserInfo(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.GetCloudUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudUserInfo`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.GetCloudUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudUserInfoRequest struct via the builder pattern


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


## GetMspCloudAccessStatus

> OperationResponseCloudAccessOpenApiVO GetMspCloudAccessStatus(ctx, mspId).Execute()

Get MSP cloud access status



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
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.GetMspCloudAccessStatus(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.GetMspCloudAccessStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspCloudAccessStatus`: OperationResponseCloudAccessOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.GetMspCloudAccessStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspCloudAccessStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseCloudAccessOpenApiVO**](OperationResponseCloudAccessOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteBindingStatus

> OperationResponseRemoteBindingOpenApiVO GetRemoteBindingStatus(ctx, omadacId).Execute()

Get remote bind status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.GetRemoteBindingStatus(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.GetRemoteBindingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteBindingStatus`: OperationResponseRemoteBindingOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.GetRemoteBindingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteBindingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseRemoteBindingOpenApiVO**](OperationResponseRemoteBindingOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindMspOwner

> OperationResponseWithoutResult UnbindMspOwner(ctx, mspId).Execute()

Unbind the MSP owner account



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
	mspId := "mspId_example" // string | MSP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.UnbindMspOwner(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.UnbindMspOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindMspOwner`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.UnbindMspOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** | MSP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindMspOwnerRequest struct via the builder pattern


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


## UnbindOwner

> OperationResponseWithoutResult UnbindOwner(ctx, omadacId).Execute()

Unbind the owner account



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudUserAPI.UnbindOwner(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudUserAPI.UnbindOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindOwner`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `CloudUserAPI.UnbindOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindOwnerRequest struct via the builder pattern


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

