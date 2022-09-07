/**
 * @generated SignedSource<<387ef79509f6e2371f71fa3d72d073fd>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type ListDetailQuery$variables = {
  id?: string | null;
};
export type ListDetailQuery$data = {
  readonly list: {
    readonly id: string;
    readonly name: string;
    readonly tags: ReadonlyArray<string> | null;
    readonly creator: {
      readonly id: string;
      readonly username: string;
      readonly matrixLinks: ReadonlyArray<string> | null;
    };
    readonly maintainers: {
      readonly edges: ReadonlyArray<{
        readonly node: {
          readonly id: string;
          readonly username: string;
          readonly matrixLinks: ReadonlyArray<string> | null;
        };
      }>;
    };
  } | null;
};
export type ListDetailQuery = {
  variables: ListDetailQuery$variables;
  response: ListDetailQuery$data;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "id"
  }
],
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v2 = [
  (v1/*: any*/),
  {
    "alias": null,
    "args": null,
    "kind": "ScalarField",
    "name": "username",
    "storageKey": null
  },
  {
    "alias": null,
    "args": null,
    "kind": "ScalarField",
    "name": "matrixLinks",
    "storageKey": null
  }
],
v3 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "id",
        "variableName": "id"
      }
    ],
    "concreteType": "List",
    "kind": "LinkedField",
    "name": "list",
    "plural": false,
    "selections": [
      (v1/*: any*/),
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "name",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "tags",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "concreteType": "User",
        "kind": "LinkedField",
        "name": "creator",
        "plural": false,
        "selections": (v2/*: any*/),
        "storageKey": null
      },
      {
        "alias": null,
        "args": [
          {
            "kind": "Literal",
            "name": "first",
            "value": 100
          }
        ],
        "concreteType": "UserConnection",
        "kind": "LinkedField",
        "name": "maintainers",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "UserEdge",
            "kind": "LinkedField",
            "name": "edges",
            "plural": true,
            "selections": [
              {
                "alias": null,
                "args": null,
                "concreteType": "User",
                "kind": "LinkedField",
                "name": "node",
                "plural": false,
                "selections": (v2/*: any*/),
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": "maintainers(first:100)"
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "ListDetailQuery",
    "selections": (v3/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "ListDetailQuery",
    "selections": (v3/*: any*/)
  },
  "params": {
    "cacheID": "25b9482991b61f15c2856c99875b239e",
    "id": null,
    "metadata": {},
    "name": "ListDetailQuery",
    "operationKind": "query",
    "text": "query ListDetailQuery(\n  $id: ID\n) {\n  list(id: $id) {\n    id\n    name\n    tags\n    creator {\n      id\n      username\n      matrixLinks\n    }\n    maintainers(first: 100) {\n      edges {\n        node {\n          id\n          username\n          matrixLinks\n        }\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "7cef889868efd30077c76c270e95474e";

export default node;
