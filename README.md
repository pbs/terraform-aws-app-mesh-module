# PBS TF App Mesh Module

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs/terraform-aws-app-mesh-module?ref=x.y.z
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

This module provisions a basic App Mesh.

Integrate this module like so:

```hcl
module "app_mesh" {
  source = "github.com/pbs/terraform-aws-app-mesh-module?ref=x.y.z"

  # Tagging Parameters
  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo

  # Optional Parameters
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.2 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.7.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_appmesh_mesh.mesh](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/appmesh_mesh) | resource |
| [aws_default_tags.common_tags](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/default_tags) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_environment"></a> [environment](#input\_environment) | Environment (sharedtools, dev, staging, qa, prod) | `string` | n/a | yes |
| <a name="input_organization"></a> [organization](#input\_organization) | Organization using this module. Used to prefix tags so that they are easily identified as being from your organization | `string` | n/a | yes |
| <a name="input_product"></a> [product](#input\_product) | Tag used to group resources according to product | `string` | n/a | yes |
| <a name="input_repo"></a> [repo](#input\_repo) | Tag used to point to the repo using this module | `string` | n/a | yes |
| <a name="input_is_egress_allowed"></a> [is\_egress\_allowed](#input\_is\_egress\_allowed) | Determines whether egress is allowed from the app mesh. | `bool` | `true` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the app mesh. Will default to product if not defined. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Extra tags | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the service mesh. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the service mesh. |
| <a name="output_name"></a> [name](#output\_name) | The name of the service mesh. |
