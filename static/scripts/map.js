document.addEventListener('DOMContentLoaded', function() {
    initializeMap();
});

function initializeMap() {
    const mapDiv = document.getElementById('map');

    if (!mapDiv) {
        return;
    }

    mapDiv.style.height = '500px';
    mapDiv.style.width = '100%';

    const joaoPessoaCoords = [-7.1195, -34.8451]; // Latitude, Longitude for João Pessoa
    const map = L.map(mapDiv).setView(joaoPessoaCoords, 12);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(map);

    map.invalidateSize();

    var greenIcon = L.icon({
        iconUrl: "static/css/images/marker-icon.png",
        iconSize:     [38, 38], // size of the icon
        popupAnchor:  [0, -20] // point from which the popup should open relative to the iconAnchor
    });

    const locationsData = JSON.parse(mapDiv.dataset.locations);
    console.log(locationsData);

    // Loop through the locations and add markers to the map
    locationsData.forEach(function(loc) {
        if (loc.latitude && loc.longitude) {
            const marker = L.marker([loc.latitude, loc.longitude], {icon: greenIcon}).addTo(map);
            marker.bindPopup(`<b>${loc.nome}</b><br>${loc.endereco}`);
        }
    });
}
