# giftweaver
secret santa bot

Assign bot webhook
``` 
https://api.telegram.org/bot<YOUR_BOT_TOKEN>/setWebhook?url=https://<app_name>.vercel.app/webhook/
```

Delete bot webhook
```
curl -X POST "https://api.telegram.org/bot<YOUR_BOT_TOKEN>/deleteWebhook"
```