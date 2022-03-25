/**
 * @generated SignedSource<<a9d897ad5a18e0863fbc64c673b079ec>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type Register = {
  username: string;
  password: string;
  mxID: string;
};
export type RegisterMutation$variables = {
  registerInput: Register;
};
export type RegisterMutation$data = {
  readonly register: string;
};
export type RegisterMutation = {
  variables: RegisterMutation$variables;
  response: RegisterMutation$data;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "registerInput"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "input",
        "variableName": "registerInput"
      }
    ],
    "kind": "ScalarField",
    "name": "register",
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "RegisterMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "RegisterMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "d49e8f643837ce5419c3a7a0a67d3bf0",
    "id": null,
    "metadata": {},
    "name": "RegisterMutation",
    "operationKind": "mutation",
    "text": "mutation RegisterMutation(\n  $registerInput: Register!\n) {\n  register(input: $registerInput)\n}\n"
  }
};
})();

(node as any).hash = "152daf2d7a917ffaafddee2b18453aa8";

export default node;
