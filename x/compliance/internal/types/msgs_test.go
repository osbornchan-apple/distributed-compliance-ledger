//nolint:testpackage
package types

//nolint:goimports
import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/integration_tests/constants"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewMsgCertifyModel(t *testing.T) {
	var msg = NewMsgCertifyModel(testconstants.VID, testconstants.PID, testconstants.CertificationDate,
		ZbCertificationType, testconstants.Reason, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "certify_model")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestMsgCertifyModelValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgCertifyModel
	}{
		{true, NewMsgCertifyModel(
			testconstants.VID, testconstants.PID, testconstants.CertificationDate,
			CertificationType(testconstants.CertificationType), testconstants.Reason, testconstants.Signer)},
		{false, NewMsgCertifyModel(
			testconstants.VID, 0, testconstants.CertificationDate,
			CertificationType(testconstants.CertificationType), testconstants.Reason, testconstants.Signer)},
		{false, NewMsgCertifyModel(
			testconstants.VID, 0, testconstants.CertificationDate,
			CertificationType(testconstants.CertificationType), testconstants.Reason, testconstants.Signer)},
		{false, NewMsgCertifyModel(
			testconstants.VID, testconstants.PID, time.Time{},
			CertificationType(testconstants.CertificationType), testconstants.Reason, testconstants.Signer)},
		{false, NewMsgCertifyModel(
			testconstants.VID, testconstants.PID, testconstants.CertificationDate,
			"", testconstants.Reason, testconstants.Signer)},
		{false, NewMsgCertifyModel(
			testconstants.VID, testconstants.PID, testconstants.CertificationDate,
			"Other Type", testconstants.Reason, testconstants.Signer)},
		{true, NewMsgCertifyModel(
			testconstants.VID, testconstants.PID, testconstants.CertificationDate,
			CertificationType(testconstants.CertificationType), "", testconstants.Signer)},
		{false, NewMsgCertifyModel(
			testconstants.VID, testconstants.PID, testconstants.CertificationDate,
			CertificationType(testconstants.CertificationType), testconstants.Reason, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsgCertifyModelGetSignBytes(t *testing.T) {
	var msg = NewMsgCertifyModel(testconstants.VID, testconstants.PID, testconstants.CertificationDate,
		CertificationType(testconstants.CertificationType), testconstants.EmptyString, testconstants.Signer)

	expected := `{"type":"compliance/CertifyModel","value":{"certification_date":"2020-01-01T00:00:00Z",` +
		`"certification_type":"zb","pid":22,"signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz","vid":1}}`

	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestNewMsgRevokeModel(t *testing.T) {
	var msg = NewMsgRevokeModel(testconstants.VID, testconstants.PID, testconstants.RevocationDate,
		CertificationType(testconstants.CertificationType), testconstants.RevocationReason, testconstants.Signer)

	require.Equal(t, msg.Route(), RouterKey)
	require.Equal(t, msg.Type(), "revoke_model")
	require.Equal(t, msg.GetSigners(), []sdk.AccAddress{testconstants.Signer})
}

func TestMsgRevokeModelValidation(t *testing.T) {
	cases := []struct {
		valid bool
		msg   MsgRevokeModel
	}{
		{true, NewMsgRevokeModel(
			testconstants.VID, testconstants.PID, testconstants.RevocationDate,
			CertificationType(testconstants.CertificationType), testconstants.RevocationReason, testconstants.Signer)},
		{false, NewMsgRevokeModel(
			0, testconstants.PID, testconstants.RevocationDate,
			CertificationType(testconstants.CertificationType), testconstants.RevocationReason, testconstants.Signer)},
		{false, NewMsgRevokeModel(
			testconstants.VID, 0, testconstants.RevocationDate,
			CertificationType(testconstants.CertificationType), testconstants.RevocationReason, testconstants.Signer)},
		{false, NewMsgRevokeModel(
			testconstants.VID, testconstants.PID, time.Time{},
			CertificationType(testconstants.CertificationType), testconstants.RevocationReason, testconstants.Signer)},
		{true, NewMsgRevokeModel(
			testconstants.VID, testconstants.PID, testconstants.RevocationDate,
			CertificationType(testconstants.CertificationType), "", testconstants.Signer)},
		{false, NewMsgRevokeModel(
			testconstants.VID, testconstants.PID, testconstants.RevocationDate,
			"", testconstants.RevocationReason, testconstants.Signer)},
		{false, NewMsgRevokeModel(
			testconstants.VID, testconstants.PID, testconstants.RevocationDate,
			CertificationType(testconstants.CertificationType), testconstants.RevocationReason, nil)},
	}

	for _, tc := range cases {
		err := tc.msg.ValidateBasic()

		if tc.valid {
			require.Nil(t, err)
		} else {
			require.NotNil(t, err)
		}
	}
}

func TestMsRevokeModelGetSignBytes(t *testing.T) {
	var msg = NewMsgRevokeModel(testconstants.VID, testconstants.PID, testconstants.RevocationDate,
		CertificationType(testconstants.CertificationType), testconstants.RevocationReason, testconstants.Signer)

	expected := `{"type":"compliance/RevokeModel","value":{"certification_type":"zb","pid":22,"reason":"Some Reason",` +
		`"revocation_date":"2020-03-03T03:30:00Z","signer":"cosmos1p72j8mgkf39qjzcmr283w8l8y9qv30qpj056uz","vid":1}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}
