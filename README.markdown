# lindsay.dev

Run a go binary in Section's Node.js edge app hosting.

# Deploying

Ensure you have [`sectionctl`](https://github.com/section/sectionctl/releases/latest) installed, and run:

```
# build the server
export GOOS=linux
make

# deploy the app
export ACCOUNT_ID=123
export APP_ID=456
sectionctl deploy --account-id $ACCOUNT_ID --app-id $APP_ID
```
