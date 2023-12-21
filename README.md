# Go REST API with Gin and PostgreSQL

Manage player scores with ease using this lightweight and efficient REST API.

## Features

- CRUD Operations: Create, read, update, and delete player information seamlessly.
- Leaderboard: Retrieve a ranked list of players based on their scores.

``` bash
cd go-gin-rest-api
go mod download
```

## Getting Started

### Prerequisites

- Golang (version 1.18 or later)
- PostgreSQL

### Installation

1. Clone the repository:

  ```bash
  git clone [https://github.com/your-username/go-gin-rest-api.git](https://github.com/your-username/go-gin-rest-api.git)
  ```

2. Install dependencies:

  ``` bash
cd go-gin-rest-api
go mod download
```

3. Database Setup:
   - Create a PostgreSQL database named gamedatabase
   - Create a table named players with the following columns:
     * player_id (integer, primary key)
     * player_name (text)
     * score (integer)
       
4. Running the API:
   - Start the API server:
     ```bash
     go run main.go
     ```

   - The API will be accessible at [](http://localhost:8080)
  
5. API Endpoints:
   |Endpoints     | Method | Description                                     |
   |--------------|--------|-------------------------------------------------|
   |/player       | GET    | Retrieve information about all players.         |
   |/player       | POST   | Add a new player with score information.        |
   |/player/{id}  | DELETE | Delete a player by their ID.                    |
   |/player/{id}  | PATCH  | Update player information: name, score, or both.|
   |/leaderboard  | GET    | Retrieve the player leaderboard sorted by score.|

   

