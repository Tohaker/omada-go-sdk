# OsgDhcpUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**LeftLeaseTime** | Pointer to **string** | Left Lease Time | [optional] 
**MacAddress** | Pointer to **string** | Mac Address | [optional] 
**Model** | Pointer to **string** | device Model | [optional] 
**ModelVersion** | Pointer to **string** | device ModelVersion | [optional] 
**Name** | Pointer to **string** | Dhcp User Name | [optional] 
**NetId** | Pointer to **string** | Network ID | [optional] 
**NetName** | Pointer to **string** | Network Name | [optional] 
**ShowingType** | Pointer to **string** | Client Type or Device Type | [optional] 
**Status** | Pointer to **int32** | Status should be a value as follows: 0: Dynamic Binding; 1: Static Binding | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 0: Device; 1: Client | [optional] 

## Methods

### NewOsgDhcpUserVO

`func NewOsgDhcpUserVO() *OsgDhcpUserVO`

NewOsgDhcpUserVO instantiates a new OsgDhcpUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgDhcpUserVOWithDefaults

`func NewOsgDhcpUserVOWithDefaults() *OsgDhcpUserVO`

NewOsgDhcpUserVOWithDefaults instantiates a new OsgDhcpUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *OsgDhcpUserVO) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *OsgDhcpUserVO) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *OsgDhcpUserVO) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *OsgDhcpUserVO) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLeftLeaseTime

`func (o *OsgDhcpUserVO) GetLeftLeaseTime() string`

GetLeftLeaseTime returns the LeftLeaseTime field if non-nil, zero value otherwise.

### GetLeftLeaseTimeOk

`func (o *OsgDhcpUserVO) GetLeftLeaseTimeOk() (*string, bool)`

GetLeftLeaseTimeOk returns a tuple with the LeftLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftLeaseTime

`func (o *OsgDhcpUserVO) SetLeftLeaseTime(v string)`

SetLeftLeaseTime sets LeftLeaseTime field to given value.

### HasLeftLeaseTime

`func (o *OsgDhcpUserVO) HasLeftLeaseTime() bool`

HasLeftLeaseTime returns a boolean if a field has been set.

### GetMacAddress

`func (o *OsgDhcpUserVO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *OsgDhcpUserVO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *OsgDhcpUserVO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *OsgDhcpUserVO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetModel

`func (o *OsgDhcpUserVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OsgDhcpUserVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OsgDhcpUserVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OsgDhcpUserVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OsgDhcpUserVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OsgDhcpUserVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OsgDhcpUserVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OsgDhcpUserVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OsgDhcpUserVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsgDhcpUserVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsgDhcpUserVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsgDhcpUserVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetId

`func (o *OsgDhcpUserVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *OsgDhcpUserVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *OsgDhcpUserVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *OsgDhcpUserVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNetName

`func (o *OsgDhcpUserVO) GetNetName() string`

GetNetName returns the NetName field if non-nil, zero value otherwise.

### GetNetNameOk

`func (o *OsgDhcpUserVO) GetNetNameOk() (*string, bool)`

GetNetNameOk returns a tuple with the NetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetName

`func (o *OsgDhcpUserVO) SetNetName(v string)`

SetNetName sets NetName field to given value.

### HasNetName

`func (o *OsgDhcpUserVO) HasNetName() bool`

HasNetName returns a boolean if a field has been set.

### GetShowingType

`func (o *OsgDhcpUserVO) GetShowingType() string`

GetShowingType returns the ShowingType field if non-nil, zero value otherwise.

### GetShowingTypeOk

`func (o *OsgDhcpUserVO) GetShowingTypeOk() (*string, bool)`

GetShowingTypeOk returns a tuple with the ShowingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowingType

`func (o *OsgDhcpUserVO) SetShowingType(v string)`

SetShowingType sets ShowingType field to given value.

### HasShowingType

`func (o *OsgDhcpUserVO) HasShowingType() bool`

HasShowingType returns a boolean if a field has been set.

### GetStatus

`func (o *OsgDhcpUserVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsgDhcpUserVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsgDhcpUserVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsgDhcpUserVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OsgDhcpUserVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsgDhcpUserVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsgDhcpUserVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OsgDhcpUserVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


