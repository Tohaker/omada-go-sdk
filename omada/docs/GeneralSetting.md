# GeneralSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dst** | Pointer to [**DstOpenApiVO**](DstOpenApiVO.md) |  | [optional] 
**Name** | **string** | General setting name should be visible ASCII, between 1 and 32 characters. | 
**Region** | **string** | Country/Region of the controller | 
**TimeZone** | **string** | For the values of timeZone, refer to section 5.1 of the Open API Access Guide. | 

## Methods

### NewGeneralSetting

`func NewGeneralSetting(name string, region string, timeZone string, ) *GeneralSetting`

NewGeneralSetting instantiates a new GeneralSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralSettingWithDefaults

`func NewGeneralSettingWithDefaults() *GeneralSetting`

NewGeneralSettingWithDefaults instantiates a new GeneralSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDst

`func (o *GeneralSetting) GetDst() DstOpenApiVO`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *GeneralSetting) GetDstOk() (*DstOpenApiVO, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *GeneralSetting) SetDst(v DstOpenApiVO)`

SetDst sets Dst field to given value.

### HasDst

`func (o *GeneralSetting) HasDst() bool`

HasDst returns a boolean if a field has been set.

### GetName

`func (o *GeneralSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeneralSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeneralSetting) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *GeneralSetting) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GeneralSetting) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GeneralSetting) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetTimeZone

`func (o *GeneralSetting) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GeneralSetting) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GeneralSetting) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


