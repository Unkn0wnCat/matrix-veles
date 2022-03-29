/**
 * @generated SignedSource<<a4ab3c123e20e4c610b878fd82d98eb5>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type ComponentDashMyRooms$variables = {
  count?: number | null;
  first?: string | null;
};
export type ComponentDashMyRooms$data = {
  readonly " $fragmentSpreads": FragmentRefs<"DashMyRoomsFragment">;
};
export type ComponentDashMyRooms = {
  variables: ComponentDashMyRooms$variables;
  response: ComponentDashMyRooms$data;
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
    "kind": "Literal",
    "name": "filter",
    "value": {
      "canEdit": true
    }
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
    "name": "ComponentDashMyRooms",
    "selections": [
      {
        "args": null,
        "kind": "FragmentSpread",
        "name": "DashMyRoomsFragment"
      }
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "ComponentDashMyRooms",
    "selections": [
      {
        "alias": null,
        "args": (v1/*: any*/),
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
        "filters": [
          "filter"
        ],
        "handle": "connection",
        "key": "ComponentDashMyRooms_rooms",
        "kind": "LinkedHandle",
        "name": "rooms"
      }
    ]
  },
  "params": {
    "cacheID": "900fab96de8f0dd453020c4443727ed4",
    "id": null,
    "metadata": {},
    "name": "ComponentDashMyRooms",
    "operationKind": "query",
    "text": "query ComponentDashMyRooms(\n  $count: Int\n  $first: String\n) {\n  ...DashMyRoomsFragment\n}\n\nfragment DashMyRoomsFragment on Query {\n  rooms(after: $first, first: $count, filter: {canEdit: true}) {\n    edges {\n      node {\n        id\n        name\n        active\n        debug\n        roomId\n        __typename\n      }\n      cursor\n    }\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "fc18a3a1a32649d6d9fd694500167a4c";

export default node;
