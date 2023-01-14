FROM node:16-alpine

WORKDIR /app

COPY . .
RUN yarn
RUN npm cache clean --force

RUN yarn build
EXPOSE 3000

CMD [ "yarn", "start","-p","3000" ]
