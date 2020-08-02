---
title: "Design"
date: 2020-08-02T09:37:25-07:00
draft: true
---

# [](#anchor) HyperChat Design Doc

[*s/chat*](https://api.short-d.com/r/chat)

**Date**: 03/29/2020

**Status**: In Progress

# [](#anchor-1) Objective

Build a generic real-time chat infrastructure that is scalable to
billions of users.

# [](#anchor-2) Background

Currently, there are 4 major large scale chat systems in the market:

-   Facebook messenger
-   Whatsapp
-   Slack
-   WeChat
-   Telegram

However, none of them provides easy to use APIs or libraries for small
businesses to create their own custom service chat system or internal
chat app.

# [](#anchor-3) Requirements

# [](#anchor-4) Functional Requirements

# [](#anchor-5) Authentication

-   single sign on

# [](#anchor-6) Users

-   find an user ( replaced with a better way of sharing unique ID )
-   save an user to favorites
-   block user
-   user status & timezone

# [](#anchor-7) Chat

-   1 vs 1 chat with any users ( Passively receive message, offline
    delivery\[transaction\], device registry, connection hub, database,
    read and received message feedback)
-   push notification
-   Editing / Delete / Reply to message

# [](#anchor-8) Group Chat

-   create group chat
-   add friend into group chat
-   remove friends from group chat
-   end the group chat

# [](#anchor-9) Non-functional Requirements

-   Very low latency
-   Consistency(everyone should see the messages in same order)
-   High availability
-   Very secure & encrypted messages
-   Support 1 billion users

# [](#anchor-10) Extended requirements

# [](#anchor-11) Messaging

-   send a picture
-   send a video
-   send emoji

# [](#anchor-12) Extra

-   offline message support
-   see friend tweets
-   video call
-   discover nearby people (Bluetooth)
-   secret chat ( timeout message )
-   survey & vote
-   send chat history as a email
-   admin for group chat and his roles and responsibilities?
-   Broadcast a message to multiple friends
-   Share business card with other users

# [](#anchor-13) Questions

-   Do we want to show the status of the users?
-   Do we want to track the timezone of the users?
-   Do we want to track user's devices?

# [](#anchor-14) UI Mockup

# [](#anchor-15) Icon

# [](#anchor-16) ![](/messenger/Images/design/Pictures/100000000000019E00000380CE9CCA811775B445.jpg){width="3.4272in"
height="7.3634in"}

# [](#anchor-17) Launch Screen

# [](#anchor-18) ![](/messenger/Images/design/Pictures/100000000000019E0000038021CAD32BB56FAC24.jpg){width="3.661in"
height="7.911in"}

# [](#anchor-19) Sign In

# [](#anchor-20) ![](/messenger/Images/design/Pictures/100000000000019E00000380D9CAFD45EFDD4E5D.jpg){width="3.4319in" height="7.411in"}

**Home Screen & Create Group Screens:**

![](/messenger/Images/design/Pictures/10000201000003BE0000028BFCB91E5A892CCC3E.png){width="6.5in"
height="4.4165in"}

# [](#anchor-21) ![](/messenger/Images/design/Pictures/1000020100000358000002100C6225AE137187A6.png){width="6.5in" height="4.0138in"}

# [](#anchor-22) Person to person chat

![](/messenger/Images/design/Pictures/10000201000001360000028BEE7E9C444E07B2BB.png){width="3.2291in" height="6.7811in"}

![](/messenger/Images/design/Pictures/1000020100000325000001CEA94DEF1A57DA184B.png){width="6.5in" height="3.7362in"}

# [](#anchor-23) Edit chat message

![](/messenger/Images/design/Pictures/1000020100000331000001CCBD1B473ADFED866B.png){width="6.5in" height="3.6665in"}

# [](#anchor-24) Delete & reply to chat message

![](/messenger/Images/design/Pictures/1000020100000336000001D428A7C5CEB2273B24.png){width="6.5in" height="3.6945in"}

# [](#anchor-25) Add friend

![](/messenger/Images/design/Pictures/100002010000030D000001C415B7ED6919E0133C.png){width="6.5in" height="3.7638in"}

# [](#anchor-26) Group management

![](/messenger/Images/design/Pictures/10000201000002FA000001A8CCE6A357FD8F1D4E.png){width="6.5in" height="3.611in"}

# [](#anchor-27) Deleted group chat

![](/messenger/Images/design/Pictures/100002010000031C000001CA42CFDDB18DB8072A.png){width="6.5in" height="3.7362in"}

