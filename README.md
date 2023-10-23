# HRIS-APP

## About
HRIS-APP is a web-based application for managing employee data for a company. This application can be accessed by several actors, including Superadmin or C-Level, Admin or HR (Human Resources), Managers, and Employees, where each actor has different access rights. Some of the features in this application include the attendance feature, employee data management features, leave and reimbursement application features, as well as employee performance monitoring features.

## Tech Stack
- Go
- Echo Framework
- MySQL
- GORM
- Docker
- GCP

## Entity Relationship Diagram
<img src="assets\images\ERD HRIS-APP.drawio.png" width= 600>

## Installation
1. Clone:

```
git clone https://github.com/HRIS-APP-TEAM-3/Backend_Golang_immersive.git
```
2. Go to the Backend_Golang_immersive directory
```
cd Backend_Golang_immersive
```

3. Jalankan perintah berikut
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

6. For file upload purposes, create and save <i>Google Application Credentials</i> in a file with the name <b>keys.json</b>. For references to <i>Google Application Credentials</i>, please check [reference](https://adityarama1210.medium.com/simple-golang-api-uploader-using-google-cloud-storage-3d5e45df74a5) or [reference](https://cloud.google.com/storage/docs/reference/libraries#client-libraries-install-go)

## API Documentation
This API documentation can be viewed on [SwaggerHub](https://app.swaggerhub.com/apis-docs/ILHAM9D27_1/HRIS-APP-Team-2/1.0.0).

## Collaborator
- Jaya Rizky Prayoga - [Github](https://github.com/Prayogarock)
- Julius Siregar - [Github](https://github.com/juliussiregar)
- Mohammad Hadi Hamdam - [Github](https://github.com/Hadi1Hamdam)