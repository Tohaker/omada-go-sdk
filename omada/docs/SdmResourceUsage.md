# SdmResourceUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **int32** | It should be a value as follows: 0:Packet Control,  1:Packet Control V6 , 2:MAC ACL Ingress,  3:IP ACL Ingress,  4:IPV6 ACL Ingress,  5:Combined ACL Ingress,  6:MAC Diffserv Ingress, 7:IP Diffserv Ingress,  8:IPV6 Diffserv Ingress,  9:MAC Diffserv Egress,  10:IP Diffserv Egress,  11:IPV6 Diffserv Egress,  12:Voice VLAN,  13:QoS VLAN,  14:VLAN VPN Ingress,  15:VLAN VPN Egress,  16:IPV4 Source Guard,  17:IPV6 Source Guard,  18:CPP Egress,  19:VRF,  20:Tunnel V4,  21:Tunnel V6 Stage0,  22:Tunnel V6 Stage1,  23:MPLS UNI,  24:MPLS NNI0,  25:MPLS NNI1,  26:VXLAN V4 UNI,  27:VXLAN V6 UNI,  28:VXLAN V4 NNI Stage0,  29:VXLAN V4 NNI Stage1,  30:VXLAN V6 NNI Stage0,  31:VXLAN V6 NNI Stage1,  32:BFD V4,  33:BFD V6,  34:VXLAN Passenger,  35:QoS Rule V4,  36:QoS Rule V6 . | [optional] 
**FreeNum** | Pointer to **int32** | The number of entries currently available for use by the feature. | [optional] 
**UsedNum** | Pointer to **int32** | The number of entries currently used by the feature. | [optional] 

## Methods

### NewSdmResourceUsage

`func NewSdmResourceUsage() *SdmResourceUsage`

NewSdmResourceUsage instantiates a new SdmResourceUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdmResourceUsageWithDefaults

`func NewSdmResourceUsageWithDefaults() *SdmResourceUsage`

NewSdmResourceUsageWithDefaults instantiates a new SdmResourceUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SdmResourceUsage) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SdmResourceUsage) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SdmResourceUsage) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SdmResourceUsage) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFreeNum

`func (o *SdmResourceUsage) GetFreeNum() int32`

GetFreeNum returns the FreeNum field if non-nil, zero value otherwise.

### GetFreeNumOk

`func (o *SdmResourceUsage) GetFreeNumOk() (*int32, bool)`

GetFreeNumOk returns a tuple with the FreeNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeNum

`func (o *SdmResourceUsage) SetFreeNum(v int32)`

SetFreeNum sets FreeNum field to given value.

### HasFreeNum

`func (o *SdmResourceUsage) HasFreeNum() bool`

HasFreeNum returns a boolean if a field has been set.

### GetUsedNum

`func (o *SdmResourceUsage) GetUsedNum() int32`

GetUsedNum returns the UsedNum field if non-nil, zero value otherwise.

### GetUsedNumOk

`func (o *SdmResourceUsage) GetUsedNumOk() (*int32, bool)`

GetUsedNumOk returns a tuple with the UsedNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedNum

`func (o *SdmResourceUsage) SetUsedNum(v int32)`

SetUsedNum sets UsedNum field to given value.

### HasUsedNum

`func (o *SdmResourceUsage) HasUsedNum() bool`

HasUsedNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


