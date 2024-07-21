// CODE GENERATED, DO NOT EDIT

package S3

import (
	"context"
	"encoding/xml"
	"fmt"
	do "github.com/hoomy-official/go-shared/pkg/net/do"
	opts "github.com/hoomy-official/go-sonos/opts"
	"net/url"
)

const (
	// namespaceSystemPropertiesService is the UPnP namespace for the Sonos service.
	namespaceSystemPropertiesService = "urn:schemas-upnp-org:service:SystemProperties:1"

	// eventSystemPropertiesService is the event path for the Sonos event service.
	eventSystemPropertiesService = "/SystemProperties/Event"

	// controlSystemPropertiesService is the action path for the Sonos service.
	controlSystemPropertiesService = "/SystemProperties/Control"
)

// SystemPropertiesService is based on UPnP and wraps a SOAP client for communicating with Sonos services.
type SystemPropertiesService struct {
	doer do.Doer
}

// NewSystemPropertiesService initializes a new SystemPropertiesService with the provided Doer.
func NewSystemPropertiesService(doer do.Doer) *SystemPropertiesService {
	return &SystemPropertiesService{doer: doer}
}

// SubscribeEvent subscribes for events from the Sonos Event service using the provided callback URL.
func (s SystemPropertiesService) SubscribeEvent(ctx context.Context, callbackURL *url.URL) error {
	return s.doer.Do(ctx, do.WithPath(eventSystemPropertiesService), do.WithMethod("SUBSCRIBE"), do.WithExtraHeaderf("NT", "upnp:event"), do.WithExtraHeaderf("CALLBACK", fmt.Sprintf("<%s>", callbackURL.String())), do.WithExtraHeaderf("TIMEOUT", "Second-1800"))
}

// SetString Save a string in the system
func (s SystemPropertiesService) SetString(ctx context.Context, variableName VariableName, stringValue VariableStringValue) error {
	type SetStringReq struct {
		VariableName VariableName
		StringValue  VariableStringValue
	}

	type SetStringNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetStringReq
	}

	type SetStringBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		SetStringNode SetStringNode `xml:"u:SetString,omitempty"`
	}

	body := SetStringBody{SetStringNode: SetStringNode{
		SetStringReq: SetStringReq{
			StringValue:  stringValue,
			VariableName: variableName,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "SetString", body, nil)...)
}

type SystemPropertiesGetStringRes struct {
	StringValue VariableStringValue
}

// GetString Get a saved string.
func (s SystemPropertiesService) GetString(ctx context.Context, variableName VariableName, getString *SystemPropertiesGetStringRes) error {
	type GetStringReq struct {
		VariableName VariableName
	}

	type GetStringNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetStringReq
	}

	type GetStringBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		GetStringNode GetStringNode `xml:"u:GetString,omitempty"`
	}

	body := GetStringBody{GetStringNode: GetStringNode{
		GetStringReq: GetStringReq{VariableName: variableName},
		Xmlns:        namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "GetString", body, getString)...)
}

// Remove Remove a saved string
func (s SystemPropertiesService) Remove(ctx context.Context, variableName VariableName) error {
	type RemoveReq struct {
		VariableName VariableName
	}

	type RemoveNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveReq
	}

	type RemoveBody struct {
		XMLName    xml.Name   `xml:"s:Body"`
		RemoveNode RemoveNode `xml:"u:Remove,omitempty"`
	}

	body := RemoveBody{RemoveNode: RemoveNode{
		RemoveReq: RemoveReq{VariableName: variableName},
		Xmlns:     namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "Remove", body, nil)...)
}

type SystemPropertiesGetWebCodeRes struct {
	WebCode VariableStringValue
}

func (s SystemPropertiesService) GetWebCode(ctx context.Context, accountType AccountType, getWebCode *SystemPropertiesGetWebCodeRes) error {
	type GetWebCodeReq struct {
		AccountType AccountType
	}

	type GetWebCodeNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetWebCodeReq
	}

	type GetWebCodeBody struct {
		XMLName        xml.Name       `xml:"s:Body"`
		GetWebCodeNode GetWebCodeNode `xml:"u:GetWebCode,omitempty"`
	}

	body := GetWebCodeBody{GetWebCodeNode: GetWebCodeNode{
		GetWebCodeReq: GetWebCodeReq{AccountType: accountType},
		Xmlns:         namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "GetWebCode", body, getWebCode)...)
}

type SystemPropertiesProvisionCredentialedTrialAccountXRes struct {
	IsExpired  IsExpired
	AccountUDN AccountUDN
}

