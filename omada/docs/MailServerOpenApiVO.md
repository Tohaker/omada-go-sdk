# MailServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **bool** | Enable this feature if the login of the mailbox requires a username and authorization code. | 
**AuthCode** | Pointer to **string** | When Authentication is enabled, enter the authorization code that enables a third party to log in to the mailbox. Note that the authorization code is not the mailbox&#39;s password. | [optional] 
**Id** | Pointer to **string** | Mail server ID. | [optional] 
**Receiver** | **string** | The email address of the receiver, which can be the same as or different from the sender&#39;s email address. | 
**Resource** | Pointer to **int32** | Mail server setting creation resource, such as: 0: new created, 1: from template, 2: override. | [optional] 
**Sender** | **string** | The email address of the sender. | 
**SmtpPort** | Pointer to **int32** | Enter the port used by the SMTP server according to the instructions of your email service provider. | [optional] 
**SmtpServer** | Pointer to **string** | Enter the domain name or IP address of the SMTP server. | [optional] 
**Ssl** | **bool** | Enable this feature, and the data will be transmitted based on the SSL protocol. | 
**Username** | Pointer to **string** | When Authentication is enabled, enter your email address as the username. | [optional] 

## Methods

### NewMailServerOpenApiVO

`func NewMailServerOpenApiVO(auth bool, receiver string, sender string, ssl bool, ) *MailServerOpenApiVO`

NewMailServerOpenApiVO instantiates a new MailServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailServerOpenApiVOWithDefaults

`func NewMailServerOpenApiVOWithDefaults() *MailServerOpenApiVO`

NewMailServerOpenApiVOWithDefaults instantiates a new MailServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *MailServerOpenApiVO) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *MailServerOpenApiVO) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *MailServerOpenApiVO) SetAuth(v bool)`

SetAuth sets Auth field to given value.


### GetAuthCode

`func (o *MailServerOpenApiVO) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *MailServerOpenApiVO) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *MailServerOpenApiVO) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *MailServerOpenApiVO) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetId

`func (o *MailServerOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MailServerOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MailServerOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MailServerOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReceiver

`func (o *MailServerOpenApiVO) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *MailServerOpenApiVO) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *MailServerOpenApiVO) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetResource

`func (o *MailServerOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MailServerOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MailServerOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MailServerOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSender

`func (o *MailServerOpenApiVO) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MailServerOpenApiVO) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MailServerOpenApiVO) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSmtpPort

`func (o *MailServerOpenApiVO) GetSmtpPort() int32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *MailServerOpenApiVO) GetSmtpPortOk() (*int32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *MailServerOpenApiVO) SetSmtpPort(v int32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *MailServerOpenApiVO) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpServer

`func (o *MailServerOpenApiVO) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *MailServerOpenApiVO) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *MailServerOpenApiVO) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *MailServerOpenApiVO) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetSsl

`func (o *MailServerOpenApiVO) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *MailServerOpenApiVO) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *MailServerOpenApiVO) SetSsl(v bool)`

SetSsl sets Ssl field to given value.


### GetUsername

`func (o *MailServerOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MailServerOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MailServerOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MailServerOpenApiVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


