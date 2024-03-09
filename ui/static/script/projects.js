function getProjects() {
    fetch('/getProjects') 
        .then(response => response.json())
        .then(data => {
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
        projectElement.classList.add('project'); // Добавляем класс 'project'
    
        projectElement.innerHTML = `
            <h2 class="project-name">Name of the project: ${project.name.String}</h2>
            <h2 class="user-id">Created By (User Id): ${project.user_id}</h2>
            <p class="category"><strong>Category:</strong> ${project.category.String}</p>
            <p class="project-type"><strong>Type:</strong> ${project.project_type.String}</p>
            <p class="year"><strong>Year:</strong> ${project.year.Int32}</p>
            <p class="age-category"><strong>Age Category:</strong> ${project.age_category.String}</p>
            <p class="duration"><strong>Duration:</strong> ${project.duration_minutes.Int32} minutes</p>
            <p class="keywords"><strong>Keywords:</strong> ${project.keywords.String}</p>
            <p class="description"><strong>Description:</strong> ${project.description.String}</p>
            <p class="director"><strong>Director:</strong> ${project.director.String}</p>
            <p class="producer"><strong>Producer:</strong> ${project.producer.String}</p>
            <hr class="separator">
        `;
        projectsList.appendChild(projectElement);
    });
    
}

window.onload = function () {
    getProjects();
};
