# MspGeneralSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dst** | Pointer to [**DstDTO**](DstDTO.md) |  | [optional] 
**Name** | Pointer to **string** | Parameter [name] should be within the range of 1–32 visible ASCII characters. | [optional] 
**TimeZone** | Pointer to **string** | For the values of timeZone, refer to section 5.1 of the Open API Access Guide. | [optional] 

## Methods

### NewMspGeneralSettingOpenApiVO

`func NewMspGeneralSettingOpenApiVO() *MspGeneralSettingOpenApiVO`

NewMspGeneralSettingOpenApiVO instantiates a new MspGeneralSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspGeneralSettingOpenApiVOWithDefaults

`func NewMspGeneralSettingOpenApiVOWithDefaults() *MspGeneralSettingOpenApiVO`

NewMspGeneralSettingOpenApiVOWithDefaults instantiates a new MspGeneralSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDst

`func (o *MspGeneralSettingOpenApiVO) GetDst() DstDTO`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *MspGeneralSettingOpenApiVO) GetDstOk() (*DstDTO, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *MspGeneralSettingOpenApiVO) SetDst(v DstDTO)`

SetDst sets Dst field to given value.

### HasDst

`func (o *MspGeneralSettingOpenApiVO) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetName

`func (o *MspGeneralSettingOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MspGeneralSettingOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MspGeneralSettingOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MspGeneralSettingOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *MspGeneralSettingOpenApiVO) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MspGeneralSettingOpenApiVO) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MspGeneralSettingOpenApiVO) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MspGeneralSettingOpenApiVO) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


