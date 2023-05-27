# Questão 3

## Ilustre, em detalhes, o funcionamento dos seguintes algoritmos com os seguintes vetores.

- aleatório = [3,6,2,5,4,3,7,1,10^9]
- decrescente = [7,6,5,4,3,3,2,1]
- crescente = [1,2,3,3,4,5,6,7]

### a. SelectionSort (in-place)
    i. vetor para ilustrar: aleatório

O algoritmo de ordenação Selection Sort é um algoritmo simples que funciona selecionando repetidamente o menor elemento de uma lista não ordenada e colocando-o no início dessa lista. O processo é repetido até que a lista inteira esteja ordenada.

Aqui está o funcionamento passo a passo do algoritmo Selection Sort para o vetor [3, 6, 2, 5, 4, 3, 7, 1, 10^9]:

Passo 1:
O primeiro elemento é 3. Assumimos que é o menor elemento.
[3, 6, 2, 5, 4, 3, 7, 1, 10^9]

Passo 2:
Verificamos o segundo elemento, 6. É maior que o atual menor elemento (3). Não fazemos nada.
[3, 6, 2, 5, 4, 3, 7, 1, 10^9]

Passo 3:
Verificamos o terceiro elemento, 2. É menor que o atual menor elemento (3). Atualizamos o menor elemento para 2.
[2, 6, 3, 5, 4, 3, 7, 1, 10^9]

Passo 4:
Verificamos o quarto elemento, 5. É maior que o atual menor elemento (2). Não fazemos nada.
[2, 6, 3, 5, 4, 3, 7, 1, 10^9]

Passo 5:
Verificamos o quinto elemento, 4. É menor que o atual menor elemento (2). Atualizamos o menor elemento para 4.
[2, 6, 3, 5, 4, 3, 7, 1, 10^9]

Passo 6:
Verificamos o sexto elemento, 3. É menor que o atual menor elemento (4). Atualizamos o menor elemento para 3.
[2, 6, 3, 5, 4, 3, 7, 1, 10^9]

Passo 7:
Verificamos o sétimo elemento, 7. É maior que o atual menor elemento (3). Não fazemos nada.
[2, 6, 3, 5, 4, 3, 7, 1, 10^9]

Passo 8:
Verificamos o oitavo elemento, 1. É menor que o atual menor elemento (3). Atualizamos o menor elemento para 1.
[2, 6, 3, 5, 4, 3, 7, 1, 10^9]

Passo 9:
Verificamos o nono elemento, 10^9. É maior que o atual menor elemento (1). Não fazemos nada.
[2, 6, 3, 5, 4, 3, 7, 1, 10^9]

Agora, o menor elemento é 1, então o colocamos no início do vetor:
[1, 6, 3, 5, 4, 3, 7, 2, 10^9]

Repetimos o mesmo processo para o restante do vetor não ordenado:

[1, 6, 3, 5, 4, 3, 7, 2, 10^9] - menor elemento é 2
[

1, 2, 3, 5, 4, 3, 7, 6, 10^9] - menor elemento é 3
[1, 2, 3, 4, 5, 3, 7, 6, 10^9] - menor elemento é 3
[1, 2, 3, 3, 5, 4, 7, 6, 10^9] - menor elemento é 4
[1, 2, 3, 3, 4, 5, 7, 6, 10^9] - menor elemento é 5
[1, 2, 3, 3, 4, 5, 7, 6, 10^9] - menor elemento é 6
[1, 2, 3, 3, 4, 5, 6, 7, 10^9] - menor elemento é 7
[1, 2, 3, 3, 4, 5, 6, 7, 10^9] - menor elemento é 10^9

Agora, o vetor está completamente ordenado em ordem crescente usando o algoritmo Selection Sort.

### b. BubbleSort
    i. vetor para ilustrar: aleatório

O algoritmo de ordenação Bubble Sort é um algoritmo simples que funciona comparando repetidamente pares de elementos adjacentes e trocando-os se estiverem na ordem errada. O processo é repetido até que nenhum elemento precise ser trocado, indicando que a lista está ordenada.

