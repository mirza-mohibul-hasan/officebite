FROM node:25-alpine AS deps

WORKDIR /app

COPY apps/web/package*.json ./
RUN npm install

FROM node:25-alpine AS dev

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY apps/web ./

EXPOSE 5173

CMD ["npm", "run", "dev"]
