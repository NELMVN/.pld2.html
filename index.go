package main 

import (
    "fmt"
    "os"
    "os/user"
    "io/ioutil"
    p "path"
    "strings"
)

const MAP_NAME                  = "FWRPG Reforged"
const WARCRAFT3_CUSTOM_MAP_DATA = "Documents\\Warcraft III\\CustomMapData\\" + MAP_NAME
const PRELOAD_FILE_TYPE         = ".pld"
const WEBPAGE_FILE_TYPE         = ".html"

func main() {
    u, _ := user.Current()
    path := u.HomeDir + "\\" + WARCRAFT3_CUSTOM_MAP_DATA

    /*
    *    Check if user has Warcraft 3 CustomMapData
    */
    if _, status := ioutil.ReadDir(path); status != nil {
        fmt.Println("[.pld2.html] Warcraft 3 CustomMapData Folder not found.")
    } else {
        fmt.Println("[.pld2.html] Warcraft 3 CustomMapData Folder found.")

        players, _ := ioutil.ReadDir(path)

        for _, playerFolders := range players {
            if playerFolders.IsDir() {
                folder := path + "\\" + playerFolders.Name()
                playerFiles, _ := ioutil.ReadDir(folder)

                for _, pld := range playerFiles {
                    if !pld.IsDir() {
                        file := pld.Name()
                        extension := p.Ext(file)

                        if extension == PRELOAD_FILE_TYPE {
                            fullpath := folder + "\\"
                            os.Rename(fullpath + file, fullpath + strings.TrimSuffix(file, extension) + ".html")
                        }
                    }
                }
            }
        }

    }
}
