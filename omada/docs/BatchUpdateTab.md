# BatchUpdateTab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to **[]int32** | Card list displayed in Tab. Each number indicates a type of card. card should be a value as follows:  0: Controller Snapshot, 101: Alerts, 102: ISP Load, 103: L2TP/PPTP VPN, 104: Channel Distribution and Usage, 105: Most Active EAPs, 106: Most Active Switches, 108: WiFi/Switching Summary, 109: Traffic Distribution, 110: Traffic Activities, 111: Clients Distribution, 112: Retry Rate/Dropped Rate, 117: Top Device CPU/Memory Usage, 118: IPSEC VPN, 119: OPEN VPN, 120: SSL VPN, 200: Most Active Clients, 201: Longest client uptime, 202: Clients Freq Distribution, 203: Clients Activities, 204: Clients Association Activities, 205: Association Failures, 206: Clients Association Time Distribution, 207: Clients SSID Distribution, 208: Clients RSSI Distribution | [optional] 
**Id** | Pointer to **string** | A unique identifier | [optional] 

## Methods

### NewBatchUpdateTab

`func NewBatchUpdateTab() *BatchUpdateTab`

NewBatchUpdateTab instantiates a new BatchUpdateTab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateTabWithDefaults

`func NewBatchUpdateTabWithDefaults() *BatchUpdateTab`

NewBatchUpdateTabWithDefaults instantiates a new BatchUpdateTab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *BatchUpdateTab) GetCards() []int32`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *BatchUpdateTab) GetCardsOk() (*[]int32, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *BatchUpdateTab) SetCards(v []int32)`

SetCards sets Cards field to given value.

### HasCards

`func (o *BatchUpdateTab) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetId

`func (o *BatchUpdateTab) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchUpdateTab) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchUpdateTab) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BatchUpdateTab) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


