/**
 * @generated SignedSource<<92d3c9bc5ecbad19fc01cc3038c80043>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type AppMainQuery$variables = {};
export type AppMainQuery$data = {
  readonly self: {
    readonly admin: boolean | null;
    readonly id: string;
    readonly username: string;
  } | null;
};
export type AppMainQuery = {
  variables: AppMainQuery$variables;
  response: AppMainQuery$data;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
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
        "name": "admin",
        "storageKey": null
      },
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
        "name": "username",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "AppMainQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "AppMainQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "b0e68d6eaa2fc3081f412c76778294fe",
    "id": null,
    "metadata": {},
    "name": "AppMainQuery",
    "operationKind": "query",
    "text": "query AppMainQuery {\n  self {\n    admin\n    id\n    username\n  }\n}\n"
  }
};
})();

(node as any).hash = "4b6ec346f1ba3fa1a215bfd2f2b5bc9e";

export default node;
