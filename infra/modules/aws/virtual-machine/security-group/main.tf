resource "aws_security_group" "security_group" {
    name        = var.security_group_config.name
    description = var.security_group_config.description
    vpc_id      = var.vpc_id

    tags = merge(
        var.tags,
        {
          "Name" = var.security_group_config.name
        }
    )
}

resource "aws_security_group_rule" "security_group_rule" {
  count = length(var.security_group_config.security_group_rule)

  type              = var.security_group_config.security_group_rule[count.index].type
  from_port         = var.security_group_config.security_group_rule[count.index].from_port
  to_port           = var.security_group_config.security_group_rule[count.index].to_port
  protocol          = var.security_group_config.security_group_rule[count.index].protocol
  cidr_blocks       = var.security_group_config.security_group_rule[count.index].cidr_blocks
  security_group_id = aws_security_group.security_group.id
}












