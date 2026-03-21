# IptvConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomConfig** | Pointer to [**IptvCustomOpenApiVO**](IptvCustomOpenApiVO.md) |  | [optional] 
**DslConfig** | Pointer to [**IptvDslOpenApiVO**](IptvDslOpenApiVO.md) |  | [optional] 
**Enable** | **bool** |  | 
**Mode** | **int32** | Mode should be a value as follows: 0:Bridge; 1:Custom | 
**PortConfig** | [**[]IptvPortConfigOpenApiVO**](IptvPortConfigOpenApiVO.md) | All available ports need to be configured. The list of port ID is the same as that returned by \&quot;Get IPTV setting\&quot; | 
**WanPortId** | **string** | WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. | 

## Methods

### NewIptvConfigOpenApiVO

`func NewIptvConfigOpenApiVO(enable bool, mode int32, portConfig []IptvPortConfigOpenApiVO, wanPortId string, ) *IptvConfigOpenApiVO`

NewIptvConfigOpenApiVO instantiates a new IptvConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvConfigOpenApiVOWithDefaults

`func NewIptvConfigOpenApiVOWithDefaults() *IptvConfigOpenApiVO`

NewIptvConfigOpenApiVOWithDefaults instantiates a new IptvConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomConfig

`func (o *IptvConfigOpenApiVO) GetCustomConfig() IptvCustomOpenApiVO`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *IptvConfigOpenApiVO) GetCustomConfigOk() (*IptvCustomOpenApiVO, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *IptvConfigOpenApiVO) SetCustomConfig(v IptvCustomOpenApiVO)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *IptvConfigOpenApiVO) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### GetDslConfig

`func (o *IptvConfigOpenApiVO) GetDslConfig() IptvDslOpenApiVO`

GetDslConfig returns the DslConfig field if non-nil, zero value otherwise.

### GetDslConfigOk

`func (o *IptvConfigOpenApiVO) GetDslConfigOk() (*IptvDslOpenApiVO, bool)`

GetDslConfigOk returns a tuple with the DslConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslConfig

`func (o *IptvConfigOpenApiVO) SetDslConfig(v IptvDslOpenApiVO)`

SetDslConfig sets DslConfig field to given value.

### HasDslConfig

`func (o *IptvConfigOpenApiVO) HasDslConfig() bool`

HasDslConfig returns a boolean if a field has been set.

### GetEnable

`func (o *IptvConfigOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IptvConfigOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IptvConfigOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMode

`func (o *IptvConfigOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IptvConfigOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IptvConfigOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetPortConfig

`func (o *IptvConfigOpenApiVO) GetPortConfig() []IptvPortConfigOpenApiVO`

GetPortConfig returns the PortConfig field if non-nil, zero value otherwise.

### GetPortConfigOk

`func (o *IptvConfigOpenApiVO) GetPortConfigOk() (*[]IptvPortConfigOpenApiVO, bool)`

GetPortConfigOk returns a tuple with the PortConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfig

`func (o *IptvConfigOpenApiVO) SetPortConfig(v []IptvPortConfigOpenApiVO)`

SetPortConfig sets PortConfig field to given value.


### GetWanPortId

`func (o *IptvConfigOpenApiVO) GetWanPortId() string`

GetWanPortId returns the WanPortId field if non-nil, zero value otherwise.

### GetWanPortIdOk

`func (o *IptvConfigOpenApiVO) GetWanPortIdOk() (*string, bool)`

GetWanPortIdOk returns a tuple with the WanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortId

`func (o *IptvConfigOpenApiVO) SetWanPortId(v string)`

SetWanPortId sets WanPortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


