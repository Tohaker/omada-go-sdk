# StatisticInfoDetailDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlignmentErrors** | Pointer to **int64** | Alignment errors | [optional] 
**CollisionsErrors** | Pointer to **int64** | Collision errors | [optional] 
**PacketsOf1023To1518Octets** | Pointer to **int64** | Packets of 1023 to 1518 octets | [optional] 
**PacketsOf128To255Octets** | Pointer to **int64** | Packets of 128 to 255 octets | [optional] 
**PacketsOf256To511Octets** | Pointer to **int64** | Packets of 256 to511 octets | [optional] 
**PacketsOf512To1023Octets** | Pointer to **int64** | Packets of 512 to 1023 octets | [optional] 
**PacketsOf64Octets** | Pointer to **int64** | Packets Of 64 Octets | [optional] 
**PacketsOf65To127Octets** | Pointer to **int64** | Packets of 65 to 127 octets | [optional] 
**ReceivedBroadcast** | Pointer to **int64** | Received broadcast | [optional] 
**ReceivedBytes** | Pointer to **int64** | Received bytes | [optional] 
**ReceivedJumbo** | Pointer to **int64** | Received jumbo | [optional] 
**ReceivedMulticast** | Pointer to **int64** | Received multicast | [optional] 
**ReceivedPackets** | Pointer to **int64** | Received packets | [optional] 
**ReceivedUnicast** | Pointer to **int64** | Received unicast | [optional] 
**SentBroadcast** | Pointer to **int64** | Sent broadcast | [optional] 
**SentBytes** | Pointer to **int64** | Sent bytes | [optional] 
**SentJumbo** | Pointer to **int64** | Sent jumbo | [optional] 
**SentMulticast** | Pointer to **int64** | Sent multicast | [optional] 
**SentPackets** | Pointer to **int64** | Sent packets | [optional] 
**SentUnicast** | Pointer to **int64** | Sent unicast | [optional] 
**ServicePortIndex** | Pointer to **int32** | Service port index | [optional] 
**UndersizePackets** | Pointer to **int64** | Undersize packets | [optional] 

## Methods

### NewStatisticInfoDetailDTO

`func NewStatisticInfoDetailDTO() *StatisticInfoDetailDTO`

NewStatisticInfoDetailDTO instantiates a new StatisticInfoDetailDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticInfoDetailDTOWithDefaults

`func NewStatisticInfoDetailDTOWithDefaults() *StatisticInfoDetailDTO`

NewStatisticInfoDetailDTOWithDefaults instantiates a new StatisticInfoDetailDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignmentErrors

`func (o *StatisticInfoDetailDTO) GetAlignmentErrors() int64`

GetAlignmentErrors returns the AlignmentErrors field if non-nil, zero value otherwise.

### GetAlignmentErrorsOk

`func (o *StatisticInfoDetailDTO) GetAlignmentErrorsOk() (*int64, bool)`

GetAlignmentErrorsOk returns a tuple with the AlignmentErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignmentErrors

`func (o *StatisticInfoDetailDTO) SetAlignmentErrors(v int64)`

SetAlignmentErrors sets AlignmentErrors field to given value.

### HasAlignmentErrors

`func (o *StatisticInfoDetailDTO) HasAlignmentErrors() bool`

HasAlignmentErrors returns a boolean if a field has been set.

### GetCollisionsErrors

`func (o *StatisticInfoDetailDTO) GetCollisionsErrors() int64`

GetCollisionsErrors returns the CollisionsErrors field if non-nil, zero value otherwise.

### GetCollisionsErrorsOk

`func (o *StatisticInfoDetailDTO) GetCollisionsErrorsOk() (*int64, bool)`

GetCollisionsErrorsOk returns a tuple with the CollisionsErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollisionsErrors

`func (o *StatisticInfoDetailDTO) SetCollisionsErrors(v int64)`

SetCollisionsErrors sets CollisionsErrors field to given value.

### HasCollisionsErrors