# [](#anchor-28) User profile

![](/messenger/Images/design/Pictures/1000020100000333000001CE457869B52D0EAA5D.png){width="6.5in" height="3.6665in"}

# [](#anchor-29) Notification

# [](#anchor-30) ![](/messenger/Images/design/Pictures/1000020100000320000001C8378BAE8B803CEBA4.png){width="6.5in" height="3.7083in"}

# [](#anchor-31) Public APIs

# [](#anchor-32) Sign in

  -----------------------------------
  getChatUserID() chatUserID:String
  -----------------------------------

We can uniquely identify a chat user with chatUserID in our system

# [](#anchor-33) Chat list

  ------------------------------------------------------------------------------------
```go
  ListChats(userID: string, pageIdx: int, pageSize: int): List&lt;ChatThumbnail&gt;\
  \
  ChatThumbnail {\
   chatID string\
   iconURL string\
   mostRecentMessage ThumbnailMessage\
  }\
  \
  ThumbnailMessage {\
  senderName string\
  body string\
  receivedAt time\
  }
```
  ------------------------------------------------------------------------------------

# [](#anchor-34) Create a peer to peer chat

 Chat is created when becoming favorites

  --------------------------------------------------------------------------------
  createPeerToPeerChat(initiatorID string, receiverUserID string): chatID string
  --------------------------------------------------------------------------------

Comment: consider creating a chatID from directly sending on the first
time.

# [](#anchor-35) Create a group chat

  ------------------------------------------------------------------------------------------------
  createGroupChat(ownerID string, groupName string, friendIDs List&lt;String&gt;): chatID string
  ------------------------------------------------------------------------------------------------

userID: This user is the admin of the group chat by default

Comment: combining peer into a group of one

# [](#anchor-36) Sending a message

  -----------------------------------------------------------------------------------
  sendMessage(senderID string, chatID string, messageBody string): messageID string
  -----------------------------------------------------------------------------------

Comment: should we add a time from client

# [](#anchor-37) System event

  --------------------------------------------------------
  receiveMessageCallback(chatID string, message Message)
  --------------------------------------------------------

This functional is invoked by eventbus when the client received a
message from the server

  ----------------------------------------------------------
  receiveAckCallback(chatID string, ackMessage AckMessage)
  ----------------------------------------------------------

  --------------------------------------------
  Message {\
  messageID string\
  senderID string\
  sentAt string\
  body string\
  // Part of proactively retrieving message\
  repliedTo \*string\
  isEdited boolean\
  isDelete boolean\
  editedAt datetime\
  }\
  \
  AckMessage {\
  messageID string\
  isRead boolean\
  isReceived boolean\
  isSent boolean\
  }
  --------------------------------------------

// Include all the people who read the message

# [](#anchor-38) Editing and deleting a message

  ------------------------------------------------------------------------------------
  editMessage(senderID string, chatID string, messageID string, messageBody string)\
  deleteMessage(senderID string, chatID string, messageID string)\
  // viewDeleteMessage might need to persist\
  \
  Update {\
   messageID: String\
   type: String\
   updatedMessage: \*String\
  }
  ------------------------------------------------------------------------------------

Type: Delete, Edit

# [](#anchor-39) Reply to a message

  ------------------------------------------------------------------------------------------------------
  replyMessage(senderID string, chatID string, messageID string, messageBody string): messageID string
  ------------------------------------------------------------------------------------------------------

// Message being replied will show up right on top of the response for
it.

// If the sender tries to edit a message already being replied by
others, the old message body will be kept on top of the reply. However,
when the user clicks on the old message, then it directs the user to the
edited message. If the message is deleted, we don’t perform any
redirects.

# [](#anchor-40) Social network discovery

// Everyone can chat with others if they can discover the receiver and
they are not blocked by receiver

Optional: **discoverPeopleNearBy**() userIDs List&lt;String&gt; //
Bluetooth

# [](#anchor-41) Favorites

  --------------------------------------------------------------
  getFavorites(userID string): favoriteIDs List&lt;String&gt;\
  addAsFavorite(userID string, favoriteI.D string)
  --------------------------------------------------------------

# [](#anchor-42) Invitation to new users

  --------------------------------------------------------------------
  getInvitationLink(userID string): invitationLink string\
  retrieveUserIDFromInvitation(invitationLink string): userID string
  --------------------------------------------------------------------

