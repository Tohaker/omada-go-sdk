# AuthorizeAPI

All URIs are relative to *https://use1-omada-northbound.tplinkcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeCode**](AuthorizeAPI.md#authorizecode) | **Post** /openapi/authorize/code | Get authorization code (step 2)
[**AuthorizeLogin**](AuthorizeAPI.md#authorizelogin) | **Post** /openapi/authorize/login | Login (Authorization Code flow step 1)
[**AuthorizeToken**](AuthorizeAPI.md#authorizetoken) | **Post** /openapi/authorize/token | Get or refresh access token



## AuthorizeCode

> AuthCodeResponse AuthorizeCode(ctx).ClientId(clientId).OmadacId(omadacId).ResponseType(responseType).CsrfToken(csrfToken).Cookie(cookie).Execute()

Get authorization code (step 2)



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
	clientId := "clientId_example" // string | OAuth client ID
	omadacId := "omadacId_example" // string | Omada Controller ID
	responseType := "responseType_example" // string | Must be \"code\"
	csrfToken := "csrfToken_example" // string | CSRF token from login response
	cookie := "cookie_example" // string | Session cookie, format: TPOMADA_SESSIONID=<sessionId>

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizeAPI.AuthorizeCode(context.Background()).ClientId(clientId).OmadacId(omadacId).ResponseType(responseType).CsrfToken(csrfToken).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeAPI.AuthorizeCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizeCode`: AuthCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizeAPI.AuthorizeCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | OAuth client ID | 
 **omadacId** | **string** | Omada Controller ID | 
 **responseType** | **string** | Must be \&quot;code\&quot; | 
 **csrfToken** | **string** | CSRF token from login response | 
 **cookie** | **string** | Session cookie, format: TPOMADA_SESSIONID&#x3D;&lt;sessionId&gt; | 

### Return type

[**AuthCodeResponse**](AuthCodeResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizeLogin

> LoginResponse AuthorizeLogin(ctx).ClientId(clientId).OmadacId(omadacId).LoginRequest(loginRequest).Execute()

Login (Authorization Code flow step 1)



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
	clientId := "clientId_example" // string | OAuth client ID
	omadacId := "omadacId_example" // string | Omada Controller ID
	loginRequest := *openapiclient.NewLoginRequest("Password_example", "Username_example") // LoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizeAPI.AuthorizeLogin(context.Background()).ClientId(clientId).OmadacId(omadacId).LoginRequest(loginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeAPI.AuthorizeLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizeLogin`: LoginResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizeAPI.AuthorizeLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | OAuth client ID | 
 **omadacId** | **string** | Omada Controller ID | 
 **loginRequest** | [**LoginRequest**](LoginRequest.md) |  | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizeToken

> TokenResponse AuthorizeToken(ctx).GrantType(grantType).TokenRequest(tokenRequest).Code(code).RefreshToken(refreshToken).Execute()

Get or refresh access token



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
	grantType := "grantType_example" // string | OAuth grant type
	tokenRequest := *openapiclient.NewTokenRequest("ClientId_example", "ClientSecret_example") // TokenRequest | 
	code := "code_example" // string | Authorization code (required for authorization_code grant) (optional)
	refreshToken := "refreshToken_example" // string | Refresh token (required for refresh_token grant) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizeAPI.AuthorizeToken(context.Background()).GrantType(grantType).TokenRequest(tokenRequest).Code(code).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeAPI.AuthorizeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizeToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizeAPI.AuthorizeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | OAuth grant type | 
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 
 **code** | **string** | Authorization code (required for authorization_code grant) | 
 **refreshToken** | **string** | Refresh token (required for refresh_token grant) | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[AccessToken](../README.md#accesstoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