`func (o *StatisticInfoDetailDTO) HasCollisionsErrors() bool`

HasCollisionsErrors returns a boolean if a field has been set.

### GetPacketsOf1023To1518Octets

`func (o *StatisticInfoDetailDTO) GetPacketsOf1023To1518Octets() int64`

GetPacketsOf1023To1518Octets returns the PacketsOf1023To1518Octets field if non-nil, zero value otherwise.

### GetPacketsOf1023To1518OctetsOk

`func (o *StatisticInfoDetailDTO) GetPacketsOf1023To1518OctetsOk() (*int64, bool)`

GetPacketsOf1023To1518OctetsOk returns a tuple with the PacketsOf1023To1518Octets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsOf1023To1518Octets

`func (o *StatisticInfoDetailDTO) SetPacketsOf1023To1518Octets(v int64)`

SetPacketsOf1023To1518Octets sets PacketsOf1023To1518Octets field to given value.

### HasPacketsOf1023To1518Octets

`func (o *StatisticInfoDetailDTO) HasPacketsOf1023To1518Octets() bool`

HasPacketsOf1023To1518Octets returns a boolean if a field has been set.

### GetPacketsOf128To255Octets

`func (o *StatisticInfoDetailDTO) GetPacketsOf128To255Octets() int64`

GetPacketsOf128To255Octets returns the PacketsOf128To255Octets field if non-nil, zero value otherwise.

### GetPacketsOf128To255OctetsOk

`func (o *StatisticInfoDetailDTO) GetPacketsOf128To255OctetsOk() (*int64, bool)`

GetPacketsOf128To255OctetsOk returns a tuple with the PacketsOf128To255Octets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsOf128To255Octets

`func (o *StatisticInfoDetailDTO) SetPacketsOf128To255Octets(v int64)`

SetPacketsOf128To255Octets sets PacketsOf128To255Octets field to given value.

### HasPacketsOf128To255Octets

`func (o *StatisticInfoDetailDTO) HasPacketsOf128To255Octets() bool`

HasPacketsOf128To255Octets returns a boolean if a field has been set.

### GetPacketsOf256To511Octets

`func (o *StatisticInfoDetailDTO) GetPacketsOf256To511Octets() int64`

GetPacketsOf256To511Octets returns the PacketsOf256To511Octets field if non-nil, zero value otherwise.

### GetPacketsOf256To511OctetsOk

`func (o *StatisticInfoDetailDTO) GetPacketsOf256To511OctetsOk() (*int64, bool)`

GetPacketsOf256To511OctetsOk returns a tuple with the PacketsOf256To511Octets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsOf256To511Octets

`func (o *StatisticInfoDetailDTO) SetPacketsOf256To511Octets(v int64)`

SetPacketsOf256To511Octets sets PacketsOf256To511Octets field to given value.

### HasPacketsOf256To511Octets

`func (o *StatisticInfoDetailDTO) HasPacketsOf256To511Octets() bool`

HasPacketsOf256To511Octets returns a boolean if a field has been set.

### GetPacketsOf512To1023Octets

`func (o *StatisticInfoDetailDTO) GetPacketsOf512To1023Octets() int64`

GetPacketsOf512To1023Octets returns the PacketsOf512To1023Octets field if non-nil, zero value otherwise.

### GetPacketsOf512To1023OctetsOk

`func (o *StatisticInfoDetailDTO) GetPacketsOf512To1023OctetsOk() (*int64, bool)`

GetPacketsOf512To1023OctetsOk returns a tuple with the PacketsOf512To1023Octets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsOf512To1023Octets

`func (o *StatisticInfoDetailDTO) SetPacketsOf512To1023Octets(v int64)`

SetPacketsOf512To1023Octets sets PacketsOf512To1023Octets field to given value.

### HasPacketsOf512To1023Octets

`func (o *StatisticInfoDetailDTO) HasPacketsOf512To1023Octets() bool`

HasPacketsOf512To1023Octets returns a boolean if a field has been set.

