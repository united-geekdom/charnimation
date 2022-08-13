import zipfile
import json
import os

def chmnread(chpath):
    os.remove("filedata/frames.txt")
    os.remove("filedata/metadata.json")
    with zipfile.ZipFile(chpath, 'r') as z:
        z.extractall("filedata")
