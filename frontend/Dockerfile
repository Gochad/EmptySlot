FROM node:18-alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install application dependencies
RUN npm ci

# Copy the rest of the application code to the working directory
COPY . .

# Expose the port on which your application will run
EXPOSE 3000

# Define the default command to run when the container starts
CMD ["npm", "start"]
