async function addSupply() {
    const form = document.getElementById('supplyForm');
    const formData = new FormData(form);

    // Создаем объект данных в соответствии с ожидаемым форматом
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
        const response = await fetch('http://localhost:8080/supply', { // Замените URL на URL вашего сервера
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(supply),
        });

        const result = await response.json();

        if (response.ok) {
            document.getElementById('response').innerText = `Поставка добавлена: ${result.message}`;
            // Обновите таблицу поставок или выполните другие действия по необходимости
        } else {
            document.getElementById('response').innerText = `Ошибка: ${result.error}`;
        }
    } catch (error) {
        console.error('Ошибка:', error);
        document.getElementById('response').innerText = 'Ошибка при отправке данных';
    }
}
