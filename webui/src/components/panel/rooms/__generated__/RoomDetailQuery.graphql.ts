/**
 * @generated SignedSource<<1663469824a6a66c6959a37976f2a05e>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type HashCheckerMode = "NOTICE" | "DELETE" | "MUTE" | "BAN" | "%future added value";
export type RoomDetailQuery$variables = {
  id?: string | null;
};
export type RoomDetailQuery$data = {
  readonly room: {
    readonly id: string;
    readonly active: boolean;
    readonly deactivated: boolean;
    readonly adminPowerLevel: number;
    readonly debug: boolean;
    readonly name: string;
    readonly roomId: string;
    readonly hashCheckerConfig: {
      readonly chatNotice: boolean;
      readonly hashCheckMode: HashCheckerMode;
      readonly subscribedLists: {
        readonly edges: ReadonlyArray<{
          readonly node: {
            readonly id: string;
            readonly name: string;
          };
        }>;
      };
    };
  } | null;
};
export type RoomDetailQuery = {
  variables: RoomDetailQuery$variables;
  response: RoomDetailQuery$data;
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
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "name",
  "storageKey": null
},
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
    "concreteType": "Room",
    "kind": "LinkedField",
    "name": "room",
    "plural": false,
    "selections": [
      (v1/*: any*/),
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
        "name": "deactivated",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "adminPowerLevel",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "debug",
        "storageKey": null
      },
      (v2/*: any*/),
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
            "concreteType": "ListConnection",
            "kind": "LinkedField",
            "name": "subscribedLists",
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
                      (v2/*: any*/)
                    ],
                    "storageKey": null
                  }
                ],
                "storageKey": null
              }
            ],
            "storageKey": "subscribedLists(first:100)"
          }
        ],
        "storageKey": null
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
    "name": "RoomDetailQuery",
    "selections": (v3/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "RoomDetailQuery",
    "selections": (v3/*: any*/)
  },
  "params": {
    "cacheID": "12bcb055f358f8be4fc522afa4daf9e3",
    "id": null,
    "metadata": {},
    "name": "RoomDetailQuery",
    "operationKind": "query",
    "text": "query RoomDetailQuery(\n  $id: ID\n) {\n  room(id: $id) {\n    id\n    active\n    deactivated\n    adminPowerLevel\n    debug\n    name\n    roomId\n    hashCheckerConfig {\n      chatNotice\n      hashCheckMode\n      subscribedLists(first: 100) {\n        edges {\n          node {\n            id\n            name\n          }\n        }\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "c5785d54de6de8e5986d7119f5dd7527";

export default node;
