function login() {
    var _data = {
        email: document.getElementById("email").value,
        password: document.getElementById("password").value
    }

    fetch('/login', {
        method: "POST",
        body: JSON.stringify(_data),
        headers: { "Content-type": "application/json; charset=UTF-8" }
    })
        .then(response => {
            if (response.ok) {
                window.open("Home.html", "_self")
            } else {
                throw new Error(response.statusText)
            }
        }).catch(e => {
            if (e == "Error: Unauthorized") {
                alert(e + ". Credentials does not match!")
                return
            }
        });
}