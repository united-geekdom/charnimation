from reader import chmnread
from frameparser import returnFrame


import sys
if len(sys.argv) == 1:
    print("Please specify a filepath")
elif len(sys.argv) > 2:
    print("Cannot understand argument")
else:
    chmnread(sys.argv[1])
    returnFrame()