# [](#anchor-43) Group management

  ------------------------------------------------------------------------
  getMembers(userID string, groupID string): userIDs List&lt;String&gt;\
  // ChatID is associated with a groupID\
  \
  addUserToGroup(adminID string, groupID string, userID string)\
  removeUser(adminID string, groupID string, userID string)\
  deleteGroup(adminID string, groudID string)\
  updateGroupInfo(adminID string, groupID string, groupInfo GroupInfo)\
  \
  // adminID is retrieved from authentication token\
  \
  GroupInfo {\
  name string\
  icon Array&lt;Byte&gt;\
  }
  ------------------------------------------------------------------------

# [](#anchor-44) Profile

  ---------------------------------------------------
  getUserInfo(userID string): userInfo UserInfo\
  updateUserInfo(userID string, userInfo UserInfo)\
  deleteAccount(userID string)\
  blockUser(userID string, blockID string)\
  \
  UserInfo {\
  name string\
  icon Array&lt;Byte&gt;\
  }
  ---------------------------------------------------

# [](#anchor-45) Notification settings

// app local setting

// Receive message. Users determine whether they are notified

# [](#anchor-46) Request missed updates / messages during offline

+-----------------------------------------------------------------------+
| getLatestTransctionID(chatID: string): String                         |
|                                                                       |
| Response {                                                            |
|                                                                       |
|  snapshot: Snapshot                                                   |
|                                                                       |
|  transactions: Array&lt;Transaction&gt;                               |
|                                                                       |
| }                                                                     |
+-----------------------------------------------------------------------+
| fetchSnapshotAndTransactions(chatID: String, clientTransactionID:     |
| String?, latestTransactionID: String): Response                       |
|                                                                       |
| --------------------------------------------------------------------- |
| ------                                                                |
|                                                                       |
| // Minor optimization                                                 |
|                                                                       |
| // Delta = Missed transactions + snapshots                            |
|                                                                       |
| // Snapshot aggregates changes between start & end transactions       |
|                                                                       |
| Delta {                                                               |
|                                                                       |
|  // Missed new messages that may potentially be edited. Deleted new   |
|                                                                       |
|  // messages won’t show up                                            |
|                                                                       |
|  newMessages: Array&lt;Message&gt;                                    |
|                                                                       |
|  // Read, received, edited or deleted old messages on the client.     |
|                                                                       |
|  updates: Array&lt;Update&gt;                                         |
|                                                                       |
| }\                                                                    |
| fetchDelta(chatID: String, clientTransactionID: String?,              |
| latestTransactionID: String): Delta                                   |
+-----------------------------------------------------------------------+

# [](#anchor-47) Database Schema

# [](#anchor-48) ![](/messenger/Images/design/Pictures/100002010000028A00000276FC117A13CBAB3498.png){width="6.5in"
height="6.3055in"}

Note:

user\_message\_status: Status can be one of \[Read, Received, Sent\]

# [](#anchor-49) Capacity Estimation & Constraint

# [](#anchor-50) Assumptions

-   500 millions daily active users
-   Peer to peer 50 msg/day
-   10 msg/day/group
-   average group size: 50
-   5 groups/user
-   100 bytes/msg

# [](#anchor-51) Storage

5e+8 (50 + 5 . 10) 100 = 5 TB/day = 3.65 PB / 2 yr

# [](#anchor-52) Bandwidth

5e8 (50 . 2 + 5 . 10 . 50) 100 / (60 . 60 . 24) = 1.5 GB/s

# [](#anchor-53) Cache Memory

# [](#anchor-54) High Level Architecture

# [](#anchor-55) Real-time messaging delivery

# [](#anchor-56) Peer to peer messaging

# [](#anchor-57) ![](/messenger/Images/design/Pictures/10000000000003CA000001F0B18B278804C2D925.jpg){width="6.3437in"
height="2.9945in"}

# [](#anchor-58) NAT Penetration

# [](#anchor-59) ![](/messenger/Images/design/Pictures/10000201000007F800000530ADE009BB559A2FBB.png){width="5.6402in"
height="3.6791in"}

# [](#anchor-60) Client/Server Architecture

# [](#anchor-61) ![](/messenger/Images/design/Pictures/10000000000004600000039927663B3F794D80D3.png){width="6.5in"
height="5.3472in"}

The system has 3 layers:

-   Message transfer layer: receive and route message inside the chat
    cluster
-   Persistence layer: save messages for future retrieval
-   Delivery layer: deliver the message to the receiving clients

