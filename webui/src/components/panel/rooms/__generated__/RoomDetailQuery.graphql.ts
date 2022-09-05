/**
 * @generated SignedSource<<58fc708823f0e42ade18c16f92c3b0d2>>
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
      readonly subscribedLists: ReadonlyArray<string> | null;
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
v1 = [
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
            "args": null,
            "kind": "ScalarField",
            "name": "subscribedLists",
            "storageKey": null
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
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "RoomDetailQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "397376ecc93b3b405d66d01b7fd4de21",
    "id": null,
    "metadata": {},
    "name": "RoomDetailQuery",
    "operationKind": "query",
    "text": "query RoomDetailQuery(\n  $id: ID\n) {\n  room(id: $id) {\n    id\n    active\n    deactivated\n    adminPowerLevel\n    debug\n    name\n    roomId\n    hashCheckerConfig {\n      chatNotice\n      hashCheckMode\n      subscribedLists\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "338b222f1e379335edc6847c401f52f5";

export default node;
