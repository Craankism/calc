# calc
simple calculator application for a webserver in go

- simple calculations (no exponents and stuff)
- basic calculator design with buttons (like windows)
- CI/CD with GitHub Actions (release-please)
- containerization with docker and docker-compose
-  

```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"term": "3 + 4 * 2"}'
```

Work Order:
1. term from website:



2. Parser:



3. Caluclation:



4. result to website


