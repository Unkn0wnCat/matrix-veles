import {
    commitMutation
} from "relay-runtime";
import {graphql} from "babel-plugin-relay/macro";
import RelayEnvironment from "../RelayEnvironment";
import {
    LoginMutation as LoginMutationType,
    LoginMutation$data,
    LoginMutation$variables
} from "./__generated__/LoginMutation.graphql";

const mutation = graphql`
    mutation LoginMutation($loginInput: Login!) {
        login(input: $loginInput)
    } 
`

const LoginMutation = (username: string, password: string): Promise<LoginMutation$data> => {
    return new Promise<LoginMutation$data>((resolve, reject) => {
        const variables: LoginMutation$variables = {
            loginInput: {
                username,
                password
            }
        }

        commitMutation<LoginMutationType>(
            RelayEnvironment({}),
            {
                mutation: mutation,
                variables,
                onCompleted: resolve,
                onError: reject
            }
        )
    })

}

export default LoginMutation