output "key_vault_id" {
  description = "ID of the Key Vault"
  value       = azurerm_key_vault.key_vault.id
}

output "key_vault_uri" {
  description = "URI of the Key Vault"
  value       = azurerm_key_vault.key_vault.vault_uri
}
