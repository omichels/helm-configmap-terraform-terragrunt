variable "environment_name" {
}

variable "my_key_value_map" {
	type = map
	default = {
		"defaultKey1" = "defaultValue1"
	}
}
