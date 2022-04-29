# Multi-Chain Storage

**Technology stack：** vue2 + vuex + vue-router + webpack + sass + element-ui + web3


## Container Installation

Running Swan Faucet as a container is the recommended way of using it.

### Prerequisites

#### Install Docker Compose

https://docs.docker.com/compose/install/
### Stable

Run the following command to run the latest stable image of MCS Web

```bash
docker build -t filswan/mcs-web .
docker run  -p 8080:8080 filswan/mcs-web
```
The service will be Available on:

http://127.0.0.1:8080

http://172.17.0.2:8080

Hit CTRL-C to stop the server

## Installation dependency

Run `npm install` to generate component.

## Development server

Run `npm run dev` for a dev server. Navigate to `http://localhost:8080/`. The app will automatically reload if you change any of the source files.

## Build project

```shell
# Build test projects
$ npm run build:test

# Build staging projects
$ npm run build:staging

# Build prod projects
$ npm run build:prod
```

The build artifacts will be stored in the `dist/` directory.

## Reference documents

- [vue](https://vuejs.bootcss.com/v2/guide/)：Vue is a progressive framework for building user interfaces.

- [vuex](https://vuex.vuejs.org/zh/)：Vuex is a state management pattern developed specifically for vue.js applications.

- [vue-router](https://router.vuejs.org/zh/)：Vue router is the official routing manager of vue.js.

- [webpack](https://webpack.js.org/concepts/)：Front end module packer.

- [element-ui](https://element.eleme.io/)：Element, a desktop component library based on Vue 2.0 for developers, designers and product managers.

- [web3.js](http://cw.hubwiz.com/card/c/web3.js-1.0/)：Web3.js is a set of JS libraries used to interact with local or remote Ethereum nodes.

## Pre order preparation

**Preparation before operation:**

   Since this project is based on nodejs, you need to make preparations for nodejs. Before running the project, please ensure that the following applications have been installed in the system:

   (1)、Node (version 10.18.0 and above). Please refer to:[Download and install node.](https://nodejs.org/en/download/)

## Further help

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).
