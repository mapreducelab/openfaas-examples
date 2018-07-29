# openfaas-examples
Examples with OpenFaaS and Big Data.

### postgres-query 
Queries PostgreSQL database, uses external dependencies with go package manager `dep`.

### Build & Push & Deploy
```
faas-cli build -f <func.yml> && fass-cli push -f <func.yml> && faas-cli deploy -f <func.yml>
```
