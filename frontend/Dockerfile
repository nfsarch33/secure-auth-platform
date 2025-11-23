# Build Stage
FROM node:18-alpine AS builder

WORKDIR /app

# Copy package files
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

# Copy source code
COPY frontend/ .

# Copy OpenAPI spec from backend
COPY backend/api/openapi.yaml ./openapi.yaml

# Generate API Client
RUN npx openapi-typescript-codegen --input ./openapi.yaml --output ./src/api --client axios

# Build the application
RUN npm run build

# Production Stage
FROM nginx:alpine

# Copy built assets from builder
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy Nginx config
COPY frontend/nginx.conf /etc/nginx/conf.d/default.conf

# Expose port
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
