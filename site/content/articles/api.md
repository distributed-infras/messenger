---
title: "Api"
date: 2020-08-02T17:00:17-07:00
draft: true
---


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

```go
createPeerToPeerChat(
  initiatorID string,
  receiverUserID string):
  chatID string
```

Comment: consider creating a chatID from directly sending on the first
time.

# [](#anchor-35) Create a group chat

```go
  createGroupChat(ownerID string, groupName string, friendIDs List&lt;String&gt;): chatID string
```

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

