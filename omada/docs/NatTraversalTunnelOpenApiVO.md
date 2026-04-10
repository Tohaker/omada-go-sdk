# NatTraversalTunnelOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppType** | **string** | App type of the remote access tunnel for the local target device. HTTP, HTTPS, SSH, TELNET. | 
**Duration** | Pointer to **int32** | Valid duration time of the remote access tunnel, 1-24 hours. | [optional] 
**LocalAddress** | **string** | IP address of the local target device. | 
**LocalPort** | **int32** | Port of the local target device&#39;s service. | 
**Name** | **string** | Name of the remote access tunnel. | 
**OpenStatus** | Pointer to **bool** | If open the remote access tunnel after create. | [optional] 

## Methods

### NewNatTraversalTunnelOpenApiVO

`func NewNatTraversalTunnelOpenApiVO(appType string, localAddress string, localPort int32, name string, ) *NatTraversalTunnelOpenApiVO`

NewNatTraversalTunnelOpenApiVO instantiates a new NatTraversalTunnelOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatTraversalTunnelOpenApiVOWithDefaults

`func NewNatTraversalTunnelOpenApiVOWithDefaults() *NatTraversalTunnelOpenApiVO`

NewNatTraversalTunnelOpenApiVOWithDefaults instantiates a new NatTraversalTunnelOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppType

`func (o *NatTraversalTunnelOpenApiVO) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *NatTraversalTunnelOpenApiVO) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *NatTraversalTunnelOpenApiVO) SetAppType(v string)`

SetAppType sets AppType field to given value.


### GetDuration

`func (o *NatTraversalTunnelOpenApiVO) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NatTraversalTunnelOpenApiVO) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NatTraversalTunnelOpenApiVO) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NatTraversalTunnelOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLocalAddress

`func (o *NatTraversalTunnelOpenApiVO) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *NatTraversalTunnelOpenApiVO) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *NatTraversalTunnelOpenApiVO) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.


### GetLocalPort

`func (o *NatTraversalTunnelOpenApiVO) GetLocalPort() int32`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *NatTraversalTunnelOpenApiVO) GetLocalPortOk() (*int32, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *NatTraversalTunnelOpenApiVO) SetLocalPort(v int32)`

SetLocalPort sets LocalPort field to given value.


### GetName

`func (o *NatTraversalTunnelOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatTraversalTunnelOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatTraversalTunnelOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOpenStatus

`func (o *NatTraversalTunnelOpenApiVO) GetOpenStatus() bool`

GetOpenStatus returns the OpenStatus field if non-nil, zero value otherwise.

### GetOpenStatusOk

`func (o *NatTraversalTunnelOpenApiVO) GetOpenStatusOk() (*bool, bool)`

GetOpenStatusOk returns a tuple with the OpenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatus

`func (o *NatTraversalTunnelOpenApiVO) SetOpenStatus(v bool)`

SetOpenStatus sets OpenStatus field to given value.

### HasOpenStatus

`func (o *NatTraversalTunnelOpenApiVO) HasOpenStatus() bool`

HasOpenStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


