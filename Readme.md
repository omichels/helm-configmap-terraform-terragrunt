# Example on injecting key-value properties file into kubernetes pod

-Helm-Charts in dir 'helm-chart'

-Terraform configuration using the helm provider in dir 'terraform'

-dir 'terragrunt' has different environments for dev/test/prod



prerequistes

- terraform v.12 or higher
- terragrunt 
- helm


example usage


$ cd terragrunt

$ terragrunt plan

$ terragrunt apply


$ kubectl exec -it alpine sh
/ # cat /etc/config/application.properties
key0=val0

