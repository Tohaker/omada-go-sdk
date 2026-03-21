# PortIpOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP address | [optional] 
**WanId** | Pointer to **string** | This field represents WAN port ID. WAN port ID can be obtained from can be obtained from &#39;Get internet basic info&#39; interface. | [optional] 

## Methods

### NewPortIpOpenApiVO

`func NewPortIpOpenApiVO() *PortIpOpenApiVO`

NewPortIpOpenApiVO instantiates a new PortIpOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortIpOpenApiVOWithDefaults

`func NewPortIpOpenApiVOWithDefaults() *PortIpOpenApiVO`

NewPortIpOpenApiVOWithDefaults instantiates a new PortIpOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *PortIpOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PortIpOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PortIpOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *PortIpOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetWanId

`func (o *PortIpOpenApiVO) GetWanId() string`

GetWanId returns the WanId field if non-nil, zero value otherwise.

### GetWanIdOk

`func (o *PortIpOpenApiVO) GetWanIdOk() (*string, bool)`

GetWanIdOk returns a tuple with the WanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanId

`func (o *PortIpOpenApiVO) SetWanId(v string)`

SetWanId sets WanId field to given value.

### HasWanId

`func (o *PortIpOpenApiVO) HasWanId() bool`

HasWanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


