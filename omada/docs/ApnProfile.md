# ApnProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** | Access point name, only for apnType static, should contain 1 to 64 characters and should meet the following requirements:     1. Should not start with \&quot;rac\&quot;, \&quot;lac\&quot;, \&quot;sgsn\&quot; or \&quot;rnc\&quot; and end with \&quot;.gprs\&quot;, ignoring case.     2. Should not contain \&quot;.*.\&quot;, \&quot;.-\&quot;, \&quot;-.\&quot; and spaces.     3. Should not contain any of the characters #!$%^&amp;*(),:;\&quot;&#39;|\\@.  | [optional] 
**ApnType** | **int32** | ApnType should be a value as follows: 0: static; 1: dynamic. | 
**ApplyToSIM** | Pointer to **int32** | 1: apply to SIM1; 2: apply to SIM2; 3: apply to SIM1 and SIM2. | [optional] 
**Authentication** | **int32** | Authentication should be a value as follows: 0: None; 1: PAP; 2: CHAP | 
**BuildIn** | Pointer to **bool** | Indicates whether this APN profile is a built-in APN profile for SIM card. | [optional] 
**Id** | Pointer to **string** | APN profile ID | [optional] 
**Name** | **string** | APN profile name, name should contain 1 to 64 characters. | 
**Password** | Pointer to **string** | Password should contain 1 to 64 characters, spaces, comma, single quotation marks and double quotation marks are not allowed. | [optional] 
**PdpType** | **int32** | PdpType should be a value as follows: 0: IPv4; 1: IPv6; 2: IPv4 &amp; IPv6. | 
**Username** | Pointer to **string** | Username should contain 1 to 64 characters, spaces, comma, single quotation marks and double quotation marks are not allowed. | [optional] 

## Methods

### NewApnProfile

`func NewApnProfile(apnType int32, authentication int32, name string, pdpType int32, ) *ApnProfile`

NewApnProfile instantiates a new ApnProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApnProfileWithDefaults

`func NewApnProfileWithDefaults() *ApnProfile`

NewApnProfileWithDefaults instantiates a new ApnProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *ApnProfile) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *ApnProfile) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *ApnProfile) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *ApnProfile) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetApnType

`func (o *ApnProfile) GetApnType() int32`

GetApnType returns the ApnType field if non-nil, zero value otherwise.

### GetApnTypeOk

`func (o *ApnProfile) GetApnTypeOk() (*int32, bool)`

GetApnTypeOk returns a tuple with the ApnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnType

`func (o *ApnProfile) SetApnType(v int32)`

SetApnType sets ApnType field to given value.


### GetApplyToSIM

`func (o *ApnProfile) GetApplyToSIM() int32`

GetApplyToSIM returns the ApplyToSIM field if non-nil, zero value otherwise.

### GetApplyToSIMOk

`func (o *ApnProfile) GetApplyToSIMOk() (*int32, bool)`

GetApplyToSIMOk returns a tuple with the ApplyToSIM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToSIM

`func (o *ApnProfile) SetApplyToSIM(v int32)`

SetApplyToSIM sets ApplyToSIM field to given value.

### HasApplyToSIM

`func (o *ApnProfile) HasApplyToSIM() bool`

HasApplyToSIM returns a boolean if a field has been set.

### GetAuthentication

`func (o *ApnProfile) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ApnProfile) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ApnProfile) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.


### GetBuildIn

`func (o *ApnProfile) GetBuildIn() bool`

GetBuildIn returns the BuildIn field if non-nil, zero value otherwise.

### GetBuildInOk

`func (o *ApnProfile) GetBuildInOk() (*bool, bool)`

GetBuildInOk returns a tuple with the BuildIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildIn

`func (o *ApnProfile) SetBuildIn(v bool)`

SetBuildIn sets BuildIn field to given value.

### HasBuildIn

`func (o *ApnProfile) HasBuildIn() bool`

HasBuildIn returns a boolean if a field has been set.

### GetId

`func (o *ApnProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApnProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApnProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApnProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApnProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApnProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApnProfile) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *ApnProfile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApnProfile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApnProfile) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApnProfile) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPdpType

`func (o *ApnProfile) GetPdpType() int32`

GetPdpType returns the PdpType field if non-nil, zero value otherwise.

### GetPdpTypeOk

`func (o *ApnProfile) GetPdpTypeOk() (*int32, bool)`

GetPdpTypeOk returns a tuple with the PdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdpType

`func (o *ApnProfile) SetPdpType(v int32)`

SetPdpType sets PdpType field to given value.


### GetUsername

`func (o *ApnProfile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApnProfile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApnProfile) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApnProfile) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


