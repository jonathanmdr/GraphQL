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
-- For create table categories
create table categories(
    id string primary key,
    name string,
    description string
);

-- For create table courses
create table courses(
    id string primary key,
    name string,
    description string,
    category_id string,
    foreign key (category_id) references categories(id)
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

# Create a new course
mutation createCourseQl {
  createCourse(input: {
    name: "My first course"
    description: "My first course description"
    categoryId: "8cdd0884-bebc-4af5-a657-b666d62cec18"
  }) {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
```

### Example of queries:
```graphql
# Find all categories with courses
query findAllCategories {
  categories {
    id
    name
    description
    courses {
      id
      name
      description
    }
  }
}

# Find all courses with category
query findAllCourses {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
```