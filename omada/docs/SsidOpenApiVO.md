# SsidOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to **int32** | SSID band. The lowest bit indicates whether 2.4G is included; the second lowest bit indicates whether 5G is included; the third lowest bit indicates whether 6G is included; 1 means included while 0 means not included. For example, 7(111) means that 2G/5G/6G are enabled; 1(001) means that 2G is enabled. (When 5G is included，it means 5G/5G1/5G2 are enabled.) | [optional] 
**Broadcast** | Pointer to **bool** | SSID broadcast config status. True: enable, false: disable. | [optional] 
**GuestNetEnable** | Pointer to **bool** | SSID guest network config status. True: enable, false: disable. | [optional] 
**Name** | Pointer to **string** | SSID name. It should contain 1 to 32 UTF-8 characters. | [optional] 
**Security** | Pointer to **int32** | SSID security mode; Security should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal; 4: PPSK without RADIUS; 5: PPSK with RADIUS. | [optional] 
**SsidId** | Pointer to **string** | SSID ID | [optional] 
**VlanEnable** | Pointer to **bool** | SSID VLAN config status. True: enable, false: disable. | [optional] 
**VlanId** | Pointer to **int32** | SSID VLAN ID. This field is required when Parameter [vlanEnable] is true; It should be within the range of 1–4094. | [optional] 
**VlanPoolIds** | Pointer to **string** | SSID VLAN POOL IDs. This field is required when Parameter [vlanEnable] is true; The numbers contain in it should be within the range of 1–4094. | [optional] 

## Methods

### NewSsidOpenApiVO

`func NewSsidOpenApiVO() *SsidOpenApiVO`

NewSsidOpenApiVO instantiates a new SsidOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidOpenApiVOWithDefaults

`func NewSsidOpenApiVOWithDefaults() *SsidOpenApiVO`

NewSsidOpenApiVOWithDefaults instantiates a new SsidOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *SsidOpenApiVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *SsidOpenApiVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *SsidOpenApiVO) SetBand(v int32)`

SetBand sets Band field to given value.

### HasBand

`func (o *SsidOpenApiVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetBroadcast

`func (o *SsidOpenApiVO) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *SsidOpenApiVO) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *SsidOpenApiVO) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *SsidOpenApiVO) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetGuestNetEnable

`func (o *SsidOpenApiVO) GetGuestNetEnable() bool`

GetGuestNetEnable returns the GuestNetEnable field if non-nil, zero value otherwise.

### GetGuestNetEnableOk

`func (o *SsidOpenApiVO) GetGuestNetEnableOk() (*bool, bool)`

GetGuestNetEnableOk returns a tuple with the GuestNetEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNetEnable

`func (o *SsidOpenApiVO) SetGuestNetEnable(v bool)`

SetGuestNetEnable sets GuestNetEnable field to given value.

### HasGuestNetEnable

`func (o *SsidOpenApiVO) HasGuestNetEnable() bool`

HasGuestNetEnable returns a boolean if a field has been set.

### GetName

`func (o *SsidOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsidOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsidOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SsidOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurity

`func (o *SsidOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SsidOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SsidOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SsidOpenApiVO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSsidId

`func (o *SsidOpenApiVO) GetSsidId() string`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *SsidOpenApiVO) GetSsidIdOk() (*string, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *SsidOpenApiVO) SetSsidId(v string)`

SetSsidId sets SsidId field to given value.

### HasSsidId

`func (o *SsidOpenApiVO) HasSsidId() bool`

HasSsidId returns a boolean if a field has been set.

### GetVlanEnable

`func (o *SsidOpenApiVO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *SsidOpenApiVO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *SsidOpenApiVO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.

### HasVlanEnable

`func (o *SsidOpenApiVO) HasVlanEnable() bool`

HasVlanEnable returns a boolean if a field has been set.

### GetVlanId

`func (o *SsidOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanPoolIds

`func (o *SsidOpenApiVO) GetVlanPoolIds() string`

GetVlanPoolIds returns the VlanPoolIds field if non-nil, zero value otherwise.

### GetVlanPoolIdsOk

`func (o *SsidOpenApiVO) GetVlanPoolIdsOk() (*string, bool)`

GetVlanPoolIdsOk returns a tuple with the VlanPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPoolIds

`func (o *SsidOpenApiVO) SetVlanPoolIds(v string)`

SetVlanPoolIds sets VlanPoolIds field to given value.

### HasVlanPoolIds

`func (o *SsidOpenApiVO) HasVlanPoolIds() bool`

HasVlanPoolIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


