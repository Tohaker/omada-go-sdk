# RogueAPScanResultEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beacon** | Pointer to **int32** | The time interval for the network device to send Beacon frames | [optional] 
**Bssid** | Pointer to **string** | The MAC address of the network device, used as the unique identifier of the network | [optional] 
**Channel** | Pointer to **int32** | Channel | [optional] 
**Id** | Pointer to **string** | ID of the entry | [optional] 
**LastSeen** | Pointer to **int64** | The timestamp corresponding to the time when the network device was last scanned, unit: ms | [optional] 
**Mode** | Pointer to **int32** | 0: 11a  1: 11b  2: 11g  3: 11na  4: 11ng  5: 11ac  6: 11axa  7: 11axg | [optional] 
**NearestAp** | Pointer to **string** | The AP name | [optional] 
**NearestApMac** | Pointer to **string** | The AP that scanned the network device | [optional] 
**RadioId** | Pointer to **int32** | (Wireless) Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz-1; 2:5GHz-2; 3: 6GHz | [optional] 
**Security** | Pointer to **int32** | The encryption method of the network. 0: None  1: WEP  2: WPA-Enterprise  3: WPA-Personal  4:WPA3-SAE , 15:Unknown, the other unspecified return values are also displayed as Unknown | [optional] 
**Signal** | Pointer to **int32** | The signal strength of the network device, unit: dBm | [optional] 
**Ssid** | Pointer to **string** | Name of SSID | [optional] 

## Methods

### NewRogueAPScanResultEntry

`func NewRogueAPScanResultEntry() *RogueAPScanResultEntry`

NewRogueAPScanResultEntry instantiates a new RogueAPScanResultEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRogueAPScanResultEntryWithDefaults

`func NewRogueAPScanResultEntryWithDefaults() *RogueAPScanResultEntry`

NewRogueAPScanResultEntryWithDefaults instantiates a new RogueAPScanResultEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeacon

`func (o *RogueAPScanResultEntry) GetBeacon() int32`

GetBeacon returns the Beacon field if non-nil, zero value otherwise.

### GetBeaconOk

`func (o *RogueAPScanResultEntry) GetBeaconOk() (*int32, bool)`

GetBeaconOk returns a tuple with the Beacon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeacon

`func (o *RogueAPScanResultEntry) SetBeacon(v int32)`

SetBeacon sets Beacon field to given value.

### HasBeacon

`func (o *RogueAPScanResultEntry) HasBeacon() bool`

HasBeacon returns a boolean if a field has been set.

### GetBssid

`func (o *RogueAPScanResultEntry) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *RogueAPScanResultEntry) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *RogueAPScanResultEntry) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *RogueAPScanResultEntry) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetChannel

`func (o *RogueAPScanResultEntry) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RogueAPScanResultEntry) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RogueAPScanResultEntry) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *RogueAPScanResultEntry) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetId

`func (o *RogueAPScanResultEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RogueAPScanResultEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RogueAPScanResultEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RogueAPScanResultEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSeen

`func (o *RogueAPScanResultEntry) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *RogueAPScanResultEntry) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *RogueAPScanResultEntry) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *RogueAPScanResultEntry) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMode

`func (o *RogueAPScanResultEntry) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RogueAPScanResultEntry) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RogueAPScanResultEntry) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RogueAPScanResultEntry) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNearestAp

`func (o *RogueAPScanResultEntry) GetNearestAp() string`

GetNearestAp returns the NearestAp field if non-nil, zero value otherwise.

### GetNearestApOk

`func (o *RogueAPScanResultEntry) GetNearestApOk() (*string, bool)`

GetNearestApOk returns a tuple with the NearestAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearestAp

`func (o *RogueAPScanResultEntry) SetNearestAp(v string)`

SetNearestAp sets NearestAp field to given value.

### HasNearestAp

`func (o *RogueAPScanResultEntry) HasNearestAp() bool`

HasNearestAp returns a boolean if a field has been set.

### GetNearestApMac

`func (o *RogueAPScanResultEntry) GetNearestApMac() string`

GetNearestApMac returns the NearestApMac field if non-nil, zero value otherwise.

### GetNearestApMacOk

`func (o *RogueAPScanResultEntry) GetNearestApMacOk() (*string, bool)`

GetNearestApMacOk returns a tuple with the NearestApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearestApMac

`func (o *RogueAPScanResultEntry) SetNearestApMac(v string)`

SetNearestApMac sets NearestApMac field to given value.

### HasNearestApMac

`func (o *RogueAPScanResultEntry) HasNearestApMac() bool`

HasNearestApMac returns a boolean if a field has been set.

### GetRadioId

`func (o *RogueAPScanResultEntry) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *RogueAPScanResultEntry) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *RogueAPScanResultEntry) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *RogueAPScanResultEntry) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetSecurity

`func (o *RogueAPScanResultEntry) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *RogueAPScanResultEntry) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *RogueAPScanResultEntry) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *RogueAPScanResultEntry) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSignal

`func (o *RogueAPScanResultEntry) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *RogueAPScanResultEntry) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *RogueAPScanResultEntry) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *RogueAPScanResultEntry) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetSsid

`func (o *RogueAPScanResultEntry) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *RogueAPScanResultEntry) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *RogueAPScanResultEntry) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *RogueAPScanResultEntry) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


