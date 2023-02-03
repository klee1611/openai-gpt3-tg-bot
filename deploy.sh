#!/usr/bin/env bash

gcloud functions deploy "openai-tg-bot" \
    --allow-unauthenticated \
    --entry-point=TgWebHook \
    --env-vars-file=.env.yaml \
    --gen2 \
    --max-instances=1 \
    --memory=128MiB \
    --no-user-output-enabled \
    --region=us-central1 \
    --runtime=go118 \
    --timeout=60s \
    --trigger-http
