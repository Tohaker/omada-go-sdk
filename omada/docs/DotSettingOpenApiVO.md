# DotSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomizedServers** | Pointer to [**[]DotCustomizedServerOpenApiVO**](DotCustomizedServerOpenApiVO.md) | Custom Service list. Up to 2 DoT default and custom Servers can be selected, setting the parameter [enable] to true indicates the selection of the custom server. | [optional] 
**DefaultServers** | Pointer to **[]int32** | Preconfigured Server List. DefaultServers should be a value as follows: 0:Google, 1：Cloudflare, 2：OpenDNS, 3：Quad9, 4: CleanBrowsing. Up to 2 DoT default and custom servers can be selected. For example, defaultServers : [0, 1] represents that you have selected default services Google and Cloudflare | [optional] 

## Methods

### NewDotSettingOpenApiVO

`func NewDotSettingOpenApiVO() *DotSettingOpenApiVO`

NewDotSettingOpenApiVO instantiates a new DotSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDotSettingOpenApiVOWithDefaults

`func NewDotSettingOpenApiVOWithDefaults() *DotSettingOpenApiVO`

NewDotSettingOpenApiVOWithDefaults instantiates a new DotSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizedServers

`func (o *DotSettingOpenApiVO) GetCustomizedServers() []DotCustomizedServerOpenApiVO`

GetCustomizedServers returns the CustomizedServers field if non-nil, zero value otherwise.

### GetCustomizedServersOk

`func (o *DotSettingOpenApiVO) GetCustomizedServersOk() (*[]DotCustomizedServerOpenApiVO, bool)`

GetCustomizedServersOk returns a tuple with the CustomizedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedServers

`func (o *DotSettingOpenApiVO) SetCustomizedServers(v []DotCustomizedServerOpenApiVO)`

SetCustomizedServers sets CustomizedServers field to given value.

### HasCustomizedServers

`func (o *DotSettingOpenApiVO) HasCustomizedServers() bool`

HasCustomizedServers returns a boolean if a field has been set.

### GetDefaultServers

`func (o *DotSettingOpenApiVO) GetDefaultServers() []int32`

GetDefaultServers returns the DefaultServers field if non-nil, zero value otherwise.

### GetDefaultServersOk

`func (o *DotSettingOpenApiVO) GetDefaultServersOk() (*[]int32, bool)`

GetDefaultServersOk returns a tuple with the DefaultServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServers

`func (o *DotSettingOpenApiVO) SetDefaultServers(v []int32)`

SetDefaultServers sets DefaultServers field to given value.

### HasDefaultServers

`func (o *DotSettingOpenApiVO) HasDefaultServers() bool`

HasDefaultServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


