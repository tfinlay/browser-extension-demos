chrome.action.onClicked.addListener((tab) => {
    console.log("Browser action clicked!", tab)

    const tabUrl = new URL(tab.url)

    chrome.runtime.sendNativeMessage("native_messaging_example_app", {
        host: tabUrl.host
    }).then(
        res => {
            if (res.error) {
                if (res.error === "unknown host") {
                    console.log("No login is saved for this host")
                } else {
                    console.error(res.error)
                }
                return
            }
            console.log("Found login for this host:", res)
            chrome.tabs.sendMessage(tab.id, res)
        },
        err => {
            console.log("Got error", err)
        }
    )
});