func (s SystemPropertiesService) ProvisionCredentialedTrialAccountX(ctx context.Context, accountType AccountType, accountId AccountID, accountPassword AccountPassword, provisionCredentialedTrialAccountX *SystemPropertiesProvisionCredentialedTrialAccountXRes) error {
	type ProvisionCredentialedTrialAccountXReq struct {
		AccountType     AccountType
		AccountID       AccountID
		AccountPassword AccountPassword
	}

	type ProvisionCredentialedTrialAccountXNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ProvisionCredentialedTrialAccountXReq
	}

	type ProvisionCredentialedTrialAccountXBody struct {
		XMLName                                xml.Name                               `xml:"s:Body"`
		ProvisionCredentialedTrialAccountXNode ProvisionCredentialedTrialAccountXNode `xml:"u:ProvisionCredentialedTrialAccountX,omitempty"`
	}

	body := ProvisionCredentialedTrialAccountXBody{ProvisionCredentialedTrialAccountXNode: ProvisionCredentialedTrialAccountXNode{
		ProvisionCredentialedTrialAccountXReq: ProvisionCredentialedTrialAccountXReq{
			AccountID:       accountId,
			AccountPassword: accountPassword,
			AccountType:     accountType,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "ProvisionCredentialedTrialAccountX", body, provisionCredentialedTrialAccountX)...)
}

type SystemPropertiesAddAccountXRes struct {
	AccountUDN AccountUDN
}

func (s SystemPropertiesService) AddAccountX(ctx context.Context, accountType AccountType, accountId AccountID, accountPassword AccountPassword, addAccountX *SystemPropertiesAddAccountXRes) error {
	type AddAccountXReq struct {
		AccountType     AccountType
		AccountID       AccountID
		AccountPassword AccountPassword
	}

	type AddAccountXNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddAccountXReq
	}

	type AddAccountXBody struct {
		XMLName         xml.Name        `xml:"s:Body"`
		AddAccountXNode AddAccountXNode `xml:"u:AddAccountX,omitempty"`
	}

	body := AddAccountXBody{AddAccountXNode: AddAccountXNode{
		AddAccountXReq: AddAccountXReq{
			AccountID:       accountId,
			AccountPassword: accountPassword,
			AccountType:     accountType,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "AddAccountX", body, addAccountX)...)
}

type SystemPropertiesAddOAuthAccountXRes struct {
	AccountUDN      AccountUDN
	AccountNickname AccountNickname
}

func (s SystemPropertiesService) AddOAuthAccountX(ctx context.Context, accountType AccountType, accountToken AccountCredential, accountKey AccountCredential, oauthDeviceId OAuthDeviceID, authorizationCode AuthorizationCode, redirectUri RedirectURI, userIdHashCode UserIdHashCode, accountTier AccountTier, addOauthAccountX *SystemPropertiesAddOAuthAccountXRes) error {
	type AddOAuthAccountXReq struct {
		AccountType       AccountType
		AccountToken      AccountCredential
		AccountKey        AccountCredential
		OAuthDeviceID     OAuthDeviceID
		AuthorizationCode AuthorizationCode
		RedirectURI       RedirectURI
		UserIdHashCode    UserIdHashCode
		AccountTier       AccountTier
	}

	type AddOAuthAccountXNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		AddOAuthAccountXReq
	}

	type AddOAuthAccountXBody struct {
		XMLName              xml.Name             `xml:"s:Body"`
		AddOAuthAccountXNode AddOAuthAccountXNode `xml:"u:AddOAuthAccountX,omitempty"`
	}

	body := AddOAuthAccountXBody{AddOAuthAccountXNode: AddOAuthAccountXNode{
		AddOAuthAccountXReq: AddOAuthAccountXReq{
			AccountKey:        accountKey,
			AccountTier:       accountTier,
			AccountToken:      accountToken,
			AccountType:       accountType,
			AuthorizationCode: authorizationCode,
			OAuthDeviceID:     oauthDeviceId,
			RedirectURI:       redirectUri,
			UserIdHashCode:    userIdHashCode,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "AddOAuthAccountX", body, addOauthAccountX)...)
}

func (s SystemPropertiesService) RemoveAccount(ctx context.Context, accountType AccountType, accountId AccountID) error {
	type RemoveAccountReq struct {
		AccountType AccountType
		AccountID   AccountID
	}

	type RemoveAccountNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RemoveAccountReq
	}

	type RemoveAccountBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		RemoveAccountNode RemoveAccountNode `xml:"u:RemoveAccount,omitempty"`
	}

	body := RemoveAccountBody{RemoveAccountNode: RemoveAccountNode{
		RemoveAccountReq: RemoveAccountReq{
			AccountID:   accountId,
			AccountType: accountType,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "RemoveAccount", body, nil)...)
}