Aqui está o funcionamento passo a passo do algoritmo Bubble Sort para o vetor [3, 6, 2, 5, 4, 3, 7, 1, 10^9]:

Passo 1:
Comparamos o primeiro elemento (3) com o segundo elemento (6). Como 3 é menor que 6, não fazemos nada.
[3, 6, 2, 5, 4, 3, 7, 1, 10^9]

Passo 2:
Comparamos o segundo elemento (6) com o terceiro elemento (2). Como 6 é maior que 2, trocamos os elementos.
[3, 2, 6, 5, 4, 3, 7, 1, 10^9]

Passo 3:
Comparamos o terceiro elemento (6) com o quarto elemento (5). Como 6 é maior que 5, não fazemos nada.
[3, 2, 6, 5, 4, 3, 7, 1, 10^9]

Passo 4:
Comparamos o quarto elemento (6) com o quinto elemento (4). Como 6 é maior que 4, não fazemos nada.
[3, 2, 6, 5, 4, 3, 7, 1, 10^9]

Passo 5:
Comparamos o quinto elemento (6) com o sexto elemento (3). Como 6 é maior que 3, não fazemos nada.
[3, 2, 6, 5, 4, 3, 7, 1, 10^9]

Passo 6:
Comparamos o sexto elemento (6) com o sétimo elemento (7). Como 6 é menor que 7, não fazemos nada.
[3, 2, 6, 5, 4, 3, 7, 1, 10^9]

Passo 7:
Comparamos o sétimo elemento (7) com o oitavo elemento (1). Como 7 é maior que 1, trocamos os elementos.
[3, 2, 6, 5, 4, 3, 1, 7, 10^9]

Passo 8:
Comparamos o oitavo elemento (7) com o nono elemento (10^9). Como 7 é menor que 10^9, não fazemos nada.
[3, 2, 6, 5, 4, 3, 1, 7, 10^9]

Repetimos o mesmo processo para a lista novamente, ignorando o último elemento que já está na posição correta.

Passo 9:
Comparamos o primeiro elemento (3) com o segundo elemento (2). Como 3 é maior que 2, trocamos os elementos.
[2, 3, 6, 5, 4, 3, 1, 7, 10^9]

Passo 10:
Comparamos o segundo elemento (3) com o terceiro elemento (6). Como

 3 é menor que 6, não fazemos nada.
[2, 3, 6, 5, 4, 3, 1, 7, 10^9]

... Continuamos repetindo o processo até que nenhum elemento precise ser trocado.

Passo 11:
[2, 3, 5, 6, 4, 3, 1, 7, 10^9]

Passo 12:
[2, 3, 5, 4, 6, 3, 1, 7, 10^9]

Passo 13:
[2, 3, 5, 4, 3, 6, 1, 7, 10^9]

Passo 14:
[2, 3, 5, 4, 3, 1, 6, 7, 10^9]

Passo 15:
[2, 3, 5, 4, 3, 1, 6, 7, 10^9]

Passo 16:
[2, 3, 4, 5, 3, 1, 6, 7, 10^9]

Passo 17:
[2, 3, 4, 3, 5, 1, 6, 7, 10^9]

Passo 18:
[2, 3, 4, 3, 1, 5, 6, 7, 10^9]

Passo 19:
[2, 3, 4, 3, 1, 5, 6, 7, 10^9]

Passo 20:
[2, 3, 3, 4, 1, 5, 6, 7, 10^9]

Passo 21:
[2, 3, 3, 1, 4, 5, 6, 7, 10^9]

Passo 22:
[2, 3, 3, 1, 4, 5, 6, 7, 10^9]

Passo 23:
[2, 3, 1, 3, 4, 5, 6, 7, 10^9]

Passo 24:
[2, 3, 1, 3, 4, 5, 6, 7, 10^9]

Passo 25:
[2, 1, 3, 3, 4, 5, 6, 7, 10^9]

Passo 26:
[2, 1, 3, 3, 4, 5, 6, 7, 10^9]

Passo 27:
[1, 2, 3, 3, 4, 5, 6, 7, 10^9]

Agora, o vetor está completamente ordenado em ordem crescente usando o algoritmo Bubble Sort.

