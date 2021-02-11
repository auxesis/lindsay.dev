# lindsay.dev

Run a go binary in Section's Node.js edge app hosting.

## Deploying

Ensure you:

- Have signed up to Section and have a Node.js app configured
- Have [`sectionctl`](https://github.com/section/sectionctl/releases/latest) installed

Then:

```
# build the server
export GOOS=linux GOARCH=amd64
make

# deploy the app
export ACCOUNT_ID=123 # use your own ACCOUNT_ID
export APP_ID=456 # use your own APP_ID
sectionctl deploy --account-id $ACCOUNT_ID --app-id $APP_ID
```

## What is this sourcery? 🧙‍♀️

Section's Node.js module has a "strict" contract that must be met to run an app:

- A [`server.conf`](https://github.com/section/nodejs-example/blob/master/server.conf) that is used by Nginx.
- A `node_modules` directory. This directory can technically be empty. 
- A `package.json` with a `scripts.start` key that includes a command to start a server: 
  ``` javascript
  {
    // ...
    "scripts": {
      "start": "node index.js"
    }
    // ...
  }
  ```
- A server process that listens on TCP 8080.
- Any binaries that are upload are 64-bit ELF

The module doesn't mandate that the command in the `scripts.start` key has to be run by `node`. 

We can build a standalone binary with Go that listens on TCP 8080, and is started by the `scripts.start` key:

``` javascript
{
  // ...
  "scripts": {
    "start": "./server"
  }
  // ...
}
```
