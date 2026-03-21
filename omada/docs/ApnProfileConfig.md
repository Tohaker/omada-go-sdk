# ApnProfileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** | Access point name, only for apnType static, should contain 1 to 62 characters and should meet the following requirements:     1. Only allow letters, numbers, hyphens and dots.     2. Should not start with \&quot;rac\&quot;, \&quot;lac\&quot;, \&quot;sgsn\&quot; or \&quot;rnc\&quot; and end with \&quot;.gprs\&quot;, ignoring case.     3. Should not contain \&quot;.-\&quot;, \&quot;-.\&quot; and spaces.  | [optional] 
**ApnType** | **int32** | ApnType should be a value as follows: 0: static; 1: dynamic. | 
**ApplyToSIM** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [applyToSim] should be a value as follows: 1: apply to SIM1; 2: apply to SIM2; 3: apply to SIM1 and SIM2. | [optional] 
**Authentication** | **int32** | Authentication should be a value as follows: 0: None; 1: PAP; 2: CHAP | 
**Name** | **string** | APN profile name, name should contain 1 to 64 characters. | 
**Password** | Pointer to **string** | Password should contain 1 to 64 characters, spaces, comma, single quotation marks and double quotation marks are not allowed. | [optional] 
**PdpType** | **int32** | PdpType should be a value as follows: 0: IPv4; 1: IPv6; 2: IPv4 &amp; IPv6. | 
**Username** | Pointer to **string** | Username should contain 1 to 64 characters, spaces, comma, single quotation marks and double quotation marks are not allowed. | [optional] 

## Methods

### NewApnProfileConfig

`func NewApnProfileConfig(apnType int32, authentication int32, name string, pdpType int32, ) *ApnProfileConfig`

NewApnProfileConfig instantiates a new ApnProfileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApnProfileConfigWithDefaults

`func NewApnProfileConfigWithDefaults() *ApnProfileConfig`

NewApnProfileConfigWithDefaults instantiates a new ApnProfileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *ApnProfileConfig) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *ApnProfileConfig) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *ApnProfileConfig) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *ApnProfileConfig) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetApnType

`func (o *ApnProfileConfig) GetApnType() int32`

GetApnType returns the ApnType field if non-nil, zero value otherwise.

### GetApnTypeOk

`func (o *ApnProfileConfig) GetApnTypeOk() (*int32, bool)`

GetApnTypeOk returns a tuple with the ApnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnType

`func (o *ApnProfileConfig) SetApnType(v int32)`

SetApnType sets ApnType field to given value.


### GetApplyToSIM

`func (o *ApnProfileConfig) GetApplyToSIM() int32`

GetApplyToSIM returns the ApplyToSIM field if non-nil, zero value otherwise.

### GetApplyToSIMOk

`func (o *ApnProfileConfig) GetApplyToSIMOk() (*int32, bool)`

GetApplyToSIMOk returns a tuple with the ApplyToSIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToSIM

`func (o *ApnProfileConfig) SetApplyToSIM(v int32)`

SetApplyToSIM sets ApplyToSIM field to given value.

### HasApplyToSIM

`func (o *ApnProfileConfig) HasApplyToSIM() bool`

HasApplyToSIM returns a boolean if a field has been set.

### GetAuthentication

`func (o *ApnProfileConfig) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ApnProfileConfig) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ApnProfileConfig) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.


### GetName

`func (o *ApnProfileConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApnProfileConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApnProfileConfig) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *ApnProfileConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApnProfileConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApnProfileConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApnProfileConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPdpType

`func (o *ApnProfileConfig) GetPdpType() int32`

GetPdpType returns the PdpType field if non-nil, zero value otherwise.

### GetPdpTypeOk

`func (o *ApnProfileConfig) GetPdpTypeOk() (*int32, bool)`

GetPdpTypeOk returns a tuple with the PdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpType

`func (o *ApnProfileConfig) SetPdpType(v int32)`

SetPdpType sets PdpType field to given value.


### GetUsername

`func (o *ApnProfileConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApnProfileConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApnProfileConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApnProfileConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


