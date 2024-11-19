# Games Night

This is a little games night leaderboard server. The aim is that there are multiple rooms within a session. Each room acts
as a "station" where a game is played. At the end of the game, winners are decided and scores updated.

## Environment Variables

The following environment variables need to be set for the application to run:

-   `FIREBASE_SA_PATH` - Path to Firebase Service Account JSON file
-   `GO_ENV` - The environment in which the application is running (e.g., development, production)

## Setup

1. **Clone the repository:**

    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. **Create a `.env` file**

    Create a `.env` file in the root directory of the project and add the following environment variables:

    ```env
    FIREBASE_SA_PATH=./path/to/firebase-service-account.json
    GO_ENV=development
    ```

3. **Install dependencies:**

    ```sh
    go mod download
    ```

4. **Run the application:**

    ```sh
    go run .
    ```
