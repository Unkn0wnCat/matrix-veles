/**
 * @generated SignedSource<<4f02a0f62b7e8664299d742dc566496a>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type Login = {
  username: string;
  password: string;
};
export type LoginMutation$variables = {
  loginInput: Login;
};
export type LoginMutation$data = {
  readonly login: string;
};
export type LoginMutation = {
  variables: LoginMutation$variables;
  response: LoginMutation$data;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "loginInput"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "input",
        "variableName": "loginInput"
      }
    ],
    "kind": "ScalarField",
    "name": "login",
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "LoginMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "LoginMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "f93ebcc6267e266343c48ca5696aa0f2",
    "id": null,
    "metadata": {},
    "name": "LoginMutation",
    "operationKind": "mutation",
    "text": "mutation LoginMutation(\n  $loginInput: Login!\n) {\n  login(input: $loginInput)\n}\n"
  }
};
})();

(node as any).hash = "f4c8734042b7e93c4d1e63db9879561a";

export default node;
