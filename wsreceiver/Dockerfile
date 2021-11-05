FROM node:16

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY tsconfig.json ./
COPY package*.json ./
COPY src /app/src
COPY dummy.ts ./
RUN ls -a

RUN npm install
RUN npm run build
# If you are building your code for production
# RUN npm ci --only=production

# Bundle app source

EXPOSE 8080

CMD [ "node", "./dist/server.js" ]