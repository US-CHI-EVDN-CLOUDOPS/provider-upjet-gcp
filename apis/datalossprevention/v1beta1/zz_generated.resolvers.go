/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-gcp/apis/bigquery/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this JobTrigger.
func (mg *JobTrigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.InspectJob); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.InspectJob[i3].Actions); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig); i6++ {
					for i7 := 0; i7 < len(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table); i7++ {
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].DatasetID),
							Extract:      reference.ExternalName(),
							Reference:    mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].DatasetIDRef,
							Selector:     mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].DatasetIDSelector,
							To: reference.To{
								List:    &v1beta1.DatasetList{},
								Managed: &v1beta1.Dataset{},
							},
						})
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].DatasetID")
						}
						mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].DatasetID = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].DatasetIDRef = rsp.ResolvedReference

					}
				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.InspectJob); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.InspectJob[i3].Actions); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig); i6++ {
					for i7 := 0; i7 < len(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table); i7++ {
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].TableID),
							Extract:      reference.ExternalName(),
							Reference:    mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].TableIDRef,
							Selector:     mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].TableIDSelector,
							To: reference.To{
								List:    &v1beta1.TableList{},
								Managed: &v1beta1.Table{},
							},
						})
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].TableID")
						}
						mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].TableID = reference.ToPtrValue(rsp.ResolvedValue)
						mg.Spec.ForProvider.InspectJob[i3].Actions[i4].Deidentify[i5].TransformationDetailsStorageConfig[i6].Table[i7].TableIDRef = rsp.ResolvedReference

					}
				}
			}
		}
	}

	return nil
}
