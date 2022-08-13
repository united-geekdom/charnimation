import os
rootdir = '/mnt/c/Users/Kiran Nithilan/music/the_who/live//'

subdirs = [sb for sb in os.listdir(rootdir) if os.path.isdir(os.path.join(rootdir, sb))]
print(subdirs)