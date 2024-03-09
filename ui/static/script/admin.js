document.addEventListener("DOMContentLoaded", function() {
    var phoneInput = document.querySelector('input[data-attribute="phone"]');

    phoneInput.addEventListener('input', function(event) {
        var inputValue = event.target.value;
        var filteredValue = '';

        for (var i = 0; i < inputValue.length; i++) {
            var char = inputValue[i];
            if (!isNaN(char) || char === '-' || char === '(' || char === ')') {
                filteredValue += char;
            }
        }

        if (filteredValue.length > 12) {
            filteredValue = filteredValue.slice(0, 12);
        }

        event.target.value = filteredValue;
    });
});
document.addEventListener('DOMContentLoaded', function () {
    getUsers(); 

    const closeButton = document.getElementById('closeButton');
    closeButton.addEventListener('click', function () {
        const windowDiv = document.getElementById('window');
        windowDiv.style.display = 'none';
    });
    const deleteButton = document.getElementById('deleteButton');
    deleteButton.addEventListener('click', function () {
        const userId = document.getElementById('userId').textContent;

        if (confirm("Are you sure you want to delete this user?")) {
            fetch(`/deleteUser?userId=${userId}`, { // Замените на эту строку
                method: 'DELETE'
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    alert("User Deleted")
                    return response.json();
                })
        }

    });


    const sendButton = document.getElementById('sendButton');
    sendButton.addEventListener('click', function () {
        const userId = document.getElementById('userId').textContent;
        const inputs = document.querySelectorAll('#window input[data-attribute]');
        const updatedUserData = {
            user_id: userId
        };

        inputs.forEach(input => {
            const attribute = input.getAttribute('data-attribute');
            if (input.type === 'checkbox') {
                updatedUserData[attribute] = input.checked;
            } else {
                updatedUserData[attribute] = input.value;
            }
        });

        fetch('/updateUserAdmin', {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(updatedUserData)
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                alert("User updated")
                return response.json();
            })
    });
});

function getUsers() {
    fetch('/getAllUsers')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            const userTable = document.getElementById('userTable');
            data.forEach(user => {
                const row = userTable.insertRow();
                row.innerHTML = `
                <td id="user${user.id}">${user.id}</td>
                <td>${user.email}</td>
                <td>${user.password}</td>
                <td>${user.name}</td>
                <td>${user.phone}</td>
                <td>${user.date_of_birth}</td>
                <td>${user.session_id}</td>
                <td>${user.is_admin}</td>
                <td>${user.created_at}</td>
                <td>${user.updated_at}</td>
                <td><button class="editButton">Edit/Delete</button></td>
            `;
            });

            const editButtons = document.querySelectorAll('.editButton');
            editButtons.forEach(button => {
                button.addEventListener('click', function () {
                    const userId = this.parentElement.parentElement.querySelector('td:first-child').innerText;
                    const windowDiv = document.getElementById('window');
                    const userIdField = windowDiv.querySelector('#userId');
                    userIdField.textContent = userId;

                    const user = data.find(u => u.id === parseInt(userId));
                    const inputs = windowDiv.querySelectorAll('input[data-attribute], textarea[data-attribute]');
                    inputs.forEach(input => {
                        const attribute = input.getAttribute('data-attribute');
                        if (input.type === 'checkbox') {
                            input.checked = user[attribute];
                        } else {
                            input.value = user[attribute];
                        }
                    });

                    windowDiv.style.display = 'block';
                });
            });
        })
        .catch(error => {
            console.error('There was a problem with your fetch operation:', error);
        });
}