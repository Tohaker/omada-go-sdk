# ApBridgeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeSsidName** | **string** | Bridge SSID name. It should contain 1 to 32 UTF-8 characters. | 
**BridgeSsidPassword** | **string** | Bridge SSID password. It should contain 8-63 printable ASCII characters. | 

## Methods

### NewApBridgeConfig

`func NewApBridgeConfig(bridgeSsidName string, bridgeSsidPassword string, ) *ApBridgeConfig`

NewApBridgeConfig instantiates a new ApBridgeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBridgeConfigWithDefaults

`func NewApBridgeConfigWithDefaults() *ApBridgeConfig`

NewApBridgeConfigWithDefaults instantiates a new ApBridgeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeSsidName

`func (o *ApBridgeConfig) GetBridgeSsidName() string`

GetBridgeSsidName returns the BridgeSsidName field if non-nil, zero value otherwise.

### GetBridgeSsidNameOk

`func (o *ApBridgeConfig) GetBridgeSsidNameOk() (*string, bool)`

GetBridgeSsidNameOk returns a tuple with the BridgeSsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeSsidName

`func (o *ApBridgeConfig) SetBridgeSsidName(v string)`

SetBridgeSsidName sets BridgeSsidName field to given value.


### GetBridgeSsidPassword

`func (o *ApBridgeConfig) GetBridgeSsidPassword() string`

GetBridgeSsidPassword returns the BridgeSsidPassword field if non-nil, zero value otherwise.

### GetBridgeSsidPasswordOk

`func (o *ApBridgeConfig) GetBridgeSsidPasswordOk() (*string, bool)`

GetBridgeSsidPasswordOk returns a tuple with the BridgeSsidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeSsidPassword

`func (o *ApBridgeConfig) SetBridgeSsidPassword(v string)`

SetBridgeSsidPassword sets BridgeSsidPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


