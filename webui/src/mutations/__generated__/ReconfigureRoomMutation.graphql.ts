/**
 * @generated SignedSource<<0e26637d8a4933e2cfcb7d11667fc67c>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type HashCheckerMode = "NOTICE" | "DELETE" | "MUTE" | "BAN" | "%future added value";
export type RoomConfigUpdate = {
  id: string;
  deactivate?: boolean | null;
  debug?: boolean | null;
  adminPowerLevel?: number | null;
  hashChecker?: HashCheckerConfigUpdate | null;
};
export type HashCheckerConfigUpdate = {
  chatNotice?: boolean | null;
  hashCheckMode?: HashCheckerMode | null;
};
export type ReconfigureRoomMutation$variables = {
  reconfigureInput: RoomConfigUpdate;
};
export type ReconfigureRoomMutation$data = {
  readonly reconfigureRoom: {
    readonly id: string;
    readonly active: boolean;
    readonly deactivated: boolean;
    readonly name: string;
    readonly roomId: string;
    readonly debug: boolean;
    readonly adminPowerLevel: number;
    readonly hashCheckerConfig: {
      readonly chatNotice: boolean;
      readonly hashCheckMode: HashCheckerMode;
      readonly subscribedLists: {
        readonly edges: ReadonlyArray<{
          readonly node: {
            readonly id: string;
            readonly name: string;
            readonly tags: ReadonlyArray<string> | null;
          };
        }>;
      } | null;
    };
  };
};
export type ReconfigureRoomMutation = {
  variables: ReconfigureRoomMutation$variables;
  response: ReconfigureRoomMutation$data;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "reconfigureInput"
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
        "name": "input",
        "variableName": "reconfigureInput"
      }
    ],
    "concreteType": "Room",
    "kind": "LinkedField",
    "name": "reconfigureRoom",
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
        "kind": "ScalarField",
        "name": "debug",
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
                      (v2/*: any*/),
                      {
                        "alias": null,
                        "args": null,
                        "kind": "ScalarField",
                        "name": "tags",
                        "storageKey": null
                      }
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
    "name": "ReconfigureRoomMutation",
    "selections": (v3/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "ReconfigureRoomMutation",
    "selections": (v3/*: any*/)
  },
  "params": {
    "cacheID": "0961d774133f607eb0b0145371e73f32",
    "id": null,
    "metadata": {},
    "name": "ReconfigureRoomMutation",
    "operationKind": "mutation",
    "text": "mutation ReconfigureRoomMutation(\n  $reconfigureInput: RoomConfigUpdate!\n) {\n  reconfigureRoom(input: $reconfigureInput) {\n    id\n    active\n    deactivated\n    name\n    roomId\n    debug\n    adminPowerLevel\n    hashCheckerConfig {\n      chatNotice\n      hashCheckMode\n      subscribedLists(first: 100) {\n        edges {\n          node {\n            id\n            name\n            tags\n          }\n        }\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "9041b8d9fb6c1da957feb5f33ec8ac1b";

export default node;