func (s SystemPropertiesService) EditAccountPasswordX(ctx context.Context, accountType AccountType, accountId AccountID, newAccountPassword AccountPassword) error {
	type EditAccountPasswordXReq struct {
		AccountType        AccountType
		AccountID          AccountID
		NewAccountPassword AccountPassword
	}

	type EditAccountPasswordXNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		EditAccountPasswordXReq
	}

	type EditAccountPasswordXBody struct {
		XMLName                  xml.Name                 `xml:"s:Body"`
		EditAccountPasswordXNode EditAccountPasswordXNode `xml:"u:EditAccountPasswordX,omitempty"`
	}

	body := EditAccountPasswordXBody{EditAccountPasswordXNode: EditAccountPasswordXNode{
		EditAccountPasswordXReq: EditAccountPasswordXReq{
			AccountID:          accountId,
			AccountType:        accountType,
			NewAccountPassword: newAccountPassword,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "EditAccountPasswordX", body, nil)...)
}

func (s SystemPropertiesService) SetAccountNicknameX(ctx context.Context, accountUdn AccountUDN, accountNickname AccountNickname) error {
	type SetAccountNicknameXReq struct {
		AccountUDN      AccountUDN
		AccountNickname AccountNickname
	}

	type SetAccountNicknameXNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		SetAccountNicknameXReq
	}

	type SetAccountNicknameXBody struct {
		XMLName                 xml.Name                `xml:"s:Body"`
		SetAccountNicknameXNode SetAccountNicknameXNode `xml:"u:SetAccountNicknameX,omitempty"`
	}

	body := SetAccountNicknameXBody{SetAccountNicknameXNode: SetAccountNicknameXNode{
		SetAccountNicknameXReq: SetAccountNicknameXReq{
			AccountNickname: accountNickname,
			AccountUDN:      accountUdn,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "SetAccountNicknameX", body, nil)...)
}

func (s SystemPropertiesService) RefreshAccountCredentialsX(ctx context.Context, accountType AccountType, accountUid AccountUID, accountToken AccountCredential, accountKey AccountCredential) error {
	type RefreshAccountCredentialsXReq struct {
		AccountType  AccountType
		AccountUID   AccountUID
		AccountToken AccountCredential
		AccountKey   AccountCredential
	}

	type RefreshAccountCredentialsXNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		RefreshAccountCredentialsXReq
	}

	type RefreshAccountCredentialsXBody struct {
		XMLName                        xml.Name                       `xml:"s:Body"`
		RefreshAccountCredentialsXNode RefreshAccountCredentialsXNode `xml:"u:RefreshAccountCredentialsX,omitempty"`
	}

	body := RefreshAccountCredentialsXBody{RefreshAccountCredentialsXNode: RefreshAccountCredentialsXNode{
		RefreshAccountCredentialsXReq: RefreshAccountCredentialsXReq{
			AccountKey:   accountKey,
			AccountToken: accountToken,
			AccountType:  accountType,
			AccountUID:   accountUid,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "RefreshAccountCredentialsX", body, nil)...)
}

func (s SystemPropertiesService) EditAccountMd(ctx context.Context, accountType AccountType, accountId AccountID, newAccountMd AccountMd) error {
	type EditAccountMdReq struct {
		AccountType  AccountType
		AccountID    AccountID
		NewAccountMd AccountMd
	}

	type EditAccountMdNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		EditAccountMdReq
	}

	type EditAccountMdBody struct {
		XMLName           xml.Name          `xml:"s:Body"`
		EditAccountMdNode EditAccountMdNode `xml:"u:EditAccountMd,omitempty"`
	}

	body := EditAccountMdBody{EditAccountMdNode: EditAccountMdNode{
		EditAccountMdReq: EditAccountMdReq{
			AccountID:    accountId,
			AccountType:  accountType,
			NewAccountMd: newAccountMd,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "EditAccountMd", body, nil)...)
}

