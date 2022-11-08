# GraphQL
Example of GraphQL application with Go.

### Setup skeleton project:
---
Follow all steps in the documentation of [gqlgen](https://gqlgen.com/getting-started/#set-up-project)

### Generating files based on `schema.graphqls`:
---
> This command should be executed if your `schema.graphqls` has modified.
```sh
go run github.com/99designs/gqlgen generate
```

### Creating SQLite database:
---
```sh
# For connect on database
sqlite3 db.sqlite
```
```sql
-- For create table
create table categories(
    id string,
    name string,
    description string
);
```

### Running server:
---
```sh
go run cmd/server/server.go
```

### Example of mutations:
```graphql
# Create a new category
mutation createCategoryQl {
  createCategory(input: {
    name: "My first category"
    description: "My first category description"
  }) {
    id
    name
    description
  }
}
```

### Example of queries:
```graphql
# Find all categories
query findAllCategories {
  categories {
    id
    name
    description
  }
}
```