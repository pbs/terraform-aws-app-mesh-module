locals {
  name          = var.name != null ? var.name : var.product
  egress_filter = var.is_egress_allowed ? "ALLOW_ALL" : "DROP_ALL"

  creator = "terraform"

  tags = merge(
    var.tags,
    {
      Name                                      = local.name
      "${var.organization}:billing:product"     = var.product
      "${var.organization}:billing:environment" = var.environment
      creator                                   = local.creator
      repo                                      = var.repo
    }
  )
}
