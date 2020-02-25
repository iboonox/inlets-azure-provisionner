output "inlets_exit_node_public_ip" {
   value = azurerm_linux_virtual_machine.inlets.private_ip_address
   description = "The public ip of the exit node"
}

