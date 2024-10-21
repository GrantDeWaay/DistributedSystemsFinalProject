# Frontend
The frontend will be designed to be as lightweight and simple as possible to prioratize performace (and to avoid wasting our time)

To create a room, a user will go to the website and click the **CREATE A ROOM**  button or something. This will do a few things:
- Bring the User A to a new page thats kinda like when you join an empty WebEx / Zoom room.
- Make a WebSocket with the signaling server, and specify that there is a new room, which will return a unique ID embedded in a URL

To join a room:
- When users B to E go to this link, they are also connected to the signaling server via WebSocket, and they pass the ID and accept the call
- Then, they will be routed to that page
- Changes are reflected on all user's ends when a new connection is established

There will be buttons to leave the room and mute other users. This will be our MVP, if we have time, I think I found a good way of testing if all of the users are in the same League of Legends game using an API, but that is probably a stretch goal.