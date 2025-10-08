# Notes API

Notes API has four APIs that can:

- create note,
- get note,
- update note,
- delete note.

## Tech Stack
1. Go
2. Docker
3. Kubernetes
4. Helm
5. Terraform
6. Prometheus
7. OpenTelemetry

## Architecture
![alt text](https://github.com/yusuffranklin/notes-api/blob/main/architecture.png?raw=true)

## How To Run
1. Build and Push Docker Image
`docker buildx build -t <your-image-repo>/notes-api:latest . --push`

2. Provision GCP Resources
- Add `GOOGLE_APPLICATION_CREDENTIALS` environment variable with GCP Service Account as its value.

`export GOOGLE_APPLICATION_CREDENTIALS=/path/to/gcp-service-account.json`

- Create a `terraform.tfvars` in `terraform/` folder and add these variables:
```
project_id = "<your-gcp-project-id>"
db_instance_name = "<db-instance-name>"
db_user = "<db-user>"
db_password = "<db-password>"
db_name = "<db-name>"
```
- After that, run these commands, make sure you are in the `terraform/` directory:
```
terraform init
terraform plan
terraform apply
```

3. Run the Application on Kubernetes
- Install Prometheus Stack
`helm install prometheus prometheus-community/kube-prometheus-stack --set grafana.enabled=false -n monitoring --create-namespace`

- Install Keda
`helm install keda kedacore/keda -n keda --create-namespace`

- Run notes-api with Helm. Make sure you are in `k8s/` directory.
`helm install notes-api .`
