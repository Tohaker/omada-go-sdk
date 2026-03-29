# NetworkWithServerVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | LAN Network ID | [optional] 
**Name** | Pointer to **string** | LAN Network Name | [optional] 
**ServerDevices** | Pointer to [**[]BriefServerDeviceVO**](BriefServerDeviceVO.md) | List of devices acting as DHCP servers in this network | [optional] 
**Type** | Pointer to **string** | LAN Network Type | [optional] 

## Methods

### NewNetworkWithServerVO

`func NewNetworkWithServerVO() *NetworkWithServerVO`

NewNetworkWithServerVO instantiates a new NetworkWithServerVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithServerVOWithDefaults

`func NewNetworkWithServerVOWithDefaults() *NetworkWithServerVO`

NewNetworkWithServerVOWithDefaults instantiates a new NetworkWithServerVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkWithServerVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkWithServerVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkWithServerVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkWithServerVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkWithServerVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkWithServerVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkWithServerVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkWithServerVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerDevices

`func (o *NetworkWithServerVO) GetServerDevices() []BriefServerDeviceVO`

GetServerDevices returns the ServerDevices field if non-nil, zero value otherwise.

### GetServerDevicesOk

`func (o *NetworkWithServerVO) GetServerDevicesOk() (*[]BriefServerDeviceVO, bool)`

GetServerDevicesOk returns a tuple with the ServerDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDevices

`func (o *NetworkWithServerVO) SetServerDevices(v []BriefServerDeviceVO)`

SetServerDevices sets ServerDevices field to given value.

### HasServerDevices

`func (o *NetworkWithServerVO) HasServerDevices() bool`

HasServerDevices returns a boolean if a field has been set.

### GetType

`func (o *NetworkWithServerVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkWithServerVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkWithServerVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkWithServerVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


