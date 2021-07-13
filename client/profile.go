package client

import (
	"context"

	"github.com/kim-sardine/kfadmin/manifest"
)

// Kubeflow Profile
// https://www.kubeflow.org/docs/components/multi-tenancy/getting-started/#manual-profile-creation

// GetProfile TBU
func (c *KfClient) GetProfile(profileName string) (manifest.Profile, error) {
	data, err := c.cs.RESTClient().
		Get().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Name(profileName).
		DoRaw(context.TODO())
	if err != nil {
		return manifest.Profile{}, err
	}

	profile, err := manifest.UnmarshalProfile(data)
	if err != nil {
		return manifest.Profile{}, err
	}

	return profile, nil
}

// GetProfiles TBU
func (c *KfClient) GetProfiles() (manifest.Profiles, error) {
	data, err := c.cs.RESTClient().
		Get().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		DoRaw(context.TODO())
	if err != nil {
		return manifest.Profiles{}, err
	}

	profiles, err := manifest.UnmarshalProfiles(data)
	if err != nil {
		return manifest.Profiles{}, err
	}

	return profiles, nil
}

// CreateProfile TBU
func (c *KfClient) CreateProfile(profile manifest.Profile) error {
	body, err := manifest.MarshalProfile(profile)
	if err != nil {
		return err
	}

	_, err = c.cs.RESTClient().
		Post().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Body(body).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// UpdateProfile TBU
func (c *KfClient) UpdateProfile(profile manifest.Profile) error {
	body, err := manifest.MarshalProfile(profile)
	if err != nil {
		return err
	}

	_, err = c.cs.RESTClient().
		Put().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Name(profile.ObjectMeta.Name).
		Body(body).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

// DeleteProfile TBU
func (c *KfClient) DeleteProfile(profileName string) error {
	_, err := c.cs.RESTClient().
		Delete().
		AbsPath("/apis/kubeflow.org/v1/profiles").
		Name(profileName).
		DoRaw(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
