output "web_app_url" {
  description = "URL of the web app"
  value       = "https://${azurerm_linux_web_app.web_app.default_hostname}"
}

output "web_app_name" {
  description = "Name of the web app"
  value       = azurerm_linux_web_app.web_app.name
}

output "web_app_id" {
  description = "ID of the web app"
  value       = azurerm_linux_web_app.web_app.id
}
