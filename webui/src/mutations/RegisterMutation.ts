import {
    commitMutation
} from "relay-runtime";
import {graphql} from "babel-plugin-relay/macro";
import RelayEnvironment from "../RelayEnvironment";
import {
    RegisterMutation as RegisterMutationType,
    RegisterMutation$data,
    RegisterMutation$variables
} from "./__generated__/RegisterMutation.graphql";

const mutation = graphql`
    mutation RegisterMutation($registerInput: Register!) {
        register(input: $registerInput)
    }
`

const RegisterMutation = (username: string, password: string, mxID: string): Promise<RegisterMutation$data> => {
    return new Promise<RegisterMutation$data>((resolve, reject) => {
        const variables: RegisterMutation$variables = {
            registerInput: {
                username,
                password,
                mxID
            }
        }

        commitMutation<RegisterMutationType>(
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

export default RegisterMutation