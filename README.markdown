# lindsay.dev

Run a go binary in Section's Node.js edge app hosting.

## Deploying

Ensure you:

- Have signed up to Section and have a Node.js app configured
- Have [`sectionctl`](https://github.com/section/sectionctl/releases/latest) installed

Then:

```
# build the server
export GOOS=linux
make

# deploy the app
export ACCOUNT_ID=123 # use your own ACCOUNT_ID
export APP_ID=456 # use your own APP_ID
sectionctl deploy --account-id $ACCOUNT_ID --app-id $APP_ID
```
