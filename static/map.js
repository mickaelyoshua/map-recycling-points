document.addEventListener('DOMContentLoaded', function() {
    // Initialize the map
    var map = L.map('map').setView([-7.1194, -34.8643], 12); // Centered on João Pessoa

    // Add a tile layer
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    var markers = L.featureGroup().addTo(map);

    function addMarkers(locations) {
        markers.clearLayers(); // Clear existing markers
        locations.forEach(function(location) {
            if (location.latitude && location.longitude) {
                var marker = L.marker([location.latitude, location.longitude]).addTo(markers);
                marker.bindPopup('<b>' + location.nome + '</b><br>' + location.endereco + '<br>Categoria: ' + location.categoria);
            }
        });
    }

    // Initial load of markers
    if (typeof locationsData !== 'undefined' && locationsData.length > 0) {
        addMarkers(locationsData);
    }

    // HTMX event listener for when the content is swapped (e.g., after filter)
    document.body.addEventListener('htmx:afterSwap', function(event) {
        // Check if the swapped content contains the map and new locationsData
        if (event.detail.target.id === 'map-container' || event.detail.target.closest('#map-container')) {
            // Re-initialize map if it was removed or update markers if data changed
            // For simplicity, we'll assume locationsData is updated globally or passed again
            // A more robust solution might involve passing data directly to this function
            if (typeof locationsData !== 'undefined' && locationsData.length > 0) {
                addMarkers(locationsData);
            }
        }
    });
});