function getProjects() {
    fetch('/getProjects') // Замените '/getProjects' на путь к вашему серверному эндпоинту для получения проектов
        .then(response => response.json())
        .then(data => {
            // Отобразить проекты на странице
            displayProjects(data);
        })
        .catch(error => {
            console.error('Error fetching projects:', error);
        });
}

function displayProjects(projects) {
    const projectsList = document.getElementById('projects-list');

    projects.forEach(project => {
        const projectElement = document.createElement('div');
        projectElement.innerHTML = `
            <h2>Name: ${project.name.String}</h2>
            <h2>Created By: ${project.user_id.String}</h2>
            <p><strong>Category:</strong> ${project.category.String}</p>
            <p><strong>Type:</strong> ${project.project_type.String}</p>
            <p><strong>Year:</strong> ${project.year.Int32}</p>
            <p><strong>Age Category:</strong> ${project.age_category.String}</p>
            <p><strong>Duration:</strong> ${project.duration_minutes.Int32} minutes</p>
            <p><strong>Keywords:</strong> ${project.keywords.String}</p>
            <p><strong>Description:</strong> ${project.description.String}</p>
            <p><strong>Director:</strong> ${project.director.String}</p>
            <p><strong>Producer:</strong> ${project.producer.String}</p>
            <hr>
        `;
        projectsList.appendChild(projectElement);
    });
}

window.onload = function () {
    getProjects();
};
