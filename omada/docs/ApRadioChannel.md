# ApRadioChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualChannel** | Pointer to **string** | Actual channel of the device | [optional] 
**AiRoamingOffset** | Pointer to **int32** | AI Roaming offset of the device | [optional] 
**BandWidth** | Pointer to **string** | BandWidth of the device | [optional] 
**BusyUtil** | Pointer to **int32** | BusyUtil of the device(Support by MTK device), value range [0, 100]. | [optional] 
**InterUtil** | Pointer to **int32** | InterUtil of the device, value range [0, 100]. | [optional] 
**MaxTxRate** | Pointer to **int32** | TxRate of the device | [optional] 
**RdMode** | Pointer to **string** | RdMode of the device | [optional] 
**Region** | Pointer to **int32** | Region code of the device | [optional] 
**RxUtil** | Pointer to **int32** | RxUtil of the device, value range [0, 100]. | [optional] 
**TxPower** | Pointer to **int32** | TxPower of the device | [optional] 
**TxUtil** | Pointer to **int32** | TxUtil of the device, value range [0, 100]. | [optional] 

## Methods

### NewApRadioChannel

`func NewApRadioChannel() *ApRadioChannel`

NewApRadioChannel instantiates a new ApRadioChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRadioChannelWithDefaults

`func NewApRadioChannelWithDefaults() *ApRadioChannel`

NewApRadioChannelWithDefaults instantiates a new ApRadioChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualChannel

`func (o *ApRadioChannel) GetActualChannel() string`

GetActualChannel returns the ActualChannel field if non-nil, zero value otherwise.

### GetActualChannelOk

`func (o *ApRadioChannel) GetActualChannelOk() (*string, bool)`

GetActualChannelOk returns a tuple with the ActualChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualChannel

`func (o *ApRadioChannel) SetActualChannel(v string)`

SetActualChannel sets ActualChannel field to given value.

### HasActualChannel

`func (o *ApRadioChannel) HasActualChannel() bool`

HasActualChannel returns a boolean if a field has been set.

### GetAiRoamingOffset

`func (o *ApRadioChannel) GetAiRoamingOffset() int32`

GetAiRoamingOffset returns the AiRoamingOffset field if non-nil, zero value otherwise.

### GetAiRoamingOffsetOk

`func (o *ApRadioChannel) GetAiRoamingOffsetOk() (*int32, bool)`

GetAiRoamingOffsetOk returns a tuple with the AiRoamingOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiRoamingOffset

`func (o *ApRadioChannel) SetAiRoamingOffset(v int32)`

SetAiRoamingOffset sets AiRoamingOffset field to given value.

### HasAiRoamingOffset

`func (o *ApRadioChannel) HasAiRoamingOffset() bool`

HasAiRoamingOffset returns a boolean if a field has been set.

### GetBandWidth

`func (o *ApRadioChannel) GetBandWidth() string`

GetBandWidth returns the BandWidth field if non-nil, zero value otherwise.

### GetBandWidthOk

`func (o *ApRadioChannel) GetBandWidthOk() (*string, bool)`

GetBandWidthOk returns a tuple with the BandWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidth

`func (o *ApRadioChannel) SetBandWidth(v string)`

SetBandWidth sets BandWidth field to given value.

### HasBandWidth

`func (o *ApRadioChannel) HasBandWidth() bool`

HasBandWidth returns a boolean if a field has been set.

### GetBusyUtil

`func (o *ApRadioChannel) GetBusyUtil() int32`

GetBusyUtil returns the BusyUtil field if non-nil, zero value otherwise.

### GetBusyUtilOk

`func (o *ApRadioChannel) GetBusyUtilOk() (*int32, bool)`

GetBusyUtilOk returns a tuple with the BusyUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyUtil

`func (o *ApRadioChannel) SetBusyUtil(v int32)`

SetBusyUtil sets BusyUtil field to given value.

### HasBusyUtil

`func (o *ApRadioChannel) HasBusyUtil() bool`

HasBusyUtil returns a boolean if a field has been set.

### GetInterUtil

`func (o *ApRadioChannel) GetInterUtil() int32`

GetInterUtil returns the InterUtil field if non-nil, zero value otherwise.

### GetInterUtilOk

`func (o *ApRadioChannel) GetInterUtilOk() (*int32, bool)`

GetInterUtilOk returns a tuple with the InterUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterUtil

`func (o *ApRadioChannel) SetInterUtil(v int32)`

SetInterUtil sets InterUtil field to given value.

### HasInterUtil

`func (o *ApRadioChannel) HasInterUtil() bool`

HasInterUtil returns a boolean if a field has been set.

### GetMaxTxRate

`func (o *ApRadioChannel) GetMaxTxRate() int32`

GetMaxTxRate returns the MaxTxRate field if non-nil, zero value otherwise.

### GetMaxTxRateOk

`func (o *ApRadioChannel) GetMaxTxRateOk() (*int32, bool)`

GetMaxTxRateOk returns a tuple with the MaxTxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTxRate

`func (o *ApRadioChannel) SetMaxTxRate(v int32)`

SetMaxTxRate sets MaxTxRate field to given value.

### HasMaxTxRate

`func (o *ApRadioChannel) HasMaxTxRate() bool`

HasMaxTxRate returns a boolean if a field has been set.

### GetRdMode

`func (o *ApRadioChannel) GetRdMode() string`

GetRdMode returns the RdMode field if non-nil, zero value otherwise.

### GetRdModeOk

`func (o *ApRadioChannel) GetRdModeOk() (*string, bool)`

GetRdModeOk returns a tuple with the RdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode

`func (o *ApRadioChannel) SetRdMode(v string)`

SetRdMode sets RdMode field to given value.

### HasRdMode

`func (o *ApRadioChannel) HasRdMode() bool`

HasRdMode returns a boolean if a field has been set.

### GetRegion

`func (o *ApRadioChannel) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApRadioChannel) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApRadioChannel) SetRegion(v int32)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApRadioChannel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRxUtil

`func (o *ApRadioChannel) GetRxUtil() int32`

GetRxUtil returns the RxUtil field if non-nil, zero value otherwise.

### GetRxUtilOk

`func (o *ApRadioChannel) GetRxUtilOk() (*int32, bool)`

GetRxUtilOk returns a tuple with the RxUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxUtil

`func (o *ApRadioChannel) SetRxUtil(v int32)`

SetRxUtil sets RxUtil field to given value.

### HasRxUtil

`func (o *ApRadioChannel) HasRxUtil() bool`

HasRxUtil returns a boolean if a field has been set.

### GetTxPower

`func (o *ApRadioChannel) GetTxPower() int32`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *ApRadioChannel) GetTxPowerOk() (*int32, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *ApRadioChannel) SetTxPower(v int32)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *ApRadioChannel) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### GetTxUtil

`func (o *ApRadioChannel) GetTxUtil() int32`

GetTxUtil returns the TxUtil field if non-nil, zero value otherwise.

### GetTxUtilOk

`func (o *ApRadioChannel) GetTxUtilOk() (*int32, bool)`

GetTxUtilOk returns a tuple with the TxUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxUtil

`func (o *ApRadioChannel) SetTxUtil(v int32)`

SetTxUtil sets TxUtil field to given value.

### HasTxUtil

`func (o *ApRadioChannel) HasTxUtil() bool`

HasTxUtil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


