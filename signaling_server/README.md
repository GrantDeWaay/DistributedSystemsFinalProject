# Signaling Server
In WebRTC (Web Real-Time Communication), a signaling server facilitates the exchange of connection information between two clients to establish a peer-to-peer (P2P) connection. This process happens before any direct media transmission occurs. The signaling server itself does not transmit media; it only handles the exchange of metadata required to set up the connection, such as network information and session parameters.

Through the signaling server, two clients can connect without directly exposing sensitive information, such as their IP addresses. The information exchanged between the clients is typically handled using **SDP** (Session Description Protocol) and **ICE** (Interactive Connectivity Establishment).

## SDP (Session Description Protocol)
**SDP** is a protocol used to describe the multimedia session parameters during WebRTC communications. It specifies details such as the type of media (audio, video, text, etc.), codecs, and other connection-related settings. In this case, since this is a voice-only application, the SDP will include details relevant to audio streams only.

## ICE (Interactive Connectivity Establishment protocol)
**ICE** is a framework used to find the best path for data to travel between two peers. It allows WebRTC to overcome issues such as NAT (Network Address Translation) or firewall restrictions by testing various network routes (e.g., STUN/TURN servers). ICE helps ensure a reliable and optimal peer-to-peer connection by determining the most efficient way to connect the clients.

# Connection Setup Process

## 1. Initiating a Connection (Offer):

- When a user (User A) wants to initiate a connection, they first generate an **offer** that contains their SDP (Session Description Protocol) and ICE candidates (potential network routes for the connection).
- This offer is sent to the **signaling server**, which assigns a unique ID to it. The offer remains on the server until another client (User B) retrieves it.
- User A typically maintains a persistent connection with the signaling server, often using a **WebSocket**, to listen for any response from User B.

## 2. Connecting Another User (Answer):

- To connect with User A, User B must retrieve the unique ID corresponding to User A's offer. User B will use their own **WebSocket** connection to communicate with the signaling server and retrieve the offer.
- After receiving User A's offer, User B generates an **answer** (their own SDP) and sends it back to the signaling server.
- The signaling server then forwards this answer to User A. Once both clients have exchanged their SDP and ICE candidates, they have the necessary information to establish a direct peer-to-peer (P2P) connection.

## 3. Managing Connections:

- The signaling server may also moderate or manage active connections. For example, it can enforce application-level rules, such as ensuring only specific users can join a session or limiting the number of participants (e.g., one-on-one calls).

## Special adjustments to our signaling server

Our signaling server will be adjusted to allow only up to 5 users to join a call at the same time.