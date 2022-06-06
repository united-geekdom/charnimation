import zipfile
import json
archive = zipfile.ZipFile('ch1/ch1.chmn', 'r')
metafile = archive.open('metadata.json', 'r')
metaread = metafile.read()
data = json.loads(metaread)
print(data["title"])
print(data["author"][0])
