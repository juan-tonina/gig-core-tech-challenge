FROM node:16

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY package*.json ./

RUN npm install


COPY tsconfig.json ./
COPY src src

RUN npm run build
# If you are building your code for production
# RUN npm ci --only=production

# Bundle app source
EXPOSE 8080

#RUN npm run start
CMD ["node", "./dist/server.js"]