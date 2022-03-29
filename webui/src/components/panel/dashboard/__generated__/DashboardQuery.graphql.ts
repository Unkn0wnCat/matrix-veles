/**
 * @generated SignedSource<<a34bf324ec46a1d657a8967d773745fe>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type DashboardQuery$variables = {
  first?: string | null;
  count?: number | null;
};
export type DashboardQuery$data = {
  readonly self: {
    readonly username: string;
    readonly id: string;
    readonly admin: boolean | null;
  } | null;
  readonly " $fragmentSpreads": FragmentRefs<"DashMyRoomsFragment" | "DashMyListsFragment">;
};
export type DashboardQuery = {
  variables: DashboardQuery$variables;
  response: DashboardQuery$data;
};

const node: ConcreteRequest = (function(){
var v0 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "count"
},
v1 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "first"
},
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v3 = {
  "alias": null,
  "args": null,
  "concreteType": "User",
  "kind": "LinkedField",
  "name": "self",
  "plural": false,
  "selections": [
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "username",
      "storageKey": null
    },
    (v2/*: any*/),
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "admin",
      "storageKey": null
    }
  ],
  "storageKey": null
},
v4 = {
  "kind": "Variable",
  "name": "after",
  "variableName": "first"
},
v5 = {
  "kind": "Variable",
  "name": "first",
  "variableName": "count"
},
v6 = [
  (v4/*: any*/),
  {
    "kind": "Literal",
    "name": "filter",
    "value": {
      "canEdit": true
    }
  },
  (v5/*: any*/)
],
v7 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "name",
  "storageKey": null
},
v8 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "__typename",
  "storageKey": null
},
v9 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "cursor",
  "storageKey": null
},
v10 = {
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
},
v11 = [
  (v4/*: any*/),
  (v5/*: any*/)
];
return {
  "fragment": {
    "argumentDefinitions": [
      (v0/*: any*/),
      (v1/*: any*/)
    ],
    "kind": "Fragment",
    "metadata": null,
    "name": "DashboardQuery",
    "selections": [
      (v3/*: any*/),
      {
        "args": null,
        "kind": "FragmentSpread",
        "name": "DashMyRoomsFragment"
      },
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
    "argumentDefinitions": [
      (v1/*: any*/),
      (v0/*: any*/)
    ],
    "kind": "Operation",
    "name": "DashboardQuery",
    "selections": [
      (v3/*: any*/),
      {
        "alias": null,
        "args": (v6/*: any*/),
        "concreteType": "RoomConnection",
        "kind": "LinkedField",
        "name": "rooms",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "RoomEdge",
            "kind": "LinkedField",
            "name": "edges",
            "plural": true,
            "selections": [
              {
                "alias": null,
                "args": null,
                "concreteType": "Room",
                "kind": "LinkedField",
                "name": "node",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  (v7/*: any*/),
                  {
                    "alias": null,
                    "args": null,
                    "kind": "ScalarField",
                    "name": "active",
                    "storageKey": null
                  },
                  {
                    "alias": null,
                    "args": null,
                    "kind": "ScalarField",
                    "name": "debug",
                    "storageKey": null
                  },
                  {
                    "alias": null,
                    "args": null,
                    "kind": "ScalarField",
                    "name": "roomId",
                    "storageKey": null
                  },
                  (v8/*: any*/)
                ],
                "storageKey": null
              },
              (v9/*: any*/)
            ],
            "storageKey": null
          },
          (v10/*: any*/)
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": (v6/*: any*/),
        "filters": [
          "filter"
        ],
        "handle": "connection",
        "key": "ComponentDashMyRooms_rooms",
        "kind": "LinkedHandle",
        "name": "rooms"
      },
      {
        "alias": null,
        "args": (v11/*: any*/),
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
                  (v2/*: any*/),
                  (v7/*: any*/),
                  (v8/*: any*/)
                ],
                "storageKey": null
              },
              (v9/*: any*/)
            ],
            "storageKey": null
          },
          (v10/*: any*/)
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": (v11/*: any*/),
        "filters": null,
        "handle": "connection",
        "key": "ComponentDashMyLists_lists",
        "kind": "LinkedHandle",
        "name": "lists"
      }
    ]
  },
  "params": {
    "cacheID": "72b0fa7a61819018fd38de7564f1f3cb",
    "id": null,
    "metadata": {},
    "name": "DashboardQuery",
    "operationKind": "query",
    "text": "query DashboardQuery(\n  $first: String\n  $count: Int\n) {\n  self {\n    username\n    id\n    admin\n  }\n  ...DashMyRoomsFragment\n  ...DashMyListsFragment\n}\n\nfragment DashMyListsFragment on Query {\n  lists(after: $first, first: $count) {\n    edges {\n      node {\n        id\n        name\n        __typename\n      }\n      cursor\n    }\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n\nfragment DashMyRoomsFragment on Query {\n  rooms(after: $first, first: $count, filter: {canEdit: true}) {\n    edges {\n      node {\n        id\n        name\n        active\n        debug\n        roomId\n        __typename\n      }\n      cursor\n    }\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "5a8ab0cabd608a79789053ba85880813";

export default node;
