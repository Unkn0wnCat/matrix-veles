/**
 * @generated SignedSource<<5a8663de6323dd94e27c8d5b53db80eb>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type EntryDetailQuery$variables = {
  id?: string | null;
};
export type EntryDetailQuery$data = {
  readonly entry: {
    readonly id: string;
    readonly tags: ReadonlyArray<string> | null;
  } | null;
};
export type EntryDetailQuery = {
  variables: EntryDetailQuery$variables;
  response: EntryDetailQuery$data;
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
    "concreteType": "Entry",
    "kind": "LinkedField",
    "name": "entry",
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
        "name": "tags",
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
    "name": "EntryDetailQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "EntryDetailQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "1f01a8cbc9fbd149cc3511fbfb5a0929",
    "id": null,
    "metadata": {},
    "name": "EntryDetailQuery",
    "operationKind": "query",
    "text": "query EntryDetailQuery(\n  $id: ID\n) {\n  entry(id: $id) {\n    id\n    tags\n  }\n}\n"
  }
};
})();

(node as any).hash = "2121d3309f0f6813613afae4a045e275";

export default node;
