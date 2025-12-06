FROM node:25.2-alpine AS development-dependencies-env
COPY . /app
WORKDIR /app
RUN npm install
CMD ["npm", "run", "build", "&&", "npm", "start"]