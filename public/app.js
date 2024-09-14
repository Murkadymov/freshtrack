const API_URL = 'http://localhost:8080/deliveries';

async function fetchDeliveries() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const deliveries = await response.json();
        displayDeliveries(deliveries);
    } catch (error) {
        console.error('Error fetching deliveries:', error);
    }
}

function displayDeliveries(deliveries) {
    const deliveriesList = document.getElementById('deliveries-list');
    deliveriesList.innerHTML = ''; // Очистка таблицы перед вставкой новых данных

    deliveries.forEach(delivery => {
        delivery.cargoes.forEach(cargo => {
            const row = document.createElement('tr');
            row.innerHTML = `
        <td>${delivery.car.number}</td>
        <td>${cargo.name}</td>
        <td>${delivery.car.driver}</td>
        <td>${delivery.from}</td>
        <td>Unknown</td> <!-- В примере данных нет производителя, можно добавить поле если нужно -->
      `;
            deliveriesList.appendChild(row);
        });
    });
}

window.onload = fetchDeliveries;
