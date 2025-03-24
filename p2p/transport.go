package p2p

// Peer is an interface that represents a remote node in the network.
type Peer interface {}

// Transport is anything that handles the communication between the nodes in the network.
// It could be a TCP connection, a UDP connection, or even a Websocket connection.
type Transport interface {}