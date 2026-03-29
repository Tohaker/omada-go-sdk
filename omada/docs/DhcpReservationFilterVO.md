# DhcpReservationFilterVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetId** | Pointer to **string** | Lan Network IDs | [optional] 
**SearchKey** | Pointer to **string** | Search Key | [optional] 
**SelectMacs** | [**BatchSelectMacsVO**](BatchSelectMacsVO.md) |  | 
**SortMap** | Pointer to **map[string]string** | Sort Direction. The keys that can be sorted are:mac, ip, net_name, description, status, clientName, name, value is asc or desc | [optional] 
**Type** | Pointer to **string** | Device types for which IP is reserved | [optional] 

## Methods

### NewDhcpReservationFilterVO

`func NewDhcpReservationFilterVO(selectMacs BatchSelectMacsVO, ) *DhcpReservationFilterVO`

NewDhcpReservationFilterVO instantiates a new DhcpReservationFilterVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpReservationFilterVOWithDefaults

`func NewDhcpReservationFilterVOWithDefaults() *DhcpReservationFilterVO`

NewDhcpReservationFilterVOWithDefaults instantiates a new DhcpReservationFilterVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetId

`func (o *DhcpReservationFilterVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *DhcpReservationFilterVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *DhcpReservationFilterVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *DhcpReservationFilterVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetSearchKey

`func (o *DhcpReservationFilterVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *DhcpReservationFilterVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *DhcpReservationFilterVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *DhcpReservationFilterVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectMacs

`func (o *DhcpReservationFilterVO) GetSelectMacs() BatchSelectMacsVO`

GetSelectMacs returns the SelectMacs field if non-nil, zero value otherwise.

### GetSelectMacsOk

`func (o *DhcpReservationFilterVO) GetSelectMacsOk() (*BatchSelectMacsVO, bool)`

GetSelectMacsOk returns a tuple with the SelectMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectMacs

`func (o *DhcpReservationFilterVO) SetSelectMacs(v BatchSelectMacsVO)`

SetSelectMacs sets SelectMacs field to given value.


### GetSortMap

`func (o *DhcpReservationFilterVO) GetSortMap() map[string]string`

GetSortMap returns the SortMap field if non-nil, zero value otherwise.

### GetSortMapOk

`func (o *DhcpReservationFilterVO) GetSortMapOk() (*map[string]string, bool)`

GetSortMapOk returns a tuple with the SortMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortMap

`func (o *DhcpReservationFilterVO) SetSortMap(v map[string]string)`

SetSortMap sets SortMap field to given value.

### HasSortMap

`func (o *DhcpReservationFilterVO) HasSortMap() bool`

HasSortMap returns a boolean if a field has been set.

### GetType

`func (o *DhcpReservationFilterVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DhcpReservationFilterVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DhcpReservationFilterVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DhcpReservationFilterVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


