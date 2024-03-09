document.addEventListener("DOMContentLoaded", function() {
    var phoneInput = document.getElementById('newPhone');

    phoneInput.addEventListener('input', function(event) {
        var inputValue = event.target.value;
        var filteredValue = '';

        for (var i = 0; i < inputValue.length; i++) {
            var char = inputValue[i];
            if (!isNaN(char) || char === '-') {
                filteredValue += char;
            }
        }

        event.target.value = filteredValue;
    });
});
function logout() {
    fetch('/logout', {
        method: 'POST',
    })
        .then(response => {
            if (response.ok) {
                alert('Logout successful');
                window.location.href = "index.html";

            } else {
                alert('Logout failed');
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });
}
function saveChanges() {
    var newName = document.getElementById('newName').value;
    var newPhone = document.getElementById('newPhone').value;
    var newDob = document.getElementById('newDob').value;
    var id = document.getElementById('id').value;

    var data = {
        id: id,
        name: newName,
        phone: newPhone,
        date_of_birth: newDob
    };

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/updateProfile', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                alert('Profile updated successfully');
            } else {
                alert('Failed to update profile');
            }
        }
    };
    xhr.send(JSON.stringify(data));
}
