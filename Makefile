fmt:
	@echo "Formatting code..."
	terraform fmt -recursive
	@echo "Formatting successful"

deploy-prod:
	@echo "Deploying to production..."
	terraform -chdir=deploy/prod get
	terraform -chdir=deploy/prod apply
	@echo "Deployment successful"
