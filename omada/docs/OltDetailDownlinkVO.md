# OltDetailDownlinkVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Ip address | [optional] 
**Ipv6** | Pointer to **string** | Ipv6 | [optional] 
**LinkStatus** | Pointer to **int32** | Link status should be a value as follows: 0:LINK_DOWN;1:LINK_UP | [optional] 
**Mac** | Pointer to **string** | Mac of the device | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Name of the device connected to the port | [optional] 
**Port** | Pointer to **string** | Port | [optional] 
**Speed** | Pointer to **int64** | Speed | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewOltDetailDownlinkVO

`func NewOltDetailDownlinkVO() *OltDetailDownlinkVO`

NewOltDetailDownlinkVO instantiates a new OltDetailDownlinkVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltDetailDownlinkVOWithDefaults

`func NewOltDetailDownlinkVOWithDefaults() *OltDetailDownlinkVO`

NewOltDetailDownlinkVOWithDefaults instantiates a new OltDetailDownlinkVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *OltDetailDownlinkVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OltDetailDownlinkVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OltDetailDownlinkVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OltDetailDownlinkVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *OltDetailDownlinkVO) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *OltDetailDownlinkVO) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *OltDetailDownlinkVO) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *OltDetailDownlinkVO) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLinkStatus

`func (o *OltDetailDownlinkVO) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *OltDetailDownlinkVO) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *OltDetailDownlinkVO) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *OltDetailDownlinkVO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetMac

`func (o *OltDetailDownlinkVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OltDetailDownlinkVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OltDetailDownlinkVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OltDetailDownlinkVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OltDetailDownlinkVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OltDetailDownlinkVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OltDetailDownlinkVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OltDetailDownlinkVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OltDetailDownlinkVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OltDetailDownlinkVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OltDetailDownlinkVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OltDetailDownlinkVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *OltDetailDownlinkVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OltDetailDownlinkVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OltDetailDownlinkVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OltDetailDownlinkVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *OltDetailDownlinkVO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OltDetailDownlinkVO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OltDetailDownlinkVO) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *OltDetailDownlinkVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSpeed

`func (o *OltDetailDownlinkVO) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OltDetailDownlinkVO) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OltDetailDownlinkVO) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OltDetailDownlinkVO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetType

`func (o *OltDetailDownlinkVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OltDetailDownlinkVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OltDetailDownlinkVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OltDetailDownlinkVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


