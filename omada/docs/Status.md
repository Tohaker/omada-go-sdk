# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | To get status information from the router, send a message starting with \&quot;LTE Router Status\&quot;, followed by Password/PIN(e.g. LTE Router Status 1234). The password is required for viewing device-related Information via SMS. | [optional] 
**Content** | Pointer to [**StatusContent**](StatusContent.md) |  | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Status) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Status) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Status) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Status) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetContent

`func (o *Status) GetContent() StatusContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Status) GetContentOk() (*StatusContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Status) SetContent(v StatusContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *Status) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


