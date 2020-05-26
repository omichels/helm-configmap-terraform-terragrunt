inputs = {

  environment_name  = basename(get_terragrunt_dir())

  my_key_value_map = {
      "apiKeyBase" = "http://dev/api"
      "keyTrinity" = "trinity"
      "keyMorpheus" = "morpheus"
      "keyNeo" = "neo"
  }
}


terraform {
  source = "../..//terraform" # NOTE: double slash is intentional and required!

}