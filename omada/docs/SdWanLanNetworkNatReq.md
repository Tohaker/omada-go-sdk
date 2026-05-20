# SdWanLanNetworkNatReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewaySubnet** | **string** | The Gateway Subnet of the lan network | 
**Id** | Pointer to **string** | The ID of the lan network | [optional] 
**NeedMap** | Pointer to **bool** | This value exists in customNetwork. If true, it indicates a conflicting route that requires mapping; routes with false should be avoided by NAT mapping. | [optional] 
**SiteId** | **string** | The site ID of the lan network | 

## Methods

### NewSdWanLanNetworkNatReq

`func NewSdWanLanNetworkNatReq(gatewaySubnet string, siteId string, ) *SdWanLanNetworkNatReq`

NewSdWanLanNetworkNatReq instantiates a new SdWanLanNetworkNatReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanLanNetworkNatReqWithDefaults

`func NewSdWanLanNetworkNatReqWithDefaults() *SdWanLanNetworkNatReq`

NewSdWanLanNetworkNatReqWithDefaults instantiates a new SdWanLanNetworkNatReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewaySubnet

`func (o *SdWanLanNetworkNatReq) GetGatewaySubnet() string`

GetGatewaySubnet returns the GatewaySubnet field if non-nil, zero value otherwise.

### GetGatewaySubnetOk

`func (o *SdWanLanNetworkNatReq) GetGatewaySubnetOk() (*string, bool)`

GetGatewaySubnetOk returns a tuple with the GatewaySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubnet

`func (o *SdWanLanNetworkNatReq) SetGatewaySubnet(v string)`

SetGatewaySubnet sets GatewaySubnet field to given value.


### GetId

`func (o *SdWanLanNetworkNatReq) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdWanLanNetworkNatReq) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdWanLanNetworkNatReq) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SdWanLanNetworkNatReq) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNeedMap

`func (o *SdWanLanNetworkNatReq) GetNeedMap() bool`

GetNeedMap returns the NeedMap field if non-nil, zero value otherwise.

### GetNeedMapOk

`func (o *SdWanLanNetworkNatReq) GetNeedMapOk() (*bool, bool)`

GetNeedMapOk returns a tuple with the NeedMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedMap

`func (o *SdWanLanNetworkNatReq) SetNeedMap(v bool)`

SetNeedMap sets NeedMap field to given value.

### HasNeedMap

`func (o *SdWanLanNetworkNatReq) HasNeedMap() bool`

HasNeedMap returns a boolean if a field has been set.

### GetSiteId

`func (o *SdWanLanNetworkNatReq) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *SdWanLanNetworkNatReq) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *SdWanLanNetworkNatReq) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


