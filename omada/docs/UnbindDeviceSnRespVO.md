# UnbindDeviceSnRespVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | ap, l2Switch, l3Switch, gateway | [optional] 
**LicenseId** | Pointer to **string** | License ID | [optional] 
**LicenseType** | Pointer to **string** | 1year, 2years, 3years, 4years, 5years, others, trial | [optional] 
**Sn** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | unbind device status(0: success;-1: failed;) | [optional] 

## Methods

### NewUnbindDeviceSnRespVO

`func NewUnbindDeviceSnRespVO() *UnbindDeviceSnRespVO`

NewUnbindDeviceSnRespVO instantiates a new UnbindDeviceSnRespVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbindDeviceSnRespVOWithDefaults

`func NewUnbindDeviceSnRespVOWithDefaults() *UnbindDeviceSnRespVO`

NewUnbindDeviceSnRespVOWithDefaults instantiates a new UnbindDeviceSnRespVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *UnbindDeviceSnRespVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UnbindDeviceSnRespVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UnbindDeviceSnRespVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UnbindDeviceSnRespVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLicenseId

`func (o *UnbindDeviceSnRespVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *UnbindDeviceSnRespVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *UnbindDeviceSnRespVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *UnbindDeviceSnRespVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseType

`func (o *UnbindDeviceSnRespVO) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *UnbindDeviceSnRespVO) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *UnbindDeviceSnRespVO) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *UnbindDeviceSnRespVO) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetSn

`func (o *UnbindDeviceSnRespVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *UnbindDeviceSnRespVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *UnbindDeviceSnRespVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *UnbindDeviceSnRespVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *UnbindDeviceSnRespVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnbindDeviceSnRespVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnbindDeviceSnRespVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnbindDeviceSnRespVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


