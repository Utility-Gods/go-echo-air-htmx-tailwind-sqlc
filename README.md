# 🚀 Go Web Template with HTMX and Tailwind CSS v4

A modern, production-ready web application template using Go, HTMX, and the latest Tailwind CSS v4.

## ✨ Features

- **Go + Echo**: High-performance web framework
- **HTMX**: Build dynamic interfaces without writing JavaScript
- **Tailwind CSS v4**: Latest version with CSS-first configuration
- **templ**: Type-safe HTML templates for Go
- **Task**: Modern task runner and build tool
- **Live Reload**: Development server with hot reloading for Go, templates, and CSS
- **PostgreSQL + sqlc**: Type-safe SQL with Go
- **Air**: Live reload for Go code

## 🛠️ Tech Stack

- **Backend**: Go 1.22+ with Echo framework
- **Frontend**: 
  - HTMX for dynamic interactions
  - Tailwind CSS v4 for styling
  - templ for type-safe templates
- **Database**: PostgreSQL with sqlc for type-safe queries
- **Development**:
  - Task for running development commands
  - Air for live reloading
  - Docker for PostgreSQL (optional)

## 📦 Prerequisites

- Go 1.22 or later
- Task (task runner)
- PostgreSQL 15 or later
- Docker (optional, for containerized database)

## 🚀 Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/your-repo-name
   cd your-repo-name
   ```

2. **Set up the database**

   Using Docker (recommended for development):
   ```bash
   task db-docker-up    # Starts PostgreSQL in Docker
   task db-setup        # Creates database and runs migrations
   ```

   Using local PostgreSQL:
   ```bash
   # Create the database
   createdb your_db_name
   
   # Run migrations
   task db-setup
   ```

3. **Install tools and dependencies**
   ```bash
   task install-tools
   task download-deps
   ```

4. **Configure environment**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

5. **Start development server**
   ```bash
   task dev
   ```

   This will start:
   - Go server with live reload on :6969
   - Tailwind CSS in watch mode
   - templ template compilation in watch mode

6. **Visit [http://localhost:6969](http://localhost:6969)**

## 📁 Project Structure

```
.
├── bin/                    # Binary tools (Tailwind, etc.)
├── cmd/                    # Application entrypoints
├── internal/
│   ├── database/          # Database related code
│   │   ├── migrations/    # SQL migrations
│   │   ├── queries/       # SQL queries
│   │   └── sqlc.yaml      # sqlc configuration
│   ├── handlers/          # HTTP handlers
│   ├── models/            # Data models
│   └── templates/         # templ templates
├── static/
│   ├── css/              # CSS files
│   │   ├── input.css    # Tailwind CSS entry
│   │   └── output.css   # Generated CSS
│   └── js/              # JavaScript files
├── .env.example          # Example environment variables
├── docker-compose.yml    # Docker services definition
└── Taskfile.yml          # Task runner configuration
```

## 🎨 Tailwind CSS v4 Setup

This template uses Tailwind CSS v4 with its new CSS-first configuration:

```css
@import "tailwindcss";
@source "./internal/templates/**/*.{templ,go}";
@plugin "./static/js/daisyui.js" {
    themes: emerald --default, dark --prefers-dark
}

/* Custom styles here */
```

Key features:
- No `tailwind.config.js` needed
- CSS-first configuration
- Built-in source detection
- Native dark mode support
- Modern color system with OKLCH

## 🔧 Development Commands

- `task dev`: Start development server with live reload
- `task build`: Build for production
- `task css-watch`: Watch and compile CSS changes
- `task css-build`: Build CSS for production
- `task templ-generate`: Generate Go code from templates
- `task templ-watch`: Watch and compile template changes

## 📝 Database Commands

- `task db-docker-up`: Start PostgreSQL in Docker
- `task db-docker-down`: Stop PostgreSQL container
- `task db-setup`: Create database and run migrations
- `task db-create-migration name=example`: Create a new migration
- `task db-migrate`: Run pending migrations
- `task db-rollback`: Rollback last migration
- `task db-status`: Check migration status
- `task db-generate`: Generate Go code from SQL queries using sqlc

## 🚀 Deployment

1. **Build for production**
   ```bash
   task build
   ```

2. **Set up the database**
   ```bash
   # Configure your production database URL in .env
   task db-migrate
   ```

3. **Run the server**
   ```bash
   ./bin/server
   ```

## 📚 Learning Resources

- [Echo Framework Documentation](https://echo.labstack.com/)
- [HTMX Documentation](https://htmx.org/docs/)
- [Tailwind CSS v4 Documentation](https://tailwindcss.com/)
- [templ Documentation](https://templ.guide/)
- [Task Documentation](https://taskfile.dev/)
- [sqlc Documentation](https://docs.sqlc.dev/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 