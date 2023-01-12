# Plain todo API

Pure Api with CRUD operations with lists and tasks with using JWT-auth.

# To run on local machine

### Docker (In process)

---

### Manually

---
First off, clone repository:
    
    https://github.com/Semaffor/go_todo/
Then, in main directory create file `.env` and put applicable values for following keys: `DB_PASSWORD`, `TOKEN_VALID_TIME_HOURS` and `SIGN_KEY`.

Recheck params in `configs/configs.yaml

Run migration file from the main directory

    migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:[port]/postgres?sslmode=disable' up

Finally start the app

    go run cmd/main.go
    
Docs

    http://localhost:8000/swagger/
    
# Databse schema

![image](https://user-images.githubusercontent.com/75733938/212204317-6b009b74-d99b-45ae-ae4b-22ccd5e093a7.png)
