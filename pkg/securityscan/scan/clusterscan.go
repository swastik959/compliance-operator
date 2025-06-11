package scan

import (
	operatorapiv1 "github.com/rancher/compliance-operator/pkg/apis/compliance.cattle.io/v1"
	"github.com/rancher/wrangler/v3/pkg/crd"
	"github.com/rancher/wrangler/v3/pkg/schemas/openapi"
)

func ClusterScanCRD() (*crd.CRD, error) {
	prototype := operatorapiv1.NewClusterScan("", "", operatorapiv1.ClusterScan{})
	schema, err := openapi.ToOpenAPIFromStruct(*prototype)
	if err != nil {
		return nil, err
	}
	return &crd.CRD{
		GVK:        prototype.GroupVersionKind(),
		PluralName: operatorapiv1.ClusterScanResourceName,
		Status:     true,
		Schema:     schema,
		Categories: []string{"securityscan"},
	}, nil
}
