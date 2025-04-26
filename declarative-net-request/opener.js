(() => {
    const pageObjects = document.getElementsByTagName("object");
    for (const obj of pageObjects) {
        if (obj.type === "application/pdf" && obj.data.startsWith('http')) {
            window.location.replace(obj.data);
        }
    }
})();