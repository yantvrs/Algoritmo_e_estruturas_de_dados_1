# Complexidade de tempo dos algoritmos de ordenação

## Selection Sort

Tem complexidade de tempo O(n^2) independente do caso. 

### Pior caso:
- O(n^2)
- Elementos organizados de forma decrescente no array.
- Comparações: O(n^2)
- Trocas : O(n)

### Melhor caso:
- O(n^2)
- Elementos organizados de forma crescente no array.



## BubbleSort:

Não é muito adequado para grande conjuntos de dados.

### Pior caso:
- Ordem decrescente
- O(n^2)

### Melhor caso:
- Ordem crescente
- Durante as interações, o algoritmo apenas compara  os elementos adjacentes.
- O(n)

Em suma, o Bubble Sort é considerado um algoritmo relativamente lento em comparação com outras opções de ordenação mais eficientes, como o Quick Sort ou o Merge Sort.


## InsertionSort

### Pior caso:
- Vetor em ordem decrescente
- Comparações: O(n^2)
- Trocas : O(n^2)

### Melhor caso:
- Vetor em ordem crescente
- Comparações: O(n)
- Trocas : zero

Em resumo, o Insertion Sort apresenta um desempenho muito eficiente no melhor caso, com uma complexidade de tempo de O(n), tornando-o adequado para conjuntos de dados pequenos ou quase ordenados. No entanto, no pior caso, sua complexidade de tempo é quadrática, O(n^2), o que o torna menos eficiente em comparação com outros algoritmos de ordenação mais eficientes, como o Merge Sort ou o Quick Sort.

## MergeSort

### Pior caso:
- O(nlg(n))
- Comparações : O(nlg(n))
- Trocas: O(nlg(n))

### Melhor caso:
- Mesma complexidade do pior caso

Em resumo, o Merge Sort é eficiente tanto no melhor caso quanto no pior caso, com uma complexidade de tempo de O(n log n). No entanto, no melhor caso, seu desempenho é otimizado devido à possibilidade de evitar comparações e trocas desnecessárias, enquanto no pior caso, o algoritmo ainda é capaz de lidar efetivamente com grandes conjuntos de dados devido à sua abordagem de divisão e fusão. Isso faz do Merge Sort uma escolha popular para a ordenação de conjuntos de dados de diversos tamanhos e características.

## QuickSort

### Pior caso:
- O(n^2)
- Ordem crescente ou decrescente

O pior caso do Quick Sort ocorre quando o pivô escolhido sempre resulta em uma partição desbalanceada. Isso acontece quando o vetor de entrada já está ordenado em ordem crescente ou decrescente. Nesse cenário, o pivô escolhido em cada iteração é o maior ou o menor elemento do subvetor, resultando em uma partição vazia ou com apenas um elemento. 

### Melhor caso:
- O(nlg(n))

O melhor caso do Quick Sort ocorre quando o pivô escolhido divide o vetor de entrada em duas partições aproximadamente iguais. Isso significa que, a cada iteração, o pivô escolhido divide o vetor em duas partes de tamanho semelhante, permitindo que o algoritmo prossiga de forma balanceada. É importante destacar que, na prática, o Quick Sort geralmente apresenta um bom desempenho médio, superando outros algoritmos de ordenação em muitos casos.

## CountingSort

### Pior caso:
- Grande diferença entre os valores mínimos e máximos do vetor de entrada, que resulta em uma grande quantidade de espaços vazios na estrutura de contagem utilizada pelo algorítmo.
- O(n+k)

k(Intervalo entre valor mínimo e máximo dos elementos do vetor)


### Melhor caso:

O melhor caso do algoritmo de ordenação Counting Sort ocorre quando o intervalo entre os valores mínimo e máximo no vetor de entrada é pequeno e os valores estão uniformemente distribuídos. Nesse cenário, o algoritmo pode criar a estrutura de contagem com rapidez e realizar a contagem dos elementos em tempo linear. 
- Quando o intervalo "k" é pequeno e proporcional ao tamanho do vetor.
- O(n+k)