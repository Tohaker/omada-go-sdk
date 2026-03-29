# WireguardKeyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to **string** | The private key of WireGuard VPN. | [optional] 
**PublicKey** | Pointer to **string** | The public key of WireGuard VPN. | [optional] 

## Methods

### NewWireguardKeyOpenApiVO

`func NewWireguardKeyOpenApiVO() *WireguardKeyOpenApiVO`

NewWireguardKeyOpenApiVO instantiates a new WireguardKeyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardKeyOpenApiVOWithDefaults

`func NewWireguardKeyOpenApiVOWithDefaults() *WireguardKeyOpenApiVO`

NewWireguardKeyOpenApiVOWithDefaults instantiates a new WireguardKeyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *WireguardKeyOpenApiVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WireguardKeyOpenApiVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WireguardKeyOpenApiVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *WireguardKeyOpenApiVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardKeyOpenApiVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardKeyOpenApiVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardKeyOpenApiVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardKeyOpenApiVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


