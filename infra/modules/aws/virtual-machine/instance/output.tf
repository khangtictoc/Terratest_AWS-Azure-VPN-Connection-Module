output private_ip {
  value = aws_instance.instance.private_ip
}

output public_ip {
  value = aws_instance.instance.public_ip
}

output instance_id {
  value = aws_instance.instance.id
}