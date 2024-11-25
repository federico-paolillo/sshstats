variable "sshstats_subscription" {
  default     = ""
  description = "Target Azure subscription identifier"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_fn_config_headerkey" {
  default     = "X-Remembertostarthemwithanuppercaseletter"
  description = "Secrete authz. header key"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_fn_config_headervalue" {
  default     = "usejustlowercaselettersandnumbersplease"
  description = "Secrete authz. header value"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_fn_config_loki_pwd" {
  default     = ""
  description = "Password to access Loki instance on Grafana Cloud to fetch logs from"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_fn_config_loki_usr" {
  default     = ""
  description = "Username to access Loki instance on Grafana Cloud to fetch logs from"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_fn_config_loki_endpoint" {
  default     = ""
  description = "Loki instance endpoint on Grafana Cloud to fetch logs from"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_az_use_cli" {
  default     = false
  description = "Turn on or off Azure CLI authentication"
  nullable    = false
  sensitive   = false
  type        = bool
}

variable "sshstats_az_client_id" {
  default     = ""
  description = "Azure Service Principal identifier"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_az_client_secret" {
  default     = ""
  description = "Azure Service Principal secret"
  nullable    = false
  sensitive   = true
  type        = string
}

variable "sshstats_az_tenant_id" {
  default     = ""
  description = "Azure Tenant identifier"
  nullable    = false
  sensitive   = true
  type        = string
}
