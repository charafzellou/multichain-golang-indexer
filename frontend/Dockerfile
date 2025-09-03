# Use an existing Node.js image as the base image
FROM node:19.5.0-alpine

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install the app dependencies
RUN npm install

# Copy the rest of the app's source code
COPY . .

# Expose the app's port to the host
EXPOSE 3000

# Specify the command to run when the container starts
CMD [ "npm", "run", "start" ]
