# quests

A TUI application for managing your adventure.

At best this is currently a developer preview. If you happen to get it working for yourself, good job and have fun with it! However this is not currently "production ready". Stay tuned though!

# installation

`go install github.com/kanennn/quests`

_better installation coming soon i promise_

# spec

## quests as units

each quest is a self contained unit represented by a folder/directory. child quests of a quest are represented by folders within this folder, and thus parent quests are the enclosing folder of the quest folder.

within each folder are three key files 
- **quest.yml** – this represents metadata of the quest such as display name, description, status, tags, type, etc
- **lore.md** – this is a markdown file representing any information about the quest the user would like to represent in markdown. this file is not directly controlled by quests.
- **legend.log** – this is the log file for the quest. this is also open ended as lore.md is, but contains timestamped information. these entries can be created manually or using the quests executable.

other files and folders may be present, and these may be browsed manually or per file by the quests executable, however no other files besides the three mentioned will be considered part of the quests structure. 



---

made with <3 by kanen stephens
