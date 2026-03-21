# \FirmwareAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUploadUpgradeFirmwareLimit**](FirmwareAPI.md#CheckUploadUpgradeFirmwareLimit) | **Post** /openapi/v1/{omadacId}/upgrade/firmware-size/{firmwareSize} | Check the limit of the firmware pool
[**CreateAutoCheckUpgrade**](FirmwareAPI.md#CreateAutoCheckUpgrade) | **Post** /openapi/v1/{omadacId}/upgrade/autoCheck | Create autocheck upgrade plan
[**CreateFailedDeviceModelPlan**](FirmwareAPI.md#CreateFailedDeviceModelPlan) | **Post** /openapi/v1/{omadacId}/logs/{upgradeLogId}/upgrade/overview/failed-model-plan | Create Upgrade Plan for Failed Device
[**CreateFirmwareUpgradePlan**](FirmwareAPI.md#CreateFirmwareUpgradePlan) | **Post** /openapi/v1/{omadacId}/firmwares/{firmwareId}/upgrade/plan | Create firmware upgrade plan based on manually uploaded firmware
[**CreateRollback**](FirmwareAPI.md#CreateRollback) | **Post** /openapi/v1/{omadacId}/logs/{upgradeLogId}/upgrade/overview/rollback | Create rollback plan
[**CreateUpgradePlan**](FirmwareAPI.md#CreateUpgradePlan) | **Post** /openapi/v1/{omadacId}/upgrade/overview/plans | Create firmware upgrade plan
[**DeleteAutoCheckUpgrade**](FirmwareAPI.md#DeleteAutoCheckUpgrade) | **Delete** /openapi/v1/{omadacId}/autoCheck/{autoCheckId}/upgrade | Delete autoCheck upgrade plan
[**DeleteUpgradeFirmware**](FirmwareAPI.md#DeleteUpgradeFirmware) | **Delete** /openapi/v1/{omadacId}/firmwares/{firmwareId}/upgrade | Delete upgrade firmware
[**DeleteUpgradePlan**](FirmwareAPI.md#DeleteUpgradePlan) | **Delete** /openapi/v1/{omadacId}/plans/{planId}/upgrade/overview | Delete firmware upgrade plan
[**EditAutoCheckUpgrade**](FirmwareAPI.md#EditAutoCheckUpgrade) | **Patch** /openapi/v1/{omadacId}/autoCheck/{autoCheckId}/upgrade | Edit autoCheck upgrade plan
[**EditFirmwareUpgrade**](FirmwareAPI.md#EditFirmwareUpgrade) | **Post** /openapi/v1/{omadacId}/upgrade/firmwares/{firmwareId} | Edit the uploaded firmware
[**EditUpgradePlan**](FirmwareAPI.md#EditUpgradePlan) | **Patch** /openapi/v1/{omadacId}/plans/{planId}/upgrade/overview | Edit firmware upgrade plan
[**GetGridAutoCheckUpgrade**](FirmwareAPI.md#GetGridAutoCheckUpgrade) | **Get** /openapi/v1/{omadacId}/upgrade/autoCheck | Get autocheck upgrade plan
[**GetGridFirmwareList**](FirmwareAPI.md#GetGridFirmwareList) | **Get** /openapi/v1/{omadacId}/upgrade/firmwares | Get the uploaded firmware list
[**GetGridFirmwarePoolList**](FirmwareAPI.md#GetGridFirmwarePoolList) | **Get** /openapi/v1/{omadacId}/upgrade/overview/firmwares | Get firmware pool list
[**GetGridPlanUpgradeModelList**](FirmwareAPI.md#GetGridPlanUpgradeModelList) | **Get** /openapi/v1/{omadacId}/upgrade/plan/models | Get the list of models
[**GetGridPlanUpgradeSiteList**](FirmwareAPI.md#GetGridPlanUpgradeSiteList) | **Post** /openapi/v1/{omadacId}/upgrade/plan/sites | Get the sites of the selected model
[**GetGridUpgradeLogs**](FirmwareAPI.md#GetGridUpgradeLogs) | **Get** /openapi/v1/{omadacId}/upgrade/overview/logs | Get firmware upgrade logs
[**GetGridUpgradePlans**](FirmwareAPI.md#GetGridUpgradePlans) | **Get** /openapi/v1/{omadacId}/upgrade/overview/plans | Get firmware upgrade plans
[**GetModelBySites**](FirmwareAPI.md#GetModelBySites) | **Post** /openapi/v1/{omadacId}/upgrade/models | Get the model of the specified site
[**GetModelFirmwareReleaseNotes**](FirmwareAPI.md#GetModelFirmwareReleaseNotes) | **Post** /openapi/v1/{omadacId}/upgrade/overview/firmwares/release-note | Get release notes information
[**GetPlanUpgradeModelInfo**](FirmwareAPI.md#GetPlanUpgradeModelInfo) | **Post** /openapi/v1/{omadacId}/upgrade/plan/firmware | Get the upgradeable information of the selected model
[**GetTargetFirmwareAllSites**](FirmwareAPI.md#GetTargetFirmwareAllSites) | **Get** /openapi/v1/{omadacId}/upgrade/firmwares/{firmwareId}/target-sites | Get all sites of target firmware
[**GetTryBetaStatus**](FirmwareAPI.md#GetTryBetaStatus) | **Get** /openapi/v1/{omadacId}/upgrade/overview/try-beta | Get try-beta switch status
[**GetUpgradeFailedDeviceFirmwareInfo**](FirmwareAPI.md#GetUpgradeFailedDeviceFirmwareInfo) | **Get** /openapi/v1/{omadacId}/logs/{upgradeLogId}/upgrade/overview/failed-model-firmware | Get firmware information about the failed device to be upgraded
[**GetUpgradeFailedDeviceInfos**](FirmwareAPI.md#GetUpgradeFailedDeviceInfos) | **Get** /openapi/v1/{omadacId}/logs/{upgradeLogId}/upgrade/overview/failed-devices | Get the list of devices which upgrade failed
[**GetcriticalModelNum**](FirmwareAPI.md#GetcriticalModelNum) | **Get** /openapi/v1/{omadacId}/upgrade/overview/critical | Get the number of critical models
[**ModifyTryBetaStatus**](FirmwareAPI.md#ModifyTryBetaStatus) | **Patch** /openapi/v1/{omadacId}/upgrade/overview/try-beta | Modify try-beta switch status
[**UploadUpgradeFirmware**](FirmwareAPI.md#UploadUpgradeFirmware) | **Post** /openapi/v1/{omadacId}/files/upgrade/firmware | Upload upgrade firmware



## CheckUploadUpgradeFirmwareLimit

> OperationResponseWithoutResult CheckUploadUpgradeFirmwareLimit(ctx, omadacId, firmwareSize).Execute()

Check the limit of the firmware pool



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
	firmwareSize := "firmwareSize_example" // string | firmwareSize

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.CheckUploadUpgradeFirmwareLimit(context.Background(), omadacId, firmwareSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.CheckUploadUpgradeFirmwareLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUploadUpgradeFirmwareLimit`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.CheckUploadUpgradeFirmwareLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**firmwareSize** | **string** | firmwareSize | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUploadUpgradeFirmwareLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAutoCheckUpgrade

> OperationResponseCreateAutoCheckResInfo CreateAutoCheckUpgrade(ctx, omadacId).AutoCheckUpgradeCreateInfo(autoCheckUpgradeCreateInfo).Execute()

Create autocheck upgrade plan



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
	autoCheckUpgradeCreateInfo := *openapiclient.NewAutoCheckUpgradeCreateInfo(int32(123), []openapiclient.ModelTypeInfoOpenApiVO{*openapiclient.NewModelTypeInfoOpenApiVO("CompoundModel_example", "ShowModel_example")}, *openapiclient.NewUpgradeBaseScheduleTimeOpenApiVO(int32(123), int32(123), int32(123)), []string{"SiteIds_example"}) // AutoCheckUpgradeCreateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.CreateAutoCheckUpgrade(context.Background(), omadacId).AutoCheckUpgradeCreateInfo(autoCheckUpgradeCreateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.CreateAutoCheckUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoCheckUpgrade`: OperationResponseCreateAutoCheckResInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.CreateAutoCheckUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoCheckUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoCheckUpgradeCreateInfo** | [**AutoCheckUpgradeCreateInfo**](AutoCheckUpgradeCreateInfo.md) |  | 

### Return type

[**OperationResponseCreateAutoCheckResInfo**](OperationResponseCreateAutoCheckResInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFailedDeviceModelPlan

> OperationResponseCreateFailedDeviceUpgradePlan CreateFailedDeviceModelPlan(ctx, omadacId, upgradeLogId).FailedCreateModelPlanUpgradeInfo(failedCreateModelPlanUpgradeInfo).Execute()

Create Upgrade Plan for Failed Device



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
	upgradeLogId := "upgradeLogId_example" // string |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs)
	failedCreateModelPlanUpgradeInfo := *openapiclient.NewFailedCreateModelPlanUpgradeInfo(int32(123)) // FailedCreateModelPlanUpgradeInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.CreateFailedDeviceModelPlan(context.Background(), omadacId, upgradeLogId).FailedCreateModelPlanUpgradeInfo(failedCreateModelPlanUpgradeInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.CreateFailedDeviceModelPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFailedDeviceModelPlan`: OperationResponseCreateFailedDeviceUpgradePlan
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.CreateFailedDeviceModelPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**upgradeLogId** | **string** |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFailedDeviceModelPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failedCreateModelPlanUpgradeInfo** | [**FailedCreateModelPlanUpgradeInfo**](FailedCreateModelPlanUpgradeInfo.md) |  | 

### Return type

[**OperationResponseCreateFailedDeviceUpgradePlan**](OperationResponseCreateFailedDeviceUpgradePlan.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirmwareUpgradePlan

> OperationResponseCreateUpgradePlanResinfo CreateFirmwareUpgradePlan(ctx, omadacId, firmwareId).PlanFirmwareUpgradeCreateInfo(planFirmwareUpgradeCreateInfo).Execute()

Create firmware upgrade plan based on manually uploaded firmware



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
	firmwareId := "firmwareId_example" // string | Firmware ID
	planFirmwareUpgradeCreateInfo := *openapiclient.NewPlanFirmwareUpgradeCreateInfo(int32(123), []string{"Sites_example"}) // PlanFirmwareUpgradeCreateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.CreateFirmwareUpgradePlan(context.Background(), omadacId, firmwareId).PlanFirmwareUpgradeCreateInfo(planFirmwareUpgradeCreateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.CreateFirmwareUpgradePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFirmwareUpgradePlan`: OperationResponseCreateUpgradePlanResinfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.CreateFirmwareUpgradePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**firmwareId** | **string** | Firmware ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirmwareUpgradePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **planFirmwareUpgradeCreateInfo** | [**PlanFirmwareUpgradeCreateInfo**](PlanFirmwareUpgradeCreateInfo.md) |  | 

### Return type

[**OperationResponseCreateUpgradePlanResinfo**](OperationResponseCreateUpgradePlanResinfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRollback

> OperationResponseWithoutResult CreateRollback(ctx, omadacId, upgradeLogId).RollbackCreateInfo(rollbackCreateInfo).Execute()

Create rollback plan



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
	upgradeLogId := "upgradeLogId_example" // string |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs)
	rollbackCreateInfo := *openapiclient.NewRollbackCreateInfo(int32(123), "TargetVersion_example") // RollbackCreateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.CreateRollback(context.Background(), omadacId, upgradeLogId).RollbackCreateInfo(rollbackCreateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.CreateRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRollback`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.CreateRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**upgradeLogId** | **string** |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rollbackCreateInfo** | [**RollbackCreateInfo**](RollbackCreateInfo.md) |  | 

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


## CreateUpgradePlan

> OperationResponseCreateUpgradePlanResinfo CreateUpgradePlan(ctx, omadacId).PlanUpgradeCreateInfo(planUpgradeCreateInfo).Execute()

Create firmware upgrade plan



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
	planUpgradeCreateInfo := *openapiclient.NewPlanUpgradeCreateInfo([]openapiclient.PlanUpgradeSelectedModel{*openapiclient.NewPlanUpgradeSelectedModel([]string{"CurrentVersion_example"}, *openapiclient.NewModelTypeInfoOpenApiVO("CompoundModel_example", "ShowModel_example"), "TargetVersion_example")}, int32(123), []string{"Sites_example"}) // PlanUpgradeCreateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.CreateUpgradePlan(context.Background(), omadacId).PlanUpgradeCreateInfo(planUpgradeCreateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.CreateUpgradePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpgradePlan`: OperationResponseCreateUpgradePlanResinfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.CreateUpgradePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpgradePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **planUpgradeCreateInfo** | [**PlanUpgradeCreateInfo**](PlanUpgradeCreateInfo.md) |  | 

### Return type

[**OperationResponseCreateUpgradePlanResinfo**](OperationResponseCreateUpgradePlanResinfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoCheckUpgrade

> OperationResponseWithoutResult DeleteAutoCheckUpgrade(ctx, omadacId, autoCheckId).Execute()

Delete autoCheck upgrade plan



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
	autoCheckId := "autoCheckId_example" // string | Auto Check ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.DeleteAutoCheckUpgrade(context.Background(), omadacId, autoCheckId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.DeleteAutoCheckUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAutoCheckUpgrade`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.DeleteAutoCheckUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**autoCheckId** | **string** | Auto Check ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoCheckUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUpgradeFirmware

> OperationResponseWithoutResult DeleteUpgradeFirmware(ctx, omadacId, firmwareId).Execute()

Delete upgrade firmware



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
	firmwareId := "firmwareId_example" // string | Firmware ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.DeleteUpgradeFirmware(context.Background(), omadacId, firmwareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.DeleteUpgradeFirmware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUpgradeFirmware`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.DeleteUpgradeFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**firmwareId** | **string** | Firmware ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUpgradeFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUpgradePlan

> OperationResponseWithoutResult DeleteUpgradePlan(ctx, omadacId, planId).Execute()

Delete firmware upgrade plan



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
	planId := "planId_example" // string | Upgrade Plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.DeleteUpgradePlan(context.Background(), omadacId, planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.DeleteUpgradePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUpgradePlan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.DeleteUpgradePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**planId** | **string** | Upgrade Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUpgradePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseWithoutResult**](OperationResponseWithoutResult.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAutoCheckUpgrade

> OperationResponseWithoutResult EditAutoCheckUpgrade(ctx, omadacId, autoCheckId).AutoCheckUpgradeCreateInfo(autoCheckUpgradeCreateInfo).Execute()

Edit autoCheck upgrade plan



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
	autoCheckId := "autoCheckId_example" // string | Auto Check ID
	autoCheckUpgradeCreateInfo := *openapiclient.NewAutoCheckUpgradeCreateInfo(int32(123), []openapiclient.ModelTypeInfoOpenApiVO{*openapiclient.NewModelTypeInfoOpenApiVO("CompoundModel_example", "ShowModel_example")}, *openapiclient.NewUpgradeBaseScheduleTimeOpenApiVO(int32(123), int32(123), int32(123)), []string{"SiteIds_example"}) // AutoCheckUpgradeCreateInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.EditAutoCheckUpgrade(context.Background(), omadacId, autoCheckId).AutoCheckUpgradeCreateInfo(autoCheckUpgradeCreateInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.EditAutoCheckUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAutoCheckUpgrade`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.EditAutoCheckUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**autoCheckId** | **string** | Auto Check ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAutoCheckUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoCheckUpgradeCreateInfo** | [**AutoCheckUpgradeCreateInfo**](AutoCheckUpgradeCreateInfo.md) |  | 

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


## EditFirmwareUpgrade

> OperationResponseWithoutResult EditFirmwareUpgrade(ctx, omadacId, firmwareId).EditUploadFirmwareInfo(editUploadFirmwareInfo).Execute()

Edit the uploaded firmware



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
	firmwareId := "firmwareId_example" // string | Firmware ID
	editUploadFirmwareInfo := *openapiclient.NewEditUploadFirmwareInfo() // EditUploadFirmwareInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.EditFirmwareUpgrade(context.Background(), omadacId, firmwareId).EditUploadFirmwareInfo(editUploadFirmwareInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.EditFirmwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditFirmwareUpgrade`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.EditFirmwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**firmwareId** | **string** | Firmware ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditFirmwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editUploadFirmwareInfo** | [**EditUploadFirmwareInfo**](EditUploadFirmwareInfo.md) |  | 

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


## EditUpgradePlan

> OperationResponseWithoutResult EditUpgradePlan(ctx, omadacId, planId).PlanUpgradeEditInfo(planUpgradeEditInfo).Execute()

Edit firmware upgrade plan



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
	planId := "planId_example" // string | Upgrade Plan ID
	planUpgradeEditInfo := *openapiclient.NewPlanUpgradeEditInfo(int32(123), []string{"Sites_example"}) // PlanUpgradeEditInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.EditUpgradePlan(context.Background(), omadacId, planId).PlanUpgradeEditInfo(planUpgradeEditInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.EditUpgradePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditUpgradePlan`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.EditUpgradePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**planId** | **string** | Upgrade Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUpgradePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **planUpgradeEditInfo** | [**PlanUpgradeEditInfo**](PlanUpgradeEditInfo.md) |  | 

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


## GetGridAutoCheckUpgrade

> OperationResponseGridVOAutoCheckUpgradeInfo GetGridAutoCheckUpgrade(ctx, omadacId).Page(page).PageSize(pageSize).SortsAutoCheckTime(sortsAutoCheckTime).SearchKey(searchKey).Execute()

Get autocheck upgrade plan



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsAutoCheckTime := "sortsAutoCheckTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field upgrade (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetGridAutoCheckUpgrade(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsAutoCheckTime(sortsAutoCheckTime).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetGridAutoCheckUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridAutoCheckUpgrade`: OperationResponseGridVOAutoCheckUpgradeInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetGridAutoCheckUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridAutoCheckUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsAutoCheckTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field upgrade | 

### Return type

[**OperationResponseGridVOAutoCheckUpgradeInfo**](OperationResponseGridVOAutoCheckUpgradeInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridFirmwareList

> OperationResponseGridVOFirmwareInfo GetGridFirmwareList(ctx, omadacId).Page(page).PageSize(pageSize).SortsModelTypeInfo(sortsModelTypeInfo).SortsUploadTime(sortsUploadTime).SearchKey(searchKey).Execute()

Get the uploaded firmware list



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsModelTypeInfo := "sortsModelTypeInfo_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsUploadTime := "sortsUploadTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field upgrade (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetGridFirmwareList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsModelTypeInfo(sortsModelTypeInfo).SortsUploadTime(sortsUploadTime).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetGridFirmwareList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridFirmwareList`: OperationResponseGridVOFirmwareInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetGridFirmwareList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridFirmwareListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsModelTypeInfo** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsUploadTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field upgrade | 

### Return type

[**OperationResponseGridVOFirmwareInfo**](OperationResponseGridVOFirmwareInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridFirmwarePoolList

> OperationResponseModelFirmwarePoolGridInfoModelLatestFwInfo GetGridFirmwarePoolList(ctx, omadacId).Channel(channel).Page(page).PageSize(pageSize).SortsModelTypeInfo(sortsModelTypeInfo).SortsReleaseTime(sortsReleaseTime).Execute()

Get firmware pool list



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
	channel := int32(56) // int32 | Channel should be a value as follows: 0: stable; 1: Release Candidate(RC); 2: Beta.
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsModelTypeInfo := "sortsModelTypeInfo_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsReleaseTime := "sortsReleaseTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetGridFirmwarePoolList(context.Background(), omadacId).Channel(channel).Page(page).PageSize(pageSize).SortsModelTypeInfo(sortsModelTypeInfo).SortsReleaseTime(sortsReleaseTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetGridFirmwarePoolList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridFirmwarePoolList`: OperationResponseModelFirmwarePoolGridInfoModelLatestFwInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetGridFirmwarePoolList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridFirmwarePoolListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **int32** | Channel should be a value as follows: 0: stable; 1: Release Candidate(RC); 2: Beta. | 
 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsModelTypeInfo** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsReleaseTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseModelFirmwarePoolGridInfoModelLatestFwInfo**](OperationResponseModelFirmwarePoolGridInfoModelLatestFwInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPlanUpgradeModelList

> OperationResponseGridVOPlanUpgradeModelList GetGridPlanUpgradeModelList(ctx, omadacId).Page(page).PageSize(pageSize).SortsModelTypeInfo(sortsModelTypeInfo).SearchKey(searchKey).Execute()

Get the list of models



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsModelTypeInfo := "sortsModelTypeInfo_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	searchKey := "searchKey_example" // string | Fuzzy query parameters, support field upgrade (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetGridPlanUpgradeModelList(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsModelTypeInfo(sortsModelTypeInfo).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetGridPlanUpgradeModelList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPlanUpgradeModelList`: OperationResponseGridVOPlanUpgradeModelList
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetGridPlanUpgradeModelList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridPlanUpgradeModelListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsModelTypeInfo** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **searchKey** | **string** | Fuzzy query parameters, support field upgrade | 

### Return type

[**OperationResponseGridVOPlanUpgradeModelList**](OperationResponseGridVOPlanUpgradeModelList.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridPlanUpgradeSiteList

> OperationResponseModelUpgradeSiteInfo GetGridPlanUpgradeSiteList(ctx, omadacId).ModelUpgradeSiteReqInfo(modelUpgradeSiteReqInfo).Execute()

Get the sites of the selected model



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
	modelUpgradeSiteReqInfo := *openapiclient.NewModelUpgradeSiteReqInfo(int32(123), int32(123), []openapiclient.ModelBaseInfo{*openapiclient.NewModelBaseInfo(*openapiclient.NewModelTypeInfoOpenApiVO("CompoundModel_example", "ShowModel_example"))}) // ModelUpgradeSiteReqInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetGridPlanUpgradeSiteList(context.Background(), omadacId).ModelUpgradeSiteReqInfo(modelUpgradeSiteReqInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetGridPlanUpgradeSiteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridPlanUpgradeSiteList`: OperationResponseModelUpgradeSiteInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetGridPlanUpgradeSiteList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridPlanUpgradeSiteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUpgradeSiteReqInfo** | [**ModelUpgradeSiteReqInfo**](ModelUpgradeSiteReqInfo.md) |  | 

### Return type

[**OperationResponseModelUpgradeSiteInfo**](OperationResponseModelUpgradeSiteInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridUpgradeLogs

> OperationResponseGridVOUpgradeLogOpenApiInfo GetGridUpgradeLogs(ctx, omadacId).Page(page).PageSize(pageSize).SortsUpgradeTime(sortsUpgradeTime).SortsModelTypeInfo(sortsModelTypeInfo).Execute()

Get firmware upgrade logs



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsUpgradeTime := "sortsUpgradeTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsModelTypeInfo := "sortsModelTypeInfo_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetGridUpgradeLogs(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsUpgradeTime(sortsUpgradeTime).SortsModelTypeInfo(sortsModelTypeInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetGridUpgradeLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridUpgradeLogs`: OperationResponseGridVOUpgradeLogOpenApiInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetGridUpgradeLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridUpgradeLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsUpgradeTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsModelTypeInfo** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseGridVOUpgradeLogOpenApiInfo**](OperationResponseGridVOUpgradeLogOpenApiInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGridUpgradePlans

> OperationResponseGridVOPlanUpgradeInfo GetGridUpgradePlans(ctx, omadacId).Page(page).PageSize(pageSize).SortsScheduledUpgradeTime(sortsScheduledUpgradeTime).SortsModelTypeInfo(sortsModelTypeInfo).Execute()

Get firmware upgrade plans



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
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.
	sortsScheduledUpgradeTime := "sortsScheduledUpgradeTime_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)
	sortsModelTypeInfo := "sortsModelTypeInfo_example" // string | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetGridUpgradePlans(context.Background(), omadacId).Page(page).PageSize(pageSize).SortsScheduledUpgradeTime(sortsScheduledUpgradeTime).SortsModelTypeInfo(sortsModelTypeInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetGridUpgradePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGridUpgradePlans`: OperationResponseGridVOPlanUpgradeInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetGridUpgradePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGridUpgradePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 
 **sortsScheduledUpgradeTime** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 
 **sortsModelTypeInfo** | **string** | Sort parameter may be one of asc or desc. Optional parameter. If it is not carried, it means it is not sorted by this field. When there are more than one, the first one takes effect | 

### Return type

[**OperationResponseGridVOPlanUpgradeInfo**](OperationResponseGridVOPlanUpgradeInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelBySites

> OperationResponseUpgradeSiteModelInfo GetModelBySites(ctx, omadacId).UpgradeSiteModelReqInfo(upgradeSiteModelReqInfo).Execute()

Get the model of the specified site



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
	upgradeSiteModelReqInfo := *openapiclient.NewUpgradeSiteModelReqInfo() // UpgradeSiteModelReqInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetModelBySites(context.Background(), omadacId).UpgradeSiteModelReqInfo(upgradeSiteModelReqInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetModelBySites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelBySites`: OperationResponseUpgradeSiteModelInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetModelBySites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelBySitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeSiteModelReqInfo** | [**UpgradeSiteModelReqInfo**](UpgradeSiteModelReqInfo.md) |  | 

### Return type

[**OperationResponseUpgradeSiteModelInfo**](OperationResponseUpgradeSiteModelInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelFirmwareReleaseNotes

> OperationResponseModelFwOemReleaseNoteInfo GetModelFirmwareReleaseNotes(ctx, omadacId).ModelFwReleaseNoteReqInfo(modelFwReleaseNoteReqInfo).Execute()

Get release notes information



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
	modelFwReleaseNoteReqInfo := *openapiclient.NewModelFwReleaseNoteReqInfo("OemId_example") // ModelFwReleaseNoteReqInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetModelFirmwareReleaseNotes(context.Background(), omadacId).ModelFwReleaseNoteReqInfo(modelFwReleaseNoteReqInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetModelFirmwareReleaseNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelFirmwareReleaseNotes`: OperationResponseModelFwOemReleaseNoteInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetModelFirmwareReleaseNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelFirmwareReleaseNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelFwReleaseNoteReqInfo** | [**ModelFwReleaseNoteReqInfo**](ModelFwReleaseNoteReqInfo.md) |  | 

### Return type

[**OperationResponseModelFwOemReleaseNoteInfo**](OperationResponseModelFwOemReleaseNoteInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanUpgradeModelInfo

> OperationResponsePlanUpgradeModelInfo GetPlanUpgradeModelInfo(ctx, omadacId).ModelUpgradeInfo(modelUpgradeInfo).Execute()

Get the upgradeable information of the selected model



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
	modelUpgradeInfo := *openapiclient.NewModelUpgradeInfo([]string{"ModelList_example"}, "SelectType_example") // ModelUpgradeInfo | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetPlanUpgradeModelInfo(context.Background(), omadacId).ModelUpgradeInfo(modelUpgradeInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetPlanUpgradeModelInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanUpgradeModelInfo`: OperationResponsePlanUpgradeModelInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetPlanUpgradeModelInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanUpgradeModelInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUpgradeInfo** | [**ModelUpgradeInfo**](ModelUpgradeInfo.md) |  | 

### Return type

[**OperationResponsePlanUpgradeModelInfo**](OperationResponsePlanUpgradeModelInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetFirmwareAllSites

> OperationResponseAllTargetSites GetTargetFirmwareAllSites(ctx, omadacId, firmwareId).Execute()

Get all sites of target firmware



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
	firmwareId := "firmwareId_example" // string | Firmware ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetTargetFirmwareAllSites(context.Background(), omadacId, firmwareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetTargetFirmwareAllSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetFirmwareAllSites`: OperationResponseAllTargetSites
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetTargetFirmwareAllSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**firmwareId** | **string** | Firmware ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetFirmwareAllSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseAllTargetSites**](OperationResponseAllTargetSites.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTryBetaStatus

> OperationResponseUpgradeSettingTryBeta GetTryBetaStatus(ctx, omadacId).Execute()

Get try-beta switch status



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
	resp, r, err := apiClient.FirmwareAPI.GetTryBetaStatus(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetTryBetaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTryBetaStatus`: OperationResponseUpgradeSettingTryBeta
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetTryBetaStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTryBetaStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseUpgradeSettingTryBeta**](OperationResponseUpgradeSettingTryBeta.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpgradeFailedDeviceFirmwareInfo

> OperationResponseFailedDeviceUpgradeFirmwareInfo GetUpgradeFailedDeviceFirmwareInfo(ctx, omadacId, upgradeLogId).Execute()

Get firmware information about the failed device to be upgraded



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
	upgradeLogId := "upgradeLogId_example" // string |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetUpgradeFailedDeviceFirmwareInfo(context.Background(), omadacId, upgradeLogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetUpgradeFailedDeviceFirmwareInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpgradeFailedDeviceFirmwareInfo`: OperationResponseFailedDeviceUpgradeFirmwareInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetUpgradeFailedDeviceFirmwareInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**upgradeLogId** | **string** |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpgradeFailedDeviceFirmwareInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OperationResponseFailedDeviceUpgradeFirmwareInfo**](OperationResponseFailedDeviceUpgradeFirmwareInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpgradeFailedDeviceInfos

> OperationResponseGridVOUpgradeFailedDeviceInfo GetUpgradeFailedDeviceInfos(ctx, omadacId, upgradeLogId).Page(page).PageSize(pageSize).Execute()

Get the list of devices which upgrade failed



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
	upgradeLogId := "upgradeLogId_example" // string |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs)
	page := int32(56) // int32 | Start page number. Start from 1.
	pageSize := int32(56) // int32 | Number of entries per page. It should be within the range of 1–1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.GetUpgradeFailedDeviceInfos(context.Background(), omadacId, upgradeLogId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetUpgradeFailedDeviceInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpgradeFailedDeviceInfos`: OperationResponseGridVOUpgradeFailedDeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetUpgradeFailedDeviceInfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 
**upgradeLogId** | **string** |  The Upgrade log ID obtained from the interface (Get firmware upgrade logs) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpgradeFailedDeviceInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Start page number. Start from 1. | 
 **pageSize** | **int32** | Number of entries per page. It should be within the range of 1–1000. | 

### Return type

[**OperationResponseGridVOUpgradeFailedDeviceInfo**](OperationResponseGridVOUpgradeFailedDeviceInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetcriticalModelNum

> OperationResponseCriticalModelNum GetcriticalModelNum(ctx, omadacId).Execute()

Get the number of critical models



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
	resp, r, err := apiClient.FirmwareAPI.GetcriticalModelNum(context.Background(), omadacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.GetcriticalModelNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetcriticalModelNum`: OperationResponseCriticalModelNum
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.GetcriticalModelNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetcriticalModelNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationResponseCriticalModelNum**](OperationResponseCriticalModelNum.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyTryBetaStatus

> OperationResponseWithoutResult ModifyTryBetaStatus(ctx, omadacId).UpgradeSettingTryBeta(upgradeSettingTryBeta).Execute()

Modify try-beta switch status



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
	upgradeSettingTryBeta := *openapiclient.NewUpgradeSettingTryBeta(false) // UpgradeSettingTryBeta | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.ModifyTryBetaStatus(context.Background(), omadacId).UpgradeSettingTryBeta(upgradeSettingTryBeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.ModifyTryBetaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyTryBetaStatus`: OperationResponseWithoutResult
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.ModifyTryBetaStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyTryBetaStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeSettingTryBeta** | [**UpgradeSettingTryBeta**](UpgradeSettingTryBeta.md) |  | 

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


## UploadUpgradeFirmware

> OperationResponseUploadFirmwareResInfo UploadUpgradeFirmware(ctx, omadacId).Description(description).ShowModel(showModel).CompoundModel(compoundModel).TargetEnable(targetEnable).TargetSites(targetSites).UploadUpgradeFirmwareRequest(uploadUpgradeFirmwareRequest).Execute()

Upload upgrade firmware



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
	description := "description_example" // string | Description of upload firmware
	showModel := "showModel_example" // string | ShowModel of upload firmware, you can also get this field throw: \"Get the model of the specified site\"
	compoundModel := "compoundModel_example" // string | CompoundModel of upload firmware, you can also get this field throw: \"Get the model of the specified site\"
	targetEnable := true // bool | Do the sites set up specified firmware, it should be a value as follows: true, false
	targetSites := "targetSites_example" // string | A string of target site IDs separated by commas, exists when 'targetEnable' is true. (optional)
	uploadUpgradeFirmwareRequest := *openapiclient.NewUploadUpgradeFirmwareRequest() // UploadUpgradeFirmwareRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirmwareAPI.UploadUpgradeFirmware(context.Background(), omadacId).Description(description).ShowModel(showModel).CompoundModel(compoundModel).TargetEnable(targetEnable).TargetSites(targetSites).UploadUpgradeFirmwareRequest(uploadUpgradeFirmwareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirmwareAPI.UploadUpgradeFirmware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadUpgradeFirmware`: OperationResponseUploadFirmwareResInfo
	fmt.Fprintf(os.Stdout, "Response from `FirmwareAPI.UploadUpgradeFirmware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**omadacId** | **string** | Omada ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadUpgradeFirmwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** | Description of upload firmware | 
 **showModel** | **string** | ShowModel of upload firmware, you can also get this field throw: \&quot;Get the model of the specified site\&quot; | 
 **compoundModel** | **string** | CompoundModel of upload firmware, you can also get this field throw: \&quot;Get the model of the specified site\&quot; | 
 **targetEnable** | **bool** | Do the sites set up specified firmware, it should be a value as follows: true, false | 
 **targetSites** | **string** | A string of target site IDs separated by commas, exists when &#39;targetEnable&#39; is true. | 
 **uploadUpgradeFirmwareRequest** | [**UploadUpgradeFirmwareRequest**](UploadUpgradeFirmwareRequest.md) |  | 

### Return type

[**OperationResponseUploadFirmwareResInfo**](OperationResponseUploadFirmwareResInfo.md)

### Authorization

[AccessToken](../README.md#AccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

