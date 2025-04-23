function getUser(userId) {
    console.log(userId);
    fetch('/getUser/' + userId)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            return response.json();
        })
        .then(data => {
            document.getElementById('nome').value = data.Nome;
            document.getElementById('idade').value = data.Idade;
        })
        .catch(error => {
            console.error('There has been a problem with your fetch operation:', error);
        });
}

function getUsers() {
    const response = fetch('/getUsers')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            return response.json();
        })
        .then(data => {
            populateTable(data);
        })  
        .catch(error => {
            console.error('There has been a problem with your fetch operation:', error);
        });
}

function addUser() {
    event.preventDefault();
    const nome = document.getElementById('nome').value;
    const idade = parseInt(document.getElementById('idade').value, 10);

    const dataToSend = JSON.stringify({
        nome: nome,
        idade: idade
    });

    console.log(dataToSend);

    fetch('/addUser', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: dataToSend
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }
        return response.json();
    })
    .then(data => {
        console.log(data);
        alert('Usuário adicionado com sucesso!');
        window.location.href = '/';
    })
    .catch(error => {
        console.error('There has been a problem with your fetch operation:', error);
    });
}

function deleteUser(userId, userName) {
    if (confirm(`Você tem certeza que deseja excluir o usuário ${userName}?`)) {
        fetch(`/deleteUser/${userId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok ' + response.statusText);
            }
            alert(`Usuário ${userName} excluído com sucesso!`);
            getUsers();
        })
        .catch(error => {
            console.error('There has been a problem with your fetch operation:', error);
        });
    }
}

function populateTable(data) {
    const userRows = document.getElementById('userRows');
    userRows.innerHTML = ''; // Limpa as linhas existentes
    data.forEach(user => {
        const row = document.createElement('tr');
        row.innerHTML = `
        <td>${user.ID}</td>
        <td>${user.Nome}</td>
        <td>${user.Idade}</td>
        <td><i class="fa fa-edit" style="font-size:36px;color:white" onclick="window.location.href='/atualizar/${user.ID}'"></i></td>
        <td><i class="fa fa-trash-o" style="font-size:36px;color:red" onclick="deleteUser(${user.ID}, '${user.Nome}')"></i></td>
        `;
        userRows.appendChild(row);
    })
}

function updateUser(userId) {
    event.preventDefault();
    const nome = document.getElementById('nome').value;
    const idade = parseInt(document.getElementById('idade').value, 10);

    const dataToSend = JSON.stringify({
        nome: nome,
        idade: idade
    });

    console.log(dataToSend);

    fetch(`/updateUser/${userId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: dataToSend
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }
        return response.json();
    })
    .then(data => {
        console.log(data);
        alert('Usuário atualizado com sucesso!');
        window.location.href = '/';
    })
    .catch(error => {
        console.error('There has been a problem with your fetch operation:', error);
    });
}