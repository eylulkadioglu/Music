# Artist Management System

This project is an Artist Management System that allows users to manage artist information, claims, and other related data. It includes functionalities for user authentication, database interactions, and email notifications.

## Features

- **User Authentication**: Secure user login and registration.
- **Artist Management**: Add, update, view, and delete artist information.
- **Claim Management**: Handle claims related to artists.
- **Database Interactions**: Efficiently manage data using a database.
- **Email Notifications**: Send email notifications for important events.

### Prerequisites

- Go 1.x
- PostgreSQL

### Steps

1. **Clone the repository**:
    ```bash
    git clone https://github.com/eylulkadioglu/Music.git
    cd Music
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Set up the database**:
    - Configure your PostgreSQL database and update the `db.go` with your database credentials.

4. **Run the application**:
    ```bash
    go run main.go
    ```

## Usage

1. **API Endpoints**:
    - The API endpoints can be accessed at `http://localhost:8080`.
    - Example endpoints:
        - `GET /artists` - Retrieve all artists.
        - `POST /artists` - Add a new artist.
        - `PUT /artists/:id` - Update an artist.
        - `DELETE /artists/:id` - Delete an artist.
        - `POST /login` - User login.
        - `POST /register` - User registration.

## API Endpoints

Below are some of the key API endpoints:

- **User Endpoints**:
    - `POST /register`: Register a new user.
    - `POST /login`: Login a user.

- **Artist Endpoints**:
    - `GET /artists`: Get a list of all artists.
    - `POST /artists`: Add a new artist.
    - `PUT /artists/:id`: Update an existing artist.
    - `DELETE /artists/:id`: Delete an artist.

- **Claim Endpoints**:
    - `GET /claims`: Get a list of all claims.
    - `POST /claims`: Add a new claim.
    - `PUT /claims/:id`: Update an existing claim.
    - `DELETE /claims/:id`: Delete a claim.


