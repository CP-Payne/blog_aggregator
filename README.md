
# RSS Blog Aggregator

## Overview

RSS Blog Aggregator is a web service designed to aggregate and manage RSS feeds, allowing users to subscribe to, manage, and follow their favourite blogs and news sources in one centralized location. This service provides a set of RESTful endpoints to handle user registration, feed management, and post retrieval.

## API Endpoints

Below is a list of available endpoints, grouped by functionality. Detailed documentation and examples for each endpoint is available in the `/doc` folder, which you can access by clicking on the endpoint name.

### Users

- [`POST /v1/users`](createUser.md) - Create a new user (returns an API key)
- [`GET /v1/users`](getUser.md) - Get user information (**Authenticated**)

### Feeds

- [`POST /v1/feeds`](createFeeds.md) - Add a new feed and automatically follow it (**Authenticated**)
- [`GET /v1/feeds`](getFeeds.md) - Return a list of all feeds

### Following Feeds

- [`POST /v1/feed_follows`](createFeedFollow.md) - Follow a feed (**Authenticated**)
- [`GET /v1/feed_follows`](getFeedFollows.md) - Retrieve a list of feeds a user is following (**Authenticated**)
- [`DELETE /v1/feed_follows/{feedFollowID}`](deleteFeedFollow.md) - Unfollow a specific feed (**Authenticated**)

### Posts

- [`GET /v1/posts`](getPosts.md) - Retrieves posts from the RSS feeds that the user is subscribed to (**Authenticated**)


## Current Data Storage Implementation

The application currently uses a PostgreSQL database.

## Installation

To get the RSS Blog Aggregator up and running on your local machine, follow these steps:

### Prerequisites

- Ensure you have Go installed on your system. This project requires Go version 1.22 or later.
- Git should be installed to clone the repository.

### Cloning the Repository

First, clone the RSS Blog Aggregator repository to your local machine using Git. Open your terminal, and run the following command:

```bash
git clone https://github.com/CP-Payne/blog_aggregator
cd blog_aggregator
```

### Installing Dependencies
The project uses `go.mod` for managing dependencies. To install all the necessary dependencies, run:

```bash
go mod download
```

This command will download and install all the required packages and dependencies for this project.

### Environment Setup

Before running the `blog_aggregator`, you need to set up the required environment variables. Create a `.env` file in the root directory of the project with the following content:

```plaintext
PORT=":<port>"
DSN="<your-postgres-connection-string>"
```

- Replace `<port>` with the port number where you want the server to run (e.g., `8080`).
- Replace `<your-postgres-connection-string>` with your PostgreSQL database connection string. The format for the connection string is:
```plaintext
protocol://username:password@host:port/database?sslmode=disable
```
Here is what you should include based on your PostgreSQL setup:

- `protocol`: This is usually `postgres`.
- `username`: Your database username.
- `password`: Your database password.
- `host`: The hostname of your database server.
- `port`: The port number your database runs on.
- `database`: The name of your database.

### Database Setup

Before running the server, you need to create the database schema. You can run the SQL scripts manually in the order specified in the `/sql/schema` folder or use `goose` for automatic migrations. Install `goose` and execute the following command within the schema directory:

```bash
goose postgres "protocol://username:password@host:port/database" up
```

### Running the Server

To start the server, simply run the following command from the root of the project:
```bash
go run ./cmd/web
```

After running the above command, the API server should be up and will listen for requests on the port specified in the `.env` file. You can access it via `http://localhost:<port>`.

### Testing Endpoints

You can test the endpoints using tools like `curl` or Postman. For example, to get a list of all feeds, you might use:

```bash
curl -X GET 'http://localhost:8080/v1/feeds'
```




## Future Enhancements

- Support pagination of the endpoints that can return many items
- Support different options for sorting and filtering posts using query parameters
- Add bookmarking or "liking" to posts
- Create a simple web UI that uses your backend API