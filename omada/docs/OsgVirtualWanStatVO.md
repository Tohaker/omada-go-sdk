# OsgVirtualWanStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnexType** | Pointer to **int32** |  | [optional] 
**DslModulationType** | Pointer to **int32** |  | [optional] 
**InternetState** | Pointer to **int32** | Virtual Wan internet state, 0-disconnected; 1-connected; | [optional] 
**LineStatus** | Pointer to **int32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Virtual wan name | [optional] 
**OnlineDetection** | Pointer to **int32** |  | [optional] 
**Proto** | Pointer to **string** | Virtual WAN ipv4 proto type. | [optional] 
**Status** | Pointer to **int32** | Virtual Wan status, 0-disconnected; 1-connected; | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**VirtualWanId** | Pointer to **string** | Virtual Wan Id | [optional] 
**WanPortIpv4Config** | Pointer to [**OsgVirtualWanIpv4ConfigVO**](OsgVirtualWanIpv4ConfigVO.md) |  | [optional] 

## Methods

### NewOsgVirtualWanStatVO

`func NewOsgVirtualWanStatVO() *OsgVirtualWanStatVO`

NewOsgVirtualWanStatVO instantiates a new OsgVirtualWanStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgVirtualWanStatVOWithDefaults

`func NewOsgVirtualWanStatVOWithDefaults() *OsgVirtualWanStatVO`

NewOsgVirtualWanStatVOWithDefaults instantiates a new OsgVirtualWanStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnexType

`func (o *OsgVirtualWanStatVO) GetAnnexType() int32`

GetAnnexType returns the AnnexType field if non-nil, zero value otherwise.

### GetAnnexTypeOk

`func (o *OsgVirtualWanStatVO) GetAnnexTypeOk() (*int32, bool)`

GetAnnexTypeOk returns a tuple with the AnnexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnexType

`func (o *OsgVirtualWanStatVO) SetAnnexType(v int32)`

SetAnnexType sets AnnexType field to given value.

### HasAnnexType

`func (o *OsgVirtualWanStatVO) HasAnnexType() bool`

HasAnnexType returns a boolean if a field has been set.

### GetDslModulationType

`func (o *OsgVirtualWanStatVO) GetDslModulationType() int32`

GetDslModulationType returns the DslModulationType field if non-nil, zero value otherwise.

### GetDslModulationTypeOk

`func (o *OsgVirtualWanStatVO) GetDslModulationTypeOk() (*int32, bool)`

GetDslModulationTypeOk returns a tuple with the DslModulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslModulationType

`func (o *OsgVirtualWanStatVO) SetDslModulationType(v int32)`

SetDslModulationType sets DslModulationType field to given value.

### HasDslModulationType

`func (o *OsgVirtualWanStatVO) HasDslModulationType() bool`

HasDslModulationType returns a boolean if a field has been set.

### GetInternetState

`func (o *OsgVirtualWanStatVO) GetInternetState() int32`

GetInternetState returns the InternetState field if non-nil, zero value otherwise.

### GetInternetStateOk

`func (o *OsgVirtualWanStatVO) GetInternetStateOk() (*int32, bool)`

GetInternetStateOk returns a tuple with the InternetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetState

`func (o *OsgVirtualWanStatVO) SetInternetState(v int32)`

SetInternetState sets InternetState field to given value.

### HasInternetState

`func (o *OsgVirtualWanStatVO) HasInternetState() bool`

HasInternetState returns a boolean if a field has been set.

### GetLineStatus

`func (o *OsgVirtualWanStatVO) GetLineStatus() int32`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *OsgVirtualWanStatVO) GetLineStatusOk() (*int32, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *OsgVirtualWanStatVO) SetLineStatus(v int32)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *OsgVirtualWanStatVO) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetMac

`func (o *OsgVirtualWanStatVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OsgVirtualWanStatVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OsgVirtualWanStatVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OsgVirtualWanStatVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *OsgVirtualWanStatVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsgVirtualWanStatVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsgVirtualWanStatVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsgVirtualWanStatVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnlineDetection

`func (o *OsgVirtualWanStatVO) GetOnlineDetection() int32`

GetOnlineDetection returns the OnlineDetection field if non-nil, zero value otherwise.

### GetOnlineDetectionOk

`func (o *OsgVirtualWanStatVO) GetOnlineDetectionOk() (*int32, bool)`

GetOnlineDetectionOk returns a tuple with the OnlineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDetection

`func (o *OsgVirtualWanStatVO) SetOnlineDetection(v int32)`

SetOnlineDetection sets OnlineDetection field to given value.

### HasOnlineDetection

`func (o *OsgVirtualWanStatVO) HasOnlineDetection() bool`

HasOnlineDetection returns a boolean if a field has been set.

### GetProto

`func (o *OsgVirtualWanStatVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *OsgVirtualWanStatVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *OsgVirtualWanStatVO) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *OsgVirtualWanStatVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetStatus

`func (o *OsgVirtualWanStatVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OsgVirtualWanStatVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OsgVirtualWanStatVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OsgVirtualWanStatVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OsgVirtualWanStatVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsgVirtualWanStatVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsgVirtualWanStatVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OsgVirtualWanStatVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualWanId

`func (o *OsgVirtualWanStatVO) GetVirtualWanId() string`

GetVirtualWanId returns the VirtualWanId field if non-nil, zero value otherwise.

### GetVirtualWanIdOk

`func (o *OsgVirtualWanStatVO) GetVirtualWanIdOk() (*string, bool)`

GetVirtualWanIdOk returns a tuple with the VirtualWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanId

`func (o *OsgVirtualWanStatVO) SetVirtualWanId(v string)`

SetVirtualWanId sets VirtualWanId field to given value.

### HasVirtualWanId

`func (o *OsgVirtualWanStatVO) HasVirtualWanId() bool`

HasVirtualWanId returns a boolean if a field has been set.

### GetWanPortIpv4Config

`func (o *OsgVirtualWanStatVO) GetWanPortIpv4Config() OsgVirtualWanIpv4ConfigVO`

GetWanPortIpv4Config returns the WanPortIpv4Config field if non-nil, zero value otherwise.

### GetWanPortIpv4ConfigOk

`func (o *OsgVirtualWanStatVO) GetWanPortIpv4ConfigOk() (*OsgVirtualWanIpv4ConfigVO, bool)`

GetWanPortIpv4ConfigOk returns a tuple with the WanPortIpv4Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv4Config

`func (o *OsgVirtualWanStatVO) SetWanPortIpv4Config(v OsgVirtualWanIpv4ConfigVO)`

SetWanPortIpv4Config sets WanPortIpv4Config field to given value.

### HasWanPortIpv4Config

`func (o *OsgVirtualWanStatVO) HasWanPortIpv4Config() bool`

HasWanPortIpv4Config returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


