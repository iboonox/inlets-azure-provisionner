variable "prefix" {
  description = "The prefix which should be used for all resources"
  default = "inlets"
}

variable "location" {
  description = "The Azure Region in which all resources should be created."
  default = "westeurope"
}

variable "vm_size_sku" {
  description = "The vm size sku name"
  default = "Standard_B1ls"
}
