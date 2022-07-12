# build stage
FROM node:lts-alpine as build-stage
ARG NODE_ENV
ARG API_BASE_URL
ENV NODE_ENV "$NODE_ENV"
ENV API_BASE_URL "$API_BASE_URL"

WORKDIR /app
COPY ["package.json", "package-lock.json*", "./"]
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80