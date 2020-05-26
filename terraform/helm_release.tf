
resource "helm_release" "helm-test" {

	name = "helm-test"
	chart = "../helm-charts/example"
  namespace = "helm-test"
	
	dynamic "set" {
		for_each = var.my_key_value_map
		content {
			name = "myKeyValueMap.${set.key}"
			value = set.value
		}
	}

	set {
		name = "environmentName"
		value = var.environment_name
	}
}
