# personal-shopper

### GraphQL commands
GraphQL is pretty simple once you get the hang of it. The root types are queries and mutations

Example Mutation
```
mutation createImageData {
  createImageData(
    imageData: {
      name: "John Doe"
    }
  ) {
    id
  }
}
```

Example Query
```
query fetchImages {
  images {
    id,
    name
  }
}
```