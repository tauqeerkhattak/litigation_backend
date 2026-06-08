Create a container first using the following command:

```bash
gcloud builds submit --tag gcr.io/litigation-app-5f964/litigation-backend
```

Deploy to GCR using the following command:

```bash
gcloud run deploy litigation-backend --image gcr.io/litigation-app-5f964/litigation-backend --platform managed --region us-east1 --allow-unauthenticated --port 8080 --env-vars-file env.yaml
```
