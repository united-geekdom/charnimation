from reader import chmnread
import frameparser


import sys
if len(sys.argv) == 1:
    fpath = input("Filepath: ")
    chmnread(fpath)
    frameparser.returnFrame(0)
    print(frameparser.findex)
elif len(sys.argv) > 2:
    print("Cannot understand argument")
else:
    chmnread(sys.argv[1])
    frameparser.returnFrame(0) 
    print(frameparser.findex)
 
