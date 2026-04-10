# ClientStatisticalDataDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Down** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**DownRate** | Pointer to **int64** | Downlink rate (Byte/s). | [optional] 
**Mac** | Pointer to **string** | Client MAC Address. | [optional] 
**RxR** | Pointer to **int64** | (Wireless) Uplink negotiation rate (bit/s). | [optional] 
**Signal** | Pointer to **int32** | (Wireless) Signal strength, unit: dBm. | [optional] 
**Time** | Pointer to **int64** | The statistical data collected timestamp, unit: second. | [optional] 
**TxFP** | Pointer to **int64** | Number of downstream failed packets. | [optional] 
**TxR** | Pointer to **int64** | (Wireless) Downlink negotiation rate (bit/s). | [optional] 
**Up** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 
**UpRate** | Pointer to **int64** | Uplink rate (Byte/s). | [optional] 
**Wireless** | Pointer to **bool** | true: Wireless client;  false: Not wireless client | [optional] 

## Methods

### NewClientStatisticalDataDetail

`func NewClientStatisticalDataDetail() *ClientStatisticalDataDetail`

NewClientStatisticalDataDetail instantiates a new ClientStatisticalDataDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStatisticalDataDetailWithDefaults

`func NewClientStatisticalDataDetailWithDefaults() *ClientStatisticalDataDetail`

NewClientStatisticalDataDetailWithDefaults instantiates a new ClientStatisticalDataDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDown

`func (o *ClientStatisticalDataDetail) GetDown() int64`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *ClientStatisticalDataDetail) GetDownOk() (*int64, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *ClientStatisticalDataDetail) SetDown(v int64)`

SetDown sets Down field to given value.

### HasDown

`func (o *ClientStatisticalDataDetail) HasDown() bool`

HasDown returns a boolean if a field has been set.

### GetDownRate

`func (o *ClientStatisticalDataDetail) GetDownRate() int64`

GetDownRate returns the DownRate field if non-nil, zero value otherwise.

### GetDownRateOk

`func (o *ClientStatisticalDataDetail) GetDownRateOk() (*int64, bool)`

GetDownRateOk returns a tuple with the DownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownRate

`func (o *ClientStatisticalDataDetail) SetDownRate(v int64)`

SetDownRate sets DownRate field to given value.

### HasDownRate

`func (o *ClientStatisticalDataDetail) HasDownRate() bool`

HasDownRate returns a boolean if a field has been set.

### GetMac

`func (o *ClientStatisticalDataDetail) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientStatisticalDataDetail) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientStatisticalDataDetail) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientStatisticalDataDetail) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetRxR

`func (o *ClientStatisticalDataDetail) GetRxR() int64`

GetRxR returns the RxR field if non-nil, zero value otherwise.

### GetRxROk

`func (o *ClientStatisticalDataDetail) GetRxROk() (*int64, bool)`

GetRxROk returns a tuple with the RxR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxR

`func (o *ClientStatisticalDataDetail) SetRxR(v int64)`

SetRxR sets RxR field to given value.

### HasRxR

`func (o *ClientStatisticalDataDetail) HasRxR() bool`

HasRxR returns a boolean if a field has been set.

### GetSignal

`func (o *ClientStatisticalDataDetail) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *ClientStatisticalDataDetail) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *ClientStatisticalDataDetail) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *ClientStatisticalDataDetail) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetTime

`func (o *ClientStatisticalDataDetail) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClientStatisticalDataDetail) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClientStatisticalDataDetail) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClientStatisticalDataDetail) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxFP

`func (o *ClientStatisticalDataDetail) GetTxFP() int64`

GetTxFP returns the TxFP field if non-nil, zero value otherwise.

### GetTxFPOk

`func (o *ClientStatisticalDataDetail) GetTxFPOk() (*int64, bool)`

GetTxFPOk returns a tuple with the TxFP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxFP

`func (o *ClientStatisticalDataDetail) SetTxFP(v int64)`

SetTxFP sets TxFP field to given value.

### HasTxFP

`func (o *ClientStatisticalDataDetail) HasTxFP() bool`

HasTxFP returns a boolean if a field has been set.

### GetTxR

`func (o *ClientStatisticalDataDetail) GetTxR() int64`

GetTxR returns the TxR field if non-nil, zero value otherwise.

### GetTxROk

`func (o *ClientStatisticalDataDetail) GetTxROk() (*int64, bool)`

GetTxROk returns a tuple with the TxR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxR

`func (o *ClientStatisticalDataDetail) SetTxR(v int64)`

SetTxR sets TxR field to given value.

### HasTxR

`func (o *ClientStatisticalDataDetail) HasTxR() bool`

HasTxR returns a boolean if a field has been set.

### GetUp

`func (o *ClientStatisticalDataDetail) GetUp() int64`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *ClientStatisticalDataDetail) GetUpOk() (*int64, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *ClientStatisticalDataDetail) SetUp(v int64)`

SetUp sets Up field to given value.

### HasUp

`func (o *ClientStatisticalDataDetail) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUpRate

`func (o *ClientStatisticalDataDetail) GetUpRate() int64`

GetUpRate returns the UpRate field if non-nil, zero value otherwise.

### GetUpRateOk

`func (o *ClientStatisticalDataDetail) GetUpRateOk() (*int64, bool)`

GetUpRateOk returns a tuple with the UpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRate

`func (o *ClientStatisticalDataDetail) SetUpRate(v int64)`

SetUpRate sets UpRate field to given value.

### HasUpRate

`func (o *ClientStatisticalDataDetail) HasUpRate() bool`

HasUpRate returns a boolean if a field has been set.

### GetWireless

`func (o *ClientStatisticalDataDetail) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *ClientStatisticalDataDetail) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *ClientStatisticalDataDetail) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *ClientStatisticalDataDetail) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


