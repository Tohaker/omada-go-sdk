# WifiCallingTrafficOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **string** | Ap MAC. | [optional] 
**ApName** | Pointer to **string** | The name of the AP device connected to the client. | [optional] 
**Band** | Pointer to **int32** | SSID band. The lowest bit indicates whether 2.4G is included; the second lowest bit indicates whether 5G is included; the third lowest bit indicates whether 6G is included; 1 means included while 0 means not included. For example, 7(111) means that 2G/5G/6G are enabled; 1(001) means that 2G is enabled. (When 5G is included，it means 5G/5G1/5G2 are enabled.) | [optional] 
**CallNum** | Pointer to **int32** | Number of calls on the same client MAC. | [optional] 
**CarrierName** | Pointer to **string** | Carrier name used by the client. | [optional] 
**ClientMac** | Pointer to **string** | Client Mac Address | [optional] 
**ClientName** | Pointer to **string** | Client name. | [optional] 
**DeviceType** | Pointer to **string** | Client device type. | [optional] 
**Domain** | Pointer to **string** | EPDG domain name or IP. | [optional] 
**EndTime** | Pointer to **int64** | End time of the voice call. | [optional] 
**Ip** | Pointer to **string** | Client ip or ipv6. | [optional] 
**Model** | Pointer to **string** | Client model. | [optional] 
**Priority** | Pointer to **int32** | Priority of ePDG in the WiFi calling profile. | [optional] 
**Ssid** | Pointer to **string** | The name of the SSID that the client is connected to. | [optional] 
**StartTime** | Pointer to **int64** | Start time of the voice call. | [optional] 
**TotalTraffic** | Pointer to **int64** | Total traffic. | [optional] 
**TrafficDown** | Pointer to **int64** | Downlink traffic related to voice calls. | [optional] 
**TrafficUp** | Pointer to **int64** | Uplink traffic related to voice calls. | [optional] 
**WifiCallingProfileName** | Pointer to **string** | Wifi calling profile name used by the client. | [optional] 

## Methods

### NewWifiCallingTrafficOpenApiVO

`func NewWifiCallingTrafficOpenApiVO() *WifiCallingTrafficOpenApiVO`

NewWifiCallingTrafficOpenApiVO instantiates a new WifiCallingTrafficOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWifiCallingTrafficOpenApiVOWithDefaults

`func NewWifiCallingTrafficOpenApiVOWithDefaults() *WifiCallingTrafficOpenApiVO`

NewWifiCallingTrafficOpenApiVOWithDefaults instantiates a new WifiCallingTrafficOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *WifiCallingTrafficOpenApiVO) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *WifiCallingTrafficOpenApiVO) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *WifiCallingTrafficOpenApiVO) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *WifiCallingTrafficOpenApiVO) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetApName

`func (o *WifiCallingTrafficOpenApiVO) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *WifiCallingTrafficOpenApiVO) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *WifiCallingTrafficOpenApiVO) SetApName(v string)`

SetApName sets ApName field to given value.

### HasApName

`func (o *WifiCallingTrafficOpenApiVO) HasApName() bool`

HasApName returns a boolean if a field has been set.

### GetBand

`func (o *WifiCallingTrafficOpenApiVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *WifiCallingTrafficOpenApiVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *WifiCallingTrafficOpenApiVO) SetBand(v int32)`

SetBand sets Band field to given value.

### HasBand

`func (o *WifiCallingTrafficOpenApiVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetCallNum

`func (o *WifiCallingTrafficOpenApiVO) GetCallNum() int32`

GetCallNum returns the CallNum field if non-nil, zero value otherwise.

### GetCallNumOk

`func (o *WifiCallingTrafficOpenApiVO) GetCallNumOk() (*int32, bool)`

GetCallNumOk returns a tuple with the CallNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallNum

`func (o *WifiCallingTrafficOpenApiVO) SetCallNum(v int32)`

SetCallNum sets CallNum field to given value.

### HasCallNum

`func (o *WifiCallingTrafficOpenApiVO) HasCallNum() bool`

HasCallNum returns a boolean if a field has been set.

### GetCarrierName

`func (o *WifiCallingTrafficOpenApiVO) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *WifiCallingTrafficOpenApiVO) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *WifiCallingTrafficOpenApiVO) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *WifiCallingTrafficOpenApiVO) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetClientMac

`func (o *WifiCallingTrafficOpenApiVO) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *WifiCallingTrafficOpenApiVO) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *WifiCallingTrafficOpenApiVO) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *WifiCallingTrafficOpenApiVO) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetClientName

`func (o *WifiCallingTrafficOpenApiVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *WifiCallingTrafficOpenApiVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *WifiCallingTrafficOpenApiVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *WifiCallingTrafficOpenApiVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDeviceType

`func (o *WifiCallingTrafficOpenApiVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WifiCallingTrafficOpenApiVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WifiCallingTrafficOpenApiVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WifiCallingTrafficOpenApiVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDomain

`func (o *WifiCallingTrafficOpenApiVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *WifiCallingTrafficOpenApiVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *WifiCallingTrafficOpenApiVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *WifiCallingTrafficOpenApiVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEndTime

`func (o *WifiCallingTrafficOpenApiVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WifiCallingTrafficOpenApiVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WifiCallingTrafficOpenApiVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WifiCallingTrafficOpenApiVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIp

`func (o *WifiCallingTrafficOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *WifiCallingTrafficOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *WifiCallingTrafficOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *WifiCallingTrafficOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetModel

`func (o *WifiCallingTrafficOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *WifiCallingTrafficOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *WifiCallingTrafficOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *WifiCallingTrafficOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPriority

`func (o *WifiCallingTrafficOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WifiCallingTrafficOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WifiCallingTrafficOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WifiCallingTrafficOpenApiVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSsid

`func (o *WifiCallingTrafficOpenApiVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WifiCallingTrafficOpenApiVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WifiCallingTrafficOpenApiVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *WifiCallingTrafficOpenApiVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStartTime

`func (o *WifiCallingTrafficOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WifiCallingTrafficOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WifiCallingTrafficOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WifiCallingTrafficOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *WifiCallingTrafficOpenApiVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *WifiCallingTrafficOpenApiVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *WifiCallingTrafficOpenApiVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *WifiCallingTrafficOpenApiVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetTrafficDown

`func (o *WifiCallingTrafficOpenApiVO) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *WifiCallingTrafficOpenApiVO) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *WifiCallingTrafficOpenApiVO) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *WifiCallingTrafficOpenApiVO) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *WifiCallingTrafficOpenApiVO) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *WifiCallingTrafficOpenApiVO) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *WifiCallingTrafficOpenApiVO) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *WifiCallingTrafficOpenApiVO) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.

### GetWifiCallingProfileName

`func (o *WifiCallingTrafficOpenApiVO) GetWifiCallingProfileName() string`

GetWifiCallingProfileName returns the WifiCallingProfileName field if non-nil, zero value otherwise.

### GetWifiCallingProfileNameOk

`func (o *WifiCallingTrafficOpenApiVO) GetWifiCallingProfileNameOk() (*string, bool)`

GetWifiCallingProfileNameOk returns a tuple with the WifiCallingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiCallingProfileName

`func (o *WifiCallingTrafficOpenApiVO) SetWifiCallingProfileName(v string)`

SetWifiCallingProfileName sets WifiCallingProfileName field to given value.

### HasWifiCallingProfileName

`func (o *WifiCallingTrafficOpenApiVO) HasWifiCallingProfileName() bool`

HasWifiCallingProfileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


