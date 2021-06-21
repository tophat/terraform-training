variable "slack_channel_name" {
  description = "Name of the slack channel"
  type        = string
  default     = ""
}

variable "slack_webhook_url" {
  description = "Url for slack webhook"
  type        = string
  default     = ""
}

variable "sns_topic_name"{
  description = "SNS topic Name"
  type        = string
  default     = ""
} 

variable "slack_username" {
    description = "slack username"
    type = string
    default = ""
}