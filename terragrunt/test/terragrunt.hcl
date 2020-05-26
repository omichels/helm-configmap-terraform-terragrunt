inputs = {

  environment_name  = basename(get_terragrunt_dir())

  my_key_value_map = {
      "apiKeyBase" = "http://test/api"
  }
}


terraform {
  source = "../..//terraform" # NOTE: double slash is intentional and required!

}