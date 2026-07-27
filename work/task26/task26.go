package main

func BuildHTTPResponse(statusLine, headers, body string) string {
    request := statusLine + "\r\n"
	if headers != "" {
        request += headers
    }
    request += "\r\n"
    
    if body != "" {
        request += body
    }
	
    return request
}