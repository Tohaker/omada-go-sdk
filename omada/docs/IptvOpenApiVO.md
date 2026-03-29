# IptvOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomConfig** | Pointer to [**IptvCustomOpenApiVO**](IptvCustomOpenApiVO.md) |  | [optional] 
**DslConfig** | Pointer to [**IptvDslOpenApiVO**](IptvDslOpenApiVO.md) |  | [optional] 
**Enable** | **bool** | Whether to enable iptv feature. Options: &#39;true&#39; or &#39;false&#39;. | 
**Mode** | **int32** | Mode should be a value as follows: 0:Bridge, 1:Custom | 
**PortConfig** | [**[]IptvPortOpenApiVO**](IptvPortOpenApiVO.md) | Config the port mode of the LAN ports to determine which port is used to support Internet service, IPTV service, or IP Phone service. | 
**WanPortId** | **string** | WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 

## Methods

### NewIptvOpenApiVO

`func NewIptvOpenApiVO(enable bool, mode int32, portConfig []IptvPortOpenApiVO, wanPortId string, ) *IptvOpenApiVO`

NewIptvOpenApiVO instantiates a new IptvOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvOpenApiVOWithDefaults

`func NewIptvOpenApiVOWithDefaults() *IptvOpenApiVO`

NewIptvOpenApiVOWithDefaults instantiates a new IptvOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomConfig

`func (o *IptvOpenApiVO) GetCustomConfig() IptvCustomOpenApiVO`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *IptvOpenApiVO) GetCustomConfigOk() (*IptvCustomOpenApiVO, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *IptvOpenApiVO) SetCustomConfig(v IptvCustomOpenApiVO)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *IptvOpenApiVO) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDslConfig

`func (o *IptvOpenApiVO) GetDslConfig() IptvDslOpenApiVO`

GetDslConfig returns the DslConfig field if non-nil, zero value otherwise.

### GetDslConfigOk

`func (o *IptvOpenApiVO) GetDslConfigOk() (*IptvDslOpenApiVO, bool)`

GetDslConfigOk returns a tuple with the DslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslConfig

`func (o *IptvOpenApiVO) SetDslConfig(v IptvDslOpenApiVO)`

SetDslConfig sets DslConfig field to given value.

### HasDslConfig

`func (o *IptvOpenApiVO) HasDslConfig() bool`

HasDslConfig returns a boolean if a field has been set.

### GetEnable

`func (o *IptvOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IptvOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IptvOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMode

`func (o *IptvOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IptvOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IptvOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetPortConfig

`func (o *IptvOpenApiVO) GetPortConfig() []IptvPortOpenApiVO`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *IptvOpenApiVO) GetPortConfigOk() (*[]IptvPortOpenApiVO, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *IptvOpenApiVO) SetPortConfig(v []IptvPortOpenApiVO)`

SetPortConfig sets PortConfig field to given value.


### GetWanPortId

`func (o *IptvOpenApiVO) GetWanPortId() string`

GetWanPortId returns the WanPortId field if non-nil, zero value otherwise.

### GetWanPortIdOk

`func (o *IptvOpenApiVO) GetWanPortIdOk() (*string, bool)`

GetWanPortIdOk returns a tuple with the WanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortId

`func (o *IptvOpenApiVO) SetWanPortId(v string)`

SetWanPortId sets WanPortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


