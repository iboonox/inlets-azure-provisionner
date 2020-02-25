resource "random_string" "vm_admin_password" {
  length  = 64
  special = true
}

data "template_file" "userdata" {
  template = "${file("${path.module}/templates/cloud-config.tpl")}"
}

# ################################################
# ## INLETS - RESOURCE GROUP 
# ################################################

## Resource group
resource "azurerm_resource_group" "inlets" {
  name     = "${var.prefix}-resources"
  location = var.location
}
# ################################################
# ## INLETS - NETWORK 
# ################################################

## Network
resource "azurerm_virtual_network" "inlets" {
  name                = "${var.prefix}-network"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.inlets.location
  resource_group_name = azurerm_resource_group.inlets.name
}

resource "azurerm_subnet" "internal" {
  name                 = "internal"
  resource_group_name  = azurerm_resource_group.inlets.name
  virtual_network_name = azurerm_virtual_network.inlets.name
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_network_interface" "inlets" {
  name                = "${var.prefix}-nic"
  resource_group_name = azurerm_resource_group.inlets.name
  location            = azurerm_resource_group.inlets.location

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.internal.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.inlets.id
  }
}

resource "azurerm_public_ip" "inlets" {
  name                = "inletsExitPublicIp1"
  location            = var.location
  resource_group_name = azurerm_resource_group.inlets.name
  allocation_method   = "Static"
}

# ################################################
# ## INLETS - EXIT NODE VM 
# ################################################
resource "azurerm_linux_virtual_machine" "inlets" {
  name                            = "${var.prefix}-vm"
  resource_group_name             = azurerm_resource_group.inlets.name
  location                        = azurerm_resource_group.inlets.location
  size                            = var.vm_size_sku
  admin_username                  = "inlets"
  admin_password                  = random_string.vm_admin_password.result
  custom_data                     = base64encode("${data.template_file.userdata.rendered}")
  disable_password_authentication = false
  network_interface_ids = [
    azurerm_network_interface.inlets.id,
  ]

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }
}
