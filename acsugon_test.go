package acsugon_sdk

import (
	ct "acsugon_sdk/constant"
	"fmt"
	"os"
	"testing"
)

var (
	ACS *ACSugon
)

const (
	User  = "***"
	Pwd   = "***"
	OrgID = "***"
)

func init() {
	acs, err := NewACSugon(ct.EndpointProd, User, Pwd, OrgID)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}

	ACS = acs
}

func TestNewACSugon(t *testing.T) {
	t.Logf("NewACSugon: %+v", ACS)
}

func TestACSugon_Login(t *testing.T) {
	acs, err := ACS.Login()
	if err != nil {
		t.Error(err)
		return
	}

	for s, token := range acs.Tokens {
		t.Logf("id: %s, token: %+v", s, token)
	}
}
