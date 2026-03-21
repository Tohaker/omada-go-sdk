# UplinkSwitchInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex: 1:Half; 2:Full. | [optional] 
**LagId** | Pointer to **string** | Lag Id | [optional] 
**LinkSpeed** | Pointer to **int32** | Port link speed, 1: 10Mbps; 2: 100Mbps; 3: 1000Mbps; 4: 2.5Gbps; 5: 10Gbps. | [optional] 
**Port** | Pointer to **int32** | Client connected port. | [optional] 
**StandardPort** | Pointer to **string** | Standard Port of Stack | [optional] 
**TrafficDown** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**TrafficUp** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 

## Methods

### NewUplinkSwitchInfo

`func NewUplinkSwitchInfo() *UplinkSwitchInfo`

NewUplinkSwitchInfo instantiates a new UplinkSwitchInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplinkSwitchInfoWithDefaults

`func NewUplinkSwitchInfoWithDefaults() *UplinkSwitchInfo`

NewUplinkSwitchInfoWithDefaults instantiates a new UplinkSwitchInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *UplinkSwitchInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *UplinkSwitchInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *UplinkSwitchInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *UplinkSwitchInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLagId

`func (o *UplinkSwitchInfo) GetLagId() string`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *UplinkSwitchInfo) GetLagIdOk() (*string, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *UplinkSwitchInfo) SetLagId(v string)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *UplinkSwitchInfo) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *UplinkSwitchInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *UplinkSwitchInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *UplinkSwitchInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *UplinkSwitchInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *UplinkSwitchInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UplinkSwitchInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UplinkSwitchInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UplinkSwitchInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardPort

`func (o *UplinkSwitchInfo) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *UplinkSwitchInfo) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *UplinkSwitchInfo) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *UplinkSwitchInfo) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetTrafficDown

`func (o *UplinkSwitchInfo) GetTrafficDown() int64`

GetTrafficDown returns the TrafficDown field if non-nil, zero value otherwise.

### GetTrafficDownOk

`func (o *UplinkSwitchInfo) GetTrafficDownOk() (*int64, bool)`

GetTrafficDownOk returns a tuple with the TrafficDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDown

`func (o *UplinkSwitchInfo) SetTrafficDown(v int64)`

SetTrafficDown sets TrafficDown field to given value.

### HasTrafficDown

`func (o *UplinkSwitchInfo) HasTrafficDown() bool`

HasTrafficDown returns a boolean if a field has been set.

### GetTrafficUp

`func (o *UplinkSwitchInfo) GetTrafficUp() int64`

GetTrafficUp returns the TrafficUp field if non-nil, zero value otherwise.

### GetTrafficUpOk

`func (o *UplinkSwitchInfo) GetTrafficUpOk() (*int64, bool)`

GetTrafficUpOk returns a tuple with the TrafficUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficUp

`func (o *UplinkSwitchInfo) SetTrafficUp(v int64)`

SetTrafficUp sets TrafficUp field to given value.

### HasTrafficUp

`func (o *UplinkSwitchInfo) HasTrafficUp() bool`

HasTrafficUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


