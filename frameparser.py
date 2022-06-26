import linecache
import json
import sys
import time
def parseFrame(l, h, au):
    toPrint = ""
    for n in range((l-1)*(h+1) + 1, l*(h+1) + 1):
        if n != (l-1)*(h+1) + 1:
            toPrint += linecache.getline("filedata/frames.txt", n)
    toPrint += f"\nBy: {au}\n"
    return toPrint
def returnFrame():
    metadata = json.load(open("filedata/metadata.json"))
    height = metadata["dimensions"][1]
    fl = metadata["frames"]
    intrv = 1/metadata["fps"]
    if len(metadata["author"]) == 1:
        author = metadata["author"][0]
    else:
        author = ", ".join(map(str, metadata["author"]))

    for c in range(fl):
        print('\033[H\033[J', end='')
        print(parseFrame(c+1, height, author), end="\r")
        time.sleep(intrv)
