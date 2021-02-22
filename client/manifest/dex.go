package manifest

import "gopkg.in/yaml.v2"

// DexDataConfig TBU
type DexDataConfig struct {
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
	EnablePasswordDB bool             `yaml:"enablePasswordDB"`
	StaticPasswords  []StaticPassword `yaml:"staticPasswords"`
	StaticClients    []StatiClient    `yaml:"staticClients"`
}

// StaticPassword TBU
type StaticPassword struct {
	Email    string `yaml:"email"`
	Hash     string `yaml:"hash"`
	Username string `yaml:"username"`
	UserID   string `yaml:"userID"`
}

// StatiClient TBU
type StatiClient struct {
	ID           string   `yaml:"id"`
	RedirectURIs []string `yaml:"redirectURIs,omitempty"`
	Name         string   `yaml:"name"`
	Secret       string   `yaml:"secret"`
}

// UnmarshalDexDataConfig TBU
func UnmarshalDexDataConfig(data string) (*DexDataConfig, error) {
	var dc *DexDataConfig = &DexDataConfig{}
	err := yaml.Unmarshal([]byte(data), dc)
	if err != nil {
		return nil, err
	}
	return dc, nil
}

// MarshalDexDataConfig TBU
func MarshalDexDataConfig(dc *DexDataConfig) (string, error) {
	data, err := yaml.Marshal(dc)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
