chrome.runtime.onMessage.addListener((message) => {
    console.log("Received message", message)

    // Find and attempt to fill form elements.
    document.querySelectorAll('input[name="password"]').forEach((el) => {
        el.value = message.password
    })
    document.querySelectorAll('input[name="username"]').forEach((el) => {
        el.value = message.username
    })
})