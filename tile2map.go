package main

import "log"
import "bufio"
import "encoding/csv"
import "os"
import "fmt"
import "strconv"

func main() {

    fmt.Println("****************************");
    fmt.Println("  CSV to SATURN MAP Format  ");
    fmt.Println("****************************\n");

    var FILENAME string
    var TILE_SIZE, MAP_HEIGHT, MAP_WIDTH, X_POS, Y_POS int;

    fmt.Println("Please enter the filename of the exported Tilemap")
    fmt.Scan(&FILENAME);

    INPUT_FILE, ERR  := os.Open(FILENAME);
    if ERR != nil { log.Fatal(ERR); }

    fmt.Println("TILE SIZE? (EX 8,16,32,64)");
    fmt.Scan(&TILE_SIZE);
    fmt.Println("MAP HEIGHT IN TILES?");
    fmt.Scan(&MAP_HEIGHT);
    fmt.Println("MAP WIDTH IN TILES?");
    fmt.Scan(&MAP_WIDTH);
    fmt.Println("\n\n");

    fmt.Println("CONVERTING TILE MAP...\n");

    CSV_MAP          := csv.NewReader(bufio.NewReader(INPUT_FILE));
    TILEMAP_DATA, _  := CSV_MAP.ReadAll();  // READ ALL DATA FROM CSV VERSION OF TILEMAP (EXPORTED FROM TILED MAP EDITOR)

    OUTPUT_FILE, err := os.OpenFile("OUTPUT.MAP", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644);
    if err != nil { log.Fatal(err); }

    // THE MEAT AND POTATOES (not much to look at)
    for i:=0;i<=MAP_HEIGHT-1;i++ {
        for j:=0;j<=MAP_WIDTH-1;j++ {
            tile_value := ((TILEMAP_DATA[i])[j])
            if (tile_value != string("0") ) {
            ENTRY := tile_value + ".TGA" + " " + strconv.Itoa(X_POS) + " " + strconv.Itoa(Y_POS);
            X_POS += TILE_SIZE; // INCREMENT X AXIS BASED ON TILE SIZE
            OUTPUT_FILE.WriteString(ENTRY); // WRITE OUT THE COMPLETED ENTRY TO OUTPUT FILE
            if (!(i == MAP_HEIGHT-1 && j == MAP_WIDTH-1)) { OUTPUT_FILE.WriteString("\r\n"); } // write end line on all but last entry in map file.
            } else { X_POS += TILE_SIZE; }

        }

    X_POS = 0; // RESET X_POS
    Y_POS += TILE_SIZE; // INCREMENT Y AXIS BY TILE SIZE

    }

    if err := OUTPUT_FILE.Close(); err != nil { print(err); }
    fmt.Println("TILEMAP CONVERTED!");
}