![](/messenger/Images/design/Pictures/10000000000009C4000006AEB2DFAE1EC49E38D5.png){width="6.5in"
height="4.4445in"}

# [](#anchor-62) ![](/messenger/Images/design/Pictures/10000000000009C4000006EF010ABA0BC5BD98FD.png){width="6.5in"
height="4.611in"}

3 layers

1)  Delivery layer
2)  Message transfer layer: Route message between servers
3)  Persistence layer

0 =&gt; Background notification: APNS, FCM

1 =&gt; Long live socket, WebSocket, MQTT, Long-poll, polling

2 =&gt; Direct RPC between server / Kafka

3 =&gt; HBase ( Column based NoSQL ) / Cassandra ( NoSQL ) / MongoDB (
NoSQL ) / Postgres ( SQL ) / MySQL ( SQL ) \[DynamoDB ( NoSQL ), Aurora
(SQL), BigTable (NoSQL), Spanner ( SQL )

1 + 2 =&gt; Ejabberd

2 + 3 =&gt; Firebase ( NoSQL )

Multiple client delivery & ACK priorities?

Assume 1 user only has 1 device for v1.

# [](#anchor-63) Recommendation

1)  Use client/server architecture for V1
2)  Define the interface of 3 layers
3)  Message Transfer Layer: Use Kafka as the global message queue for V1
    ( Voted with 7 Kafka and 2 RPC)
4)  Use WebSocket as real-time delivery protocol for V1 because it’s
    easier to implement ( Voted with Raw Socket: 1, WebSocket: 4, MQTT:
    4, Long-pull: 2)
5)  Use NoSQL as backend data store ( Vote NoSQL:7 SQL: 3 )

Message status update?

# [](#anchor-64) **Detailed Component Design**

# [](#anchor-65) Real-time messaging
delivery![](/messenger/Images/design/Pictures/10000000000009C400000559C74D4DD02C7AA6D3.png){width="5.6516in"
height="3.089in"}

Connection Hub

Input

-   Business logic

> {

message:

clientID:

> }

-   Device registry

> Health check ( 100 ms )

-   Client

> Request for connection

Output

-   Connect with client

# [](#anchor-66) Notification when app is not running

![](/messenger/Images/design/Pictures/1000000000000779000004B0F5B0AD7F039FECF1.png){width="4.8075in"
height="3.0201in"}

# [](#anchor-67) Message transport

![](/messenger/Images/design/Pictures/10000000000009BE000007C40506BB232BBC5EBF.png){width="6.5in"
height="5.1807in"}

# [](#anchor-68) 

# [](#anchor-69) 

# [](#anchor-70) 

# [](#anchor-71) Receive & read message feedback

Online

![](/messenger/Images/design/Pictures/1000000000000751000006636BBE088957D232B9.png){width="4.9555in"
height="4.328in"}

Offline

![](/messenger/Images/design/Pictures/1000000000000801000005BCC8E58124E84D3104.png){width="6.5in"
height="4.6528in"}

# [](#anchor-72) Unique ID generation

-   user
-   chat
-   messages in each chat ( counter service )

# [](#anchor-73) Message update

![](/messenger/Images/design/Pictures/1000000000000386000001722B770B11FB314984.png){width="6.5in"
height="2.6665in"}

Mutable State vs Transaction

# [](#anchor-74) Edit message

Modify the previous message

# [](#anchor-75) Delete message

# [](#anchor-76) Reply to message

Create a new message

# [](#anchor-77) Persistence

# [](#anchor-78) Protocol negotiation

# [](#anchor-79) Send message detail

![](/messenger/Images/design/Pictures/10000000000004A20000055D8287F9BBC3A39D01.png){width="6.5in"
height="7.528in"}

# [](#anchor-80) System Behavior Priority

1)  Audio & video chat connection initiation
2)  Real time message delivery
3)  Fetch missed messages & updates when going back online
4)  Download file, audio, video from static media server

# [](#anchor-81) Message Delivery Priority

Allow only one active device at time

1)  Web notification preferred
2)  Mobile notification preferred
3)  Smart watch notification preferred

Or based on user settings

# [](#anchor-82) Data Storage

-   User data in DB
-   Uploaded image, audio & video (file server)
-   Client cache

    -   Authentication token

    -   chatID =&gt; lastTransactionID (key-value pair)

    -   peer status

    -   chatHistory

    -   downloaded media

# [](#anchor-83) **Data Partitioning**

# [](#anchor-84) **Fault Tolerance**

# [](#anchor-85) Point of failures

1)  Network congested

