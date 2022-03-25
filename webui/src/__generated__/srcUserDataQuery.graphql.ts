/**
 * @generated SignedSource<<557f6510fc295600d6e0d8aa1cdad3af>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type srcUserDataQuery$variables = {};
export type srcUserDataQuery$data = {
  readonly self: {
    readonly id: string;
    readonly username: string;
    readonly matrixLinks: ReadonlyArray<string> | null;
    readonly admin: boolean | null;
  };
};
export type srcUserDataQuery = {
  variables: srcUserDataQuery$variables;
  response: srcUserDataQuery$data;
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
        "name": "id",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "username",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "matrixLinks",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "admin",
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
    "name": "srcUserDataQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "srcUserDataQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "5a22c874dfb788e95810ee63ec866c39",
    "id": null,
    "metadata": {},
    "name": "srcUserDataQuery",
    "operationKind": "query",
    "text": "query srcUserDataQuery {\n  self {\n    id\n    username\n    matrixLinks\n    admin\n  }\n}\n"
  }
};
})();

(node as any).hash = "22115af4048ca4bee32c21e03d4cd3da";

export default node;
