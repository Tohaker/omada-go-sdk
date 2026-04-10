# ClientNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthStatus** | Pointer to **int32** | Authentication status should be a value as follows: &lt;br/&gt;0: CONNECTED // Access without any authentication method; &lt;br/&gt;1: PENDING // Access to Portal, but authentication failed; &lt;br/&gt;2: AUTHORIZED // Pass through portal, pass other authentication without portal; &lt;br/&gt;3: AUTH-FREE // No portal authentication required. | [optional] 
**ClientType** | Pointer to **string** | Client Type: iphone, ipod, android, pc, printer, tv... | [optional] 
**DevRxRate** | Pointer to **int64** | Client real-time downloadRate | [optional] 
**DevTxRate** | Pointer to **int64** | Client real-time uploadRate | [optional] 
**Guest** | Pointer to **bool** | (Wireless) Whether it is Guest (used to display the wireless Guest client icon). | [optional] 
**HealthScore** | Pointer to **int32** | Health Score, 1~3: poor; 4~7: fair; 0: no data; 8~10 good. | [optional] 
**Ip** | Pointer to **string** | Client IP. | [optional] 
**Mac** | Pointer to **string** | Client MAC Address. | [optional] 
**Manager** | Pointer to **bool** | Whether it is the device currently accessing the Controller itself. | [optional] 
**Model** | Pointer to **string** | Model of client device. | [optional] 
**Name** | Pointer to **string** | Client name. | [optional] 
**UpApInfo** | Pointer to [**UplinkAPInfo**](UplinkAPInfo.md) |  | [optional] 
**UpDeviceType** | Pointer to **int32** | Uplink device type, 0: AP; 1: Switch; 2: Gateway. | [optional] 
**UpOswInfo** | Pointer to [**UplinkSwitchInfo**](UplinkSwitchInfo.md) |  | [optional] 
**Wireless** | Pointer to **bool** | true: Wireless client;  false: Not wireless client | [optional] 

## Methods

### NewClientNodeInfo

`func NewClientNodeInfo() *ClientNodeInfo`

NewClientNodeInfo instantiates a new ClientNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientNodeInfoWithDefaults

`func NewClientNodeInfoWithDefaults() *ClientNodeInfo`

NewClientNodeInfoWithDefaults instantiates a new ClientNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthStatus

`func (o *ClientNodeInfo) GetAuthStatus() int32`

GetAuthStatus returns the AuthStatus field if non-nil, zero value otherwise.

### GetAuthStatusOk

`func (o *ClientNodeInfo) GetAuthStatusOk() (*int32, bool)`

GetAuthStatusOk returns a tuple with the AuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthStatus

`func (o *ClientNodeInfo) SetAuthStatus(v int32)`

SetAuthStatus sets AuthStatus field to given value.

### HasAuthStatus

`func (o *ClientNodeInfo) HasAuthStatus() bool`

HasAuthStatus returns a boolean if a field has been set.

### GetClientType

`func (o *ClientNodeInfo) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *ClientNodeInfo) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *ClientNodeInfo) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *ClientNodeInfo) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetDevRxRate

`func (o *ClientNodeInfo) GetDevRxRate() int64`

GetDevRxRate returns the DevRxRate field if non-nil, zero value otherwise.

### GetDevRxRateOk

`func (o *ClientNodeInfo) GetDevRxRateOk() (*int64, bool)`

GetDevRxRateOk returns a tuple with the DevRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevRxRate

`func (o *ClientNodeInfo) SetDevRxRate(v int64)`

SetDevRxRate sets DevRxRate field to given value.

### HasDevRxRate

`func (o *ClientNodeInfo) HasDevRxRate() bool`

HasDevRxRate returns a boolean if a field has been set.

### GetDevTxRate

`func (o *ClientNodeInfo) GetDevTxRate() int64`

GetDevTxRate returns the DevTxRate field if non-nil, zero value otherwise.

### GetDevTxRateOk

`func (o *ClientNodeInfo) GetDevTxRateOk() (*int64, bool)`

GetDevTxRateOk returns a tuple with the DevTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTxRate

`func (o *ClientNodeInfo) SetDevTxRate(v int64)`

SetDevTxRate sets DevTxRate field to given value.

### HasDevTxRate

`func (o *ClientNodeInfo) HasDevTxRate() bool`

HasDevTxRate returns a boolean if a field has been set.

### GetGuest

`func (o *ClientNodeInfo) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ClientNodeInfo) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ClientNodeInfo) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ClientNodeInfo) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHealthScore

`func (o *ClientNodeInfo) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *ClientNodeInfo) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *ClientNodeInfo) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *ClientNodeInfo) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetIp

`func (o *ClientNodeInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientNodeInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientNodeInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientNodeInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ClientNodeInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientNodeInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientNodeInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientNodeInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *ClientNodeInfo) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ClientNodeInfo) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ClientNodeInfo) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ClientNodeInfo) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetModel

`func (o *ClientNodeInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClientNodeInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClientNodeInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClientNodeInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *ClientNodeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientNodeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientNodeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientNodeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpApInfo

`func (o *ClientNodeInfo) GetUpApInfo() UplinkAPInfo`

GetUpApInfo returns the UpApInfo field if non-nil, zero value otherwise.

### GetUpApInfoOk

`func (o *ClientNodeInfo) GetUpApInfoOk() (*UplinkAPInfo, bool)`

GetUpApInfoOk returns a tuple with the UpApInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpApInfo

`func (o *ClientNodeInfo) SetUpApInfo(v UplinkAPInfo)`

SetUpApInfo sets UpApInfo field to given value.

### HasUpApInfo

`func (o *ClientNodeInfo) HasUpApInfo() bool`

HasUpApInfo returns a boolean if a field has been set.

### GetUpDeviceType

`func (o *ClientNodeInfo) GetUpDeviceType() int32`

GetUpDeviceType returns the UpDeviceType field if non-nil, zero value otherwise.

### GetUpDeviceTypeOk

`func (o *ClientNodeInfo) GetUpDeviceTypeOk() (*int32, bool)`

GetUpDeviceTypeOk returns a tuple with the UpDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpDeviceType

`func (o *ClientNodeInfo) SetUpDeviceType(v int32)`

SetUpDeviceType sets UpDeviceType field to given value.

### HasUpDeviceType

`func (o *ClientNodeInfo) HasUpDeviceType() bool`

HasUpDeviceType returns a boolean if a field has been set.

### GetUpOswInfo

`func (o *ClientNodeInfo) GetUpOswInfo() UplinkSwitchInfo`

GetUpOswInfo returns the UpOswInfo field if non-nil, zero value otherwise.

### GetUpOswInfoOk

`func (o *ClientNodeInfo) GetUpOswInfoOk() (*UplinkSwitchInfo, bool)`

GetUpOswInfoOk returns a tuple with the UpOswInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpOswInfo

`func (o *ClientNodeInfo) SetUpOswInfo(v UplinkSwitchInfo)`

SetUpOswInfo sets UpOswInfo field to given value.

### HasUpOswInfo

`func (o *ClientNodeInfo) HasUpOswInfo() bool`

HasUpOswInfo returns a boolean if a field has been set.

### GetWireless

`func (o *ClientNodeInfo) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *ClientNodeInfo) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *ClientNodeInfo) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *ClientNodeInfo) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


