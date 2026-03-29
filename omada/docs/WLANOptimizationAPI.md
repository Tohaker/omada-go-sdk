# \WLANOptimizationAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyPlanningHistory**](WLANOptimizationAPI.md#ApplyPlanningHistory) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory/{id} | apply Planning History
[**CancelRadioFrequencyPlanning**](WLANOptimizationAPI.md#CancelRadioFrequencyPlanning) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/rfPlanning/cancel | cancel Radio Frequency Planning
[**DeletePlanningHistory**](WLANOptimizationAPI.md#DeletePlanningHistory) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory/{id} | delete Planning History
[**GetExperienceIndex**](WLANOptimizationAPI.md#GetExperienceIndex) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/experienceIndex | get Experience Index
[**GetGirdPlanningHistorys**](WLANOptimizationAPI.md#GetGirdPlanningHistorys) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory | get Gird Planning Historys
[**GetPlanningHistory**](WLANOptimizationAPI.md#GetPlanningHistory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/planningHistory/{id} | Get Planning History
[**GetRFPlanningDeployHistory**](WLANOptimizationAPI.md#GetRFPlanningDeployHistory) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfplanning/history | get RF Planning Deploy History
[**GetRadioFrequencyPlanningConfig**](WLANOptimizationAPI.md#GetRadioFrequencyPlanningConfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning | get Radio Frequency Planning Config
[**GetRadioFrequencyPlanningConfigTemplate**](WLANOptimizationAPI.md#GetRadioFrequencyPlanningConfigTemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rfPlanning/config | get Radio Frequency Planning Config Template
[**GetRadioFrequencyPlanningResult**](WLANOptimizationAPI.md#GetRadioFrequencyPlanningResult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/result | get Radio Frequency Planning Result
[**ModifyExcludeAps**](WLANOptimizationAPI.md#ModifyExcludeAps) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/excludeAps | modify Exclude Aps
[**ModifyRFPlanningDeployConfig**](WLANOptimizationAPI.md#ModifyRFPlanningDeployConfig) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/rfPlanning/config | modify RF Planning Deploy Config
[**ModifyRFPlanningDeployConfigTemplate**](WLANOptimizationAPI.md#ModifyRFPlanningDeployConfigTemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/rfPlanning/config | modify RF Planning Deploy Config Template
[**StartOptimization**](WLANOptimizationAPI.md#StartOptimization) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/rfPlanning/optimization | start Optimization
[**StartRrmOptimization**](WLANOptimizationAPI.md#StartRrmOptimization) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/cmd/rfPlanning/rrmOptimization | start Rrm Optimization



## ApplyPlanningHistory

> OperationResponse ApplyPlanningHistory(ctx, omadacId, siteId, id).AppliedConfig(appliedConfig).Execute()

apply Planning History



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
	id := "id_example" // string | planning history ID
	appliedConfig := *openapiclient.NewAppliedConfig() // AppliedConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ApplyPlanningHistory(context.Background(), omadacId, siteId, id).AppliedConfig(appliedConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ApplyPlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyPlanningHistory`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ApplyPlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | planning history ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyPlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **appliedConfig** | [**AppliedConfig**](AppliedConfig.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelRadioFrequencyPlanning

> OperationResponse CancelRadioFrequencyPlanning(ctx, omadacId, siteId).Execute()

cancel Radio Frequency Planning



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
	resp, r, err := apiClient.WLANOptimizationAPI.CancelRadioFrequencyPlanning(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.CancelRadioFrequencyPlanning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRadioFrequencyPlanning`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.CancelRadioFrequencyPlanning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRadioFrequencyPlanningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlanningHistory

> OperationResponse DeletePlanningHistory(ctx, omadacId, siteId, id).Execute()

delete Planning History



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
	id := "id_example" // string | planning history ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.DeletePlanningHistory(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.DeletePlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlanningHistory`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.DeletePlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | planning history ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperienceIndex

> ExperienceIndex GetExperienceIndex(ctx, omadacId, siteId).Execute()

get Experience Index



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetExperienceIndex(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetExperienceIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperienceIndex`: ExperienceIndex
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetExperienceIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperienceIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExperienceIndex**](ExperienceIndex.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGirdPlanningHistorys

> GridVORFPlanningHistory GetGirdPlanningHistorys(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

get Gird Planning Historys



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetGirdPlanningHistorys(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetGirdPlanningHistorys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGirdPlanningHistorys`: GridVORFPlanningHistory
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetGirdPlanningHistorys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGirdPlanningHistorysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 



### Return type

[**GridVORFPlanningHistory**](GridVORFPlanningHistory.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanningHistory

> OperationResponsePlanningHistoryDetail GetPlanningHistory(ctx, omadacId, siteId, id).Execute()

Get Planning History



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
	id := "id_example" // string | planning history ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.GetPlanningHistory(context.Background(), omadacId, siteId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetPlanningHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanningHistory`: OperationResponsePlanningHistoryDetail
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetPlanningHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**id** | **string** | planning history ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanningHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponsePlanningHistoryDetail**](OperationResponsePlanningHistoryDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRFPlanningDeployHistory

> RFPlanningDeployResult GetRFPlanningDeployHistory(ctx, omadacId, siteId).Execute()

get RF Planning Deploy History



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetRFPlanningDeployHistory(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRFPlanningDeployHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRFPlanningDeployHistory`: RFPlanningDeployResult
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRFPlanningDeployHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRFPlanningDeployHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RFPlanningDeployResult**](RFPlanningDeployResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadioFrequencyPlanningConfig

> OperationResponsePlanningHistoryDetail GetRadioFrequencyPlanningConfig(ctx, omadacId, siteId).Execute()

get Radio Frequency Planning Config



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetRadioFrequencyPlanningConfig(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRadioFrequencyPlanningConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadioFrequencyPlanningConfig`: OperationResponsePlanningHistoryDetail
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRadioFrequencyPlanningConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadioFrequencyPlanningConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponsePlanningHistoryDetail**](OperationResponsePlanningHistoryDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadioFrequencyPlanningConfigTemplate

> OperationResponsePlanningHistoryDetail GetRadioFrequencyPlanningConfigTemplate(ctx, omadacId, siteTemplateId).Execute()

get Radio Frequency Planning Config Template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.GetRadioFrequencyPlanningConfigTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRadioFrequencyPlanningConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadioFrequencyPlanningConfigTemplate`: OperationResponsePlanningHistoryDetail
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRadioFrequencyPlanningConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadioFrequencyPlanningConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponsePlanningHistoryDetail**](OperationResponsePlanningHistoryDetail.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRadioFrequencyPlanningResult

> OperationResponseRFPlanningResult GetRadioFrequencyPlanningResult(ctx, omadacId, siteId).Execute()

get Radio Frequency Planning Result



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
	resp, r, err := apiClient.WLANOptimizationAPI.GetRadioFrequencyPlanningResult(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.GetRadioFrequencyPlanningResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRadioFrequencyPlanningResult`: OperationResponseRFPlanningResult
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.GetRadioFrequencyPlanningResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRadioFrequencyPlanningResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseRFPlanningResult**](OperationResponseRFPlanningResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyExcludeAps

> OperationResponse ModifyExcludeAps(ctx, omadacId, siteId).ExcludedAPsConfig(excludedAPsConfig).Execute()

modify Exclude Aps



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
	excludedAPsConfig := *openapiclient.NewExcludedAPsConfig(false) // ExcludedAPsConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ModifyExcludeAps(context.Background(), omadacId, siteId).ExcludedAPsConfig(excludedAPsConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ModifyExcludeAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyExcludeAps`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ModifyExcludeAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyExcludeApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludedAPsConfig** | [**ExcludedAPsConfig**](ExcludedAPsConfig.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRFPlanningDeployConfig

> OperationResponse ModifyRFPlanningDeployConfig(ctx, omadacId, siteId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()

modify RF Planning Deploy Config



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
	rFPlanningDeployConfig := *openapiclient.NewRFPlanningDeployConfig(int32(123)) // RFPlanningDeployConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ModifyRFPlanningDeployConfig(context.Background(), omadacId, siteId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ModifyRFPlanningDeployConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRFPlanningDeployConfig`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ModifyRFPlanningDeployConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRFPlanningDeployConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rFPlanningDeployConfig** | [**RFPlanningDeployConfig**](RFPlanningDeployConfig.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyRFPlanningDeployConfigTemplate

> OperationResponse ModifyRFPlanningDeployConfigTemplate(ctx, omadacId, siteTemplateId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()

modify RF Planning Deploy Config Template



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
	siteTemplateId := "siteTemplateId_example" // string | Site Template ID
	rFPlanningDeployConfig := *openapiclient.NewRFPlanningDeployConfig(int32(123)) // RFPlanningDeployConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.ModifyRFPlanningDeployConfigTemplate(context.Background(), omadacId, siteTemplateId).RFPlanningDeployConfig(rFPlanningDeployConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.ModifyRFPlanningDeployConfigTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyRFPlanningDeployConfigTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.ModifyRFPlanningDeployConfigTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRFPlanningDeployConfigTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rFPlanningDeployConfig** | [**RFPlanningDeployConfig**](RFPlanningDeployConfig.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOptimization

> OperationResponse StartOptimization(ctx, omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()

start Optimization



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
	optimizationStrategy := *openapiclient.NewOptimizationStrategy(int32(123)) // OptimizationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.StartOptimization(context.Background(), omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.StartOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOptimization`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.StartOptimization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **optimizationStrategy** | [**OptimizationStrategy**](OptimizationStrategy.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRrmOptimization

> OperationResponse StartRrmOptimization(ctx, omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()

start Rrm Optimization



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
	optimizationStrategy := *openapiclient.NewOptimizationStrategy(int32(123)) // OptimizationStrategy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WLANOptimizationAPI.StartRrmOptimization(context.Background(), omadacId, siteId).OptimizationStrategy(optimizationStrategy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WLANOptimizationAPI.StartRrmOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartRrmOptimization`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WLANOptimizationAPI.StartRrmOptimization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartRrmOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **optimizationStrategy** | [**OptimizationStrategy**](OptimizationStrategy.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

