# OLTONURegisterAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthenticationConfig**](OLTONURegisterAPI.md#addauthenticationconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/authentications/add | Create authentication config
[**AddOnuAutoAuthenticationRuleConfig**](OLTONURegisterAPI.md#addonuautoauthenticationruleconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/rules/add | Create ONU auto authentication rule
[**AuthOnu**](OLTONURegisterAPI.md#authonu) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/autofind/onus/auth | Authenticate ONU
[**ClearOnuAutofindConfig**](OLTONURegisterAPI.md#clearonuautofindconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/autofind/onus/clear | Batch delete existing ONU auto find config
[**DeleteAuthenticationConfig**](OLTONURegisterAPI.md#deleteauthenticationconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/authentications/delete | Batch delete existing authentication config
[**DeleteOnuAutoAuthenticationRuleConfig**](OLTONURegisterAPI.md#deleteonuautoauthenticationruleconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/rules/delete | Delete an existing ONU auto authentication rule
[**EditAuthenticationConfig**](OLTONURegisterAPI.md#editauthenticationconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/authentications/edit | Modify an existing authentication config
[**EditOnuAdminStatus**](OLTONURegisterAPI.md#editonuadminstatus) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/authentications/admin-status/edit | Modify an existing ONU admin status
[**EditOnuAutoAuthenticationConfig**](OLTONURegisterAPI.md#editonuautoauthenticationconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/global/configs/edit | Modify an existing ONU auto authentication config
[**EditOnuAutoAuthenticationRuleConfig**](OLTONURegisterAPI.md#editonuautoauthenticationruleconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/rules/edit | Modify an existing ONU auto authentication rule
[**EditOnuPonAutoAuthenticationConfig**](OLTONURegisterAPI.md#editonuponautoauthenticationconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/port-configs/edit | Modify an existing ONU pon auto authentication config
[**EditOnuRegisterAutofindConfig**](OLTONURegisterAPI.md#editonuregisterautofindconfig) | **Post** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/autofind/configs/edit | Modify an existing ONU register auto find config
[**GetAuthenticationConfigList**](OLTONURegisterAPI.md#getauthenticationconfiglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/authentications/list | Get authentication config list
[**GetAuthenticationConfigPage**](OLTONURegisterAPI.md#getauthenticationconfigpage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/authentications/page | Get authentication config page
[**GetOnuAutoAuthenticationConfig**](OLTONURegisterAPI.md#getonuautoauthenticationconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/global/configs/get | Get ONU auto authentication config
[**GetOnuAutoAuthenticationRuleConfigList**](OLTONURegisterAPI.md#getonuautoauthenticationruleconfiglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/rules/list | Get ONU auto authentication rule list
[**GetOnuAutoAuthenticationRuleConfigPage**](OLTONURegisterAPI.md#getonuautoauthenticationruleconfigpage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/rules/page | Get ONU auto authentication rule page
[**GetOnuAutofindConfigList**](OLTONURegisterAPI.md#getonuautofindconfiglist) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/autofinds/list | Get ONU auto find config list
[**GetOnuAutofindConfigPage**](OLTONURegisterAPI.md#getonuautofindconfigpage) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/autofinds/page | get ONU auto find config page
[**GetOnuPonAutoAuthenticationConfig**](OLTONURegisterAPI.md#getonuponautoauthenticationconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/auto-authentication/port-configs/get | Get ONU pon auto authentication config
[**GetOnuRegisterAutofindConfig**](OLTONURegisterAPI.md#getonuregisterautofindconfig) | **Get** /openapi/v1/{omadacId}/sites/{siteId}/olts/{deviceMac}/pon/onu-register/autofind/configs/get | Get ONU register auto find config



## AddAuthenticationConfig

> OperationResponseDeviceResponseBodyVoid AddAuthenticationConfig(ctx, omadacId, siteId, deviceMac).AuthenticationConfigAddDTO(authenticationConfigAddDTO).Execute()

Create authentication config



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
	authenticationConfigAddDTO := *openapiclient.NewAuthenticationConfigAddDTO("PortId_example") // AuthenticationConfigAddDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.AddAuthenticationConfig(context.Background(), omadacId, siteId, deviceMac).AuthenticationConfigAddDTO(authenticationConfigAddDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.AddAuthenticationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAuthenticationConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.AddAuthenticationConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddAuthenticationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authenticationConfigAddDTO** | [**AuthenticationConfigAddDTO**](AuthenticationConfigAddDTO.md) |  | 

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


## AddOnuAutoAuthenticationRuleConfig

> OperationResponseDeviceResponseBodyVoid AddOnuAutoAuthenticationRuleConfig(ctx, omadacId, siteId, deviceMac).PonAutoAuthenticationRuleConfigDTO(ponAutoAuthenticationRuleConfigDTO).Execute()

Create ONU auto authentication rule



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
	ponAutoAuthenticationRuleConfigDTO := *openapiclient.NewPonAutoAuthenticationRuleConfigDTO("LineProfile_example", "PonPort_example", "ServiceProfile_example") // PonAutoAuthenticationRuleConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.AddOnuAutoAuthenticationRuleConfig(context.Background(), omadacId, siteId, deviceMac).PonAutoAuthenticationRuleConfigDTO(ponAutoAuthenticationRuleConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.AddOnuAutoAuthenticationRuleConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOnuAutoAuthenticationRuleConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.AddOnuAutoAuthenticationRuleConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddOnuAutoAuthenticationRuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ponAutoAuthenticationRuleConfigDTO** | [**PonAutoAuthenticationRuleConfigDTO**](PonAutoAuthenticationRuleConfigDTO.md) |  | 

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


## AuthOnu

> OperationResponseDeviceResponseBodyVoid AuthOnu(ctx, omadacId, siteId, deviceMac).OnuActivationConfigDTO(onuActivationConfigDTO).Execute()

Authenticate ONU



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
	onuActivationConfigDTO := *openapiclient.NewOnuActivationConfigDTO("AuthenticationMethod_example", []string{"Keys_example"}, "LineProfile_example", "ServiceProfile_example") // OnuActivationConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.AuthOnu(context.Background(), omadacId, siteId, deviceMac).OnuActivationConfigDTO(onuActivationConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.AuthOnu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthOnu`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.AuthOnu`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAuthOnuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onuActivationConfigDTO** | [**OnuActivationConfigDTO**](OnuActivationConfigDTO.md) |  | 

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


## ClearOnuAutofindConfig

> OperationResponseDeviceResponseBodyVoid ClearOnuAutofindConfig(ctx, omadacId, siteId, deviceMac).StringKeyListRequest(stringKeyListRequest).Execute()

Batch delete existing ONU auto find config



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
	stringKeyListRequest := *openapiclient.NewStringKeyListRequest([]string{"Keys_example"}) // StringKeyListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.ClearOnuAutofindConfig(context.Background(), omadacId, siteId, deviceMac).StringKeyListRequest(stringKeyListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.ClearOnuAutofindConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearOnuAutofindConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.ClearOnuAutofindConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiClearOnuAutofindConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **stringKeyListRequest** | [**StringKeyListRequest**](StringKeyListRequest.md) |  | 

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


## DeleteAuthenticationConfig

> OperationResponseDeviceResponseBodyVoid DeleteAuthenticationConfig(ctx, omadacId, siteId, deviceMac).StringKeyListRequest(stringKeyListRequest).Execute()

Batch delete existing authentication config



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
	stringKeyListRequest := *openapiclient.NewStringKeyListRequest([]string{"Keys_example"}) // StringKeyListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.DeleteAuthenticationConfig(context.Background(), omadacId, siteId, deviceMac).StringKeyListRequest(stringKeyListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.DeleteAuthenticationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAuthenticationConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.DeleteAuthenticationConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteAuthenticationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **stringKeyListRequest** | [**StringKeyListRequest**](StringKeyListRequest.md) |  | 

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


## DeleteOnuAutoAuthenticationRuleConfig

> OperationResponseDeviceResponseBodyVoid DeleteOnuAutoAuthenticationRuleConfig(ctx, omadacId, siteId, deviceMac).PonAutoAuthenticationRuleConfigDeleteDTO(ponAutoAuthenticationRuleConfigDeleteDTO).Execute()

Delete an existing ONU auto authentication rule



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
	ponAutoAuthenticationRuleConfigDeleteDTO := *openapiclient.NewPonAutoAuthenticationRuleConfigDeleteDTO([]string{"Keys_example"}, "PonPort_example") // PonAutoAuthenticationRuleConfigDeleteDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.DeleteOnuAutoAuthenticationRuleConfig(context.Background(), omadacId, siteId, deviceMac).PonAutoAuthenticationRuleConfigDeleteDTO(ponAutoAuthenticationRuleConfigDeleteDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.DeleteOnuAutoAuthenticationRuleConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOnuAutoAuthenticationRuleConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.DeleteOnuAutoAuthenticationRuleConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteOnuAutoAuthenticationRuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ponAutoAuthenticationRuleConfigDeleteDTO** | [**PonAutoAuthenticationRuleConfigDeleteDTO**](PonAutoAuthenticationRuleConfigDeleteDTO.md) |  | 

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


## EditAuthenticationConfig

> OperationResponseDeviceResponseBodyVoid EditAuthenticationConfig(ctx, omadacId, siteId, deviceMac).AuthenticationConfigEditDTO(authenticationConfigEditDTO).Execute()

Modify an existing authentication config



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
	authenticationConfigEditDTO := *openapiclient.NewAuthenticationConfigEditDTO("Key_example", "PortId_example") // AuthenticationConfigEditDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.EditAuthenticationConfig(context.Background(), omadacId, siteId, deviceMac).AuthenticationConfigEditDTO(authenticationConfigEditDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.EditAuthenticationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAuthenticationConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.EditAuthenticationConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditAuthenticationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authenticationConfigEditDTO** | [**AuthenticationConfigEditDTO**](AuthenticationConfigEditDTO.md) |  | 

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


## EditOnuAdminStatus

> OperationResponseDeviceResponseBodyVoid EditOnuAdminStatus(ctx, omadacId, siteId, deviceMac).OnuAdminStatusEditDTO(onuAdminStatusEditDTO).Execute()

Modify an existing ONU admin status



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
	onuAdminStatusEditDTO := *openapiclient.NewOnuAdminStatusEditDTO("AdminStatus_example", []string{"Keys_example"}) // OnuAdminStatusEditDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.EditOnuAdminStatus(context.Background(), omadacId, siteId, deviceMac).OnuAdminStatusEditDTO(onuAdminStatusEditDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.EditOnuAdminStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuAdminStatus`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.EditOnuAdminStatus`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditOnuAdminStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **onuAdminStatusEditDTO** | [**OnuAdminStatusEditDTO**](OnuAdminStatusEditDTO.md) |  | 

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


## EditOnuAutoAuthenticationConfig

> OperationResponseDeviceResponseBodyVoid EditOnuAutoAuthenticationConfig(ctx, omadacId, siteId, deviceMac).AutoAuthenticationConfigDTO(autoAuthenticationConfigDTO).Execute()

Modify an existing ONU auto authentication config



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
	autoAuthenticationConfigDTO := *openapiclient.NewAutoAuthenticationConfigDTO("AutoAuthenticationStatus_example") // AutoAuthenticationConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.EditOnuAutoAuthenticationConfig(context.Background(), omadacId, siteId, deviceMac).AutoAuthenticationConfigDTO(autoAuthenticationConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.EditOnuAutoAuthenticationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuAutoAuthenticationConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.EditOnuAutoAuthenticationConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditOnuAutoAuthenticationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **autoAuthenticationConfigDTO** | [**AutoAuthenticationConfigDTO**](AutoAuthenticationConfigDTO.md) |  | 

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


## EditOnuAutoAuthenticationRuleConfig

> OperationResponseDeviceResponseBodyVoid EditOnuAutoAuthenticationRuleConfig(ctx, omadacId, siteId, deviceMac).PonAutoAuthenticationRuleConfigEditDTO(ponAutoAuthenticationRuleConfigEditDTO).Execute()

Modify an existing ONU auto authentication rule



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
	ponAutoAuthenticationRuleConfigEditDTO := *openapiclient.NewPonAutoAuthenticationRuleConfigEditDTO("Key_example", "LineProfile_example", "PonPort_example", int32(123), "ServiceProfile_example") // PonAutoAuthenticationRuleConfigEditDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.EditOnuAutoAuthenticationRuleConfig(context.Background(), omadacId, siteId, deviceMac).PonAutoAuthenticationRuleConfigEditDTO(ponAutoAuthenticationRuleConfigEditDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.EditOnuAutoAuthenticationRuleConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuAutoAuthenticationRuleConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.EditOnuAutoAuthenticationRuleConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditOnuAutoAuthenticationRuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ponAutoAuthenticationRuleConfigEditDTO** | [**PonAutoAuthenticationRuleConfigEditDTO**](PonAutoAuthenticationRuleConfigEditDTO.md) |  | 

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


## EditOnuPonAutoAuthenticationConfig

> OperationResponseDeviceResponseBodyVoid EditOnuPonAutoAuthenticationConfig(ctx, omadacId, siteId, deviceMac).PonAutoAuthenticationConfigDTO(ponAutoAuthenticationConfigDTO).Execute()

Modify an existing ONU pon auto authentication config



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
	ponAutoAuthenticationConfigDTO := *openapiclient.NewPonAutoAuthenticationConfigDTO("PonPort_example") // PonAutoAuthenticationConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.EditOnuPonAutoAuthenticationConfig(context.Background(), omadacId, siteId, deviceMac).PonAutoAuthenticationConfigDTO(ponAutoAuthenticationConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.EditOnuPonAutoAuthenticationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuPonAutoAuthenticationConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.EditOnuPonAutoAuthenticationConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditOnuPonAutoAuthenticationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ponAutoAuthenticationConfigDTO** | [**PonAutoAuthenticationConfigDTO**](PonAutoAuthenticationConfigDTO.md) |  | 

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


## EditOnuRegisterAutofindConfig

> OperationResponseDeviceResponseBodyVoid EditOnuRegisterAutofindConfig(ctx, omadacId, siteId, deviceMac).AutofindConfigDTO(autofindConfigDTO).Execute()

Modify an existing ONU register auto find config



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
	autofindConfigDTO := *openapiclient.NewAutofindConfigDTO("AgingTimeStatus_example", int32(123)) // AutofindConfigDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.EditOnuRegisterAutofindConfig(context.Background(), omadacId, siteId, deviceMac).AutofindConfigDTO(autofindConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.EditOnuRegisterAutofindConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditOnuRegisterAutofindConfig`: OperationResponseDeviceResponseBodyVoid
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.EditOnuRegisterAutofindConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditOnuRegisterAutofindConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **autofindConfigDTO** | [**AutofindConfigDTO**](AutofindConfigDTO.md) |  | 

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


## GetAuthenticationConfigList

> OperationResponseListAuthenticationConfigDTO GetAuthenticationConfigList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get authentication config list



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
	dto := *openapiclient.NewOnuRegisterAuthenPageQueryRequestDTO("PonPort_example") // OnuRegisterAuthenPageQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.GetAuthenticationConfigList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetAuthenticationConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticationConfigList`: OperationResponseListAuthenticationConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetAuthenticationConfigList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAuthenticationConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuRegisterAuthenPageQueryRequestDTO**](OnuRegisterAuthenPageQueryRequestDTO.md) |  | 

### Return type

[**OperationResponseListAuthenticationConfigDTO**](OperationResponseListAuthenticationConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticationConfigPage

> OperationResponsePageResponseAuthenticationConfigDTO GetAuthenticationConfigPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get authentication config page



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
	dto := *openapiclient.NewOnuRegisterAuthenPageQueryRequestDTO("PonPort_example") // OnuRegisterAuthenPageQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.GetAuthenticationConfigPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetAuthenticationConfigPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticationConfigPage`: OperationResponsePageResponseAuthenticationConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetAuthenticationConfigPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAuthenticationConfigPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuRegisterAuthenPageQueryRequestDTO**](OnuRegisterAuthenPageQueryRequestDTO.md) |  | 

### Return type

[**OperationResponsePageResponseAuthenticationConfigDTO**](OperationResponsePageResponseAuthenticationConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuAutoAuthenticationConfig

> OperationResponseAutoAuthenticationConfigDTO GetOnuAutoAuthenticationConfig(ctx, omadacId, siteId, deviceMac).Execute()

Get ONU auto authentication config



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
	resp, r, err := apiClient.OLTONURegisterAPI.GetOnuAutoAuthenticationConfig(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetOnuAutoAuthenticationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuAutoAuthenticationConfig`: OperationResponseAutoAuthenticationConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetOnuAutoAuthenticationConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOnuAutoAuthenticationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseAutoAuthenticationConfigDTO**](OperationResponseAutoAuthenticationConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuAutoAuthenticationRuleConfigList

> OperationResponseListPonAutoAuthenticationRuleConfigDTO GetOnuAutoAuthenticationRuleConfigList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get ONU auto authentication rule list



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
	dto := *openapiclient.NewOnuRegisterAuthenticationPageQueryRequestDTO("PonPort_example") // OnuRegisterAuthenticationPageQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.GetOnuAutoAuthenticationRuleConfigList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetOnuAutoAuthenticationRuleConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuAutoAuthenticationRuleConfigList`: OperationResponseListPonAutoAuthenticationRuleConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetOnuAutoAuthenticationRuleConfigList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOnuAutoAuthenticationRuleConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuRegisterAuthenticationPageQueryRequestDTO**](OnuRegisterAuthenticationPageQueryRequestDTO.md) |  | 

### Return type

[**OperationResponseListPonAutoAuthenticationRuleConfigDTO**](OperationResponseListPonAutoAuthenticationRuleConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuAutoAuthenticationRuleConfigPage

> OperationResponsePageResponsePonAutoAuthenticationRuleConfigDTO GetOnuAutoAuthenticationRuleConfigPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get ONU auto authentication rule page



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
	dto := *openapiclient.NewOnuRegisterAuthenticationPageQueryRequestDTO("PonPort_example") // OnuRegisterAuthenticationPageQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.GetOnuAutoAuthenticationRuleConfigPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetOnuAutoAuthenticationRuleConfigPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuAutoAuthenticationRuleConfigPage`: OperationResponsePageResponsePonAutoAuthenticationRuleConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetOnuAutoAuthenticationRuleConfigPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOnuAutoAuthenticationRuleConfigPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuRegisterAuthenticationPageQueryRequestDTO**](OnuRegisterAuthenticationPageQueryRequestDTO.md) |  | 

### Return type

[**OperationResponsePageResponsePonAutoAuthenticationRuleConfigDTO**](OperationResponsePageResponsePonAutoAuthenticationRuleConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuAutofindConfigList

> OperationResponseListOnuAutofindConfigDTO GetOnuAutofindConfigList(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get ONU auto find config list



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
	dto := *openapiclient.NewOnuRegisterAutoFindPageQueryRequestDTO("PonPort_example") // OnuRegisterAutoFindPageQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.GetOnuAutofindConfigList(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetOnuAutofindConfigList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuAutofindConfigList`: OperationResponseListOnuAutofindConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetOnuAutofindConfigList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOnuAutofindConfigListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuRegisterAutoFindPageQueryRequestDTO**](OnuRegisterAutoFindPageQueryRequestDTO.md) |  | 

### Return type

[**OperationResponseListOnuAutofindConfigDTO**](OperationResponseListOnuAutofindConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuAutofindConfigPage

> OperationResponsePageResponseOnuAutofindConfigDTO GetOnuAutofindConfigPage(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

get ONU auto find config page



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
	dto := *openapiclient.NewOnuRegisterAutoFindPageQueryRequestDTO("PonPort_example") // OnuRegisterAutoFindPageQueryRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.GetOnuAutofindConfigPage(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetOnuAutofindConfigPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuAutofindConfigPage`: OperationResponsePageResponseOnuAutofindConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetOnuAutofindConfigPage`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOnuAutofindConfigPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**OnuRegisterAutoFindPageQueryRequestDTO**](OnuRegisterAutoFindPageQueryRequestDTO.md) |  | 

### Return type

[**OperationResponsePageResponseOnuAutofindConfigDTO**](OperationResponsePageResponseOnuAutofindConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuPonAutoAuthenticationConfig

> OperationResponsePonAutoAuthenticationConfigDTO GetOnuPonAutoAuthenticationConfig(ctx, omadacId, siteId, deviceMac).Dto(dto).Execute()

Get ONU pon auto authentication config



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
	dto := *openapiclient.NewPonPortRequestDTO("PonPort_example") // PonPortRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OLTONURegisterAPI.GetOnuPonAutoAuthenticationConfig(context.Background(), omadacId, siteId, deviceMac).Dto(dto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetOnuPonAutoAuthenticationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuPonAutoAuthenticationConfig`: OperationResponsePonAutoAuthenticationConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetOnuPonAutoAuthenticationConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOnuPonAutoAuthenticationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dto** | [**PonPortRequestDTO**](PonPortRequestDTO.md) |  | 

### Return type

[**OperationResponsePonAutoAuthenticationConfigDTO**](OperationResponsePonAutoAuthenticationConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnuRegisterAutofindConfig

> OperationResponseAutofindConfigDTO GetOnuRegisterAutofindConfig(ctx, omadacId, siteId, deviceMac).Execute()

Get ONU register auto find config



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
	resp, r, err := apiClient.OLTONURegisterAPI.GetOnuRegisterAutofindConfig(context.Background(), omadacId, siteId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OLTONURegisterAPI.GetOnuRegisterAutofindConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnuRegisterAutofindConfig`: OperationResponseAutofindConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `OLTONURegisterAPI.GetOnuRegisterAutofindConfig`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOnuRegisterAutofindConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OperationResponseAutofindConfigDTO**](OperationResponseAutofindConfigDTO.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

