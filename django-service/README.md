# Learning Module Django REST API

A Django REST Framework API service for managing learning modules with token-based authentication.

## Tech Stack

- Python 3.8+
- Django 5.2.1
- Django REST Framework 3.16.0
- SQLite (development) / PostgreSQL (production-ready)
- Token Authentication

## Project Structure

```
django-service/
├── api/
│   ├── migrations/        # Database migrations
│   ├── models.py         # Data models
│   ├── serializers.py    # Data serializers
│   ├── urls.py          # URL routing
│   └── views.py         # API views
├── django-service/       # Project settings
├── manage.py            # Django management script
└── requirements.txt     # Python dependencies
```

## Prerequisites

- Python 3.8 or higher
- pip (Python package manager)
- virtualenv (recommended)

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd django-service
   ```

2. Create and activate a virtual environment: (optional, I haven't done this)
   ```bash
   # Windows
   python -m venv venv
   venv\Scripts\activate

   # Unix/MacOS
   python -m venv venv
   source venv/bin/activate
   ```

3. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

4. Set up the database:
   ```bash
   python manage.py makemigrations
   python manage.py migrate
   ```

5. Create a superuser (admin):
   ```bash
   python manage.py createsuperuser
   ```

6. Start the development server:
   ```bash
   python manage.py runserver
   ```

The server will start at `http://localhost:8000/`

## API Endpoints

Base URL: `http://localhost:8000/api/v1/`

### Authentication

#### Register a new user
```http
POST /register/
Content-Type: application/json

{
    "username": "your_username",
    "password": "your_password",
    "email": "your_email@example.com"  // optional
}
```

#### Get authentication token
```http
POST /api-token-auth/
Content-Type: application/json

{
    "username": "your_username",
    "password": "your_password"
}
```

Response:
```json
{
    "token": "your_auth_token_here"
}
```

For all authenticated requests, include the token in the header:
```
Authorization: Token your_auth_token_here
```
note "Token " this is very important for your authentication header to work

### Learning Modules API

#### List all modules
```http
GET /modules/
Authorization: Token your_auth_token_here
```

#### Create a new module
```http
POST /modules/
Authorization: Token your_auth_token_here
Content-Type: application/json

{
    "subject": "Mathematics",
    "grade": "10th",
    "chapter": "Algebra",
    "content": {
        "sections": ["Introduction", "Basic Concepts"],
        "exercises": ["Exercise 1", "Exercise 2"]
    }
}
```

#### Get a specific module
```http
GET /modules/{id}/
Authorization: Token your_auth_token_here
```

#### Update a module
```http
PUT /modules/{id}/
Authorization: Token your_auth_token_here
Content-Type: application/json

{
    "subject": "Updated Mathematics",
    "grade": "11th",
    "chapter": "Advanced Algebra",
    "content": {
        "sections": ["Advanced Concepts"],
        "exercises": ["Exercise 3"]
    }
}
```

#### Delete a module
```http
DELETE /modules/{id}/
Authorization: Token your_auth_token_here
```

## Response Format

### Module Object
```json
{
    "id": 1,
    "subject": "Mathematics",
    "grade": "10th",
    "chapter": "Algebra",
    "content": {
        "sections": ["Introduction", "Basic Concepts"],
        "exercises": ["Exercise 1", "Exercise 2"]
    },
    "created_at": "2024-03-14T10:00:00Z",
    "updated_at": "2024-03-14T10:00:00Z",
    "created_by": "username"
}
```

## Development

### Running Tests
```bash
python manage.py test
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
