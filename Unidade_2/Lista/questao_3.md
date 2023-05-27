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