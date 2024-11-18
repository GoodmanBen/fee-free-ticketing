# If a Windows-based developer:
set shell := ["powershell.exe", "-c"]
set dotenv-load

run-local:
    go run main/main.go

deploy-cloud-run:
    echo "nothing"
