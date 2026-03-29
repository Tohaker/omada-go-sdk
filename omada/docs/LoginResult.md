# LoginResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrfToken** | Pointer to **string** | CSRF token for subsequent requests | [optional] 
**SessionId** | Pointer to **string** | Session ID to use as TPOMADA_SESSIONID cookie | [optional] 

## Methods

### NewLoginResult

`func NewLoginResult() *LoginResult`

NewLoginResult instantiates a new LoginResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginResultWithDefaults

`func NewLoginResultWithDefaults() *LoginResult`

NewLoginResultWithDefaults instantiates a new LoginResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrfToken

`func (o *LoginResult) GetCsrfToken() string`

GetCsrfToken returns the CsrfToken field if non-nil, zero value otherwise.

### GetCsrfTokenOk

`func (o *LoginResult) GetCsrfTokenOk() (*string, bool)`

GetCsrfTokenOk returns a tuple with the CsrfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfToken

`func (o *LoginResult) SetCsrfToken(v string)`

SetCsrfToken sets CsrfToken field to given value.

### HasCsrfToken

`func (o *LoginResult) HasCsrfToken() bool`

HasCsrfToken returns a boolean if a field has been set.

### GetSessionId

`func (o *LoginResult) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *LoginResult) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *LoginResult) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *LoginResult) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


