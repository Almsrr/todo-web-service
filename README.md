
# Todo Web Service

A small Go project that performs the four basic operations in any system: CRUD. It counts with its own web and mobile app.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Environment Variables

To run this project, you will need a .env file with the following environment variables:

`DB_USER`, `DB_PASS`, `DB_HOST`, `DB_PORT`, `DB_NAME`.

There's `.env.sample` file in the repository. Duplicate it and replace the values of the variables.

## Database

Create the database before running the project. In `database/` directory you can find a `.sql` file. Execute this query in a MYSQL database.

## Run Locally

Clone the project

```bash
  git clone https://github.com/almsrr/todo-web-service.git
```

Go to the project directory

```bash
  cd todo-web-service
```

Install AIR

```bash
  go install github.com/air-verse/air@latest
```

Start the server

```bash
  air
```

## Authors

- [@almsrr](https://www.github.com/almsrr)
