# League of Legends Distributed Voice Chat Application (5-peer WebRTC)

Despite the repeated demand from players, League of Legends currently does not support voice chat in their ranked solo queue matches, which is the game’s primary competitive mode. For nearly 10+ years, players have needed to turn to third-party applications not built for the game, which causes issues such as voice chat raiding and DDOSing.

We want to build a 5-peer voice chat application that League of Legends players can use to communicate and strategize during matches without needing to rely on Discord, which is not designed for strangers to talk temporarily.

This application will be built by utilizing WebRTC, a distributed Peer-to-peer protocol designed by Google that enables users (in this case, players) to communicate over a web browser without exposing their network information to strangers.

To build the web application, we will use the React library to build the user interface and logic, and a yet-to-be-determined language will be used for the simple Signal server and SFU to forward the media channels to clients, reducing the latency.

Check the README.md in each folder for more information about each component