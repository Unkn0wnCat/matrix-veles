/**
 * @generated SignedSource<<4b6e4870f3f95cad6fceb76ce82c2ca8>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment, RefetchableFragment } from 'relay-runtime';
export type HashCheckerMode = "NOTICE" | "DELETE" | "MUTE" | "BAN" | "%future added value";
import { FragmentRefs } from "relay-runtime";
export type DashboardQueryLists$data = {
  readonly rooms: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly id: string;
        readonly active: boolean;
        readonly debug: boolean;
        readonly hashCheckerConfig: {
          readonly chatNotice: boolean;
          readonly hashCheckMode: HashCheckerMode;
        };
      };
    }>;
  } | null;
  readonly " $fragmentType": "DashboardQueryLists";
};
export type DashboardQueryLists$key = {
  readonly " $data"?: DashboardQueryLists$data;
  readonly " $fragmentSpreads": FragmentRefs<"DashboardQueryLists">;
};

const node: ReaderFragment = (function(){
var v0 = [
  "rooms"
];
return {
  "argumentDefinitions": [
    {
      "kind": "RootArgument",
      "name": "count"
    },
    {
      "kind": "RootArgument",
      "name": "first"
    }
  ],
  "kind": "Fragment",
  "metadata": {
    "connection": [
      {
        "count": "count",
        "cursor": "first",
        "direction": "forward",
        "path": (v0/*: any*/)
      }
    ],
    "refetch": {
      "connection": {
        "forward": {
          "count": "count",
          "cursor": "first"
        },
        "backward": null,
        "path": (v0/*: any*/)
      },
      "fragmentPathInResult": [],
      "operation": require('./DashboardListsQuery.graphql')
    }
  },
  "name": "DashboardQueryLists",
  "selections": [
    {
      "alias": "rooms",
      "args": [
        {
          "kind": "Literal",
          "name": "filter",
          "value": {
            "canEdit": true
          }
        }
      ],
      "concreteType": "RoomConnection",
      "kind": "LinkedField",
      "name": "__DashboardQueryLists_rooms_connection",
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
                  "concreteType": "HashCheckerConfig",
                  "kind": "LinkedField",
                  "name": "hashCheckerConfig",
                  "plural": false,
                  "selections": [
                    {
                      "alias": null,
                      "args": null,
                      "kind": "ScalarField",
                      "name": "chatNotice",
                      "storageKey": null
                    },
                    {
                      "alias": null,
                      "args": null,
                      "kind": "ScalarField",
                      "name": "hashCheckMode",
                      "storageKey": null
                    }
                  ],
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
      "storageKey": "__DashboardQueryLists_rooms_connection(filter:{\"canEdit\":true})"
    }
  ],
  "type": "Query",
  "abstractKey": null
};
})();

(node as any).hash = "56aca4bdff70334e5602099f0fe53156";

export default node;
