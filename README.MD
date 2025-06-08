# URL Shortener Service

A lightweight and fast URL shortener service
---

## 📦 Project Structure

```
url-shortener/
├── cmd/
|   |── server/            # Entry point
|   |── config/            # config handler
├── internal/
│   ├── handler/           # HTTP handlers
│   ├── service/           # Business logic
│   └── storage/           # In-memory storage
├── test/                  # tests
├── go.mod / go.sum
├── Dockerfile             # Container support
└── README.md
```

---

## 🛠 Setup & Run

### ⚙️ Configuration

To customize the service, create a `config.yaml` file in the project root:

```yaml
port: 8080
host: "http://localhost:8080"
debug_server: true
```

### Steps to generate `config.yaml`:

1. Create a new file named `config.yaml` in the project root directory.
2. Copy and paste the example configuration above into the file.
3. Adjust the values as needed for your environment (e.g., change `port` or `host`).
4. Save the file.

The service will automatically load `config.yaml` on startup if present.

### Run Locally:

```bash
go run ./cmd/server
```

### Run with docker:

```bash
docker build -t url-shortener .
docker run -p 8080:8080 url-shortener
```
---

## ✅ Endpoints

| Method | Endpoint           | Description               |
| ------ | ------------------ | ------------------------- |
| POST   | `/shorten`         | Shortens a given URL      |
| GET    | `/r/:code`         | Redirects to original URL |
| GET    | `/metrics/domains` | Top 3 most used domains   |

---

### Test API:

```bash
# Shorten a URL
curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"url": "https://youtube.com/abc"}'

# Redirect using short code
curl -v http://localhost:8080/r/abc123

# Domain usage metrics
curl http://localhost:8080/metrics/domains
```
### Running Tests

```bash
go test ./internal/service -v
```

---



---