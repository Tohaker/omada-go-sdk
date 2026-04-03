# FormAuthDataAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthForm**](FormAuthDataAPI.md#createauthform) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys | Create a new authentication survey
[**DeleteAuthForm**](FormAuthDataAPI.md#deleteauthform) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys/{surveyId} | Delete an existing authentication survey
[**DeleteFormAuthResult**](FormAuthDataAPI.md#deleteformauthresult) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys/auth-results/{surveyResultId} | Delete an authentication survey result for given surveyResultId
[**DeleteSelectedFormAuthResult**](FormAuthDataAPI.md#deleteselectedformauthresult) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys/{surveyId}/auth-results | Delete selected authentication survey results for given surveyId
[**ExportAuthForms**](FormAuthDataAPI.md#exportauthforms) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/files/hotspot/form-auths/{formId}/export | Export the Form Authentication Result List
[**GetAuthForm**](FormAuthDataAPI.md#getauthform) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys/{surveyId} | Get an authentication survey for given surveyId
[**GetAuthForms**](FormAuthDataAPI.md#getauthforms) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys | Get authentication survey list
[**GetFormAuthResult**](FormAuthDataAPI.md#getformauthresult) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys/auth-results/{surveyResultId} | Get an authentication survey result for given surveyResultId
[**GetFormAuthResults**](FormAuthDataAPI.md#getformauthresults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys/{surveyId}/auth-results | Get authentication result lists for given survey
[**GetGridFormAuthResults**](FormAuthDataAPI.md#getgridformauthresults) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/form-auths/{formId}/auth-results | Get Form Authentication Result
[**ModifyAuthForm**](FormAuthDataAPI.md#modifyauthform) | **Patch** /openapi/v1/{omadacId}/sites/{siteId}/hotspot/surveys/{surveyId} | Modify an authentication survey for given surveyId



## CreateAuthForm

> OperationResponseCreatedResIdOpenApiVO CreateAuthForm(ctx, omadacId, siteId).CreateFormAuthOpenApiVO(createFormAuthOpenApiVO).Execute()

Create a new authentication survey



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
	createFormAuthOpenApiVO := *openapiclient.NewCreateFormAuthOpenApiVO(*openapiclient.NewAuthTimeOpenApiVO(int32(123)), []openapiclient.FormAuthCardOpenApiVO{*openapiclient.NewFormAuthCardOpenApiVO([]string{"Choices_example"}, false, "Title_example", int32(123))}, "Name_example", "Note_example", false, "Title_example") // CreateFormAuthOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.CreateAuthForm(context.Background(), omadacId, siteId).CreateFormAuthOpenApiVO(createFormAuthOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.CreateAuthForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuthForm`: OperationResponseCreatedResIdOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.CreateAuthForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createFormAuthOpenApiVO** | [**CreateFormAuthOpenApiVO**](CreateFormAuthOpenApiVO.md) |  | 

### Return type

[**OperationResponseCreatedResIdOpenApiVO**](OperationResponseCreatedResIdOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthForm

> OperationResponseWithoutResult DeleteAuthForm(ctx, omadacId, siteId, surveyId).Execute()

Delete an existing authentication survey



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
	surveyId := "surveyId_example" // string | Auth survey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.DeleteAuthForm(context.Background(), omadacId, siteId, surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.DeleteAuthForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAuthForm`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.DeleteAuthForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**surveyId** | **string** | Auth survey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthFormRequest struct via the builder pattern


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


## DeleteFormAuthResult

> OperationResponseWithoutResult DeleteFormAuthResult(ctx, omadacId, siteId, surveyResultId).Execute()

Delete an authentication survey result for given surveyResultId



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
	surveyResultId := "surveyResultId_example" // string | Survey authentication result ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.DeleteFormAuthResult(context.Background(), omadacId, siteId, surveyResultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.DeleteFormAuthResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFormAuthResult`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.DeleteFormAuthResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**surveyResultId** | **string** | Survey authentication result ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormAuthResultRequest struct via the builder pattern


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


## DeleteSelectedFormAuthResult

> OperationResponseWithoutResult DeleteSelectedFormAuthResult(ctx, omadacId, siteId, surveyId).FormAuthResultSelector(formAuthResultSelector).Execute()

Delete selected authentication survey results for given surveyId



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
	surveyId := "surveyId_example" // string | Survey authentication ID
	formAuthResultSelector := *openapiclient.NewFormAuthResultSelector("Type_example") // FormAuthResultSelector | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.DeleteSelectedFormAuthResult(context.Background(), omadacId, siteId, surveyId).FormAuthResultSelector(formAuthResultSelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.DeleteSelectedFormAuthResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSelectedFormAuthResult`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.DeleteSelectedFormAuthResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**surveyId** | **string** | Survey authentication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSelectedFormAuthResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **formAuthResultSelector** | [**FormAuthResultSelector**](FormAuthResultSelector.md) |  | 

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


## ExportAuthForms

> map[string]interface{} ExportAuthForms(ctx, omadacId, siteId, formId).ExportFormVO(exportFormVO).Execute()

Export the Form Authentication Result List



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
	formId := "formId_example" // string | Form auth Id.
	exportFormVO := *openapiclient.NewExportFormOpenApiVO("Ids_example", "Type_example") // ExportFormOpenApiVO | Export Form Parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.ExportAuthForms(context.Background(), omadacId, siteId, formId).ExportFormVO(exportFormVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.ExportAuthForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAuthForms`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.ExportAuthForms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**formId** | **string** | Form auth Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAuthFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **exportFormVO** | [**ExportFormOpenApiVO**](ExportFormOpenApiVO.md) | Export Form Parameters | 

### Return type

**map[string]interface{}**

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthForm

> OperationResponseFormAuthOpenApiVO GetAuthForm(ctx, omadacId, siteId, surveyId).Execute()

Get an authentication survey for given surveyId



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
	surveyId := "surveyId_example" // string | Auth survey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.GetAuthForm(context.Background(), omadacId, siteId, surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.GetAuthForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthForm`: OperationResponseFormAuthOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.GetAuthForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**surveyId** | **string** | Auth survey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseFormAuthOpenApiVO**](OperationResponseFormAuthOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthForms

> OperationResponseListFormAuthOpenApiVO GetAuthForms(ctx, omadacId, siteId).SortsName(sortsName).SortsCreateTime(sortsCreateTime).SortsAnswerNum(sortsAnswerNum).Execute()

Get authentication survey list



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
	sortsName := "sortsName_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsCreateTime := "sortsCreateTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsAnswerNum := "sortsAnswerNum_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.GetAuthForms(context.Background(), omadacId, siteId).SortsName(sortsName).SortsCreateTime(sortsCreateTime).SortsAnswerNum(sortsAnswerNum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.GetAuthForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthForms`: OperationResponseListFormAuthOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.GetAuthForms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortsName** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsCreateTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsAnswerNum** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseListFormAuthOpenApiVO**](OperationResponseListFormAuthOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormAuthResult

> OperationResponseFormAuthResultOpenApiVO GetFormAuthResult(ctx, omadacId, siteId, surveyResultId).Execute()

Get an authentication survey result for given surveyResultId



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
	surveyResultId := "surveyResultId_example" // string | Survey authentication result ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.GetFormAuthResult(context.Background(), omadacId, siteId, surveyResultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.GetFormAuthResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormAuthResult`: OperationResponseFormAuthResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.GetFormAuthResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**surveyResultId** | **string** | Survey authentication result ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormAuthResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseFormAuthResultOpenApiVO**](OperationResponseFormAuthResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormAuthResults

> OperationResponseGridVOFormAuthResultOpenApiVO GetFormAuthResults(ctx, omadacId, siteId, surveyId).Page(page).PageSize(pageSize).SortsFormAuth(sortsFormAuth).Execute()

Get authentication result lists for given survey



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
	surveyId := "surveyId_example" // string | Auth survey ID
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsFormAuth := "sortsFormAuth_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.GetFormAuthResults(context.Background(), omadacId, siteId, surveyId).Page(page).PageSize(pageSize).SortsFormAuth(sortsFormAuth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.GetFormAuthResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormAuthResults`: OperationResponseGridVOFormAuthResultOpenApiVO
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.GetFormAuthResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**surveyId** | **string** | Auth survey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormAuthResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsFormAuth** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseGridVOFormAuthResultOpenApiVO**](OperationResponseGridVOFormAuthResultOpenApiVO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridFormAuthResults

> OperationResponse GetGridFormAuthResults(ctx, omadacId, siteId, formId).QueryDataVO(queryDataVO).Execute()

Get Form Authentication Result



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
	formId := "formId_example" // string | Form auth Id.
	queryDataVO := *openapiclient.NewOpenApiQueryDataVO(int32(123), int32(123)) // OpenApiQueryDataVO | Query Data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.GetGridFormAuthResults(context.Background(), omadacId, siteId, formId).QueryDataVO(queryDataVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.GetGridFormAuthResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridFormAuthResults`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.GetGridFormAuthResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**formId** | **string** | Form auth Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridFormAuthResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **queryDataVO** | [**OpenApiQueryDataVO**](OpenApiQueryDataVO.md) | Query Data | 

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


## ModifyAuthForm

> OperationResponseWithoutResult ModifyAuthForm(ctx, omadacId, siteId, surveyId).CreateFormAuthOpenApiVO(createFormAuthOpenApiVO).Execute()

Modify an authentication survey for given surveyId



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
	surveyId := "surveyId_example" // string | Survey authentication ID
	createFormAuthOpenApiVO := *openapiclient.NewCreateFormAuthOpenApiVO(*openapiclient.NewAuthTimeOpenApiVO(int32(123)), []openapiclient.FormAuthCardOpenApiVO{*openapiclient.NewFormAuthCardOpenApiVO([]string{"Choices_example"}, false, "Title_example", int32(123))}, "Name_example", "Note_example", false, "Title_example") // CreateFormAuthOpenApiVO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAuthDataAPI.ModifyAuthForm(context.Background(), omadacId, siteId, surveyId).CreateFormAuthOpenApiVO(createFormAuthOpenApiVO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAuthDataAPI.ModifyAuthForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAuthForm`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FormAuthDataAPI.ModifyAuthForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**surveyId** | **string** | Survey authentication ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAuthFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createFormAuthOpenApiVO** | [**CreateFormAuthOpenApiVO**](CreateFormAuthOpenApiVO.md) |  | 

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

