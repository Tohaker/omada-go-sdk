# SendMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingCode** | Pointer to **string** | Calling code should contain 2 to 5 characters. Calling code must be entered when entering the country code. For the values of Calling code, refer to section 5.4.1 of the Open API Access Guide. | [optional] 
**Content** | Pointer to **string** | When parameter [type] is 0, parameter [content] should not be null. | [optional] 
**Receiver** | **string** | Receiver number. | 
**SimCard** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCard] shoud not be null.1: SIM1; 2: SIM2. | [optional] 
**Test** | Pointer to [**SimQuotaSetting**](SimQuotaSetting.md) |  | [optional] 
**Type** | **int32** | Send type. 0: formal; 1: test | 

## Methods

### NewSendMessage

`func NewSendMessage(receiver string, type_ int32, ) *SendMessage`

NewSendMessage instantiates a new SendMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageWithDefaults

`func NewSendMessageWithDefaults() *SendMessage`

NewSendMessageWithDefaults instantiates a new SendMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallingCode

`func (o *SendMessage) GetCallingCode() string`

GetCallingCode returns the CallingCode field if non-nil, zero value otherwise.

### GetCallingCodeOk

`func (o *SendMessage) GetCallingCodeOk() (*string, bool)`

GetCallingCodeOk returns a tuple with the CallingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingCode

`func (o *SendMessage) SetCallingCode(v string)`

SetCallingCode sets CallingCode field to given value.

### HasCallingCode

`func (o *SendMessage) HasCallingCode() bool`

HasCallingCode returns a boolean if a field has been set.

### GetContent

`func (o *SendMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SendMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SendMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SendMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetReceiver

`func (o *SendMessage) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *SendMessage) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *SendMessage) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetSimCard

`func (o *SendMessage) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *SendMessage) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *SendMessage) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *SendMessage) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.

### GetTest

`func (o *SendMessage) GetTest() SimQuotaSetting`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *SendMessage) GetTestOk() (*SimQuotaSetting, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *SendMessage) SetTest(v SimQuotaSetting)`

SetTest sets Test field to given value.

### HasTest

`func (o *SendMessage) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetType

`func (o *SendMessage) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendMessage) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendMessage) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


