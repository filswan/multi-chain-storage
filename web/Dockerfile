FROM node:19-alpine
# install simple http server for serving static content
RUN npm install -g http-server
# python2 support

RUN apk add --update \
  python3 \
  python3-dev \
  py-pip \
  build-base \
  git \
  openssh-client \
&& pip install virtualenv \
&& rm -rf /var/cache/apk/* 
# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)

COPY . .
ENV NODE_OPTIONS="--openssl-legacy-provider"

# build app for production with minification
RUN npm run build:prod

EXPOSE 8080
CMD [ "http-server", "dist" ]
