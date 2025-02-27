package main

import (
        "bufio"
        "fmt"
        "os"
        "regexp"
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
                return initialen[:len(initialen)-1]
        }
        return initialen
}

func main() {
        reader := bufio.NewReader(os.Stdin)

        for {
                fmt.Print("Voer een naam in: ")
                naam, _ := reader.ReadString('\n')
                naam = strings.TrimSpace(naam)

                if naam == "" {
                        fmt.Println("Fout: Voer een naam in.")
                } else if regexp.MustCompile(`\d`).MatchString(naam) { // Correcte reguliere expressie
                        fmt.Println("Fout: Een naam mag geen cijfers bevatten.")
                } else {
                        fmt.Println(berekenInitialen(naam))
                }

                fmt.Print("Type 'opnieuw' om opnieuw te proberen, of op Enter om af te sluiten: ")
                input, _ := reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if strings.ToLower(input) != "opnieuw" {
                        break
                }
        }
}