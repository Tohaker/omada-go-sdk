# BatchWifiInterferenceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandWidth** | Pointer to **int32** | Band Width | [optional] 
**Beacon** | Pointer to **int32** | beacon Interval | [optional] 
**Bssid** | Pointer to **string** | bssid | [optional] 
**Channel** | Pointer to **int32** | channel | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**LastSeen** | Pointer to **int64** | Last detect time. | [optional] 
**Mode** | Pointer to **int32** | phyMode | [optional] 
**NearestAp** | Pointer to **string** | MAC of the AP device closest to that wifi interference. | [optional] 
**RadioId** | Pointer to **int32** | radio id | [optional] 
**Security** | Pointer to **int32** | security type. Security should be a value as follows:0: None;1: WEP;2: WPA;3: WPA2;4: WPA/WPA2;5: WPE3-OWE;6: WPA3-EnterPrise;7: WPA2-PSK/WPA3-SAE;8: WPA/WPA2-PSK;9: WPA/WPA2-EnterPrise. | [optional] 
**Signal** | Pointer to **int32** | Signal strength | [optional] 
**Ssid** | Pointer to **string** | ssid | [optional] 

## Methods

### NewBatchWifiInterferenceOpenApiVO

`func NewBatchWifiInterferenceOpenApiVO() *BatchWifiInterferenceOpenApiVO`

NewBatchWifiInterferenceOpenApiVO instantiates a new BatchWifiInterferenceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWifiInterferenceOpenApiVOWithDefaults

`func NewBatchWifiInterferenceOpenApiVOWithDefaults() *BatchWifiInterferenceOpenApiVO`

NewBatchWifiInterferenceOpenApiVOWithDefaults instantiates a new BatchWifiInterferenceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandWidth

`func (o *BatchWifiInterferenceOpenApiVO) GetBandWidth() int32`

GetBandWidth returns the BandWidth field if non-nil, zero value otherwise.

### GetBandWidthOk

`func (o *BatchWifiInterferenceOpenApiVO) GetBandWidthOk() (*int32, bool)`

GetBandWidthOk returns a tuple with the BandWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidth

`func (o *BatchWifiInterferenceOpenApiVO) SetBandWidth(v int32)`

SetBandWidth sets BandWidth field to given value.

### HasBandWidth

`func (o *BatchWifiInterferenceOpenApiVO) HasBandWidth() bool`

HasBandWidth returns a boolean if a field has been set.

### GetBeacon

`func (o *BatchWifiInterferenceOpenApiVO) GetBeacon() int32`

GetBeacon returns the Beacon field if non-nil, zero value otherwise.

### GetBeaconOk

`func (o *BatchWifiInterferenceOpenApiVO) GetBeaconOk() (*int32, bool)`

GetBeaconOk returns a tuple with the Beacon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeacon

`func (o *BatchWifiInterferenceOpenApiVO) SetBeacon(v int32)`

SetBeacon sets Beacon field to given value.

### HasBeacon

`func (o *BatchWifiInterferenceOpenApiVO) HasBeacon() bool`

HasBeacon returns a boolean if a field has been set.

### GetBssid

`func (o *BatchWifiInterferenceOpenApiVO) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *BatchWifiInterferenceOpenApiVO) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *BatchWifiInterferenceOpenApiVO) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *BatchWifiInterferenceOpenApiVO) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetChannel

`func (o *BatchWifiInterferenceOpenApiVO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *BatchWifiInterferenceOpenApiVO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *BatchWifiInterferenceOpenApiVO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *BatchWifiInterferenceOpenApiVO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetId

`func (o *BatchWifiInterferenceOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchWifiInterferenceOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchWifiInterferenceOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BatchWifiInterferenceOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSeen

`func (o *BatchWifiInterferenceOpenApiVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *BatchWifiInterferenceOpenApiVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *BatchWifiInterferenceOpenApiVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *BatchWifiInterferenceOpenApiVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMode

`func (o *BatchWifiInterferenceOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BatchWifiInterferenceOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BatchWifiInterferenceOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *BatchWifiInterferenceOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNearestAp

`func (o *BatchWifiInterferenceOpenApiVO) GetNearestAp() string`

GetNearestAp returns the NearestAp field if non-nil, zero value otherwise.

### GetNearestApOk

`func (o *BatchWifiInterferenceOpenApiVO) GetNearestApOk() (*string, bool)`

GetNearestApOk returns a tuple with the NearestAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearestAp

`func (o *BatchWifiInterferenceOpenApiVO) SetNearestAp(v string)`

SetNearestAp sets NearestAp field to given value.

### HasNearestAp

`func (o *BatchWifiInterferenceOpenApiVO) HasNearestAp() bool`

HasNearestAp returns a boolean if a field has been set.

### GetRadioId

`func (o *BatchWifiInterferenceOpenApiVO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *BatchWifiInterferenceOpenApiVO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *BatchWifiInterferenceOpenApiVO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *BatchWifiInterferenceOpenApiVO) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetSecurity

`func (o *BatchWifiInterferenceOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *BatchWifiInterferenceOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *BatchWifiInterferenceOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *BatchWifiInterferenceOpenApiVO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSignal

`func (o *BatchWifiInterferenceOpenApiVO) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *BatchWifiInterferenceOpenApiVO) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *BatchWifiInterferenceOpenApiVO) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *BatchWifiInterferenceOpenApiVO) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetSsid

`func (o *BatchWifiInterferenceOpenApiVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *BatchWifiInterferenceOpenApiVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *BatchWifiInterferenceOpenApiVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *BatchWifiInterferenceOpenApiVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


