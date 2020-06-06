export function postLogin(email, password) {
  fetch("http://localhost:3000/api/v1/users", {
    method: "POST",
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: email,
      password: password,
    }),
  }).then((response) => {
    response
      .json()
      .then((jsonData) => {
        console.log(jsonData);
      })
      .catch((err) => {
        console.log(err);
      });
  });
}

export function postRegister(username, email, password) {
  fetch("http://localhost:3000/api/v1/users", {
    method: "POST",
    mode: "cors",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username: username,
      email: email,
      password: password,
    }),
  }).then((response) => {
    response
      .json()
      .then((jsonData) => {
        console.log(jsonData);
      })
      .catch((err) => {
        console.log(err);
      });
  });
}
