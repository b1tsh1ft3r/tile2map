# tile2map
Tile map converter for the Sega saturn. Converts between Csv export from Tiled Map editor to a custom Sega saturn map format that can be easily read at run time by the system.

The tool assumes that the cell value on the tilemap has a complimenting image file name called that value for the saturn.

Example of a line in the tilemap:

1,0,0,1,0

1.TGA or 0.TGA would be files that would need to exist for saturn jo-engine as it will convert the tile map to a formated
listed of images with x/y coordinates for jo-engine to process.

