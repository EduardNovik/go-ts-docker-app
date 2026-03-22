# ------------------------------
# 1. BASE IMAGE (dependency layer)
# ------------------------------

FROM node:25-alpine3.22 AS base


# Install Go
RUN apk add --no-cache git bash curl && \
    rm -rf /usr/local/go && \
    curl -fsSL https://go.dev/dl/go1.25.5.linux-amd64.tar.gz -o /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz
# ✅ Explicitly add Go to PATH so child processes (pnpm, turbo) can find it

# ENV in a Dockerfile persists into every subsequent layer and into runtime — including subprocesses 
# spawned by pnpm/turbo — whereas RUN export PATH=... would only live for that single RUN command.



# Add /usr/local/go/bin to PATH
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOPATH="/root/go"
ENV PATH="${GOPATH}/bin:${PATH}"
WORKDIR /app

# install pnpm
RUN npm install -g pnpm


# ------------------------------
# 2. DEPENDENCIES LAYER
# ------------------------------
# We copy only dependency files first (package.json, lockfile, etc.)
# because Docker caches layers.
# If we copied the whole project before RUN pnpm install,
# then every small change in the source code would invalidate the cache
# and pnpm install would run again (which is slow).
# By installing dependencies first and copying the rest of the code later,
# Docker reuses the cached install layer and builds much faster.

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml turbo.json ./


# ------------------------------
# 3. COPY SOURCE
# ------------------------------
COPY apps ./apps
COPY packages ./packages

# install deps (cached unless lockfile changes)
RUN pnpm install --frozen-lockfile



FROM base AS build

# ------------------------------
# 4. BUILD STEP (build frontend)
# ------------------------------

RUN pnpm turbo run build

# ------------------------------
# 5. START APP
# ------------------------------
# EXPOSE does NOT map ports. It only documents them.
EXPOSE 8080


# This line will not work: turbo is not a Node file it’s a CLI command
# CMD [ "node", "turbo run dev" ]

CMD ["pnpm", "dev"]

# docker build -t app-go
# docker images

# Maps a port inside the container to a port on your host machine.
# Without this, anything running in the container is only accessible inside the container.
# docker run --name app-go -p 3000:3000 -p 8080:8080 app-go:latest

# docker ps -a 