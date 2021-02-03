package manifest

import "gopkg.in/yaml.v2"

// DexConfigManifest TBU
type DexConfigManifest struct {
	Issuer  string `yaml:"issuer"`
	Storage struct {
		Type   string `yaml:"type"`
		Config struct {
			InCluster bool `yaml:"inCluster"`
		} `yaml:"config"`
	} `yaml:"storage"`
	Web struct {
		HTTP string `yaml:"http"`
	} `yaml:"web"`
	Logger struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"loggger"`
	Oauth2 struct {
		SkipApprovalScreen bool `yaml:"skipApprovalScreen"`
	} `yaml:"oauth2"`
	EnablePasswordDB bool                     `yaml:"enablePasswordDB"`
	StaticPasswords  []StaticPasswordManifest `yaml:"staticPasswords"`
	StaticClients    []struct {
		ID           string   `yaml:"id"`
		RedirectURIs []string `yaml:"redirectURIs"`
		Name         string   `yaml:"name"`
		Secret       string   `yaml:"secret"`
	} `yaml:"staticClients"`
}

// StaticPasswordManifest TBU
type StaticPasswordManifest struct {
	Email    string `yaml:"email"`
	Hash     string `yaml:"hash"`
	Username string `yaml:"username"`
	UserID   string `yaml:"userID"`
}

// UnmarshalDexConfig TBU
func UnmarshalDexConfig(data string) DexConfigManifest {
	var dc DexConfigManifest
	err := yaml.Unmarshal([]byte(data), &dc)
	if err != nil {
		panic(err)
	}
	return dc
}

// MarshalDexConfig TBU
func MarshalDexConfig(dc DexConfigManifest) string {
	data, err := yaml.Marshal(&dc)
	if err != nil {
		panic(err)
	}
	return string(data)
}
