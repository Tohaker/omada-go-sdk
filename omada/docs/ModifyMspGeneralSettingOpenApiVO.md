# ModifyMspGeneralSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dst** | Pointer to [**ModifyDstDTO**](ModifyDstDTO.md) |  | [optional] 
**Name** | Pointer to **string** | Parameter [name] should be within the range of 1–32 visible ASCII characters. | [optional] 
**TimeZone** | Pointer to **string** | For the values of timeZone, refer to section 5.1 of the Open API Access Guide. | [optional] 

## Methods

### NewModifyMspGeneralSettingOpenApiVO

`func NewModifyMspGeneralSettingOpenApiVO() *ModifyMspGeneralSettingOpenApiVO`

NewModifyMspGeneralSettingOpenApiVO instantiates a new ModifyMspGeneralSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMspGeneralSettingOpenApiVOWithDefaults

`func NewModifyMspGeneralSettingOpenApiVOWithDefaults() *ModifyMspGeneralSettingOpenApiVO`

NewModifyMspGeneralSettingOpenApiVOWithDefaults instantiates a new ModifyMspGeneralSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDst

`func (o *ModifyMspGeneralSettingOpenApiVO) GetDst() ModifyDstDTO`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *ModifyMspGeneralSettingOpenApiVO) GetDstOk() (*ModifyDstDTO, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *ModifyMspGeneralSettingOpenApiVO) SetDst(v ModifyDstDTO)`

SetDst sets Dst field to given value.

### HasDst

`func (o *ModifyMspGeneralSettingOpenApiVO) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetName

`func (o *ModifyMspGeneralSettingOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyMspGeneralSettingOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyMspGeneralSettingOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyMspGeneralSettingOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *ModifyMspGeneralSettingOpenApiVO) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ModifyMspGeneralSettingOpenApiVO) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ModifyMspGeneralSettingOpenApiVO) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ModifyMspGeneralSettingOpenApiVO) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


