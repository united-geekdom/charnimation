import zipfile
import json
def chmnread(chpath):
    archive = zipfile.ZipFile(chpath, 'r')
    metafile = archive.open('metadata.json', 'r')
    metaread = metafile.read()
    data = json.loads(metaread)
    print(data["title"])
    print(data["author"][0])
    
