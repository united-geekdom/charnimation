import zipfile
import io
import json
def chmnread(chpath):
    with zipfile.ZipFile(chpath, 'r') as z:
        with z.open('metadata.json', 'r') as m:
            metaread = m.read()
            data = json.loads(metaread)

        with io.TextIOWrapper(z.open("frames.txt"), encoding="utf-8") as fs:
            toPrint = fs.read()
            print(toPrint)
