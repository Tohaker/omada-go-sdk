# SpeedTestCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildMac** | **string** | The parameter [childMac] can be the MAC address of the Main AP or Client AP. Ensure that the two devices to be tested are in a Main-Client relationship. | 

## Methods

### NewSpeedTestCommand

`func NewSpeedTestCommand(childMac string, ) *SpeedTestCommand`

NewSpeedTestCommand instantiates a new SpeedTestCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeedTestCommandWithDefaults

`func NewSpeedTestCommandWithDefaults() *SpeedTestCommand`

NewSpeedTestCommandWithDefaults instantiates a new SpeedTestCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildMac

`func (o *SpeedTestCommand) GetChildMac() string`

GetChildMac returns the ChildMac field if non-nil, zero value otherwise.

### GetChildMacOk

`func (o *SpeedTestCommand) GetChildMacOk() (*string, bool)`

GetChildMacOk returns a tuple with the ChildMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildMac

`func (o *SpeedTestCommand) SetChildMac(v string)`

SetChildMac sets ChildMac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