### c. InsertionSort( in-place )
    i. vetor para ilustrar: aleatório

O algoritmo de ordenação Insertion Sort é um algoritmo que percorre a lista da esquerda para a direita, e à medida que avança, insere cada elemento na posição correta entre os elementos já ordenados.

Aqui está o funcionamento passo a passo do algoritmo Insertion Sort para o vetor [3, 6, 2, 5, 4, 3, 7, 1, 10^9]:

Passo 1:
Começamos com o primeiro elemento (3) considerado como ordenado.
[3, 6, 2, 5, 4, 3, 7, 1, 10^9]

Passo 2:
Avançamos para o segundo elemento (6) e o comparamos com o elemento anterior (3). Como 6 é maior que 3, não fazemos nenhuma alteração.
[3, 6, 2, 5, 4, 3, 7, 1, 10^9]

Passo 3:
Avançamos para o terceiro elemento (2) e o comparamos com o elemento anterior (6). Como 2 é menor que 6, deslocamos o elemento 6 uma posição para a direita e inserimos o 2 em sua posição correta.
[2, 3, 6, 5, 4, 3, 7, 1, 10^9]

Passo 4:
Avançamos para o quarto elemento (5) e o comparamos com o elemento anterior (6). Como 5 é menor que 6, deslocamos o elemento 6 uma posição para a direita. Em seguida, comparamos o 5 com o elemento anterior (3) e vemos que é maior que 3. Inserimos o 5 entre o 3 e o 6.
[2, 3, 5, 6, 4, 3, 7, 1, 10^9]

Passo 5:
Avançamos para o quinto elemento (4) e o comparamos com o elemento anterior (6). Como 4 é menor que 6, deslocamos o elemento 6 uma posição para a direita. Em seguida, comparamos o 4 com o elemento anterior (5) e vemos que é menor que 5. Deslocamos o elemento 5 uma posição para a direita. Finalmente, inserimos o 4 na posição correta entre o 3 e o 5.
[2, 3, 4, 5, 6, 3, 7, 1, 10^9]

Passo 6:
Avançamos para o sexto elemento (3) e o comparamos com o elemento anterior (6). Como 3 é menor que 6, deslocamos o elemento 6 uma posição para a direita. Em seguida, comparamos o 3 com o elemento anterior (5) e vemos que é menor que 5. Deslocamos o elemento 5 uma posição para a direita. Continuamos comparando o 3 com os elementos anteriores (4, 3) e deslocando-os quando necessário. Finalmente, inserimos o 3 na posição correta entre o 2 e o 4.
[2, 3, 3, 4, 5, 6, 7, 1, 10^9]

Passo 7:
Avançamos para

 o sétimo elemento (7) e o comparamos com o elemento anterior (6). Como 7 é maior que 6, não fazemos nenhuma alteração.
[2, 3, 3, 4, 5, 6, 7, 1, 10^9]

Passo 8:
Avançamos para o oitavo elemento (1) e o comparamos com o elemento anterior (7). Como 1 é menor que 7, deslocamos o elemento 7 uma posição para a direita. Continuamos comparando o 1 com os elementos anteriores (6, 5, 4, 3, 3, 2) e deslocando-os quando necessário. Finalmente, inserimos o 1 na posição correta no início da lista.
[1, 2, 3, 3, 4, 5, 6, 7, 10^9]

Passo 9:
Avançamos para o nono elemento (10^9) e o comparamos com o elemento anterior (7). Como 10^9 é maior que 7, não fazemos nenhuma alteração.
[1, 2, 3, 3, 4, 5, 6, 7, 10^9]

Agora, o vetor está completamente ordenado em ordem crescente usando o algoritmo Insertion Sort.

### d. MergeSort
    i. vetor para ilustrar: decrescente

O algoritmo de ordenação Merge Sort é um algoritmo de divisão e conquista que divide repetidamente a lista em metades menores, ordena cada metade separadamente e, em seguida, combina as metades ordenadas para obter a lista final ordenada.

Aqui está o funcionamento passo a passo do algoritmo Merge Sort para o vetor [7, 6, 5, 4, 3, 3, 2, 1]:

