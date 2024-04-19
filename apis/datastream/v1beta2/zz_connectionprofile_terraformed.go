// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ConnectionProfile
func (mg *ConnectionProfile) GetTerraformResourceType() string {
	return "google_datastream_connection_profile"
}

// GetConnectionDetailsMapping for this ConnectionProfile
func (tr *ConnectionProfile) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"forward_ssh_connectivity[*].password": "spec.forProvider.forwardSshConnectivity[*].passwordSecretRef", "forward_ssh_connectivity[*].private_key": "spec.forProvider.forwardSshConnectivity[*].privateKeySecretRef", "mysql_profile[*].password": "spec.forProvider.mysqlProfile[*].passwordSecretRef", "mysql_profile[*].ssl_config[*].ca_certificate": "spec.forProvider.mysqlProfile[*].sslConfig[*].caCertificateSecretRef", "mysql_profile[*].ssl_config[*].client_certificate": "spec.forProvider.mysqlProfile[*].sslConfig[*].clientCertificateSecretRef", "mysql_profile[*].ssl_config[*].client_key": "spec.forProvider.mysqlProfile[*].sslConfig[*].clientKeySecretRef", "oracle_profile[*].password": "spec.forProvider.oracleProfile[*].passwordSecretRef", "postgresql_profile[*].password": "spec.forProvider.postgresqlProfile[*].passwordSecretRef"}
}

// GetObservation of this ConnectionProfile
func (tr *ConnectionProfile) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConnectionProfile
func (tr *ConnectionProfile) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConnectionProfile
func (tr *ConnectionProfile) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConnectionProfile
func (tr *ConnectionProfile) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConnectionProfile
func (tr *ConnectionProfile) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ConnectionProfile
func (tr *ConnectionProfile) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this ConnectionProfile
func (tr *ConnectionProfile) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this ConnectionProfile using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConnectionProfile) LateInitialize(attrs []byte) (bool, error) {
	params := &ConnectionProfileParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConnectionProfile) GetTerraformSchemaVersion() int {
	return 0
}
