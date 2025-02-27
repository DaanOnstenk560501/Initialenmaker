package main

import (
	"bufio"
	"fmt"
	"log"
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

func error(message string, logMessage string) {
	fmt.Println("Fout:", message)
	log.Println(logMessage)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	for {
		fmt.Print("Voer een naam in: ")
		naam, _ := reader.ReadString('\n')
		naam = strings.TrimSpace(naam)

		reSpaties := regexp.MustCompile(`\s+`)
		naam = reSpaties.ReplaceAllString(naam, " ")

		if naam == "" {
			error("Voer een naam in.", "Lege naam ingevoerd.")
		} else if regexp.MustCompile(`\d`).MatchString(naam) {
			error("Een naam mag geen cijfers bevatten.", fmt.Sprintf("Naam bevat cijfers: %s", naam))
		} else {

			reOngeldig := regexp.MustCompile(`[^a-zA-Z\s]`)
			if reOngeldig.MatchString(naam) {
				error("De naam bevat ongeldige tekens. Als de naam een apostrof bevat hoeft u deze niet in te vullen", fmt.Sprintf("Naam bevat ongeldige tekens: %s", naam))
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
