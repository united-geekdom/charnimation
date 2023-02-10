# `Skimmer`

WIP reference implementation of the `.chnm` file format

## Building and Testing

`./build.sh file width height`

options for `file`  
`loadscreen.chnm` width:`3` height:`3`  
`travelingo.chnm` width:`6` height:`4`

## Adding new animations

Drop all `.chnm` files into `data/`. These files will be copies over to the build directory when you run `./build.sh` next time/

## Creating new animations

Each frame is a single line of constant length, no escape sequences and indents. Note down the dimensions of your frame.

## Troubleshooting

If passed a nonexistent file or invalid parameters, the program will not display anything, and after quitting your cursor will remain hidden. To fix this, start up a new terminal session, or run `skimmer` with valid parameters

### Copyright Attribution

`loadscreen.chnm` by fisik_yum\
`travelingo.chnm` by united-geekdom (adapted for this implementation by fisik_yum)
