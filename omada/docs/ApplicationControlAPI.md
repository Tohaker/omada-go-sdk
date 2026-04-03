# ApplicationControlAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAssignRestriction**](ApplicationControlAPI.md#addassignrestriction) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/restrictions | Create new restriction assigned to lan network
[**AddAssignRestrictionTemplate**](ApplicationControlAPI.md#addassignrestrictiontemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/restrictions | Create new restriction assigned to lan network
[**AddFilter**](ApplicationControlAPI.md#addfilter) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/filters | Create new filter
[**AddFilterTemplate**](ApplicationControlAPI.md#addfiltertemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/filters | Create new filter
[**AddRule**](ApplicationControlAPI.md#addrule) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/rules | Create new rule
[**AddRuleTemplate**](ApplicationControlAPI.md#addruletemplate) | **Post** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/rules | Create new rule
[**ClearDpiData**](ApplicationControlAPI.md#cleardpidata) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/stat/dpi | Clear DPI data
[**DeleteAssignRestrictions**](ApplicationControlAPI.md#deleteassignrestrictions) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/restrictions/{restrictionId} | Delete an existing restriction
[**DeleteAssignRestrictionsTemplate**](ApplicationControlAPI.md#deleteassignrestrictionstemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/restrictions/{restrictionId} | Delete an existing restriction
[**DeleteFilters**](ApplicationControlAPI.md#deletefilters) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/filters/{filterId} | Delete an existing filter
[**DeleteFiltersTemplate**](ApplicationControlAPI.md#deletefilterstemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/filters/{filterId} | Delete an existing filter
[**DeleteRules**](ApplicationControlAPI.md#deleterules) | **Delete** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/rules/{ruleId} | Delete an existing rule
[**DeleteRulesTemplate**](ApplicationControlAPI.md#deleterulestemplate) | **Delete** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/rules/{ruleId} | Delete an existing rule
[**EditApplicationControlStatus**](ApplicationControlAPI.md#editapplicationcontrolstatus) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/status | Modify application control general settings
[**EditApplicationControlStatusTemplate**](ApplicationControlAPI.md#editapplicationcontrolstatustemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/status | Modify application control general settings
[**EditAssignRestrictions**](ApplicationControlAPI.md#editassignrestrictions) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/restrictions/{restrictionId} | Modify an existing restriction
[**EditAssignRestrictionsTemplate**](ApplicationControlAPI.md#editassignrestrictionstemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/restrictions/{restrictionId} | Modify an existing restriction
[**EditFilters**](ApplicationControlAPI.md#editfilters) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/filters/{filterId} | Modify an existing filter
[**EditFiltersTemplate**](ApplicationControlAPI.md#editfilterstemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/filters/{filterId} | Modify an existing filter
[**EditRule**](ApplicationControlAPI.md#editrule) | **Put** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/rules/{ruleId} | Modify an existing rule
[**EditRulesTemplate**](ApplicationControlAPI.md#editrulestemplate) | **Put** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/rules/{ruleId} | Modify an existing rule
[**GetAllowApp**](ApplicationControlAPI.md#getallowapp) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/allowApp | Get allow app
[**GetApplicationControlStatus**](ApplicationControlAPI.md#getapplicationcontrolstatus) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/status | Get application control general settings
[**GetApplicationControlStatusTemplate**](ApplicationControlAPI.md#getapplicationcontrolstatustemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/status | Get application control general settings
[**GetApplications**](ApplicationControlAPI.md#getapplications) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/applications | Get application list
[**GetApplicationsTemplate**](ApplicationControlAPI.md#getapplicationstemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/applications | Get application list
[**GetAssignRestrictions**](ApplicationControlAPI.md#getassignrestrictions) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/restrictions | Get application control assigned restriction list
[**GetAssignRestrictionsTemplate**](ApplicationControlAPI.md#getassignrestrictionstemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/restrictions | Get application control assigned restriction list
[**GetBlockApp**](ApplicationControlAPI.md#getblockapp) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/blockApp | Get block app
[**GetCategoryAppInfo**](ApplicationControlAPI.md#getcategoryappinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/categoryAppInfo/{familyId} | Get app traffic in specific category
[**GetCategoryInfo**](ApplicationControlAPI.md#getcategoryinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/categoryInfo | Get category info
[**GetCategoryUserInfo**](ApplicationControlAPI.md#getcategoryuserinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/categoryUserInfo/{familyId} | Get client traffic in specific category
[**GetClientTraffic**](ApplicationControlAPI.md#getclienttraffic) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/clientTraffic | Get clients traffic
[**GetFamilies**](ApplicationControlAPI.md#getfamilies) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/families | Get family list
[**GetFamiliesTemplate**](ApplicationControlAPI.md#getfamiliestemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/families | Get family list
[**GetFilters**](ApplicationControlAPI.md#getfilters) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/filters | Get filter list
[**GetFiltersTemplate**](ApplicationControlAPI.md#getfilterstemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/filters | Get filter list
[**GetMostActiveApplications**](ApplicationControlAPI.md#getmostactiveapplications) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/mostActiveAppTraffic | Get most active applications
[**GetRules**](ApplicationControlAPI.md#getrules) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/applicationControl/rules | Get rule list
[**GetRulesTemplate**](ApplicationControlAPI.md#getrulestemplate) | **Get** /openapi/v1/{omadacId}/sitetemplates/{siteTemplateId}/applicationControl/rules | Get rule list
[**GetSpecificAppInfo**](ApplicationControlAPI.md#getspecificappinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/specificAppInfo/{applicationId} | Get client traffic in specific app
[**GetSpecificClientInfo**](ApplicationControlAPI.md#getspecificclientinfo) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/dashboard/specificClientInfo/{clientMac} | Get app traffic in specific client



## AddAssignRestriction

> OperationResponseRestrictionEntity AddAssignRestriction(ctx, omadacId, siteId).RestrictionEntity(restrictionEntity).Execute()

Create new restriction assigned to lan network



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
	restrictionEntity := *openapiclient.NewRestrictionEntity(int32(123), "NetworkName_example") // RestrictionEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.AddAssignRestriction(context.Background(), omadacId, siteId).RestrictionEntity(restrictionEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.AddAssignRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAssignRestriction`: OperationResponseRestrictionEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.AddAssignRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAssignRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restrictionEntity** | [**RestrictionEntity**](RestrictionEntity.md) |  | 

### Return type

[**OperationResponseRestrictionEntity**](OperationResponseRestrictionEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAssignRestrictionTemplate

> OperationResponseRestrictionEntity AddAssignRestrictionTemplate(ctx, omadacId, siteTemplateId).RestrictionEntity(restrictionEntity).Execute()

Create new restriction assigned to lan network



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
	restrictionEntity := *openapiclient.NewRestrictionEntity(int32(123), "NetworkName_example") // RestrictionEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.AddAssignRestrictionTemplate(context.Background(), omadacId, siteTemplateId).RestrictionEntity(restrictionEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.AddAssignRestrictionTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAssignRestrictionTemplate`: OperationResponseRestrictionEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.AddAssignRestrictionTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAssignRestrictionTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restrictionEntity** | [**RestrictionEntity**](RestrictionEntity.md) |  | 

### Return type

[**OperationResponseRestrictionEntity**](OperationResponseRestrictionEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFilter

> OperationResponseFilterEntity AddFilter(ctx, omadacId, siteId).AddFilterEntity(addFilterEntity).Execute()

Create new filter



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
	addFilterEntity := *openapiclient.NewAddFilterEntity("Description_example", "FilterName_example", []int32{int32(123)}) // AddFilterEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.AddFilter(context.Background(), omadacId, siteId).AddFilterEntity(addFilterEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.AddFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFilter`: OperationResponseFilterEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.AddFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addFilterEntity** | [**AddFilterEntity**](AddFilterEntity.md) |  | 

### Return type

[**OperationResponseFilterEntity**](OperationResponseFilterEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFilterTemplate

> OperationResponseFilterEntity AddFilterTemplate(ctx, omadacId, siteTemplateId).AddFilterEntity(addFilterEntity).Execute()

Create new filter



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
	addFilterEntity := *openapiclient.NewAddFilterEntity("Description_example", "FilterName_example", []int32{int32(123)}) // AddFilterEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.AddFilterTemplate(context.Background(), omadacId, siteTemplateId).AddFilterEntity(addFilterEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.AddFilterTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFilterTemplate`: OperationResponseFilterEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.AddFilterTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFilterTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addFilterEntity** | [**AddFilterEntity**](AddFilterEntity.md) |  | 

### Return type

[**OperationResponseFilterEntity**](OperationResponseFilterEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRule

> OperationResponseRuleEntity AddRule(ctx, omadacId, siteId).AddRuleEntity(addRuleEntity).Execute()

Create new rule



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
	addRuleEntity := *openapiclient.NewAddRuleEntity([]int32{int32(123)}, false, "RuleName_example", "Schedule_example", "SelectType_example") // AddRuleEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.AddRule(context.Background(), omadacId, siteId).AddRuleEntity(addRuleEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.AddRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRule`: OperationResponseRuleEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.AddRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addRuleEntity** | [**AddRuleEntity**](AddRuleEntity.md) |  | 

### Return type

[**OperationResponseRuleEntity**](OperationResponseRuleEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRuleTemplate

> OperationResponseRuleEntity AddRuleTemplate(ctx, omadacId, siteTemplateId).AddRuleEntity(addRuleEntity).Execute()

Create new rule



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
	addRuleEntity := *openapiclient.NewAddRuleEntity([]int32{int32(123)}, false, "RuleName_example", "Schedule_example", "SelectType_example") // AddRuleEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.AddRuleTemplate(context.Background(), omadacId, siteTemplateId).AddRuleEntity(addRuleEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.AddRuleTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRuleTemplate`: OperationResponseRuleEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.AddRuleTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRuleTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addRuleEntity** | [**AddRuleEntity**](AddRuleEntity.md) |  | 

### Return type

[**OperationResponseRuleEntity**](OperationResponseRuleEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearDpiData

> OperationResponseListFamilyEntity ClearDpiData(ctx, omadacId, siteId).ClearDpiDataRequest(clearDpiDataRequest).Execute()

Clear DPI data



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
	clearDpiDataRequest := *openapiclient.NewClearDpiDataRequest(int64(123), int64(123)) // ClearDpiDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.ClearDpiData(context.Background(), omadacId, siteId).ClearDpiDataRequest(clearDpiDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.ClearDpiData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearDpiData`: OperationResponseListFamilyEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.ClearDpiData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearDpiDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clearDpiDataRequest** | [**ClearDpiDataRequest**](ClearDpiDataRequest.md) |  | 

### Return type

[**OperationResponseListFamilyEntity**](OperationResponseListFamilyEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssignRestrictions

> OperationResponse DeleteAssignRestrictions(ctx, omadacId, siteId, restrictionId).Execute()

Delete an existing restriction



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
	restrictionId := "restrictionId_example" // string | Restriction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.DeleteAssignRestrictions(context.Background(), omadacId, siteId, restrictionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.DeleteAssignRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssignRestrictions`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.DeleteAssignRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**restrictionId** | **string** | Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssignRestrictionsRequest struct via the builder pattern


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


## DeleteAssignRestrictionsTemplate

> OperationResponse DeleteAssignRestrictionsTemplate(ctx, omadacId, siteTemplateId, restrictionId).Execute()

Delete an existing restriction



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
	restrictionId := "restrictionId_example" // string | Restriction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.DeleteAssignRestrictionsTemplate(context.Background(), omadacId, siteTemplateId, restrictionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.DeleteAssignRestrictionsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssignRestrictionsTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.DeleteAssignRestrictionsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**restrictionId** | **string** | Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssignRestrictionsTemplateRequest struct via the builder pattern


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


## DeleteFilters

> OperationResponse DeleteFilters(ctx, omadacId, siteId, filterId).Execute()

Delete an existing filter



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
	filterId := "filterId_example" // string | Filter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.DeleteFilters(context.Background(), omadacId, siteId, filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.DeleteFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFilters`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.DeleteFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**filterId** | **string** | Filter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiltersRequest struct via the builder pattern


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


## DeleteFiltersTemplate

> OperationResponse DeleteFiltersTemplate(ctx, omadacId, siteTemplateId, filterId).Execute()

Delete an existing filter



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
	filterId := "filterId_example" // string | Filter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.DeleteFiltersTemplate(context.Background(), omadacId, siteTemplateId, filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.DeleteFiltersTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiltersTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.DeleteFiltersTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**filterId** | **string** | Filter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiltersTemplateRequest struct via the builder pattern


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


## DeleteRules

> OperationResponse DeleteRules(ctx, omadacId, siteId, ruleId).Execute()

Delete an existing rule



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
	ruleId := "ruleId_example" // string | Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.DeleteRules(context.Background(), omadacId, siteId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.DeleteRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRules`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.DeleteRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRulesRequest struct via the builder pattern


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


## DeleteRulesTemplate

> OperationResponse DeleteRulesTemplate(ctx, omadacId, siteTemplateId, ruleId).Execute()

Delete an existing rule



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
	ruleId := "ruleId_example" // string | Rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.DeleteRulesTemplate(context.Background(), omadacId, siteTemplateId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.DeleteRulesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRulesTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.DeleteRulesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRulesTemplateRequest struct via the builder pattern


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


## EditApplicationControlStatus

> OperationResponse EditApplicationControlStatus(ctx, omadacId, siteId).DpiSettings(dpiSettings).Execute()

Modify application control general settings



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
	dpiSettings := *openapiclient.NewDpiSettings(false, false) // DpiSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditApplicationControlStatus(context.Background(), omadacId, siteId).DpiSettings(dpiSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditApplicationControlStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditApplicationControlStatus`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditApplicationControlStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationControlStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dpiSettings** | [**DpiSettings**](DpiSettings.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditApplicationControlStatusTemplate

> OperationResponse EditApplicationControlStatusTemplate(ctx, omadacId, siteTemplateId).DpiSettings(dpiSettings).Execute()

Modify application control general settings



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
	dpiSettings := *openapiclient.NewDpiSettings(false, false) // DpiSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditApplicationControlStatusTemplate(context.Background(), omadacId, siteTemplateId).DpiSettings(dpiSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditApplicationControlStatusTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditApplicationControlStatusTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditApplicationControlStatusTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditApplicationControlStatusTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dpiSettings** | [**DpiSettings**](DpiSettings.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAssignRestrictions

> OperationResponse EditAssignRestrictions(ctx, omadacId, siteId, restrictionId).RestrictionEntity(restrictionEntity).Execute()

Modify an existing restriction



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
	restrictionId := "restrictionId_example" // string | Restriction ID
	restrictionEntity := *openapiclient.NewRestrictionEntity(int32(123), "NetworkName_example") // RestrictionEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditAssignRestrictions(context.Background(), omadacId, siteId, restrictionId).RestrictionEntity(restrictionEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditAssignRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAssignRestrictions`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditAssignRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**restrictionId** | **string** | Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAssignRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restrictionEntity** | [**RestrictionEntity**](RestrictionEntity.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAssignRestrictionsTemplate

> OperationResponse EditAssignRestrictionsTemplate(ctx, omadacId, siteTemplateId, restrictionId).RestrictionEntity(restrictionEntity).Execute()

Modify an existing restriction



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
	restrictionId := "restrictionId_example" // string | Restriction ID
	restrictionEntity := *openapiclient.NewRestrictionEntity(int32(123), "NetworkName_example") // RestrictionEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditAssignRestrictionsTemplate(context.Background(), omadacId, siteTemplateId, restrictionId).RestrictionEntity(restrictionEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditAssignRestrictionsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAssignRestrictionsTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditAssignRestrictionsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**restrictionId** | **string** | Restriction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAssignRestrictionsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restrictionEntity** | [**RestrictionEntity**](RestrictionEntity.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditFilters

> OperationResponse EditFilters(ctx, omadacId, siteId, filterId).EditFilterEntity(editFilterEntity).Execute()

Modify an existing filter



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
	filterId := "filterId_example" // string | Filter ID
	editFilterEntity := *openapiclient.NewEditFilterEntity("FilterName_example", []int32{int32(123)}) // EditFilterEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditFilters(context.Background(), omadacId, siteId, filterId).EditFilterEntity(editFilterEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditFilters`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**filterId** | **string** | Filter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **editFilterEntity** | [**EditFilterEntity**](EditFilterEntity.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditFiltersTemplate

> OperationResponse EditFiltersTemplate(ctx, omadacId, siteTemplateId, filterId).EditFilterEntity(editFilterEntity).Execute()

Modify an existing filter



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
	filterId := "filterId_example" // string | Filter ID
	editFilterEntity := *openapiclient.NewEditFilterEntity("FilterName_example", []int32{int32(123)}) // EditFilterEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditFiltersTemplate(context.Background(), omadacId, siteTemplateId, filterId).EditFilterEntity(editFilterEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditFiltersTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditFiltersTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditFiltersTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**filterId** | **string** | Filter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditFiltersTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **editFilterEntity** | [**EditFilterEntity**](EditFilterEntity.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRule

> OperationResponse EditRule(ctx, omadacId, siteId, ruleId).EditRuleEntity(editRuleEntity).Execute()

Modify an existing rule



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
	ruleId := "ruleId_example" // string | Rule ID
	editRuleEntity := *openapiclient.NewEditRuleEntity([]int32{int32(123)}, false, "RuleName_example", "Schedule_example") // EditRuleEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditRule(context.Background(), omadacId, siteId, ruleId).EditRuleEntity(editRuleEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRule`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **editRuleEntity** | [**EditRuleEntity**](EditRuleEntity.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRulesTemplate

> OperationResponse EditRulesTemplate(ctx, omadacId, siteTemplateId, ruleId).EditRuleEntity(editRuleEntity).Execute()

Modify an existing rule



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
	ruleId := "ruleId_example" // string | Rule ID
	editRuleEntity := *openapiclient.NewEditRuleEntity([]int32{int32(123)}, false, "RuleName_example", "Schedule_example") // EditRuleEntity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.EditRulesTemplate(context.Background(), omadacId, siteTemplateId, ruleId).EditRuleEntity(editRuleEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.EditRulesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRulesTemplate`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.EditRulesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 
**ruleId** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRulesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **editRuleEntity** | [**EditRuleEntity**](EditRuleEntity.md) |  | 

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowApp

> OperationResponseApplicationGridVOApplicationStatTraffic GetAllowApp(ctx, omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()

Get allow app



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetAllowApp(context.Background(), omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetAllowApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllowApp`: OperationResponseApplicationGridVOApplicationStatTraffic
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetAllowApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseApplicationGridVOApplicationStatTraffic**](OperationResponseApplicationGridVOApplicationStatTraffic.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationControlStatus

> OperationResponseDpiSettings GetApplicationControlStatus(ctx, omadacId, siteId).Execute()

Get application control general settings



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
	resp, r, err := apiClient.ApplicationControlAPI.GetApplicationControlStatus(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetApplicationControlStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationControlStatus`: OperationResponseDpiSettings
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetApplicationControlStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationControlStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDpiSettings**](OperationResponseDpiSettings.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationControlStatusTemplate

> OperationResponseDpiSettings GetApplicationControlStatusTemplate(ctx, omadacId, siteTemplateId).Execute()

Get application control general settings



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
	resp, r, err := apiClient.ApplicationControlAPI.GetApplicationControlStatusTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetApplicationControlStatusTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationControlStatusTemplate`: OperationResponseDpiSettings
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetApplicationControlStatusTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationControlStatusTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseDpiSettings**](OperationResponseDpiSettings.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> OperationResponseGridVOApplicationEntity GetApplications(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersFamilyId(filtersFamilyId).Execute()

Get application list



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field application name (optional)
	filtersFamilyId := int32(56) // int32 | Filter query parameters, support field 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetApplications(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersFamilyId(filtersFamilyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplications`: OperationResponseGridVOApplicationEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field application name | 
 **filtersFamilyId** | **int32** | Filter query parameters, support field 0 | 

### Return type

[**OperationResponseGridVOApplicationEntity**](OperationResponseGridVOApplicationEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationsTemplate

> OperationResponseGridVOApplicationEntity GetApplicationsTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersFamilyId(filtersFamilyId).Execute()

Get application list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field application name (optional)
	filtersFamilyId := int32(56) // int32 | Filter query parameters, support field 0 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetApplicationsTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).FiltersFamilyId(filtersFamilyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetApplicationsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationsTemplate`: OperationResponseGridVOApplicationEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetApplicationsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field application name | 
 **filtersFamilyId** | **int32** | Filter query parameters, support field 0 | 

### Return type

[**OperationResponseGridVOApplicationEntity**](OperationResponseGridVOApplicationEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssignRestrictions

> OperationResponseGridVORestrictionEntity GetAssignRestrictions(ctx, omadacId, siteId).Page(page).PageSize(pageSize).Execute()

Get application control assigned restriction list



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
	resp, r, err := apiClient.ApplicationControlAPI.GetAssignRestrictions(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetAssignRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssignRestrictions`: OperationResponseGridVORestrictionEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetAssignRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVORestrictionEntity**](OperationResponseGridVORestrictionEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssignRestrictionsTemplate

> OperationResponseGridVORestrictionEntity GetAssignRestrictionsTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()

Get application control assigned restriction list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetAssignRestrictionsTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetAssignRestrictionsTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssignRestrictionsTemplate`: OperationResponseGridVORestrictionEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetAssignRestrictionsTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignRestrictionsTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVORestrictionEntity**](OperationResponseGridVORestrictionEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockApp

> OperationResponseApplicationGridVOApplicationStatTraffic GetBlockApp(ctx, omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()

Get block app



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetBlockApp(context.Background(), omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetBlockApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockApp`: OperationResponseApplicationGridVOApplicationStatTraffic
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetBlockApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseApplicationGridVOApplicationStatTraffic**](OperationResponseApplicationGridVOApplicationStatTraffic.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryAppInfo

> OperationResponseListApplicationStatTraffic GetCategoryAppInfo(ctx, omadacId, siteId, familyId).Start(start).End(end).Execute()

Get app traffic in specific category



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
	familyId := "familyId_example" // string | Family Id
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetCategoryAppInfo(context.Background(), omadacId, siteId, familyId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetCategoryAppInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryAppInfo`: OperationResponseListApplicationStatTraffic
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetCategoryAppInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**familyId** | **string** | Family Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryAppInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListApplicationStatTraffic**](OperationResponseListApplicationStatTraffic.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryInfo

> OperationResponseApplicationGridVOCategoryTraffic GetCategoryInfo(ctx, omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()

Get category info



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetCategoryInfo(context.Background(), omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetCategoryInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryInfo`: OperationResponseApplicationGridVOCategoryTraffic
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetCategoryInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseApplicationGridVOCategoryTraffic**](OperationResponseApplicationGridVOCategoryTraffic.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoryUserInfo

> OperationResponseListClientTrafficWithApplicationDetail GetCategoryUserInfo(ctx, omadacId, siteId, familyId).Start(start).End(end).Execute()

Get client traffic in specific category



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
	familyId := "familyId_example" // string | Family Id
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetCategoryUserInfo(context.Background(), omadacId, siteId, familyId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetCategoryUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoryUserInfo`: OperationResponseListClientTrafficWithApplicationDetail
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetCategoryUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**familyId** | **string** | Family Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListClientTrafficWithApplicationDetail**](OperationResponseListClientTrafficWithApplicationDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientTraffic

> OperationResponseGridVOClientTrafficWithApplicationDetail GetClientTraffic(ctx, omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()

Get clients traffic



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetClientTraffic(context.Background(), omadacId, siteId).Start(start).End(end).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetClientTraffic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientTraffic`: OperationResponseGridVOClientTrafficWithApplicationDetail
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetClientTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOClientTrafficWithApplicationDetail**](OperationResponseGridVOClientTrafficWithApplicationDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFamilies

> OperationResponseListFamilyEntity GetFamilies(ctx, omadacId, siteId).Execute()

Get family list



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
	resp, r, err := apiClient.ApplicationControlAPI.GetFamilies(context.Background(), omadacId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetFamilies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFamilies`: OperationResponseListFamilyEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetFamilies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFamiliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListFamilyEntity**](OperationResponseListFamilyEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFamiliesTemplate

> OperationResponseListFamilyEntity GetFamiliesTemplate(ctx, omadacId, siteTemplateId).Execute()

Get family list



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
	resp, r, err := apiClient.ApplicationControlAPI.GetFamiliesTemplate(context.Background(), omadacId, siteTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetFamiliesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFamiliesTemplate`: OperationResponseListFamilyEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetFamiliesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFamiliesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseListFamilyEntity**](OperationResponseListFamilyEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilters

> OperationResponseGridVOFilterEntity GetFilters(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get filter list



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field filter name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetFilters(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilters`: OperationResponseGridVOFilterEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field filter name | 

### Return type

[**OperationResponseGridVOFilterEntity**](OperationResponseGridVOFilterEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiltersTemplate

> OperationResponseGridVOFilterEntity GetFiltersTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get filter list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field filter name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetFiltersTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetFiltersTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiltersTemplate`: OperationResponseGridVOFilterEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetFiltersTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiltersTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field filter name | 

### Return type

[**OperationResponseGridVOFilterEntity**](OperationResponseGridVOFilterEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMostActiveApplications

> OperationResponseListApplicationTrafficWithClientCount GetMostActiveApplications(ctx, omadacId, siteId).Start(start).End(end).Execute()

Get most active applications



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
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetMostActiveApplications(context.Background(), omadacId, siteId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetMostActiveApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMostActiveApplications`: OperationResponseListApplicationTrafficWithClientCount
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetMostActiveApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMostActiveApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListApplicationTrafficWithClientCount**](OperationResponseListApplicationTrafficWithClientCount.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRules

> OperationResponseGridVORuleEntity GetRules(ctx, omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get rule list



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
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field rule name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetRules(context.Background(), omadacId, siteId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRules`: OperationResponseGridVORuleEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field rule name | 

### Return type

[**OperationResponseGridVORuleEntity**](OperationResponseGridVORuleEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesTemplate

> OperationResponseGridVORuleEntity GetRulesTemplate(ctx, omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()

Get rule list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field rule name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetRulesTemplate(context.Background(), omadacId, siteTemplateId).Page(page).PageSize(pageSize).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetRulesTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRulesTemplate`: OperationResponseGridVORuleEntity
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetRulesTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteTemplateId** | **string** | Site Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **searchKey** | **string** | Fuzzy query parameters, support field rule name | 

### Return type

[**OperationResponseGridVORuleEntity**](OperationResponseGridVORuleEntity.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificAppInfo

> OperationResponseListClientTrafficWithApplicationDetail GetSpecificAppInfo(ctx, omadacId, siteId, applicationId).Start(start).End(end).Execute()

Get client traffic in specific app



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
	applicationId := "applicationId_example" // string | Application Id
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetSpecificAppInfo(context.Background(), omadacId, siteId, applicationId).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetSpecificAppInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecificAppInfo`: OperationResponseListClientTrafficWithApplicationDetail
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetSpecificAppInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**applicationId** | **string** | Application Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificAppInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListClientTrafficWithApplicationDetail**](OperationResponseListClientTrafficWithApplicationDetail.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecificClientInfo

> OperationResponseListApplicationStatTraffic GetSpecificClientInfo(ctx, omadacId, siteId, clientMac).Start(start).End(end).Execute()

Get app traffic in specific client



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
	clientMac := "clientMac_example" // string | Client Mac
	start := int64(789) // int64 | Start timestamp, in seconds, such as 1682000000
	end := int64(789) // int64 | End timestamp, in seconds, such as 1682000000

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationControlAPI.GetSpecificClientInfo(context.Background(), omadacId, siteId, clientMac).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationControlAPI.GetSpecificClientInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecificClientInfo`: OperationResponseListApplicationStatTraffic
	fmt.Fprintf(os.Stdout, "Response from `ApplicationControlAPI.GetSpecificClientInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**siteId** | **string** | Site ID | 
**clientMac** | **string** | Client Mac | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecificClientInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int64** | Start timestamp, in seconds, such as 1682000000 | 
 **end** | **int64** | End timestamp, in seconds, such as 1682000000 | 

### Return type

[**OperationResponseListApplicationStatTraffic**](OperationResponseListApplicationStatTraffic.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

