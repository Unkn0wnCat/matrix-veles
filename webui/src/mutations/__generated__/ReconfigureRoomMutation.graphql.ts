/**
 * @generated SignedSource<<c6a0a35dc80ab80f3fc8e42c6fe8a0f6>>
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
      readonly subscribedLists: ReadonlyArray<string> | null;
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
v1 = [
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
    "name": "ReconfigureRoomMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "ReconfigureRoomMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "4186d7d18c6230e79d890ca0043dc104",
    "id": null,
    "metadata": {},
    "name": "ReconfigureRoomMutation",
    "operationKind": "mutation",
    "text": "mutation ReconfigureRoomMutation(\n  $reconfigureInput: RoomConfigUpdate!\n) {\n  reconfigureRoom(input: $reconfigureInput) {\n    id\n    active\n    deactivated\n    name\n    roomId\n    debug\n    adminPowerLevel\n    hashCheckerConfig {\n      chatNotice\n      hashCheckMode\n      subscribedLists\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "e4ccf905922a56c92f44336d58c96a53";

export default node;
