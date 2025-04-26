const onHeadersReceived = (requestDetails) => {
    console.log(`Intercepted request of type ${requestDetails.type} to ${requestDetails.url}`);
    const headers = requestDetails.responseHeaders ?? []
    const contentType = headers.find(
        (header) => header.name.toLowerCase() === 'content-type'
    )

    if (contentType !== undefined && contentType.value.toLowerCase() === 'application/pdf') {
        console.log("Intercepted PDF request...")
        const dispositionHeaderIndex = headers.findIndex(
            (header) => header.name.toLowerCase() === 'content-disposition'
        )
        if (dispositionHeaderIndex !== undefined && /^attachment/.test(headers[dispositionHeaderIndex].value)) {
            console.log(`Caught Content-Disposition: ${headers[dispositionHeaderIndex].value}`)
            headers.splice(dispositionHeaderIndex, 1)
        }
        return {responseHeaders: headers}
    }
}

chrome.webRequest.onHeadersReceived.addListener(
    onHeadersReceived,
    {
        urls:  ["*://*.learn.canterbury.ac.nz/*", "*://127.0.0.1/*"],
        types: ["main_frame"]
    },
    ["blocking", "responseHeaders"]
)

