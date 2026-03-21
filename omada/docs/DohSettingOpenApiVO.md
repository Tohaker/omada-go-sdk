# DohSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomizedServers** | Pointer to [**[]DohCustomizedServerOpenApiVO**](DohCustomizedServerOpenApiVO.md) | Custom Service list. Up to 2 DoH default and custom Servers can be selected, setting the parameter [enable] to true indicates the selection of the custom server. | [optional] 
**DefaultServers** | Pointer to **[]int32** | Preconfigured Server List. DefaultServers should be a value as follows: 0:Google, 1：Cloudflare, 4：CleanBrowsing, 5：Quad9_1, 6: Quad9_2. Up to 2 DoH default and custom servers can be selected. For example, defaultServers : [0, 1] represents that you have selected default services Google and Cloudflare | [optional] 

## Methods

### NewDohSettingOpenApiVO

`func NewDohSettingOpenApiVO() *DohSettingOpenApiVO`

NewDohSettingOpenApiVO instantiates a new DohSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDohSettingOpenApiVOWithDefaults

`func NewDohSettingOpenApiVOWithDefaults() *DohSettingOpenApiVO`

NewDohSettingOpenApiVOWithDefaults instantiates a new DohSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizedServers

`func (o *DohSettingOpenApiVO) GetCustomizedServers() []DohCustomizedServerOpenApiVO`

GetCustomizedServers returns the CustomizedServers field if non-nil, zero value otherwise.

### GetCustomizedServersOk

`func (o *DohSettingOpenApiVO) GetCustomizedServersOk() (*[]DohCustomizedServerOpenApiVO, bool)`

GetCustomizedServersOk returns a tuple with the CustomizedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedServers

`func (o *DohSettingOpenApiVO) SetCustomizedServers(v []DohCustomizedServerOpenApiVO)`

SetCustomizedServers sets CustomizedServers field to given value.

### HasCustomizedServers

`func (o *DohSettingOpenApiVO) HasCustomizedServers() bool`

HasCustomizedServers returns a boolean if a field has been set.

### GetDefaultServers

`func (o *DohSettingOpenApiVO) GetDefaultServers() []int32`

GetDefaultServers returns the DefaultServers field if non-nil, zero value otherwise.

### GetDefaultServersOk

`func (o *DohSettingOpenApiVO) GetDefaultServersOk() (*[]int32, bool)`

GetDefaultServersOk returns a tuple with the DefaultServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServers

`func (o *DohSettingOpenApiVO) SetDefaultServers(v []int32)`

SetDefaultServers sets DefaultServers field to given value.

### HasDefaultServers

`func (o *DohSettingOpenApiVO) HasDefaultServers() bool`

HasDefaultServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


