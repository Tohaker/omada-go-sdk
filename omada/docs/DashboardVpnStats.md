# DashboardVpnStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site | [optional] 
**Name** | Pointer to **string** | VPN name | [optional] 
**RxData** | Pointer to **int64** | The total rx traffic of the tunnels included in the server/client, in Byte | [optional] 
**Status** | Pointer to **bool** | Whether VPN is enabled | [optional] 
**Tunnels** | Pointer to **int32** | The number of tunnels included in the VPN server | [optional] 
**TxData** | Pointer to **int64** | The total tx traffic of the channels included in the server/client, in Byte | [optional] 

## Methods

### NewDashboardVpnStats

`func NewDashboardVpnStats() *DashboardVpnStats`

NewDashboardVpnStats instantiates a new DashboardVpnStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardVpnStatsWithDefaults

`func NewDashboardVpnStatsWithDefaults() *DashboardVpnStats`

NewDashboardVpnStatsWithDefaults instantiates a new DashboardVpnStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *DashboardVpnStats) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DashboardVpnStats) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DashboardVpnStats) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *DashboardVpnStats) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetName

`func (o *DashboardVpnStats) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardVpnStats) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardVpnStats) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DashboardVpnStats) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRxData

`func (o *DashboardVpnStats) GetRxData() int64`

GetRxData returns the RxData field if non-nil, zero value otherwise.

### GetRxDataOk

`func (o *DashboardVpnStats) GetRxDataOk() (*int64, bool)`

GetRxDataOk returns a tuple with the RxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxData

`func (o *DashboardVpnStats) SetRxData(v int64)`

SetRxData sets RxData field to given value.

### HasRxData

`func (o *DashboardVpnStats) HasRxData() bool`

HasRxData returns a boolean if a field has been set.

### GetStatus

`func (o *DashboardVpnStats) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DashboardVpnStats) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DashboardVpnStats) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DashboardVpnStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTunnels

`func (o *DashboardVpnStats) GetTunnels() int32`

GetTunnels returns the Tunnels field if non-nil, zero value otherwise.

### GetTunnelsOk

`func (o *DashboardVpnStats) GetTunnelsOk() (*int32, bool)`

GetTunnelsOk returns a tuple with the Tunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnels

`func (o *DashboardVpnStats) SetTunnels(v int32)`

SetTunnels sets Tunnels field to given value.

### HasTunnels

`func (o *DashboardVpnStats) HasTunnels() bool`

HasTunnels returns a boolean if a field has been set.

### GetTxData

`func (o *DashboardVpnStats) GetTxData() int64`

GetTxData returns the TxData field if non-nil, zero value otherwise.

### GetTxDataOk

`func (o *DashboardVpnStats) GetTxDataOk() (*int64, bool)`

GetTxDataOk returns a tuple with the TxData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxData

`func (o *DashboardVpnStats) SetTxData(v int64)`

SetTxData sets TxData field to given value.

### HasTxData

`func (o *DashboardVpnStats) HasTxData() bool`

HasTxData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


