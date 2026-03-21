# SnmpSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunityString** | Pointer to **string** | Community string, valid when parameter [snmpV1V2CEnable] is true. The communityString should contain at least 10 characters, using a combination of numbers, letters or special characters.  The communityString should not contain consecutive identical characters. | [optional] 
**Password** | Pointer to **string** | The password should contain at least 10 characters, using a combination of numbers, letters or special characters.  The password should not contain consecutive identical characters.  Username and Password should not be the same. | [optional] 
**SnmpV1V2CEnable** | **bool** | SNMPv1 &amp; SNMPv2c enable status | 
**SnmpV3Enable** | **bool** | SNMPv3 enable status | 
**Username** | Pointer to **string** | Username, valid when parameter [snmpV3Enable] is true. Username should contain 1 to 30 characters | [optional] 

## Methods

### NewSnmpSettingOpenApiVO

`func NewSnmpSettingOpenApiVO(snmpV1V2CEnable bool, snmpV3Enable bool, ) *SnmpSettingOpenApiVO`

NewSnmpSettingOpenApiVO instantiates a new SnmpSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpSettingOpenApiVOWithDefaults

`func NewSnmpSettingOpenApiVOWithDefaults() *SnmpSettingOpenApiVO`

NewSnmpSettingOpenApiVOWithDefaults instantiates a new SnmpSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunityString

`func (o *SnmpSettingOpenApiVO) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *SnmpSettingOpenApiVO) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *SnmpSettingOpenApiVO) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *SnmpSettingOpenApiVO) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetPassword

`func (o *SnmpSettingOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SnmpSettingOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SnmpSettingOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SnmpSettingOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSnmpV1V2CEnable

`func (o *SnmpSettingOpenApiVO) GetSnmpV1V2CEnable() bool`

GetSnmpV1V2CEnable returns the SnmpV1V2CEnable field if non-nil, zero value otherwise.

### GetSnmpV1V2CEnableOk

`func (o *SnmpSettingOpenApiVO) GetSnmpV1V2CEnableOk() (*bool, bool)`

GetSnmpV1V2CEnableOk returns a tuple with the SnmpV1V2CEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpV1V2CEnable

`func (o *SnmpSettingOpenApiVO) SetSnmpV1V2CEnable(v bool)`

SetSnmpV1V2CEnable sets SnmpV1V2CEnable field to given value.


### GetSnmpV3Enable

`func (o *SnmpSettingOpenApiVO) GetSnmpV3Enable() bool`

GetSnmpV3Enable returns the SnmpV3Enable field if non-nil, zero value otherwise.

### GetSnmpV3EnableOk

`func (o *SnmpSettingOpenApiVO) GetSnmpV3EnableOk() (*bool, bool)`

GetSnmpV3EnableOk returns a tuple with the SnmpV3Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpV3Enable

`func (o *SnmpSettingOpenApiVO) SetSnmpV3Enable(v bool)`

SetSnmpV3Enable sets SnmpV3Enable field to given value.


### GetUsername

`func (o *SnmpSettingOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SnmpSettingOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SnmpSettingOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SnmpSettingOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


