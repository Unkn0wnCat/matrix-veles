import {Environment, Network, RecordSource, Store} from 'relay-runtime';
import fetchGraphQL from "./app/fetchGraphQL";
import {RequestParameters} from "relay-runtime/lib/util/RelayConcreteNode";
import {Variables} from "relay-runtime/lib/util/RelayRuntimeTypes";
import {AnyAction, EnhancedStore} from "@reduxjs/toolkit";
import {ThunkMiddlewareFor} from "@reduxjs/toolkit/dist/getDefaultMiddleware";

const fetchRelay = (auth: GQLAuthObj) => {
    return async (params: RequestParameters, variables: Variables) => {
        return fetchGraphQL(auth, params.text, variables)
    }
}

export type GQLAuthObj = {
    store?: EnhancedStore<any, AnyAction, [ThunkMiddlewareFor<any>]>
    auth?: string
}

const RelayEnvironment = (auth: GQLAuthObj) => new Environment({
    network: Network.create(fetchRelay(auth)),
    store: new Store(new RecordSource()),
});

export default RelayEnvironment