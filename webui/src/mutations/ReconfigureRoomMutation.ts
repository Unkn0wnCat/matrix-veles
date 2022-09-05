import {
    commitMutation
} from "relay-runtime";
import {graphql} from "babel-plugin-relay/macro";
import RelayEnvironment from "../RelayEnvironment";
import {
    ReconfigureRoomMutation as ReconfigureRoomMutationType,
    ReconfigureRoomMutation$data,
    ReconfigureRoomMutation$variables
} from "./__generated__/ReconfigureRoomMutation.graphql";
import {useMutation} from "react-relay";

const mutation = graphql`
    mutation ReconfigureRoomMutation($reconfigureInput: RoomConfigUpdate!) {
        reconfigureRoom(input: $reconfigureInput) {
            id
            active
            deactivated
            name
            roomId
            debug
            adminPowerLevel
            hashCheckerConfig {
                chatNotice
                hashCheckMode
                subscribedLists
            }
        }
    } 
`

export const useReconfigureRoomMutation = () => useMutation<ReconfigureRoomMutationType>(mutation)

const ReconfigureRoomMutation = (input: ReconfigureRoomMutation$variables): Promise<ReconfigureRoomMutation$data> => {
    return new Promise<ReconfigureRoomMutation$data>((resolve, reject) => {
        const variables: ReconfigureRoomMutation$variables = input

        commitMutation<ReconfigureRoomMutationType>(
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

export default ReconfigureRoomMutation