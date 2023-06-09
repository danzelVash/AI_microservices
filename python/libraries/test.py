def sort_matrix_by_characteristic(matrix: [[int]]):
    characteristics = []  # список характеристик строк
    for row in matrix:
        characteristic = sum(filter(lambda x: x > 0 and x % 2 == 0, row))  # вычисляем характеристику
        characteristics.append(characteristic)

    # создаем список пар (характеристика, индекс строки)
    pairs = [(characteristics[i], i) for i in range(len(characteristics))]
    # сортируем список пар по возрастанию характеристик
    sorted_pairs = sorted(pairs)

    # создаем новую матрицу, переставляя строки в соответствии с порядком сортировки
    new_matrix = [matrix[sorted_pairs[i][1]] for i in range(len(sorted_pairs))]

    return new_matrix


m = [[1, 2, 0], [-4, 5, 6], [7, -8, 9], [10, -11, 12]]
sorted_matrix = sort_matrix_by_characteristic(m)
print(sorted_matrix)