func (s SystemPropertiesService) DoPostUpdateTasks(ctx context.Context) error {
	type DoPostUpdateTasksReq struct{}

	type DoPostUpdateTasksNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		DoPostUpdateTasksReq
	}

	type DoPostUpdateTasksBody struct {
		XMLName               xml.Name              `xml:"s:Body"`
		DoPostUpdateTasksNode DoPostUpdateTasksNode `xml:"u:DoPostUpdateTasks,omitempty"`
	}

	body := DoPostUpdateTasksBody{DoPostUpdateTasksNode: DoPostUpdateTasksNode{
		DoPostUpdateTasksReq: DoPostUpdateTasksReq{},
		Xmlns:                namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "DoPostUpdateTasks", body, nil)...)
}

func (s SystemPropertiesService) ResetThirdPartyCredentials(ctx context.Context) error {
	type ResetThirdPartyCredentialsReq struct{}

	type ResetThirdPartyCredentialsNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ResetThirdPartyCredentialsReq
	}

	type ResetThirdPartyCredentialsBody struct {
		XMLName                        xml.Name                       `xml:"s:Body"`
		ResetThirdPartyCredentialsNode ResetThirdPartyCredentialsNode `xml:"u:ResetThirdPartyCredentials,omitempty"`
	}

	body := ResetThirdPartyCredentialsBody{ResetThirdPartyCredentialsNode: ResetThirdPartyCredentialsNode{
		ResetThirdPartyCredentialsReq: ResetThirdPartyCredentialsReq{},
		Xmlns:                         namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "ResetThirdPartyCredentials", body, nil)...)
}

func (s SystemPropertiesService) EnableRDM(ctx context.Context, rdmvalue RDMEnabled) error {
	type EnableRDMReq struct {
		RDMValue RDMEnabled
	}

	type EnableRDMNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		EnableRDMReq
	}

	type EnableRDMBody struct {
		XMLName       xml.Name      `xml:"s:Body"`
		EnableRDMNode EnableRDMNode `xml:"u:EnableRDM,omitempty"`
	}

	body := EnableRDMBody{EnableRDMNode: EnableRDMNode{
		EnableRDMReq: EnableRDMReq{RDMValue: rdmvalue},
		Xmlns:        namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "EnableRDM", body, nil)...)
}

type SystemPropertiesGetRDMRes struct {
	RDMValue RDMEnabled
}

func (s SystemPropertiesService) GetRDM(ctx context.Context, getRdm *SystemPropertiesGetRDMRes) error {
	type GetRDMReq struct{}

	type GetRDMNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		GetRDMReq
	}

	type GetRDMBody struct {
		XMLName    xml.Name   `xml:"s:Body"`
		GetRDMNode GetRDMNode `xml:"u:GetRDM,omitempty"`
	}

	body := GetRDMBody{GetRDMNode: GetRDMNode{
		GetRDMReq: GetRDMReq{},
		Xmlns:     namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "GetRDM", body, getRdm)...)
}

type SystemPropertiesReplaceAccountXRes struct {
	NewAccountUDN AccountUDN
}

func (s SystemPropertiesService) ReplaceAccountX(ctx context.Context, accountUdn AccountUDN, newAccountId AccountID, newAccountPassword AccountPassword, accountToken AccountCredential, accountKey AccountCredential, oauthDeviceId OAuthDeviceID, replaceAccountX *SystemPropertiesReplaceAccountXRes) error {
	type ReplaceAccountXReq struct {
		AccountUDN         AccountUDN
		NewAccountID       AccountID
		NewAccountPassword AccountPassword
		AccountToken       AccountCredential
		AccountKey         AccountCredential
		OAuthDeviceID      OAuthDeviceID
	}

	type ReplaceAccountXNode struct {
		Xmlns string `xml:"xmlns:u,attr"`
		ReplaceAccountXReq
	}

	type ReplaceAccountXBody struct {
		XMLName             xml.Name            `xml:"s:Body"`
		ReplaceAccountXNode ReplaceAccountXNode `xml:"u:ReplaceAccountX,omitempty"`
	}

	body := ReplaceAccountXBody{ReplaceAccountXNode: ReplaceAccountXNode{
		ReplaceAccountXReq: ReplaceAccountXReq{
			AccountKey:         accountKey,
			AccountToken:       accountToken,
			AccountUDN:         accountUdn,
			NewAccountID:       newAccountId,
			NewAccountPassword: newAccountPassword,
			OAuthDeviceID:      oauthDeviceId,
		},
		Xmlns: namespaceSystemPropertiesService,
	}}

	return s.doer.Do(
		ctx, opts.WithSonosUPnPOptions(controlSystemPropertiesService, namespaceSystemPropertiesService, "ReplaceAccountX", body, replaceAccountX)...)
}
