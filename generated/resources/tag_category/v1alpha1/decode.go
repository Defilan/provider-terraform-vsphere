/*
	Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	"fmt"

	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*TagCategory)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeTagCategory(r, ctyValue)
}

func DecodeTagCategory(prev *TagCategory, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeTagCategory_AssociableTypes(&new.Spec.ForProvider, valMap)
	DecodeTagCategory_Cardinality(&new.Spec.ForProvider, valMap)
	DecodeTagCategory_Description(&new.Spec.ForProvider, valMap)
	DecodeTagCategory_Name(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeTagCategory_AssociableTypes(p *TagCategoryParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["associable_types"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.AssociableTypes = goVals
}

//primitiveTypeDecodeTemplate
func DecodeTagCategory_Cardinality(p *TagCategoryParameters, vals map[string]cty.Value) {
	p.Cardinality = ctwhy.ValueAsString(vals["cardinality"])
}

//primitiveTypeDecodeTemplate
func DecodeTagCategory_Description(p *TagCategoryParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeTagCategory_Name(p *TagCategoryParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}
