Package collectlinks is useful for only one task:

Given a response from `http.Get` it will use parse the page and
return to you a slice of all the href links found.

Usage:

    package main
    import (
      "github.com/jackdanger/collectlinks"
      "net/http"
      "fmt"
    )

    func main() {
      resp, _ := http.Get("http://motherfuckingwebsite.com")
      links := collectlinks.All(resp.Body)
      fmt.Println(links)
    }


Running that will output:

   [http://twitter.com/thebarrytone http://txti.es]



# Mis Libros:

[![libros futuro es devops ](https://github.com/culturadevops/recursos/blob/master/portada-futuro-es-devops.png)](https://amzn.to/3S8AGG9) [![libros herramientas devops](https://github.com/culturadevops/recursos/blob/master/portada-herramientasdevops.png)](https://amzn.to/3ga1c4E)

# Mi canal de cultura Devops

[![canal de youtube sobre devops ](https://github.com/culturadevops/recursos/blob/master/logo-culturadevops.png)](https://www.youtube.com/channel/UCfJ67eVA7DkKbbIF5ceJDMA?sub_confirmation=1) 
