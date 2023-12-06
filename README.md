# SOCIAL-MEDIA-APP

## About
SocialMedia-App is a web-based application designed for seamless social interaction. Users can access the platform through a secure login and registration process. The app allows users to create posts, consisting of text and uploaded photos, facilitating personalized content sharing. Users can express appreciation for posts through a like function and engage in discussions by adding comments. This straightforward and user-friendly interface aims to provide an enjoyable and interactive social media experience.

## Tech Stack
- Go
- Echo Framework
- MySQL
- GORM
- Docker
- GCP

## Entity Relationship Diagram

## Installation
1. Clone:

```
git clone 
```
2. Go to the Backend_Golang_immersive directory
```
cd Backend_Golang_immersive
```

3. Run the following command
- Enter the package name you want to <b>package-name</b>
```
go mod init package-name
go mod tidy
```

4. Create a file in .env format (local.env for local development)

5. Write as follows in the .env file. Adjust to your needs
```
export JWT_KEY = 'your-jwt-key'
export DBUSER = 'your-db-username'
export DBPASS = 'your-db-password'
export DBHOST = 'your-db-host'
export DBPORT = 'your-db-port'
export DBNAME = 'your-db-name'
```

6. For file upload purposes, create and save <i>Google Application Credentials</i> in a file with the name <b>keys.json</b>. For references to 

## API Documentation
This API documentation can be viewed on [SwaggerHub](https://app.swaggerhub.com/apis-docs/JULIUSSIREGAR1011/Social-Media/1.0.0).

## Collaborator
- Julius Siregar - [Github](https://github.com/juliussiregar)
- 