output keypair_name {
    value = aws_key_pair.my_keypair.key_name
    description = "The name of the keypair"
}

output ssh_private_key {
    value = tls_private_key.pk.private_key_pem
    description = "The private key of the keypair"
}