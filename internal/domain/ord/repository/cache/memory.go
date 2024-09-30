package cahe

import (
	"TestBroker/internal/domain/ord/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FileOrderRepository struct {
	filePath string
}

func NewFileOrderRepository(filePath string) *FileOrderRepository {
	return &FileOrderRepository{filePath: filePath}
}

func (r *FileOrderRepository) GetOrders() ([]entity.Order, error) {
	var orders []entity.Order

	// Открываем файл
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	// Чтение содержимого файла
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	// Десериализация JSON данных в массив объектов
	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		return nil, fmt.Errorf("ошибка при декодировании JSON: %w", err)
	}

	return orders, nil
}
