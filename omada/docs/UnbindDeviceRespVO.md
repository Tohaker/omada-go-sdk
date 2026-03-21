# UnbindDeviceRespVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category should be a value as follows: ap; l2Switch; l3Switch; gateway | [optional] 
**LicenseId** | Pointer to **string** | License ID | [optional] 
**LicenseType** | Pointer to **string** | License type should be a value as follows: 1year; 2years; 3years; 4years; 5years; others; trial(Cloud Based Controller), permanent; trial(Local Controller) | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | Unbind device status should be a value as follows: 0: success;-1: failed | [optional] 

## Methods

### NewUnbindDeviceRespVO

`func NewUnbindDeviceRespVO() *UnbindDeviceRespVO`

NewUnbindDeviceRespVO instantiates a new UnbindDeviceRespVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbindDeviceRespVOWithDefaults

`func NewUnbindDeviceRespVOWithDefaults() *UnbindDeviceRespVO`

NewUnbindDeviceRespVOWithDefaults instantiates a new UnbindDeviceRespVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *UnbindDeviceRespVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UnbindDeviceRespVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UnbindDeviceRespVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UnbindDeviceRespVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLicenseId

`func (o *UnbindDeviceRespVO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *UnbindDeviceRespVO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *UnbindDeviceRespVO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *UnbindDeviceRespVO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseType

`func (o *UnbindDeviceRespVO) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *UnbindDeviceRespVO) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *UnbindDeviceRespVO) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *UnbindDeviceRespVO) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetMac

`func (o *UnbindDeviceRespVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *UnbindDeviceRespVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *UnbindDeviceRespVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *UnbindDeviceRespVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *UnbindDeviceRespVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnbindDeviceRespVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnbindDeviceRespVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnbindDeviceRespVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


