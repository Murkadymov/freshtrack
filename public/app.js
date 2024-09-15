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
            document.getElementById('response').innerText = `Поставка добавлена: ${result.message}`;
            form.reset(); // очищаем форму после успешной отправки
        } else {
            document.getElementById('response').innerText = `Ошибка: ${result.error}`;
        }
    } catch (error) {
        console.error('Ошибка:', error);
        document.getElementById('response').innerText = 'Ошибка при отправке данных';
    }
}