### GetPacketsOf64Octets

`func (o *StatisticInfoDetailDTO) GetPacketsOf64Octets() int64`

GetPacketsOf64Octets returns the PacketsOf64Octets field if non-nil, zero value otherwise.

### GetPacketsOf64OctetsOk

`func (o *StatisticInfoDetailDTO) GetPacketsOf64OctetsOk() (*int64, bool)`

GetPacketsOf64OctetsOk returns a tuple with the PacketsOf64Octets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsOf64Octets

`func (o *StatisticInfoDetailDTO) SetPacketsOf64Octets(v int64)`

SetPacketsOf64Octets sets PacketsOf64Octets field to given value.

### HasPacketsOf64Octets

`func (o *StatisticInfoDetailDTO) HasPacketsOf64Octets() bool`

HasPacketsOf64Octets returns a boolean if a field has been set.

### GetPacketsOf65To127Octets

`func (o *StatisticInfoDetailDTO) GetPacketsOf65To127Octets() int64`

GetPacketsOf65To127Octets returns the PacketsOf65To127Octets field if non-nil, zero value otherwise.

### GetPacketsOf65To127OctetsOk

`func (o *StatisticInfoDetailDTO) GetPacketsOf65To127OctetsOk() (*int64, bool)`

GetPacketsOf65To127OctetsOk returns a tuple with the PacketsOf65To127Octets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsOf65To127Octets

`func (o *StatisticInfoDetailDTO) SetPacketsOf65To127Octets(v int64)`

SetPacketsOf65To127Octets sets PacketsOf65To127Octets field to given value.

### HasPacketsOf65To127Octets

`func (o *StatisticInfoDetailDTO) HasPacketsOf65To127Octets() bool`

HasPacketsOf65To127Octets returns a boolean if a field has been set.

### GetReceivedBroadcast

`func (o *StatisticInfoDetailDTO) GetReceivedBroadcast() int64`

GetReceivedBroadcast returns the ReceivedBroadcast field if non-nil, zero value otherwise.

### GetReceivedBroadcastOk

`func (o *StatisticInfoDetailDTO) GetReceivedBroadcastOk() (*int64, bool)`

GetReceivedBroadcastOk returns a tuple with the ReceivedBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBroadcast

`func (o *StatisticInfoDetailDTO) SetReceivedBroadcast(v int64)`

SetReceivedBroadcast sets ReceivedBroadcast field to given value.

### HasReceivedBroadcast

`func (o *StatisticInfoDetailDTO) HasReceivedBroadcast() bool`

HasReceivedBroadcast returns a boolean if a field has been set.

### GetReceivedBytes

`func (o *StatisticInfoDetailDTO) GetReceivedBytes() int64`

GetReceivedBytes returns the ReceivedBytes field if non-nil, zero value otherwise.

### GetReceivedBytesOk

`func (o *StatisticInfoDetailDTO) GetReceivedBytesOk() (*int64, bool)`

GetReceivedBytesOk returns a tuple with the ReceivedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBytes

`func (o *StatisticInfoDetailDTO) SetReceivedBytes(v int64)`

SetReceivedBytes sets ReceivedBytes field to given value.

### HasReceivedBytes

`func (o *StatisticInfoDetailDTO) HasReceivedBytes() bool`

HasReceivedBytes returns a boolean if a field has been set.

### GetReceivedJumbo

`func (o *StatisticInfoDetailDTO) GetReceivedJumbo() int64`

GetReceivedJumbo returns the ReceivedJumbo field if non-nil, zero value otherwise.

### GetReceivedJumboOk

`func (o *StatisticInfoDetailDTO) GetReceivedJumboOk() (*int64, bool)`

GetReceivedJumboOk returns a tuple with the ReceivedJumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedJumbo

`func (o *StatisticInfoDetailDTO) SetReceivedJumbo(v int64)`

SetReceivedJumbo sets ReceivedJumbo field to given value.

### HasReceivedJumbo

`func (o *StatisticInfoDetailDTO) HasReceivedJumbo() bool`

