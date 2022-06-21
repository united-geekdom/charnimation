import linecache
import json
import sys
def parseFrame(l, h):
    toPrint = ""
    for n in range((l-1)*(h+1) + 1, l*(h+1) + 1):
        if n != (l-1)*(h+1) + 1:
            toPrint += linecache.getline("filedata/frames.txt", n)
    return toPrint
def returnFrame():
    metadata = json.load(open("filedata/metadata.json"))
    height = metadata["dimensions"][1]
    fl = metadata["frames"]
    for c in range(fl):
        sys.stdout.write("\r" + parseFrame(c+1, height))
        sys.stdout.flush()
