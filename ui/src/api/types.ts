import client from "./apollo";
import type {
        ApolloQueryResult, ObservableQuery, WatchQueryOptions
      } from "@apollo/client";
import { readable } from "svelte/store";
import type { Readable } from "svelte/store";
import gql from "graphql-tag"
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type CreateNode = {
  name: Scalars['String'];
  description: Scalars['String'];
  isFolder: Scalars['Boolean'];
  mimeType?: Maybe<Scalars['String']>;
  file?: Maybe<Scalars['String']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  createNode: Node;
};


export type MutationCreateNodeArgs = {
  input: CreateNode;
};

export type Node = {
  __typename?: 'Node';
  id: Scalars['ID'];
  name: Scalars['String'];
  description: Scalars['String'];
  isFolder: Scalars['Boolean'];
  mimeType?: Maybe<Scalars['String']>;
  children?: Maybe<Array<Node>>;
};

export type Query = {
  __typename?: 'Query';
  getNode: Node;
};


export type QueryGetNodeArgs = {
  path: Scalars['String'];
};

export type GetNodeQueryVariables = Exact<{
  path: Scalars['String'];
}>;


export type GetNodeQuery = { __typename?: 'Query', getNode: { __typename?: 'Node', id: string, name: string, description: string, isFolder: boolean, mimeType?: Maybe<string>, children?: Maybe<Array<{ __typename?: 'Node', id: string, name: string, description: string, isFolder: boolean, mimeType?: Maybe<string> }>> } };


export const GetNodeDoc = gql`
    query GetNode($path: String!) {
  getNode(path: $path) {
    id
    name
    description
    isFolder
    mimeType
    children {
      id
      name
      description
      isFolder
      mimeType
    }
  }
}
    `;
export const GetNode = (
            options: Omit<
              WatchQueryOptions<GetNodeQueryVariables>, 
              "query"
            >
          ): Readable<
            ApolloQueryResult<GetNodeQuery> & {
              query: ObservableQuery<
                GetNodeQuery,
                GetNodeQueryVariables
              >;
            }
          > => {
            const q = client.watchQuery({
              query: GetNodeDoc,
              ...options,
            });
            const result = readable<
              ApolloQueryResult<GetNodeQuery> & {
                query: ObservableQuery<
                  GetNodeQuery,
                  GetNodeQueryVariables
                >;
              }
            >(
              undefined,
              (set) => {
                q.subscribe((v) => {
                  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
                  set({ ...v, query: q });
                });
              }
            );
            return result;
          }
        