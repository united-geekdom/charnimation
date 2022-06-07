import zipfile
import json
def chmnread(chpath):
    archive = zipfile.ZipFile(chpath, 'r')
    metafile = archive.open('metadata.json', 'r')
    metaread = metafile.read()
    data = json.loads(metaread)
    framefile = archive.open('frames.json', 'r')
    frameread = framefile.read()
    frames = json.loads(frameread)
    print("Title: " + data["title"])
    print("Author: " + data["author"][0])
    for n in frames:
        print(n)
