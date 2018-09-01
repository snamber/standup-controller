# standup-controlling

A not-so-serious, but (for certain companies) convenient, stand-up based way to do approximate, monthly project controlling

## How

In case you write your standups (if you don't do that yet, don't start and discard this tool), write them in this format:

```
What did I do?
project: message
project: message

What will I do?
message
message
```
i.e. use the `project: message` pattern in the section of things that you actually did, and don't use it anywhere else.

At the end of the month select all your standups of the respective month, and copy them into a text file (e.g. `standups-august.txt`).
In case you have this information somewhere, determine for how many hours you need to account for, e.g. 200. Then call `standup-controlling --standups path/to/standups-august.txt --hours 200` to get your approximate controlling output:

```
project: hours1
project: hours2
```

## Installation

go get -u github.com/snamber/standup-controlling

## Disclaimer

Either you know why this tool exists, and it might make a lot of sense for you, or it doesn't make any sense for you. In that case, thanks for reading this far; like and subscribe.

In either case: Don't get carried away by numbers, focus on the mission.
