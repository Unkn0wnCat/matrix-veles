/**
 * @generated SignedSource<<6d344122998b6cbd0c9e6d2b2f8d23bd>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type ComponentDashMyLists$variables = {
  count?: number | null;
  first?: string | null;
};
export type ComponentDashMyLists$data = {
  readonly " $fragmentSpreads": FragmentRefs<"DashMyListsFragment">;
};
export type ComponentDashMyLists = {
  variables: ComponentDashMyLists$variables;
  response: ComponentDashMyLists$data;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "count"
  },
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "first"
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "after",
    "variableName": "first"
  },
  {
    "kind": "Variable",
    "name": "first",
    "variableName": "count"
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "ComponentDashMyLists",
    "selections": [
      {
        "args": null,
        "kind": "FragmentSpread",
        "name": "DashMyListsFragment"
      }
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "ComponentDashMyLists",
    "selections": [
      {
        "alias": null,
        "args": (v1/*: any*/),
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
                  {
                    "alias": null,
                    "args": null,
                    "kind": "ScalarField",
                    "name": "id",
                    "storageKey": null
                  },
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
                    "name": "__typename",
                    "storageKey": null
                  }
                ],
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "cursor",
                "storageKey": null
              }
            ],
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "concreteType": "PageInfo",
            "kind": "LinkedField",
            "name": "pageInfo",
            "plural": false,
            "selections": [
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "endCursor",
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "hasNextPage",
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": (v1/*: any*/),
        "filters": null,
        "handle": "connection",
        "key": "ComponentDashMyLists_lists",
        "kind": "LinkedHandle",
        "name": "lists"
      }
    ]
  },
  "params": {
    "cacheID": "56319a434d00fcd6406c4e9aa88de1fa",
    "id": null,
    "metadata": {},
    "name": "ComponentDashMyLists",
    "operationKind": "query",
    "text": "query ComponentDashMyLists(\n  $count: Int\n  $first: String\n) {\n  ...DashMyListsFragment\n}\n\nfragment DashMyListsFragment on Query {\n  lists(after: $first, first: $count) {\n    edges {\n      node {\n        id\n        name\n        __typename\n      }\n      cursor\n    }\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "6f41b7c821add7cf4c8c37b343bd6d2d";

export default node;
