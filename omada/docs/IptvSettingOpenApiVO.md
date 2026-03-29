# IptvSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgmpSetting** | [**IgmpOpenApiVO**](IgmpOpenApiVO.md) |  | 
**IptvSetting** | Pointer to [**IptvOpenApiVO**](IptvOpenApiVO.md) |  | [optional] 

## Methods

### NewIptvSettingOpenApiVO

`func NewIptvSettingOpenApiVO(igmpSetting IgmpOpenApiVO, ) *IptvSettingOpenApiVO`

NewIptvSettingOpenApiVO instantiates a new IptvSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvSettingOpenApiVOWithDefaults

`func NewIptvSettingOpenApiVOWithDefaults() *IptvSettingOpenApiVO`

NewIptvSettingOpenApiVOWithDefaults instantiates a new IptvSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgmpSetting

`func (o *IptvSettingOpenApiVO) GetIgmpSetting() IgmpOpenApiVO`

GetIgmpSetting returns the IgmpSetting field if non-nil, zero value otherwise.

### GetIgmpSettingOk

`func (o *IptvSettingOpenApiVO) GetIgmpSettingOk() (*IgmpOpenApiVO, bool)`

GetIgmpSettingOk returns a tuple with the IgmpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSetting

`func (o *IptvSettingOpenApiVO) SetIgmpSetting(v IgmpOpenApiVO)`

SetIgmpSetting sets IgmpSetting field to given value.


### GetIptvSetting

`func (o *IptvSettingOpenApiVO) GetIptvSetting() IptvOpenApiVO`

GetIptvSetting returns the IptvSetting field if non-nil, zero value otherwise.

### GetIptvSettingOk

`func (o *IptvSettingOpenApiVO) GetIptvSettingOk() (*IptvOpenApiVO, bool)`

GetIptvSettingOk returns a tuple with the IptvSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIptvSetting

`func (o *IptvSettingOpenApiVO) SetIptvSetting(v IptvOpenApiVO)`

SetIptvSetting sets IptvSetting field to given value.

### HasIptvSetting

`func (o *IptvSettingOpenApiVO) HasIptvSetting() bool`

HasIptvSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


