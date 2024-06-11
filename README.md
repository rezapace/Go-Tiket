# Go-Tiket 🎫
 app ticket
## Description 📓
It is a backend application using the Golang programming language that functions to purchase train tickets. This application is made using the Clean Architecture concept and uses the PostgreSQL database as data storage.

## How to Run 📓
1. Clone repository

```bash
git clone https://github.com/rezapace/go-tiket
```

2. Enter the repository directory

```bash
cd go-tiket
```

3. Check the env file and adjust it to the database configuration being used

4. Enable the PostgreSQL database

5. Run the command to create the database and migrate it.
    
```bash
migrate -database "postgres://postgres:pass@localhost:5432/weather-app?sslmode=disable" -path db/migration-golang up
```
### Note:
- make sure you have installed the migrate cli
```bash
scoop install migrate
```
- make sure you have installed the postgresql driver

6. Jalankan aplikasi

```bash
go mod tidy
go run cmd/server/main.go
```

## Detail Proyek 📓

**Project Theme:** Aplikasi Ticketing."

**Description:** This platform is used for buying and selling concert and event tickets. Users can register as buyers and search for concert schedules according to their needs, as well as pay for tickets online. The platform also provides accurate information about ongoing events.

**Features:**

- User Registration.
- Implementation of In-App Notification (in-app notification, not push notification).
- User Profile.
- Transaction History.
- Search and Filter.

**Technologies Used:**

- Programming language: Golang.
- The use of other programming languages or third-party libraries is not allowed.

**Architecture:**

- Using Model-View-Controller (MVC) basic architecture with two layers (backend and frontend).
- Implementation of search, page sharing, and sorting features.

**Available Roles:**

- Admin.
- Buyer.

**Development and Deployment:**

- Development using Vercel for staging and production.
- Git workflow with two main branches: "master" for production and "develop/staging" for testing.
- Deliverables include an API that can be used by the front end, Swagger documentation, and a project repository.

## API Documentation 🔥
Dokumentasi API dapat diakses pada file
```bash
output\Backend Ticketing.postman_collection.json
```

```bash
https://github.com/rezapace/Go-Tiket/blob/main/output/Backend%20Ticketing.postman_collection.json
```


## Image Collection 📷
<table>
  <tr>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/flow%201.png?raw=true" alt="Figma 1"></td>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/flow%202.png?raw=true" alt="Figma 2"></td>
  </tr>
</table>
<table>
  <tr>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/postman.jpg?raw=true" alt="Postman"></td>
    <td><img src="https://github.com/rezapace/Go-Tiket/blob/main/Materi/navichat.jpg?raw=true" alt="Database"></td>
  </tr>
</table>
