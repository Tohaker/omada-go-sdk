# GatewayInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | Pointer to **string** | device mac. | [optional] 
**SiteId** | Pointer to **string** | site id. | [optional] 
**Uptime** | Pointer to **string** | up time. | [optional] 
**WanStatus** | Pointer to [**[]GatewayWanStatus**](GatewayWanStatus.md) | wan status infos. | [optional] 

## Methods

### NewGatewayInfos

`func NewGatewayInfos() *GatewayInfos`

NewGatewayInfos instantiates a new GatewayInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayInfosWithDefaults

`func NewGatewayInfosWithDefaults() *GatewayInfos`

NewGatewayInfosWithDefaults instantiates a new GatewayInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *GatewayInfos) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *GatewayInfos) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *GatewayInfos) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *GatewayInfos) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetSiteId

`func (o *GatewayInfos) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GatewayInfos) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GatewayInfos) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GatewayInfos) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetUptime

`func (o *GatewayInfos) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GatewayInfos) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GatewayInfos) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GatewayInfos) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetWanStatus

`func (o *GatewayInfos) GetWanStatus() []GatewayWanStatus`

GetWanStatus returns the WanStatus field if non-nil, zero value otherwise.

### GetWanStatusOk

`func (o *GatewayInfos) GetWanStatusOk() (*[]GatewayWanStatus, bool)`

GetWanStatusOk returns a tuple with the WanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanStatus

`func (o *GatewayInfos) SetWanStatus(v []GatewayWanStatus)`

SetWanStatus sets WanStatus field to given value.

### HasWanStatus

`func (o *GatewayInfos) HasWanStatus() bool`

HasWanStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


