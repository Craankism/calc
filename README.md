# calc
simple calculator application for a webserver in go

- simple calculations (no exponents and stuff)
- basic calculator design with buttons (windows calculator as reference)
- CI/CD with GitHub Actions (release-please)
- containerization with docker and docker-compose

![Image](docs/operating_principle.svg)

[Term](#term)<br>
[Parser](#parser)<br>
[Calculation](#calculation)<br>
[Result](#result)<br>

# Term
```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"term": "3 + 4 * 2"}'
```
# Parser
Convert Term to Reverse Polish Notation (RPN)
# Calculation

# Result
Result to website

```bash
curl -X POST http://localhost:8080/calc \
  -H "Content-Type: application/json" \
  -d '{"term": "3 + 4 * 2"}'
```