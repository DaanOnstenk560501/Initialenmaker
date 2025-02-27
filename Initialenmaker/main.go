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

                // Verwijder meerdere spaties
                reSpaties := regexp.MustCompile(`\s+`)
                naam = reSpaties.ReplaceAllString(naam, " ")

                if naam == "" {
                        fmt.Println("Fout: Voer een naam in.")
                } else if regexp.MustCompile(`\d`).MatchString(naam) {
                        fmt.Println("Fout: Een naam mag geen cijfers bevatten.")
                } else {
                        // Controleer op ongeldige tekens (alles behalve letters en spaties)
                        reOngeldig := regexp.MustCompile(`[^a-zA-Z\s]`)
                        if reOngeldig.MatchString(naam) {
                                fmt.Println("Fout: De naam bevat ongeldige tekens. Indien uw naam een apostrof of ander teken bevat hoeft u deze niet in te vullen")
                        } else {
                                fmt.Println(berekenInitialen(naam))
                        }
                }

                fmt.Print("Type 'opnieuw' om opnieuw te proberen, of op Enter om af te sluiten: ")
                input, _ := reader.ReadString('\n')
                input = strings.TrimSpace(input)

                if strings.ToLower(input) != "opnieuw" {
                        break
                }
        }
}