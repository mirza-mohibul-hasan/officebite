FROM node:25-alpine AS deps

WORKDIR /app

COPY apps/web/package*.json ./
RUN npm ci

FROM node:25-alpine AS dev

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY apps/web ./

EXPOSE 5173

CMD ["npm", "run", "dev"]

FROM deps AS build

WORKDIR /app
COPY apps/web ./
RUN npm run build

FROM nginx:1.27-alpine AS production

COPY infrastructure/docker/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
