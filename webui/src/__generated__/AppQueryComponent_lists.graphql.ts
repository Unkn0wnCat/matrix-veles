/**
 * @generated SignedSource<<1ad106fe6b3ac770c83852304ee79afe>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { Fragment, ReaderFragment } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type AppQueryComponent_lists$data = {
  readonly lists: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly name: string;
        readonly id: string;
        readonly tags: ReadonlyArray<string> | null;
      };
    }>;
  };
  readonly " $fragmentType": "AppQueryComponent_lists";
};
export type AppQueryComponent_lists$key = {
  readonly " $data"?: AppQueryComponent_lists$data;
  readonly " $fragmentSpreads": FragmentRefs<"AppQueryComponent_lists">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "AppQueryComponent_lists",
  "selections": [
    {
      "alias": null,
      "args": null,
      "concreteType": "ListConnection",
      "kind": "LinkedField",
      "name": "lists",
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
          ],
          "storageKey": null
        }
      ],
      "storageKey": null
    }
  ],
  "type": "Query",
  "abstractKey": null
};

(node as any).hash = "f7682e7e8dcd1a0ad8275bc5c03dbde8";

export default node;