2)  Network temporarily disconnected

3)  Network permanently disconnected

4)  Server crashes

5)  Server restart

6)  Disk out of storage

7)  Disk failures

8)  Service pauses

    a)  Garbage collector stops the world
    b)  Memory swapping to disk
    c)  Different process hogging the CPU

9)  Lose of logs due to crashes ( acceptable because it’s only for
    > internal debugging & monitoring )

# [](#anchor-86) **Security**

# [](#anchor-87) **Caching**

# [](#anchor-88) **Testing**

# [](#anchor-89) **Appendix**

[*System Diagrams*](https://api.short-d.com/r/chat-diagram)

# [](#anchor-90) **References**

# [](#anchor-91) **Contributions**

Please add your name if you have already contributed in the following
section. If you are planning to contribute, please mark \[TODO\] beside
your name.

# [](#anchor-92) P2P Architecture

# [](#anchor-93) C/S Architecture

[*Yang Liu*](https://www.linkedin.com/in/yang-liu-416b8669/)

[*Ajay Tanikonda*](https://www.linkedin.com/in/atanikonda)

[*Sejal Shinde*](http://www.linkedin.com/in/sej-ks)

[*Frankie Liu*](https://www.linkedin.com/in/frankieyliu)

*Rohith Nedunuri*\[TODO\]

# [](#anchor-94) API Design

[*Yang Liu*](https://www.linkedin.com/in/yang-liu-416b8669/)

[*Anil Kumar Nandamuri*](https://www.linkedin.com/in/anilnandamuri/)

[*Rohith Nedunuri*](https://www.linkedin.com/in/rohithnedunuri/)\
[*Sejal Shinde*](http://www.linkedin.com/in/sej-ks)

[*Frankie Liu*](https://www.linkedin.com/in/frankieyliu)

[*Vinod Krishnan*](https://www.linkedin.com/in/vinsinin/) \[TODO\]

# [](#anchor-95) Database Schema

[*Yang Liu*](https://www.linkedin.com/in/yang-liu-416b8669/)

[*Anil Kumar Nandamuri*](https://www.linkedin.com/in/anilnandamuri/)

[*Rohith Nedunuri*](https://www.linkedin.com/in/rohithnedunuri/)\
[*Sejal Shinde*](http://www.linkedin.com/in/sej-ks)

[*Frankie Liu*](https://www.linkedin.com/in/frankieyliu)

[*Vinod Krishnan*](https://www.linkedin.com/in/vinsinin/) \[TODO\]

# [](#anchor-96) Message Queue ( Kafka )

[*Rohith
Nedunuri*](https://www.linkedin.com/in/rohithnedunuri/)\[TODO\]\
[*Sejal Shinde*](http://www.linkedin.com/in/sej-ks) \[TODO\]

[*Abhijit kajranjkar*](https://www.linkedin.com/in/abhijitkaranjkar/)
\[TODO\]

[*Vinod Krishnan*](https://www.linkedin.com/in/vinsinin/) \[TODO\]

# [](#anchor-97) Fault Tolerance

[*Yang Liu*](https://www.linkedin.com/in/yang-liu-416b8669/)

[*Rohith
Nedunuri*](https://www.linkedin.com/in/rohithnedunuri/)\[TODO\]\
[*Sejal Shinde*](http://www.linkedin.com/in/sej-ks) \[TODO\]

# [](#anchor-98) Transaction System

[*Yang Liu*](https://www.linkedin.com/in/yang-liu-416b8669/)

[*Frankie Liu*](https://www.linkedin.com/in/frankieyliu)\
[*Sejal Shinde*](http://www.linkedin.com/in/sej-ks) \[TODO\]

[*Vinod Krishnan*](https://www.linkedin.com/in/vinsinin/) \[TODO\]

[*Rohith Nedunuri*](https://www.linkedin.com/in/rohithnedunuri/)\[TODO\]

# [](#anchor-99) UI Mockup

[*Yang Liu*](https://www.linkedin.com/in/yang-liu-416b8669/)

[*Ajay Tanikonda*](https://www.linkedin.com/in/atanikonda)\
[*Sejal Shinde*](http://www.linkedin.com/in/sej-ks)

[*Frankie Liu*](https://www.linkedin.com/in/frankieyliu)

Documentation

[*Anil Kumar Nandamuri*](https://www.linkedin.com/in/anilnandamuri/)

[*Frankie Liu*](https://www.linkedin.com/in/frankieyliu)

Coding

Project/Product Management


