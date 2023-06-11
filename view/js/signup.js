function signUp() {
    var _data = {
        
        email: document.getElementById("email").value,
        username: document.getElementById("username").value,
        password: document.getElementById("password").value,
        pw : document.getElementById("password2").value,
    }
    if (_data.password !== _data.pw) {
        alert("PASSWORD doesn't match!")
        return
    }

    fetch('/signup', {
        method: "POST",
        body: JSON.stringify(_data),
        headers: { "Content-type": "application/json; charset=UTF-8" }
    })
        .then(response => {
            if (response.status == 201) {
                window.open("Login.html", "_self")
            }
        });
}
