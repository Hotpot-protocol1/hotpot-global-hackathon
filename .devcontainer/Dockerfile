FROM mcr.microsoft.com/devcontainers/javascript-node:1-20-bullseye

# Create app directory
WORKDIR /app

# Install client app dependencies
COPY ./frontend/package*.json ./frontend/
RUN cd frontend && yarn install

# Install server app dependencies
COPY ./backend/package*.json ./backend/
RUN cd backend && yarn install

# Bundle app source
COPY . .

# Build the client app
# RUN cd client && npm run build

# Command to start server app
# CMD ["npm", "run", "start", "--prefix", "server"]