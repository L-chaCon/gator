# Gator - RSS Feed Aggregator

A powerful CLI tool for aggregating and managing RSS feeds. Gator allows you to collect RSS feeds from across the internet, store posts in a PostgreSQL database, and browse aggregated content directly from your terminal.

## Features

- **Feed Management**: Add and manage RSS feeds from any website
- **User System**: Register users and manage authentication
- **Follow System**: Follow and unfollow feeds added by other users  
- **Post Aggregation**: Automatically fetch and store new posts from followed feeds
- **Browse Posts**: View summaries of aggregated posts with links to full articles
- **PostgreSQL Storage**: Persistent storage of feeds, users, and posts

## Prerequisites

Before installing Gator, make sure you have the following installed:

- **Go 1.23+**: [Download Go](https://golang.org/dl/)
- **PostgreSQL**: [Download PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the Gator CLI using Go's built-in package manager:

```bash
go install github.com/L-chaCon/gator@latest
```

After installation, the `gator` binary will be available in your `$GOPATH/bin` directory (make sure this is in your PATH).

## Setup

### 1. Database Setup

First, create a PostgreSQL database for Gator:

```sql
CREATE DATABASE gator;
```

### 2. Configuration File

Create a configuration file at `~/.gatorconfig.json` (or in your preferred config location):

```json
{
  "db_url": "postgres://username:password@localhost/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username` and `password` with your PostgreSQL credentials.

### 3. Database Migration

Run any necessary database migrations to set up the required tables (this step may vary depending on your migration setup).

## Usage

### User Management

**Register a new user:**
```bash
gator register <username>
```

**Login as an existing user:**
```bash
gator login <username>
```

**View all registered users:**
```bash
gator users
```

**Reset the database:**
```bash
gator reset
```

### Feed Management

**Add a new RSS feed:**
```bash
gator addfeed <feed_name> <feed_url>
```

**View all available feeds:**
```bash
gator feeds
```

**Follow a feed:**
```bash
gator follow <feed_url>
```

**View feeds you're following:**
```bash
gator following
```

**Unfollow a feed:**
```bash
gator unfollow <feed_url>
```

### Content Browsing

**Browse aggregated posts:**
```bash
gator browse [limit]
```

View summaries of posts from feeds you follow, with optional limit on number of posts displayed.

**Start feed aggregation service:**
```bash
gator agg <time_between_requests>
```

Starts a background service that periodically fetches new posts from all feeds.

## Example Workflow

1. **Register and login:**
   ```bash
   gator register john_doe
   gator login john_doe
   ```

2. **Add some feeds:**
   ```bash
   gator addfeed "Go Blog" https://blog.golang.org/feed.atom
   gator addfeed "Hacker News" https://hnrss.org/frontpage
   ```

3. **Follow the feeds:**
   ```bash
   gator follow https://blog.golang.org/feed.atom
   gator follow https://hnrss.org/frontpage
   ```

4. **Start aggregation:**
   ```bash
   gator agg 60s
   ```

5. **Browse posts:**
   ```bash
   gator browse 10
   ```

## Development

### Running from Source

For development, you can run the application directly:

```bash
git clone https://github.com/L-chaCon/gator.git
cd gator
go run . <command> [args...]
```

### Building

To build a production binary:

```bash
go build -o gator .
```

The resulting binary can be run without the Go toolchain installed.

## Architecture

Gator is built with:

- **Go**: Core application logic and CLI interface
- **PostgreSQL**: Data persistence for users, feeds, and posts  
- **SQLC**: Type-safe SQL query generation
- **Goose**: Database migrations
- **RSS Parsing**: Automatic fetching and parsing of RSS/Atom feeds

## Support

If you encounter any issues or have questions, please open an issue on the GitHub repository.
