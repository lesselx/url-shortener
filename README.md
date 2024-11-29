# URL Shortener

My name is Agnes Audya Tiara P, and here is my documentation.

## List of Endpoints

### 1. Generate a Short Link
- **Endpoint:** `/api/short_urls`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
        "original_url": "string"
    }
    ```
- **Response:**
    ```json
    {
        "status_code": 201,
        "data": {
            "short_url": "string",
            "created_at": "time"
        }
    }
    ```

### 2. Redirect to the Original Link with its Shortlink
- **Endpoint:** `/api/short_urls/:alias`
- **Method:** `GET, HEAD`

## Running the Application

To run the application, navigate to the project directory and execute this command with default port, host, and env:

```sh
go run ./cmd/api 
```

or If using customize port, host, or env 

```sh
go run ./cmd/api -port=4000 -environment=development -host=http://localhost 
```

No need to set up a database. The application uses the sqlite driver from [glebarez/sqlite](https://github.com/glebarez/sqlite), which will generate a database file automatically.