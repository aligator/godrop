# Example queries

Here are some examples to copy paste into the [GraphQL Playground](http://localhost:8080)

Query:
```graphql
query GetNode($path: String!) {
  getNode(path: $path) {
    id,
    name,
    description,
    isFolder,
    mimeType,
    children {
     	id,
      name,
      description,
      isFolder,
      mimeType
    }
  }
}
```

Query variables:
```json
{
  "path": "/"
}
```