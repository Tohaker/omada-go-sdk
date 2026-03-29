# CallForwardingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | **int32** | Forwarding condition. 0- no answer, 1- unconditional | 
**DestNumber** | **string** | The Destination Telephone Number that incoming calls will be redirected to. | 
**Enable** | **bool** | Enable this rule or not | 
**ForwardVia** | Pointer to **string** | \&quot;Auto\&quot; by default. | [optional] 
**FromNumbers** | Pointer to **[]string** | Any incoming calls from these numbers will be forwarded. | [optional] 
**FromPersons** | Pointer to **[]string** | Any incoming calls from these contacts will be forwarded. | [optional] 
**OmadacId** | Pointer to **string** | Omadac ID | [optional] 
**RuleId** | Pointer to **string** | Call forwarding rule id | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**ToDevices** | Pointer to **[]int32** | Any incoming calls to these devices will be forwarded. | [optional] 
**ToNumbers** | Pointer to **[]string** | Any incoming calls to these numbers will be forwarded. | [optional] 
**Type** | **int32** | The call type to be forwarded. 0: All Incoming Calls - If this option is selected, all incoming calls will be forwarded. 1: Calls to the Telephone Number - If this option is selected, select a telephone number from the list. Any incoming calls to this number will be forwarded. 2: Calls to the Phone - If this option is selected, select a telephony device from the list. Any incoming calls to this device will be forwarded. 3: Calls from a Person in the Telephone Book - If this option is selected, select a contact from the list. Any incoming calls from this contact will be forwarded. 4: Calls from the Telephone Number - If this option is selected, enter a specific telephone number. Any incoming calls from this number will be forwarded. | 

## Methods

### NewCallForwardingRule

`func NewCallForwardingRule(condition int32, destNumber string, enable bool, type_ int32, ) *CallForwardingRule`

NewCallForwardingRule instantiates a new CallForwardingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallForwardingRuleWithDefaults

`func NewCallForwardingRuleWithDefaults() *CallForwardingRule`

NewCallForwardingRuleWithDefaults instantiates a new CallForwardingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *CallForwardingRule) GetCondition() int32`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CallForwardingRule) GetConditionOk() (*int32, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CallForwardingRule) SetCondition(v int32)`

SetCondition sets Condition field to given value.


### GetDestNumber

`func (o *CallForwardingRule) GetDestNumber() string`

GetDestNumber returns the DestNumber field if non-nil, zero value otherwise.

### GetDestNumberOk

`func (o *CallForwardingRule) GetDestNumberOk() (*string, bool)`

GetDestNumberOk returns a tuple with the DestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestNumber

`func (o *CallForwardingRule) SetDestNumber(v string)`

SetDestNumber sets DestNumber field to given value.


### GetEnable

`func (o *CallForwardingRule) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CallForwardingRule) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CallForwardingRule) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetForwardVia

`func (o *CallForwardingRule) GetForwardVia() string`

GetForwardVia returns the ForwardVia field if non-nil, zero value otherwise.

### GetForwardViaOk

`func (o *CallForwardingRule) GetForwardViaOk() (*string, bool)`

GetForwardViaOk returns a tuple with the ForwardVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardVia

`func (o *CallForwardingRule) SetForwardVia(v string)`

SetForwardVia sets ForwardVia field to given value.

### HasForwardVia

`func (o *CallForwardingRule) HasForwardVia() bool`

HasForwardVia returns a boolean if a field has been set.

### GetFromNumbers

`func (o *CallForwardingRule) GetFromNumbers() []string`

GetFromNumbers returns the FromNumbers field if non-nil, zero value otherwise.

### GetFromNumbersOk

`func (o *CallForwardingRule) GetFromNumbersOk() (*[]string, bool)`

GetFromNumbersOk returns a tuple with the FromNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNumbers

`func (o *CallForwardingRule) SetFromNumbers(v []string)`

SetFromNumbers sets FromNumbers field to given value.

### HasFromNumbers

`func (o *CallForwardingRule) HasFromNumbers() bool`

HasFromNumbers returns a boolean if a field has been set.

### GetFromPersons

`func (o *CallForwardingRule) GetFromPersons() []string`

GetFromPersons returns the FromPersons field if non-nil, zero value otherwise.

### GetFromPersonsOk

`func (o *CallForwardingRule) GetFromPersonsOk() (*[]string, bool)`

GetFromPersonsOk returns a tuple with the FromPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPersons

`func (o *CallForwardingRule) SetFromPersons(v []string)`

SetFromPersons sets FromPersons field to given value.

### HasFromPersons

`func (o *CallForwardingRule) HasFromPersons() bool`

HasFromPersons returns a boolean if a field has been set.

### GetOmadacId

`func (o *CallForwardingRule) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *CallForwardingRule) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *CallForwardingRule) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *CallForwardingRule) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetRuleId

`func (o *CallForwardingRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *CallForwardingRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *CallForwardingRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *CallForwardingRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSiteId

`func (o *CallForwardingRule) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CallForwardingRule) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CallForwardingRule) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CallForwardingRule) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetToDevices

`func (o *CallForwardingRule) GetToDevices() []int32`

GetToDevices returns the ToDevices field if non-nil, zero value otherwise.

### GetToDevicesOk

`func (o *CallForwardingRule) GetToDevicesOk() (*[]int32, bool)`

GetToDevicesOk returns a tuple with the ToDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDevices

`func (o *CallForwardingRule) SetToDevices(v []int32)`

SetToDevices sets ToDevices field to given value.

### HasToDevices

`func (o *CallForwardingRule) HasToDevices() bool`

HasToDevices returns a boolean if a field has been set.

### GetToNumbers

`func (o *CallForwardingRule) GetToNumbers() []string`

GetToNumbers returns the ToNumbers field if non-nil, zero value otherwise.

### GetToNumbersOk

`func (o *CallForwardingRule) GetToNumbersOk() (*[]string, bool)`

GetToNumbersOk returns a tuple with the ToNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToNumbers

`func (o *CallForwardingRule) SetToNumbers(v []string)`

SetToNumbers sets ToNumbers field to given value.

### HasToNumbers

`func (o *CallForwardingRule) HasToNumbers() bool`

HasToNumbers returns a boolean if a field has been set.

### GetType

`func (o *CallForwardingRule) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CallForwardingRule) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CallForwardingRule) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


