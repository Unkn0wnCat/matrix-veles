import eventEmitter from "./eventEmitter";

let channel: BroadcastChannel | null = null;

export type BroadcastMessage<payload> = {
    action: string;
    payload: payload;
};

const sendMessage = (message: BroadcastMessage<any>) => {
    if (!channel) return;

    channel.postMessage(message);
};

const BroadcastEmitter = eventEmitter<{
    message: BroadcastMessage<any>;
}>();

const handleMessage = (ev: MessageEvent) => {
    const data = ev.data as BroadcastMessage<any>;

    if (data.action === "hello_world") {
        console.log("New tab opened and connected.");
        return;
    }

    BroadcastEmitter.emit("message", data);
};

if (typeof window.BroadcastChannel !== "undefined") {
    channel = new BroadcastChannel("veles_main");

    channel.onmessage = handleMessage;

    sendMessage({
        action: "hello_world",
        payload: null,
    });
}

export { sendMessage };

export default BroadcastEmitter;