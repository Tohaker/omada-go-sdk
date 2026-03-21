# EthUnit1PortAppDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Display the configured port description.Description should contain 1-32 bits numbers, Upper and lower letters, -@_:/. . | [optional] 
**Duplex** | Pointer to **int32** | Port duplex negotiation mode.Duplex should be a value as follows: 2:FULL,0:AUTO | [optional] 
**DuplexLink** | Pointer to **int32** | Actual port duplex mode.DuplexLink should be a value as follows: 0:DISABLE;1:ENABLE | [optional] 
**FlowControl** | Pointer to **int32** | Port flow control function switch.FlowControl should be a value as follows: 0:DISABLE;1:ENABLE | [optional] 
**Lag** | Pointer to **string** | The LAG to which the port belongs. | [optional] 
**LinkStatus** | Pointer to **int32** | Connection status of the PON port. LinkStatus should be a value as follows:0:LINK_DOWN;1:LINK_UP | [optional] 
**MediaType** | Pointer to **int32** | Port media type | [optional] 
**Port** | **string** | OLT physical port | 
**Speed** | Pointer to **int32** | Port speed negotiation mode.Speed should be a value as follows:0;10;100;1000;2500;10000.0 represents Auto, and all other values are in Mbps. | [optional] 
**SpeedLink** | Pointer to **int32** | Port speed negotiation mode: 0 represents Auto, and all other values are in Mbps. | [optional] 
**SpeedMax** | Pointer to **int32** | The maximum rate that this port can achieve, in Mbps.SpeedMax should be null | [optional] 
**Status** | Pointer to **int32** | Port switch status.Status should be a value as follows:0:DISABLE;1:ENABLE | [optional] 
**Type** | Pointer to **int32** | Port type.Type should be null.Type is a value as follows:0:COPPER;1:COMBO;2:SFP;3:SFP+;4:RJ45 | [optional] 

## Methods

### NewEthUnit1PortAppDTO

`func NewEthUnit1PortAppDTO(port string, ) *EthUnit1PortAppDTO`

NewEthUnit1PortAppDTO instantiates a new EthUnit1PortAppDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthUnit1PortAppDTOWithDefaults

`func NewEthUnit1PortAppDTOWithDefaults() *EthUnit1PortAppDTO`

NewEthUnit1PortAppDTOWithDefaults instantiates a new EthUnit1PortAppDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EthUnit1PortAppDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EthUnit1PortAppDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EthUnit1PortAppDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EthUnit1PortAppDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuplex

`func (o *EthUnit1PortAppDTO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *EthUnit1PortAppDTO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *EthUnit1PortAppDTO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *EthUnit1PortAppDTO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDuplexLink

`func (o *EthUnit1PortAppDTO) GetDuplexLink() int32`

GetDuplexLink returns the DuplexLink field if non-nil, zero value otherwise.

### GetDuplexLinkOk

`func (o *EthUnit1PortAppDTO) GetDuplexLinkOk() (*int32, bool)`

GetDuplexLinkOk returns a tuple with the DuplexLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexLink

`func (o *EthUnit1PortAppDTO) SetDuplexLink(v int32)`

SetDuplexLink sets DuplexLink field to given value.

### HasDuplexLink

`func (o *EthUnit1PortAppDTO) HasDuplexLink() bool`

HasDuplexLink returns a boolean if a field has been set.

### GetFlowControl

`func (o *EthUnit1PortAppDTO) GetFlowControl() int32`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *EthUnit1PortAppDTO) GetFlowControlOk() (*int32, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *EthUnit1PortAppDTO) SetFlowControl(v int32)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *EthUnit1PortAppDTO) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetLag

`func (o *EthUnit1PortAppDTO) GetLag() string`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *EthUnit1PortAppDTO) GetLagOk() (*string, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *EthUnit1PortAppDTO) SetLag(v string)`

SetLag sets Lag field to given value.

### HasLag

`func (o *EthUnit1PortAppDTO) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetLinkStatus

`func (o *EthUnit1PortAppDTO) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *EthUnit1PortAppDTO) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *EthUnit1PortAppDTO) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *EthUnit1PortAppDTO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetMediaType

`func (o *EthUnit1PortAppDTO) GetMediaType() int32`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *EthUnit1PortAppDTO) GetMediaTypeOk() (*int32, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *EthUnit1PortAppDTO) SetMediaType(v int32)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *EthUnit1PortAppDTO) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetPort

`func (o *EthUnit1PortAppDTO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EthUnit1PortAppDTO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EthUnit1PortAppDTO) SetPort(v string)`

SetPort sets Port field to given value.


### GetSpeed

`func (o *EthUnit1PortAppDTO) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *EthUnit1PortAppDTO) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *EthUnit1PortAppDTO) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *EthUnit1PortAppDTO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSpeedLink

`func (o *EthUnit1PortAppDTO) GetSpeedLink() int32`

GetSpeedLink returns the SpeedLink field if non-nil, zero value otherwise.

### GetSpeedLinkOk

`func (o *EthUnit1PortAppDTO) GetSpeedLinkOk() (*int32, bool)`

GetSpeedLinkOk returns a tuple with the SpeedLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedLink

`func (o *EthUnit1PortAppDTO) SetSpeedLink(v int32)`

SetSpeedLink sets SpeedLink field to given value.

### HasSpeedLink

`func (o *EthUnit1PortAppDTO) HasSpeedLink() bool`

HasSpeedLink returns a boolean if a field has been set.

### GetSpeedMax

`func (o *EthUnit1PortAppDTO) GetSpeedMax() int32`

GetSpeedMax returns the SpeedMax field if non-nil, zero value otherwise.

### GetSpeedMaxOk

`func (o *EthUnit1PortAppDTO) GetSpeedMaxOk() (*int32, bool)`

GetSpeedMaxOk returns a tuple with the SpeedMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMax

`func (o *EthUnit1PortAppDTO) SetSpeedMax(v int32)`

SetSpeedMax sets SpeedMax field to given value.

### HasSpeedMax

`func (o *EthUnit1PortAppDTO) HasSpeedMax() bool`

HasSpeedMax returns a boolean if a field has been set.

### GetStatus

`func (o *EthUnit1PortAppDTO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EthUnit1PortAppDTO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EthUnit1PortAppDTO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EthUnit1PortAppDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EthUnit1PortAppDTO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EthUnit1PortAppDTO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EthUnit1PortAppDTO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *EthUnit1PortAppDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