HasReceivedJumbo returns a boolean if a field has been set.

### GetReceivedMulticast

`func (o *StatisticInfoDetailDTO) GetReceivedMulticast() int64`

GetReceivedMulticast returns the ReceivedMulticast field if non-nil, zero value otherwise.

### GetReceivedMulticastOk

`func (o *StatisticInfoDetailDTO) GetReceivedMulticastOk() (*int64, bool)`

GetReceivedMulticastOk returns a tuple with the ReceivedMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedMulticast

`func (o *StatisticInfoDetailDTO) SetReceivedMulticast(v int64)`

SetReceivedMulticast sets ReceivedMulticast field to given value.

### HasReceivedMulticast

`func (o *StatisticInfoDetailDTO) HasReceivedMulticast() bool`

HasReceivedMulticast returns a boolean if a field has been set.

### GetReceivedPackets

`func (o *StatisticInfoDetailDTO) GetReceivedPackets() int64`

GetReceivedPackets returns the ReceivedPackets field if non-nil, zero value otherwise.

### GetReceivedPacketsOk

`func (o *StatisticInfoDetailDTO) GetReceivedPacketsOk() (*int64, bool)`

GetReceivedPacketsOk returns a tuple with the ReceivedPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedPackets

`func (o *StatisticInfoDetailDTO) SetReceivedPackets(v int64)`

SetReceivedPackets sets ReceivedPackets field to given value.

### HasReceivedPackets

`func (o *StatisticInfoDetailDTO) HasReceivedPackets() bool`

HasReceivedPackets returns a boolean if a field has been set.

### GetReceivedUnicast

`func (o *StatisticInfoDetailDTO) GetReceivedUnicast() int64`

GetReceivedUnicast returns the ReceivedUnicast field if non-nil, zero value otherwise.

### GetReceivedUnicastOk

`func (o *StatisticInfoDetailDTO) GetReceivedUnicastOk() (*int64, bool)`

GetReceivedUnicastOk returns a tuple with the ReceivedUnicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedUnicast

`func (o *StatisticInfoDetailDTO) SetReceivedUnicast(v int64)`

SetReceivedUnicast sets ReceivedUnicast field to given value.

### HasReceivedUnicast

`func (o *StatisticInfoDetailDTO) HasReceivedUnicast() bool`

HasReceivedUnicast returns a boolean if a field has been set.

### GetSentBroadcast

`func (o *StatisticInfoDetailDTO) GetSentBroadcast() int64`

GetSentBroadcast returns the SentBroadcast field if non-nil, zero value otherwise.

### GetSentBroadcastOk

`func (o *StatisticInfoDetailDTO) GetSentBroadcastOk() (*int64, bool)`

GetSentBroadcastOk returns a tuple with the SentBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentBroadcast

`func (o *StatisticInfoDetailDTO) SetSentBroadcast(v int64)`

SetSentBroadcast sets SentBroadcast field to given value.

### HasSentBroadcast

`func (o *StatisticInfoDetailDTO) HasSentBroadcast() bool`

HasSentBroadcast returns a boolean if a field has been set.

### GetSentBytes

`func (o *StatisticInfoDetailDTO) GetSentBytes() int64`

GetSentBytes returns the SentBytes field if non-nil, zero value otherwise.

### GetSentBytesOk

`func (o *StatisticInfoDetailDTO) GetSentBytesOk() (*int64, bool)`

GetSentBytesOk returns a tuple with the SentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentBytes

`func (o *StatisticInfoDetailDTO) SetSentBytes(v int64)`

SetSentBytes sets SentBytes field to given value.

### HasSentBytes

`func (o *StatisticInfoDetailDTO) HasSentBytes() bool`

HasSentBytes returns a boolean if a field has been set.

### GetSentJumbo

`func (o *StatisticInfoDetailDTO) GetSentJumbo() int64`

GetSentJumbo returns the SentJumbo field if non-nil, zero value otherwise.

### GetSentJumboOk

