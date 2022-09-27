# Dev
FROM node:latest AS dev

WORKDIR /usr/src/app

ENV NODE_ENV production

COPY package*.json ./

RUN yarn --no-progress

RUN yarn add react-scripts -g --silent

COPY . .

RUN yarn build

EXPOSE 8080

CMD [ "npm", "run", "dev" ]
