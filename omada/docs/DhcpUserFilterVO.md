# DhcpUserFilterVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetId** | Pointer to **string** | Lan Network IDs | [optional] 
**SearchKey** | Pointer to **string** | Search Key | [optional] 
**SelectIps** | **[]string** | Selected Ips | 
**SelectMacs** | [**BatchSelectMacsVO**](BatchSelectMacsVO.md) |  | 
**ServerMac** | Pointer to **string** | Dhcp Server Macs | [optional] 
**ServerStackId** | Pointer to **string** | Dhcp Server StackIds | [optional] 
**Type** | Pointer to **string** | Filter Type of Dhcp User: \&quot;device\&quot;, \&quot;client\&quot; or \&quot;device, client\&quot; | [optional] 

## Methods

### NewDhcpUserFilterVO

`func NewDhcpUserFilterVO(selectIps []string, selectMacs BatchSelectMacsVO, ) *DhcpUserFilterVO`

NewDhcpUserFilterVO instantiates a new DhcpUserFilterVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpUserFilterVOWithDefaults

`func NewDhcpUserFilterVOWithDefaults() *DhcpUserFilterVO`

NewDhcpUserFilterVOWithDefaults instantiates a new DhcpUserFilterVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetId

`func (o *DhcpUserFilterVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *DhcpUserFilterVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *DhcpUserFilterVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *DhcpUserFilterVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetSearchKey

`func (o *DhcpUserFilterVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *DhcpUserFilterVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *DhcpUserFilterVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *DhcpUserFilterVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectIps

`func (o *DhcpUserFilterVO) GetSelectIps() []string`

GetSelectIps returns the SelectIps field if non-nil, zero value otherwise.

### GetSelectIpsOk

`func (o *DhcpUserFilterVO) GetSelectIpsOk() (*[]string, bool)`

GetSelectIpsOk returns a tuple with the SelectIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectIps

`func (o *DhcpUserFilterVO) SetSelectIps(v []string)`

SetSelectIps sets SelectIps field to given value.


### GetSelectMacs

`func (o *DhcpUserFilterVO) GetSelectMacs() BatchSelectMacsVO`

GetSelectMacs returns the SelectMacs field if non-nil, zero value otherwise.

### GetSelectMacsOk

`func (o *DhcpUserFilterVO) GetSelectMacsOk() (*BatchSelectMacsVO, bool)`

GetSelectMacsOk returns a tuple with the SelectMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectMacs

`func (o *DhcpUserFilterVO) SetSelectMacs(v BatchSelectMacsVO)`

SetSelectMacs sets SelectMacs field to given value.


### GetServerMac

`func (o *DhcpUserFilterVO) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *DhcpUserFilterVO) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *DhcpUserFilterVO) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *DhcpUserFilterVO) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *DhcpUserFilterVO) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *DhcpUserFilterVO) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *DhcpUserFilterVO) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *DhcpUserFilterVO) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetType

`func (o *DhcpUserFilterVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DhcpUserFilterVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DhcpUserFilterVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DhcpUserFilterVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


