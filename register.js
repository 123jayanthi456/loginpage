

function register() {

    const name = document.getElementById('username').value;
    const phnno = document.getElementById('phoneno').value;
    const email = document.getElementById('email').value;
    const city = document.getElementById('city').value;
    const state = document.getElementById('state').value;
    const errormsg = document.getElementById('result');
  
    const data = {
      name:name,
      phnno:phnno,
      email:email,
      city:city,
      state:state,
    };
  
    fetch("http://localhost:8081/user", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then(response => {
        if (response.ok) {
          errormsg.innerHTML = 'Created successfully';
        } else {
          errormsg.innerHTML = 'Error: ' + response.statusText;
        }
      })
      .catch(error => {
        errormsg.innerHTML = 'Error: ' + error.message;
      });
  
    return false;
  }
  