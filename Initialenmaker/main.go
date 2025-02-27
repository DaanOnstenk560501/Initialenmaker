package main

import (
        "bufio"
        "fmt"
        "os"
        "strings"
)

func berekenInitialen(naam string) string {
        tussenvoegsels := []string{"van", "den", "der", "de", "ten", "ter"}
        woorden := strings.Split(strings.ToLower(naam), " ")
        initialen := ""

        for _, woord := range woorden {
                isTussenvoegsel := false
                for _, tv := range tussenvoegsels {
                        if woord == tv {
                                isTussenvoegsel = true
                                break
                        }
                }
                if !isTussenvoegsel {
                        initialen += strings.ToUpper(string(woord[0])) + "."
                }
        }

        if len(initialen) > 0 {
                return initialen[:len(initialen)-1] // Verwijder de laatste punt
        }
        return initialen
}

func main() {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Voer een naam in: ")
        naam, _ := reader.ReadString('\n')
        naam = strings.TrimSpace(naam) // Verwijder de newline

        fmt.Println(berekenInitialen(naam))
}