import {GQLAuthObj} from "../RelayEnvironment";

async function fetchGraphQL(auth: GQLAuthObj, text: any, variables: any) {
    let headers: Record<string, string> = {
        'Content-Type': 'application/json',
    }

    if(auth.store && auth.store.getState().auth?.jwt) {
        headers["Authorization"] = `Bearer ${auth.store.getState().auth.jwt}`
    }

    if(auth.auth) {
        headers["Authorization"] = `Bearer ${auth.auth}`
    }

    const response = await fetch('/api/query', {
        method: 'POST',
        headers: headers,
        body: JSON.stringify({
            query: text,
            variables,
        }),
    });

    // Get the response as JSON
    return await response.json();
}

export default fetchGraphQL;