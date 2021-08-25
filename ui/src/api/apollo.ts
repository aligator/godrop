import {
    ApolloClient,
    InMemoryCache,
} from "@apollo/client/core";

const client = new ApolloClient({
    uri: 'http://localhost:8080/query',
    cache: new InMemoryCache()
});

export default client