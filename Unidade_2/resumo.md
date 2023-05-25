# Selection Sort

O Selection Sort é um algoritmo de ordenação simples que consiste em encontrar repetidamente o menor elemento de uma lista e colocá-lo no início. Esse processo é repetido para cada elemento restante, até que a lista esteja completamente ordenada. É um método de classificação por seleção.

## Funcionamento do Selection Sort

1. Inicialmente, considera-se toda a lista como não ordenada.
2. A cada iteração, busca-se o elemento mínimo na lista não ordenada.
3. Troca-se o elemento mínimo encontrado com o primeiro elemento não ordenado.
4. Repete-se os passos 2 e 3 para o restante da lista não ordenada, considerando-a uma lista menor a cada iteração.

O algoritmo seleciona o menor elemento em cada iteração e o coloca na posição correta. Após cada iteração, o número de elementos não ordenados diminui, até que a lista esteja completamente ordenada.


## Complexidade de Tempo

A complexidade de tempo do Selection Sort é O(n^2), onde n é o número de elementos na lista a ser ordenada. Isso ocorre porque são necessárias duas iterações aninhadas para cada elemento da lista.

Apesar de sua simplicidade, o Selection Sort não é eficiente para listas grandes, uma vez que o número de comparações e trocas é considerável. Algoritmos como o Quick Sort ou o Merge Sort são geralmente preferidos para grandes conjuntos de dados.

## Vantagens e Desvantagens

Vantagens do Selection Sort:
- Implementação simples e fácil de entender.
- Requer um número mínimo de trocas.

Desvantagens do Selection Sort:
- Desempenho relativamente lento para listas grandes.
- Não é estável, ou seja, a ordem relativa de elementos iguais não é preservada.
- Sempre executa o mesmo número de comparações, independentemente da lista estar parcialmente ordenada.

O Selection Sort é uma forma básica de ordenação e é útil para ensinar os conceitos fundamentais de algoritmos de classificação. No entanto, para aplicações práticas, existem algoritmos mais eficientes disponíveis.


