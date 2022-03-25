/**
 * @generated SignedSource<<de5d7bef0bc8b76c051f6343d17a8f36>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type srcListsQuery$variables = {};
export type srcListsQuery$data = {
  readonly lists: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly name: string;
        readonly tags: ReadonlyArray<string> | null;
        readonly entries: {
          readonly edges: ReadonlyArray<{
            readonly node: {
              readonly id: string;
              readonly hashValue: string;
            };
          }>;
        } | null;
      };
    }>;
  };
};
export type srcListsQuery = {
  variables: srcListsQuery$variables;
  response: srcListsQuery$data;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "kind": "Literal",
    "name": "first",
    "value": 20
  }
],
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "name",
  "storageKey": null
},
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "tags",
  "storageKey": null
},
v3 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v4 = {
  "alias": null,
  "args": [
    {
      "kind": "Literal",
      "name": "first",
      "value": 5
    }
  ],
  "concreteType": "EntryConnection",
  "kind": "LinkedField",
  "name": "entries",
  "plural": false,
  "selections": [
    {
      "alias": null,
      "args": null,
      "concreteType": "EntryEdge",
      "kind": "LinkedField",
      "name": "edges",
      "plural": true,
      "selections": [
        {
          "alias": null,
          "args": null,
          "concreteType": "Entry",
          "kind": "LinkedField",
          "name": "node",
          "plural": false,
          "selections": [
            (v3/*: any*/),
            {
              "alias": null,
              "args": null,
              "kind": "ScalarField",
              "name": "hashValue",
              "storageKey": null
            }
          ],
          "storageKey": null
        }
      ],
      "storageKey": null
    }
  ],
  "storageKey": "entries(first:5)"
};
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "srcListsQuery",
    "selections": [
      {
        "alias": null,
        "args": (v0/*: any*/),
        "concreteType": "ListConnection",
        "kind": "LinkedField",
        "name": "lists",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "ListEdge",
            "kind": "LinkedField",
            "name": "edges",
            "plural": true,
            "selections": [
              {
                "alias": null,
                "args": null,
                "concreteType": "List",
                "kind": "LinkedField",
                "name": "node",
                "plural": false,
                "selections": [
                  (v1/*: any*/),
                  (v2/*: any*/),
                  (v4/*: any*/)
                ],
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": "lists(first:20)"
      }
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "srcListsQuery",
    "selections": [
      {
        "alias": null,
        "args": (v0/*: any*/),
        "concreteType": "ListConnection",
        "kind": "LinkedField",
        "name": "lists",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "ListEdge",
            "kind": "LinkedField",
            "name": "edges",
            "plural": true,
            "selections": [
              {
                "alias": null,
                "args": null,
                "concreteType": "List",
                "kind": "LinkedField",
                "name": "node",
                "plural": false,
                "selections": [
                  (v1/*: any*/),
                  (v2/*: any*/),
                  (v4/*: any*/),
                  (v3/*: any*/)
                ],
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": "lists(first:20)"
      }
    ]
  },
  "params": {
    "cacheID": "27095434122932f164ae88c05f4cc6b9",
    "id": null,
    "metadata": {},
    "name": "srcListsQuery",
    "operationKind": "query",
    "text": "query srcListsQuery {\n  lists(first: 20) {\n    edges {\n      node {\n        name\n        tags\n        entries(first: 5) {\n          edges {\n            node {\n              id\n              hashValue\n            }\n          }\n        }\n        id\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "ac81a444112c67683994c05979dedb6f";

export default node;
