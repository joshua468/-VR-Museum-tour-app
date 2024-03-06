document.addEventListener("DOMContentLoaded", function () {
    // Fetch exhibit data from the server
    fetch('http://localhost:8080/exhibits')
        .then(response => response.json())
        .then(data => {
            const exhibitsContainer = document.getElementById('exhibitsContainer');
            data.exhibits.forEach(exhibit => {
                const exhibitEntity = document.createElement('a-entity');
                exhibitEntity.setAttribute('position', exhibit.position);
                exhibitEntity.setAttribute('rotation', exhibit.rotation);
                exhibitEntity.setAttribute('scale', exhibit.scale);
                exhibitEntity.setAttribute('gltf-model', exhibit.model);
                exhibitsContainer.appendChild(exhibitEntity);
            });
        })
        .catch(error => console.error('Error fetching exhibit data:', error));
});
