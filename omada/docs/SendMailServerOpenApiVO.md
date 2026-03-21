# SendMailServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthEnable** | **bool** | Whether the SMTP email server uses basic authentication. | 
**Password** | Pointer to **string** | The user&#39;s email password used for authentication. | [optional] 
**Port** | **int32** | The port of the SMTP mailbox server,xxxx  should be within the range of 1-65535 | 
**Receiver** | Pointer to **string** | The email address where the message will be received. | [optional] 
**SenderAddress** | Pointer to **string** | The email address from which the message was sent. | [optional] 
**SmtpEnable** | Pointer to **bool** | Enable SMTP mail server | [optional] 
**SmtpServer** | **string** | The domain name of the SMTP mailbox server. | 
**SslEnable** | **bool** | Whether the SMTP mailbox server uses SSL encryption. | 
**Username** | Pointer to **string** | The user&#39;s email address used for authentication. | [optional] 

## Methods

### NewSendMailServerOpenApiVO

`func NewSendMailServerOpenApiVO(authEnable bool, port int32, smtpServer string, sslEnable bool, ) *SendMailServerOpenApiVO`

NewSendMailServerOpenApiVO instantiates a new SendMailServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMailServerOpenApiVOWithDefaults

`func NewSendMailServerOpenApiVOWithDefaults() *SendMailServerOpenApiVO`

NewSendMailServerOpenApiVOWithDefaults instantiates a new SendMailServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthEnable

`func (o *SendMailServerOpenApiVO) GetAuthEnable() bool`

GetAuthEnable returns the AuthEnable field if non-nil, zero value otherwise.

### GetAuthEnableOk

`func (o *SendMailServerOpenApiVO) GetAuthEnableOk() (*bool, bool)`

GetAuthEnableOk returns a tuple with the AuthEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEnable

`func (o *SendMailServerOpenApiVO) SetAuthEnable(v bool)`

SetAuthEnable sets AuthEnable field to given value.


### GetPassword

`func (o *SendMailServerOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SendMailServerOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SendMailServerOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SendMailServerOpenApiVO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *SendMailServerOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SendMailServerOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SendMailServerOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.


### GetReceiver

`func (o *SendMailServerOpenApiVO) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *SendMailServerOpenApiVO) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *SendMailServerOpenApiVO) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *SendMailServerOpenApiVO) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetSenderAddress

`func (o *SendMailServerOpenApiVO) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *SendMailServerOpenApiVO) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *SendMailServerOpenApiVO) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *SendMailServerOpenApiVO) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### GetSmtpEnable

`func (o *SendMailServerOpenApiVO) GetSmtpEnable() bool`

GetSmtpEnable returns the SmtpEnable field if non-nil, zero value otherwise.

### GetSmtpEnableOk

`func (o *SendMailServerOpenApiVO) GetSmtpEnableOk() (*bool, bool)`

GetSmtpEnableOk returns a tuple with the SmtpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpEnable

`func (o *SendMailServerOpenApiVO) SetSmtpEnable(v bool)`

SetSmtpEnable sets SmtpEnable field to given value.

### HasSmtpEnable

`func (o *SendMailServerOpenApiVO) HasSmtpEnable() bool`

HasSmtpEnable returns a boolean if a field has been set.

### GetSmtpServer

`func (o *SendMailServerOpenApiVO) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *SendMailServerOpenApiVO) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *SendMailServerOpenApiVO) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.


### GetSslEnable

`func (o *SendMailServerOpenApiVO) GetSslEnable() bool`

GetSslEnable returns the SslEnable field if non-nil, zero value otherwise.

### GetSslEnableOk

`func (o *SendMailServerOpenApiVO) GetSslEnableOk() (*bool, bool)`

GetSslEnableOk returns a tuple with the SslEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnable

`func (o *SendMailServerOpenApiVO) SetSslEnable(v bool)`

SetSslEnable sets SslEnable field to given value.


### GetUsername

`func (o *SendMailServerOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SendMailServerOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SendMailServerOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SendMailServerOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


