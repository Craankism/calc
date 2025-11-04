# calc
simple calculator on local host website

- simple calculations (no exponents and stuff)
- basic calculator design (like windows)
- CI/CD with GitHub Actions (release-please)

```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"term": "3 + 4 * 2"}'
```