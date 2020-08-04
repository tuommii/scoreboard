FROM node:lts-alpine AS frontend

# install simple http server for serving static content
# RUN npm install -g @vue/cli

RUN mkdir -p /frontend/node_modules && chown -R node:node /frontend
# make the 'frontend' folder the current working directory
WORKDIR /frontend

# copy both 'package.json' and 'package-lock.json' (if available)
COPY /frontend/package*.json ./

# install project dependencies
RUN npm install
RUN npm audit fix

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY  /frontend .

# build app for production with minification
RUN npm run build


FROM golang:alpine AS server

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .
RUN go mod download

# Build the application
RUN go build -o run cmd/scoreboard/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/run .

# Build a small image
FROM scratch

COPY --from=server /dist/run /
COPY /assets/designs.json /assets/designs.json
# Static dir is public
COPY --from=frontend /frontend/dist /public/

# Command to run
ENTRYPOINT ["/run"]
