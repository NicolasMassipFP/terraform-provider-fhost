package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ConvertToHCL2(ctx context.Context, attrs map[string]schema.Attribute, blockDefs map[string]schema.Block) map[string]schema.Attribute {
	for name, block := range ConvertBlocksToHCL2(ctx, blockDefs) {
		attrs[name] = block
	}
	return attrs
}

func ConvertBlocksToHCL2(ctx context.Context, blockDefs map[string]schema.Block) map[string]schema.Attribute {
	attributes := make(map[string]schema.Attribute)

	for name, block := range blockDefs {
		switch b := block.(type) {
		case schema.SingleNestedBlock:
			attributes[name] = schema.SingleNestedAttribute{
				Description:         b.Description,
				MarkdownDescription: b.MarkdownDescription,
				CustomType:          b.CustomType,
				Attributes:          b.Attributes,
				Optional:            true,
				Computed:            false,
			}
		case schema.ListNestedBlock:
			attributes[name] = schema.ListNestedAttribute{
				Description:         b.Description,
				MarkdownDescription: b.MarkdownDescription,
				CustomType:          b.CustomType,
				NestedObject: schema.NestedAttributeObject{
					Attributes: b.NestedObject.Attributes,
				},
				Optional: true,
				Computed: false,
			}
		case schema.SetNestedBlock:
			attributes[name] = schema.SetNestedAttribute{
				Description:         b.Description,
				MarkdownDescription: b.MarkdownDescription,
				CustomType:          b.CustomType,
				NestedObject: schema.NestedAttributeObject{
					Attributes: b.NestedObject.Attributes,
				},
				Optional: true,
				Computed: false,
			}
		}
	}

	return attributes
}
