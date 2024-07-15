FROM node:20-alpine

WORKDIR /usr/src/app
ENV NODE_ENV=production

COPY package*.json ./
# Can’t use `npm ci` due to https://github.com/npm/cli/issues/4828
RUN npm i --omit=dev --no-fund

# Copy Indiekit config over
COPY . .

ENV SECRET=
ENV PASSWORD_SECRET=
ENV MONGO_URL=
ENV SITE=https://www.byjp.me/
ENV PORT=80
EXPOSE 80
CMD [ "npm", "run", "start" ]
