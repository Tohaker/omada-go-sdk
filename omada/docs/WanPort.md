# WanPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex | [optional] 
**HealthLevel** | Pointer to **int32** | Health Level | [optional] 
**InternetState** | Pointer to **int32** | Port Internet State | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed | [optional] 
**Name** | Pointer to **string** | Port Name | [optional] 
**OnlineDetection** | Pointer to **int32** | Port Online Detection | [optional] 
**Port** | Pointer to **string** | Port Id | [optional] 
**Status** | Pointer to **int32** | Port Status | [optional] 
**TxRate** | Pointer to **int64** | Rx Rate | [optional] 
**Wireless** | Pointer to **bool** | Whether the port is wireless. | [optional] 

## Methods

### NewWanPort

`func NewWanPort() *WanPort`

NewWanPort instantiates a new WanPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanPortWithDefaults

`func NewWanPortWithDefaults() *WanPort`

NewWanPortWithDefaults instantiates a new WanPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *WanPort) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *WanPort) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *WanPort) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *WanPort) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetHealthLevel

`func (o *WanPort) GetHealthLevel() int32`

GetHealthLevel returns the HealthLevel field if non-nil, zero value otherwise.

### GetHealthLevelOk

`func (o *WanPort) GetHealthLevelOk() (*int32, bool)`

GetHealthLevelOk returns a tuple with the HealthLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthLevel

`func (o *WanPort) SetHealthLevel(v int32)`

SetHealthLevel sets HealthLevel field to given value.

### HasHealthLevel

`func (o *WanPort) HasHealthLevel() bool`

HasHealthLevel returns a boolean if a field has been set.

### GetInternetState

`func (o *WanPort) GetInternetState() int32`

GetInternetState returns the InternetState field if non-nil, zero value otherwise.

### GetInternetStateOk

`func (o *WanPort) GetInternetStateOk() (*int32, bool)`

GetInternetStateOk returns a tuple with the InternetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetState

`func (o *WanPort) SetInternetState(v int32)`

SetInternetState sets InternetState field to given value.

### HasInternetState

`func (o *WanPort) HasInternetState() bool`

HasInternetState returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *WanPort) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *WanPort) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *WanPort) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *WanPort) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetName

`func (o *WanPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WanPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WanPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WanPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnlineDetection

`func (o *WanPort) GetOnlineDetection() int32`

GetOnlineDetection returns the OnlineDetection field if non-nil, zero value otherwise.

### GetOnlineDetectionOk

`func (o *WanPort) GetOnlineDetectionOk() (*int32, bool)`

GetOnlineDetectionOk returns a tuple with the OnlineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDetection

`func (o *WanPort) SetOnlineDetection(v int32)`

SetOnlineDetection sets OnlineDetection field to given value.

### HasOnlineDetection

`func (o *WanPort) HasOnlineDetection() bool`

HasOnlineDetection returns a boolean if a field has been set.

### GetPort

`func (o *WanPort) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WanPort) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WanPort) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *WanPort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *WanPort) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WanPort) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WanPort) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WanPort) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTxRate

`func (o *WanPort) GetTxRate() int64`

GetTxRate returns the TxRate field if non-nil, zero value otherwise.

### GetTxRateOk

`func (o *WanPort) GetTxRateOk() (*int64, bool)`

GetTxRateOk returns a tuple with the TxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRate

`func (o *WanPort) SetTxRate(v int64)`

SetTxRate sets TxRate field to given value.

### HasTxRate

`func (o *WanPort) HasTxRate() bool`

HasTxRate returns a boolean if a field has been set.

### GetWireless

`func (o *WanPort) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *WanPort) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *WanPort) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *WanPort) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


