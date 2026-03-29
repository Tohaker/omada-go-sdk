# ClientConnectionEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** | Client connect event attributes map, such as [{dev_mac}:{AA-BB-CC-DD-EE-FF}]. | [optional] 
**Content** | Pointer to **string** | Client connect event description, such as: Connected to [ap:{dev_mac}] with SSID {ssid_name}. | [optional] 
**EventType** | Pointer to **int32** | Event type should be a value as follows: &lt;br/&gt;0: CONNECT_WIRELESS // Connected to [ap:{dev_mac}] with SSID {ssid_name}; &lt;br/&gt;1: DISCONNECT_WIRELESS // Disconnected from [ap:{dev_mac}]; &lt;br/&gt;2: ROAMING // Roaming from [ap:{old_dev_mac}] to [ap:{dev_mac}]; &lt;br/&gt;3: BLOCK_WIRELESS // Blocked by admin {admin_name}; &lt;br/&gt;100: CONNECT_WIRED // Connected to [{dev_type}:{dev_mac}]; &lt;br/&gt;101: DISCONNECT_WIRED // Disconnected from [{dev_type}:{dev_mac}]; &lt;br/&gt;102: BLOCK_WIRED // Blocked by admin {admin_name}; &lt;br/&gt;200: AUTH_SUCCESS // Authorized by {auth_type} with {username}; &lt;br/&gt;201: AUTH_FAIL // Failed Authorized by {auth_type} with {username}; &lt;br/&gt;202: UNBLOCK // Unblocked by admin {admin_name}; | [optional] 
**Time** | Pointer to **int64** | Event timestamp, unit: millisecond. | [optional] 

## Methods

### NewClientConnectionEvents

`func NewClientConnectionEvents() *ClientConnectionEvents`

NewClientConnectionEvents instantiates a new ClientConnectionEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConnectionEventsWithDefaults

`func NewClientConnectionEventsWithDefaults() *ClientConnectionEvents`

NewClientConnectionEventsWithDefaults instantiates a new ClientConnectionEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *ClientConnectionEvents) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ClientConnectionEvents) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ClientConnectionEvents) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ClientConnectionEvents) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetContent

`func (o *ClientConnectionEvents) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ClientConnectionEvents) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ClientConnectionEvents) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ClientConnectionEvents) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEventType

`func (o *ClientConnectionEvents) GetEventType() int32`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ClientConnectionEvents) GetEventTypeOk() (*int32, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ClientConnectionEvents) SetEventType(v int32)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ClientConnectionEvents) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetTime

`func (o *ClientConnectionEvents) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClientConnectionEvents) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClientConnectionEvents) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClientConnectionEvents) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


