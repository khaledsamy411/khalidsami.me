package main

import (
        "bufio"
        "encoding/json"
        "fmt"
        "log"
        "os"
        "strings"
)

func main() {
        responseHeaders := make(map[string]string)
        // Read a bunch of headers from Stdin.
        r := bufio.NewReader(os.Stdin)
        for {
                l, p, err := r.ReadLine()
                // If you hit an error, just log it and go home.
                if err != nil {
                        if len(responseHeaders) == 0 {
                                log.Fatalf("Error reading headers: %s", err.Error())
                        } else {
                                break
                        }
                }
                if p {
                        log.Fatalf("Too long. cURL is trolling")
                }
                // Split on a colon. It's required.
                kv := strings.SplitN(string(l), ":", 2)
                if len(kv) != 2 {
                        responseHeaders[kv[0]] = ""
                        continue
                }
                // Remove all colons. Trolly headers won't matter...
                headerName := strings.ToLower(strings.TrimSpace(kv[0]))
                headerValue := strings.ToLower(strings.TrimSpace(kv[1]))
                responseHeaders[headerName] = headerValue
        }
        // We want something understandable. Let's find only vulnerable.
        hasBadAcao := false
        hasAcac := false
        hostName := ""
        for headerName, headerValue := range responseHeaders {
                if headerName == "access-control-allow-origin" {
                        if strings.Contains(headerValue, "evil.com") {
                                hostName = strings.Replace(headerValue, "evil.com", "", 1)
                                hasBadAcao = true
                        }
                }
                if headerName == "access-control-allow-credentials" {
                        if strings.Contains(headerValue, "true") {
                                hasAcac = true
                        }
                }
        }
        if hasBadAcao && hasAcac {
                responseHeaders["x-hostname"] = hostName
                pwned, err := json.Marshal(responseHeaders)
                if err != nil {
                        log.Fatalf("Couldn't marshal: %s", err.Error())
                }
                fmt.Println(string(pwned))
        }
}
