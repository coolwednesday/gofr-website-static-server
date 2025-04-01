# Stage 1: Build the website
FROM node:18 AS builder

WORKDIR /app

# Copy package files first (relative to PROJECT ROOT)
COPY ./gofr-website/website/package.json .
COPY ./gofr-website/website/package-lock.json* .

# Install dependencies
RUN npm install

# Copy entire website codebase (relative to PROJECT ROOT)
COPY ./gofr-website/website/ .

# Copy documentation files (relative to PROJECT ROOT)
COPY  http-server/docs/quick-start /app/src/app/docs/quick-start
COPY  http-server/docs/public/ /app/public
COPY  http-server/docs/advanced-guide /app/src/app/docs/advanced-guide
COPY  http-server/docs/datasources /app/src/app/docs/datasources
COPY  http-server/docs/references /app/src/app/docs/references
COPY  http-server/docs/page.md /app/src/app/docs
COPY  http-server/docs/navigation.js /app/src/lib
COPY  http-server/docs/events.json /app/src/app/events
COPY  http-server/docs/testimonials.json /app/utils

# Build
ENV NODE_ENV=production
RUN npm run build

COPY http-server/go.mod .
COPY http-server/go.sum .
COPY http-server/main.go .

FROM golang:1.24-alpine AS go-builder

WORKDIR /app

# Copy Go files before building
COPY http-server/go.mod .
COPY http-server/go.sum .
COPY http-server/main.go .

# Install dependencies and build
RUN go mod tidy && \
go build -o main main.go

# Stage 3: Final image
FROM alpine:3.19

WORKDIR /app

# Copy built assets
COPY --from=builder /app/out ./website
COPY --from=go-builder /app/main .

EXPOSE 8000
CMD ["/app/main"]

