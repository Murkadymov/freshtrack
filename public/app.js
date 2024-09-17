async function addSupply() {
    const form = document.getElementById('supplyForm');

    // Проверяем, есть ли пустые поля
    if (!form.checkValidity()) {
        document.getElementById('response').innerText = 'Все поля должны быть заполнены!';
        return;
    }

    const formData = new FormData(form);

    const supply = {
        driver: {
            driverNumber: formData.get('driverNumber'),
            tractorNumber: formData.get('truckNumber'),
            trailNumber: formData.get('trailerNumber')
        },
        goods: {
            cargo: formData.get('product')
        },
        manufacturer: {
            name: formData.get('manufacturer'),
            origin: formData.get('country')
        }
    };

    try {
        const response = await fetch('http://localhost:8080/supply', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(supply),
        });

        const result = await response.json();

        if (response.ok) {
            document.getElementById('response').innerText = `Поставка добавлена ${result.message}`;
            form.reset(); // очищаем форму после успешной отправки

            loadSupplies(); // Обновляем таблицу после добавления новой поставки
        } else {
            document.getElementById('response').innerText = `Ошибка: ${result.error}`;
        }
    } catch (error) {
        console.error('Ошибка:', error);
        document.getElementById('response').innerText = 'Ошибка при отправке данных';
    }
}

async function loadSupplies() {
    try {
        const response = await fetch('http://localhost:8080/supply', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        const result = await response.json();

        if (response.ok) {
            const supplies = result.details; // Доступ к массиву details

            const tableBody = document.getElementById('deliveries-list');
            tableBody.innerHTML = ''; // Очистка таблицы перед вставкой новых данных

            supplies.forEach(supply => {
                const row = document.createElement('tr');

                // Создаем ячейки для каждого поля
                const tractorNumberCell = document.createElement('td');
                tractorNumberCell.innerText = supply.driver.tractorNumber;

                const trailNumberCell = document.createElement('td');
                trailNumberCell.innerText = supply.driver.trailNumber;

                const driverNumberCell = document.createElement('td');
                driverNumberCell.innerText = supply.driver.driverNumber;

                const cargoCell = document.createElement('td');
                cargoCell.innerText = supply.goods.cargo;

                const countryCell = document.createElement('td');
                countryCell.innerText = supply.manufacturer.origin;

                const manufacturerCell = document.createElement('td');
                manufacturerCell.innerText = supply.manufacturer.name;

                // Добавляем все ячейки в строку
                row.appendChild(tractorNumberCell);
                row.appendChild(trailNumberCell);
                row.appendChild(driverNumberCell);
                row.appendChild(cargoCell);
                row.appendChild(countryCell);
                row.appendChild(manufacturerCell);

                // Добавляем строку в таблицу
                tableBody.appendChild(row);
            });
        } else {
            document.getElementById('response').innerText = `Ошибка: ${result.error}`;
        }
    } catch (error) {
        console.error('Ошибка:', error);
        document.getElementById('response').innerText = 'Ошибка при получении данных';
    }
}

document.addEventListener('DOMContentLoaded', function() {
    loadSupplies(); // Загружаем список поставок при загрузке страницы
});
