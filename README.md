# giftweaver
secret santa bot

Assign bot webhook
``` 
https://api.telegram.org/bot<YOUR_BOT_TOKEN>/setWebhook?url=https://giftweaver-api.vercel.app/
```

Delete bot webhook
```
curl -X POST "https://api.telegram.org/bot<YOUR_BOT_TOKEN>/deleteWebhook"
```


should telegram setWebhook be set to giftweaver-api or giftweaver-ui