`func (o *StatisticInfoDetailDTO) GetSentJumboOk() (*int64, bool)`

GetSentJumboOk returns a tuple with the SentJumbo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentJumbo

`func (o *StatisticInfoDetailDTO) SetSentJumbo(v int64)`

SetSentJumbo sets SentJumbo field to given value.

### HasSentJumbo

`func (o *StatisticInfoDetailDTO) HasSentJumbo() bool`

HasSentJumbo returns a boolean if a field has been set.

### GetSentMulticast

`func (o *StatisticInfoDetailDTO) GetSentMulticast() int64`

GetSentMulticast returns the SentMulticast field if non-nil, zero value otherwise.

### GetSentMulticastOk

`func (o *StatisticInfoDetailDTO) GetSentMulticastOk() (*int64, bool)`

GetSentMulticastOk returns a tuple with the SentMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentMulticast

`func (o *StatisticInfoDetailDTO) SetSentMulticast(v int64)`

SetSentMulticast sets SentMulticast field to given value.

### HasSentMulticast

`func (o *StatisticInfoDetailDTO) HasSentMulticast() bool`

HasSentMulticast returns a boolean if a field has been set.

### GetSentPackets

`func (o *StatisticInfoDetailDTO) GetSentPackets() int64`

GetSentPackets returns the SentPackets field if non-nil, zero value otherwise.

### GetSentPacketsOk

`func (o *StatisticInfoDetailDTO) GetSentPacketsOk() (*int64, bool)`

GetSentPacketsOk returns a tuple with the SentPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentPackets

`func (o *StatisticInfoDetailDTO) SetSentPackets(v int64)`

SetSentPackets sets SentPackets field to given value.

### HasSentPackets

`func (o *StatisticInfoDetailDTO) HasSentPackets() bool`

HasSentPackets returns a boolean if a field has been set.

### GetSentUnicast

`func (o *StatisticInfoDetailDTO) GetSentUnicast() int64`

GetSentUnicast returns the SentUnicast field if non-nil, zero value otherwise.

### GetSentUnicastOk

`func (o *StatisticInfoDetailDTO) GetSentUnicastOk() (*int64, bool)`

GetSentUnicastOk returns a tuple with the SentUnicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentUnicast

`func (o *StatisticInfoDetailDTO) SetSentUnicast(v int64)`

SetSentUnicast sets SentUnicast field to given value.

### HasSentUnicast

`func (o *StatisticInfoDetailDTO) HasSentUnicast() bool`

HasSentUnicast returns a boolean if a field has been set.

### GetServicePortIndex

`func (o *StatisticInfoDetailDTO) GetServicePortIndex() int32`

GetServicePortIndex returns the ServicePortIndex field if non-nil, zero value otherwise.

### GetServicePortIndexOk

`func (o *StatisticInfoDetailDTO) GetServicePortIndexOk() (*int32, bool)`

GetServicePortIndexOk returns a tuple with the ServicePortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortIndex

`func (o *StatisticInfoDetailDTO) SetServicePortIndex(v int32)`

SetServicePortIndex sets ServicePortIndex field to given value.

### HasServicePortIndex

`func (o *StatisticInfoDetailDTO) HasServicePortIndex() bool`

HasServicePortIndex returns a boolean if a field has been set.

### GetUndersizePackets

`func (o *StatisticInfoDetailDTO) GetUndersizePackets() int64`

GetUndersizePackets returns the UndersizePackets field if non-nil, zero value otherwise.

### GetUndersizePacketsOk

`func (o *StatisticInfoDetailDTO) GetUndersizePacketsOk() (*int64, bool)`

GetUndersizePacketsOk returns a tuple with the UndersizePackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndersizePackets

`func (o *StatisticInfoDetailDTO) SetUndersizePackets(v int64)`

SetUndersizePackets sets UndersizePackets field to given value.

### HasUndersizePackets

`func (o *StatisticInfoDetailDTO) HasUndersizePackets() bool`

HasUndersizePackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