Passo 1: Divisão
Dividimos o vetor original em metades menores até que cada subvetor contenha apenas um elemento.

[7, 6, 5, 4, 3, 3, 2, 1] (Divisão 1)
[7, 6, 5, 4] | [3, 3, 2, 1] (Divisão 2)
[7, 6] | [5, 4] | [3, 3] | [2, 1] (Divisão 3)
[7] | [6] | [5] | [4] | [3] | [3] | [2] | [1] (Divisão final)

Passo 2: Combinação (Merge)
Começamos combinando os subvetores ordenados em pares, mesclando-os em um único subvetor ordenado. Em seguida, repetimos o processo para combinar os subvetores maiores, até que tenhamos a lista final ordenada.

[6, 7] | [4, 5] | [3, 3] | [1, 2] (Combinação 1)
[4, 5, 6, 7] | [1, 2, 3, 3] (Combinação 2)
[1, 2, 3, 3, 4, 5, 6, 7] (Combinação final)

Agora, o vetor está completamente ordenado em ordem crescente usando o algoritmo Merge Sort.

### e. QuickSort (sem randomização de pivô)
    i. vetor para ilustrar: crescente

Vamos ilustrar o funcionamento do algoritmo QuickSort para o vetor [1, 2, 3, 3, 4, 5, 6, 7] sem randomização do pivô. Nesse caso, vamos considerar o último elemento do vetor como pivô.

Passo 1:
Selecionamos o último elemento, 7, como pivô.
[1, 2, 3, 3, 4, 5, 6, 7]

Passo 2:
Reorganizamos os elementos do vetor de forma que todos os elementos menores que o pivô fiquem à sua esquerda, e todos os elementos maiores ou iguais ao pivô fiquem à sua direita.
[1, 2, 3, 3, 4, 5, 6, 7]

Passo 3:
Dividimos o vetor em duas partes com base na posição final do pivô.
[1, 2, 3, 3] e [4, 5, 6, 7]

Passo 4:
Recursivamente aplicamos o algoritmo QuickSort nas duas partes do vetor.
[1, 2, 3, 3] (parte esquerda)
[4, 5, 6, 7] (parte direita)

Passo 5:
Para a parte esquerda, selecionamos o último elemento, 3, como pivô.
[1, 2, 3, 3]

Passo 6:
Reorganizamos os elementos menores que o pivô à sua esquerda e os maiores ou iguais ao pivô à sua direita.
[1, 2, 3, 3]

Passo 7:
Dividimos a parte esquerda do vetor em duas partes com base na posição final do pivô.
[1, 2] e [3, 3]

Passo 8:
Recursivamente aplicamos o algoritmo QuickSort nas duas partes da parte esquerda do vetor.
[1, 2] (parte esquerda)
[3, 3] (parte direita)

Passo 9:
Para a parte esquerda da parte esquerda do vetor, selecionamos o primeiro elemento, 1, como pivô.
[1, 2]

Passo 10:
Reorganizamos os elementos menores que o pivô à sua esquerda e os maiores ou iguais ao pivô à sua direita.
[1, 2]

Neste ponto, a parte esquerda da parte esquerda do vetor está completamente ordenada.

Passo 11:
Para a parte direita da parte esquerda do vetor, selecionamos o último elemento, 3, como pivô.
[3, 3]

Passo 12:
Reorganizamos os elementos menores que o pivô à sua esquerda e os maiores ou iguais ao pivô à sua direita.
[3, 3]

Neste ponto, a parte direita da parte esquerda do vetor está completamente ordenada.

Passo 13:
Para a parte direita do vetor, selecionamos o último elemento, 7, como pivô.
[4, 5, 6, 7]

Passo 14:
Reorganizamos os elementos menores que o pivô à sua esquerda e os maiores ou iguais ao pivô à sua direita.
[4, 5, 6, 7]

Neste ponto, a parte direita do

 vetor está completamente ordenada.

Ao finalizar todos os passos, temos o vetor [1, 2, 3, 3, 4, 5, 6, 7] completamente ordenado usando o algoritmo QuickSort.