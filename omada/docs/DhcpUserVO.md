# DhcpUserVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**IpLong** | Pointer to **int64** |  | [optional] 
**LeftLeaseTime** | Pointer to **string** | Left Lease Time | [optional] 
**MacAddress** | Pointer to **string** | Mac Address | [optional] 
**Model** | Pointer to **string** | device Model | [optional] 
**ModelVersion** | Pointer to **string** | device ModelVersion | [optional] 
**Name** | Pointer to **string** | Dhcp User Name | [optional] 
**NetId** | Pointer to **string** | Network ID | [optional] 
**NetName** | Pointer to **string** | Network Name | [optional] 
**ServerMac** | Pointer to **string** | Dhcp Server Device Mac | [optional] 
**ServerName** | Pointer to **string** | Dhcp Server Device Name | [optional] 
**ServerStackId** | Pointer to **string** | Dhcp Server Stack ID | [optional] 
**ShowingType** | Pointer to **string** | Client Type or Device Type | [optional] 
**Status** | Pointer to **int32** | Status should be a value as follows: 0: Dynamic Binding; 1: Static Binding | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 0: Device; 1: Client | [optional] 

## Methods

### NewDhcpUserVO

`func NewDhcpUserVO() *DhcpUserVO`

NewDhcpUserVO instantiates a new DhcpUserVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpUserVOWithDefaults

`func NewDhcpUserVOWithDefaults() *DhcpUserVO`

NewDhcpUserVOWithDefaults instantiates a new DhcpUserVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *DhcpUserVO) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DhcpUserVO) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DhcpUserVO) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DhcpUserVO) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpLong

`func (o *DhcpUserVO) GetIpLong() int64`

GetIpLong returns the IpLong field if non-nil, zero value otherwise.

### GetIpLongOk

`func (o *DhcpUserVO) GetIpLongOk() (*int64, bool)`

GetIpLongOk returns a tuple with the IpLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpLong

`func (o *DhcpUserVO) SetIpLong(v int64)`

SetIpLong sets IpLong field to given value.

### HasIpLong

`func (o *DhcpUserVO) HasIpLong() bool`

HasIpLong returns a boolean if a field has been set.

### GetLeftLeaseTime

`func (o *DhcpUserVO) GetLeftLeaseTime() string`

GetLeftLeaseTime returns the LeftLeaseTime field if non-nil, zero value otherwise.

### GetLeftLeaseTimeOk

`func (o *DhcpUserVO) GetLeftLeaseTimeOk() (*string, bool)`

GetLeftLeaseTimeOk returns a tuple with the LeftLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftLeaseTime

`func (o *DhcpUserVO) SetLeftLeaseTime(v string)`

SetLeftLeaseTime sets LeftLeaseTime field to given value.

### HasLeftLeaseTime

`func (o *DhcpUserVO) HasLeftLeaseTime() bool`

HasLeftLeaseTime returns a boolean if a field has been set.

### GetMacAddress

`func (o *DhcpUserVO) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DhcpUserVO) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DhcpUserVO) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DhcpUserVO) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetModel

`func (o *DhcpUserVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DhcpUserVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DhcpUserVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DhcpUserVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DhcpUserVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DhcpUserVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DhcpUserVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DhcpUserVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DhcpUserVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DhcpUserVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DhcpUserVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DhcpUserVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetId

`func (o *DhcpUserVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *DhcpUserVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *DhcpUserVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *DhcpUserVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNetName

`func (o *DhcpUserVO) GetNetName() string`

GetNetName returns the NetName field if non-nil, zero value otherwise.

### GetNetNameOk

`func (o *DhcpUserVO) GetNetNameOk() (*string, bool)`

GetNetNameOk returns a tuple with the NetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetName

`func (o *DhcpUserVO) SetNetName(v string)`

SetNetName sets NetName field to given value.

### HasNetName

`func (o *DhcpUserVO) HasNetName() bool`

HasNetName returns a boolean if a field has been set.

### GetServerMac

`func (o *DhcpUserVO) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *DhcpUserVO) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *DhcpUserVO) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *DhcpUserVO) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpUserVO) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpUserVO) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpUserVO) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpUserVO) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerStackId

`func (o *DhcpUserVO) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *DhcpUserVO) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *DhcpUserVO) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *DhcpUserVO) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetShowingType

`func (o *DhcpUserVO) GetShowingType() string`

GetShowingType returns the ShowingType field if non-nil, zero value otherwise.

### GetShowingTypeOk

`func (o *DhcpUserVO) GetShowingTypeOk() (*string, bool)`

GetShowingTypeOk returns a tuple with the ShowingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowingType

`func (o *DhcpUserVO) SetShowingType(v string)`

SetShowingType sets ShowingType field to given value.

### HasShowingType

`func (o *DhcpUserVO) HasShowingType() bool`

HasShowingType returns a boolean if a field has been set.

### GetStatus

`func (o *DhcpUserVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DhcpUserVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DhcpUserVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DhcpUserVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DhcpUserVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DhcpUserVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DhcpUserVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DhcpUserVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


