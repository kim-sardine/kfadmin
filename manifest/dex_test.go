package manifest

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var (
	testStaticPassword = StaticPassword{
		Email:    "testEmail@test.test",
		Hash:     "testHash",
		Username: "testUsername",
		UserID:   "testUserId",
	}
	testDexDataConfig = DexDataConfig{
		Issuer:          "testIssuer",
		StaticPasswords: []StaticPassword{testStaticPassword},
		StaticClients:   []StatiClient{{ID: "testId"}},
	}
	testMarshaledDexDataConfig = "issuer: testIssuer\nstorage:\n  type: \"\"\n  config:\n    inCluster: false\nweb:\n  http: \"\"\nloggger:\n  level: \"\"\n  format: \"\"\noauth2:\n  skipApprovalScreen: false\nenablePasswordDB: false\nstaticPasswords:\n- email: testEmail@test.test\n  hash: testHash\n  username: testUsername\n  userID: testUserId\nstaticClients:\n- id: testId\n  name: \"\"\n  secret: \"\"\n"
)

func TestMarshalDexConfig(t *testing.T) {
	data, _ := MarshalDexDataConfig(&testDexDataConfig)

	assert.Equal(t, string(data), testMarshaledDexDataConfig)
}

func TestUnmarshalDexConfig(t *testing.T) {
	dc, _ := UnmarshalDexDataConfig(testMarshaledDexDataConfig)

	assert.Equal(t, cmp.Equal(*dc, testDexDataConfig), true)
}
