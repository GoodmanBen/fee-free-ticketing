# If a Windows-based developer:
set shell := ["powershell.exe", "-c"]
set dotenv-load

run-local:
    go run main/main.go

set-project:
    gcloud config set project $GCP_PROJECT

create-artifact-registry-repo:
    gcloud artifacts repositories create fee-free-ticketing --repository-format=docker --location=us-east4 --description="Repository for fee-free-ticketing images"

build-and-push-image:
    gcloud auth configure-docker us-east4-docker.pkg.dev
    docker build -t us-east4-docker.pkg.dev/${PROJECT_ID}/fee-free-ticketing/fee-free-image:latest .
    docker push us-east4-docker.pkg.dev/${PROJECT_ID}/fee-free-ticketing/fee-free-image:latest

deploy-cloud-run:
    gcloud run deploy fee-free-ticketing --image=us-east4-docker.pkg.dev/${PROJECT_ID}/fee-free-ticketing/fee-free-image:latest \
    --description="Fee free ticketing deployment" --ingress="all" --min-instances=1 --max-instances=3 --region=us-east4 \
    --update-env-vars=FROMEMAIL=TODO,FROMNAME="TODO" \
    --update-secrets "STRIPEWEBHOOKSECRET=STRIPE_WEBHOOK_SECRET:latest","SENDGRIDAPITOKEN=SENDGRID_API_KEY:latest","SENDGRIDEMAILTEMPLATEID=SENDGRIDEMAILTEMPLATEID:latest"
