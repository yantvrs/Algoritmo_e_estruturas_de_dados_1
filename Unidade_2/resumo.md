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

# Bubble Sort

O Bubble Sort é um algoritmo de ordenação simples e ineficiente. Ele compara repetidamente pares de elementos adjacentes e os troca se estiverem na ordem errada. O algoritmo percorre a lista várias vezes, movendo o maior elemento para a posição correta a cada iteração. Esse processo é repetido até que a lista esteja completamente ordenada.

## Funcionamento do Bubble Sort

O Bubble Sort segue os seguintes passos:

1. Começa percorrendo a lista do início ao fim.
2. Compara o elemento atual com o próximo elemento.
3. Se o elemento atual for maior do que o próximo elemento, realiza a troca.
4. Continua percorrendo a lista até o penúltimo elemento.
5. Repete os passos 2 a 4 para cada iteração até que nenhuma troca seja realizada em uma iteração completa, indicando que a lista está ordenada.


## Complexidade de Tempo

A complexidade de tempo do Bubble Sort é O(n^2), onde n é o número de elementos na lista. Isso ocorre porque o algoritmo precisa percorrer a lista várias vezes, comparando pares de elementos adjacentes e realizando trocas. Mesmo em sua forma otimizada, o Bubble Sort ainda possui um desempenho ineficiente para listas grandes.

## Vantagens e Desvantagens

Vantagens do Bubble Sort:
- Implementação simples e fácil de entender.
- Funciona bem para listas pequenas ou quase ordenadas.

Desvantagens do Bubble Sort:
- Complexidade de tempo quadrática O(n^2).
- Desempenho ineficiente para listas grandes ou desordenadas.
- Não é estável, ou seja, a ordem relativa de elementos iguais pode ser alterada.

O Bubble Sort é considerado um dos algoritmos de ordenação mais simples, mas também um dos menos eficientes. É mais adequado para fins educacionais ou para ordenar pequenas quantidades de dados quando a simplicidade e clareza são mais importantes do que o desempenho.

# Insertion Sort

O Insertion Sort é um algoritmo de ordenação simples que consiste em construir uma lista ordenada à medida que os elementos são inseridos. Ele percorre a lista da esquerda para a direita, inserindo cada elemento na posição correta em relação aos elementos já ordenados.

## Funcionamento do Insertion Sort

1. Inicialmente, considera-se o primeiro elemento da lista como uma lista ordenada.
2. Para cada elemento subsequente, encontra-se a posição correta na lista ordenada para inserir esse elemento.
3. Move-se todos os elementos maiores do que o elemento a ser inserido uma posição à frente para abrir espaço.
4. Insere-se o elemento na posição correta.

O algoritmo continua esse processo até que todos os elementos da lista sejam inseridos em suas posições corretas.


## Complexidade de Tempo

A complexidade de tempo do Insertion Sort é O(n^2), onde n é o número de elementos na lista a ser ordenada. Isso ocorre porque, em média, o algoritmo precisa percorrer metade da lista para inserir um novo elemento.

Embora tenha uma complexidade quadrática, o Insertion Sort é eficiente para listas pequenas ou quase ordenadas, pois requer menos comparações e trocas do que outros algoritmos mais complexos.

## Vantagens e Desvantagens

Vantagens do Insertion Sort:
- Implementação simples e fácil de entender.
- Eficiente para listas pequenas ou quase ordenadas.
- Requer um número mínimo de trocas quando a lista está parcialmente ordenada.

Desvantagens do Insertion Sort:
- Desempenho relativamente lento para listas grandes ou completamente desordenadas.
- Não é estável, ou seja, a ordem relativa de elementos iguais não é preservada.
- Requer uma quantidade fixa de comparações mesmo quando a lista está parcialmente ordenada.

O Insertion Sort é um algoritmo de ordenação útil em certos cenários, especialmente quando a lista é pequena ou quase ordenada. No entanto, para listas maiores e menos ordenadas, algoritmos mais eficientes, como o Merge Sort ou Quick Sort, são geralmente preferidos.

# Merge Sort

O Merge Sort é um algoritmo de ordenação eficiente que utiliza a estratégia "dividir para conquistar". Ele divide repetidamente a lista não ordenada em sublistas menores, ordena essas sublistas e, em seguida, mescla-as para obter uma lista ordenada.

## Funcionamento do Merge Sort

O Merge Sort segue os seguintes passos:

1. Divide a lista não ordenada em duas sublistas de tamanho aproximadamente igual.
2. Recursivamente, ordena cada sublista aplicando o Merge Sort.
3. Mescla as duas sublistas ordenadas para produzir uma lista única e ordenada.

Esse processo de divisão e mesclagem é repetido até que a lista esteja completamente ordenada.


## Complexidade de Tempo

A complexidade de tempo do Merge Sort é O(n log n), onde n é o número de elementos na lista a ser ordenada. Isso ocorre porque o algoritmo divide a lista ao meio em cada etapa de recursão e combina as sublistas de forma eficiente.

O Merge Sort é conhecido por sua eficiência em listas grandes e é amplamente utilizado na prática. Comparado a algoritmos de complexidade quadrática, como o Bubble Sort ou o Insertion Sort, o Merge Sort apresenta um desempenho muito melhor.

## Vantagens e Desvantagens

Vantagens do Merge Sort:
- Eficiente para listas grandes e desordenadas.
- Complexidade de tempo garantida de O(n log n).
- É um algoritmo estável, preservando a ordem relativa de elementos iguais.

Desvantagens do Merge Sort:
- Requer memória adicional para a mesclagem de sublistas.
- Implementação mais complexa do que algoritmos de ordenação mais simples, como o Bubble Sort.

O Merge Sort é um algoritmo poderoso de ordenação, especialmente para grandes conjuntos de dados. Sua complexidade de tempo garantida e estabilidade o tornam uma escolha popular em muitas aplicações.

# Quick Sort

O Quick Sort é um algoritmo de ordenação eficiente que utiliza a estratégia "dividir para conquistar". Ele seleciona um elemento chamado de "pivô" e rearranja a lista de forma que todos os elementos menores que o pivô estejam à sua esquerda, e os elementos maiores estejam à sua direita. Em seguida, aplica recursivamente o Quick Sort nas duas sublistas resultantes.

## Funcionamento do Quick Sort

O Quick Sort segue os seguintes passos:

1. Seleciona um elemento da lista como pivô. Existem várias estratégias para escolher o pivô, como selecionar o primeiro elemento, o último elemento ou um elemento aleatório.
2. Reorganiza a lista de forma que todos os elementos menores que o pivô estejam à sua esquerda e todos os elementos maiores estejam à sua direita.
3. Aplica o Quick Sort recursivamente nas sublistas à esquerda e à direita do pivô.

Esse processo é repetido até que cada sublista contenha apenas um elemento, resultando em uma lista totalmente ordenada.


## Complexidade de Tempo

A complexidade de tempo do Quick Sort é em média O(n log n), onde n é o número de elementos na lista a ser ordenada. No entanto, em pior caso, o Quick Sort pode ter uma complexidade quadrática de O(n^2) quando o pivô escolhido não divide a lista de forma balanceada. Para evitar isso, são usadas técnicas como escolher o pivô de forma aleatória ou usar o método de mediana de três.

O Quick Sort é amplamente utilizado devido à sua eficiência média e desempenho geralmente superior a outros algoritmos de ordenação, como o Bubble Sort ou o Insertion Sort.

## Vantagens e Desvantagens

Vantagens do Quick Sort:
- Eficiente para listas grandes e desordenadas.
- Complexidade de tempo média de O(n log n).
- Em geral, possui um desempenho superior a outros algoritmos de ordenação.

Desvantagens do Quick Sort:
- A complexidade de tempo pior caso pode ser O(n^2) se não forem aplicadas estratégias para escolher o pivô.
- Não é um algoritmo estável, ou seja, a ordem relativa de elementos iguais não é necessariamente preservada.

O Quick Sort é um dos algoritmos de ordenação mais populares e amplamente utilizados. Sua eficiência

# Counting Sort

O Counting Sort é um algoritmo de ordenação eficiente utilizado para ordenar elementos quando o intervalo dos valores é conhecido e relativamente pequeno. Ele conta o número de ocorrências de cada elemento na lista original e, em seguida, reconstrói a lista ordenada com base nas contagens acumuladas.

## Funcionamento do Counting Sort

O Counting Sort segue os seguintes passos:

1. Encontra o valor máximo (max) na lista original.
2. Cria um array de contagem (count) de tamanho (max+1), inicializado com zeros.
3. Conta o número de ocorrências de cada elemento na lista original e armazena as contagens no array de contagem.
4. Calcula as contagens acumuladas, onde cada elemento do array de contagem é a soma das contagens anteriores.
5. Cria um array auxiliar (sorted) do mesmo tamanho que a lista original.
6. Percorre a lista original e, para cada elemento, insere-o na posição correta no array auxiliar com base nas contagens acumuladas.
7. A lista ordenada é obtida a partir do array auxiliar.

## Pseudocódigo

Aqui está o pseudocódigo do algoritmo Counting Sort:


## Complexidade de Tempo

A complexidade de tempo do Counting Sort é O(n + k), onde n é o número de elementos na lista original e k é o intervalo dos valores. O Counting Sort tem um desempenho linear, tornando-o um algoritmo eficiente em muitos cenários, especialmente quando o intervalo dos valores é relativamente pequeno.

## Vantagens e Desvantagens

Vantagens do Counting Sort:
- Eficiente para listas com intervalo de valores pequeno.
- Desempenho linear O(n + k).
- Não requer comparações entre os elementos.

Desvantagens do Counting Sort:
- Requer memória adicional para o array de contagem.
- Não é adequado para listas com valores negativos ou intervalos muito grandes.
- Não é um algoritmo estável, ou seja, a ordem relativa de elementos iguais não é necessariamente preservada.

O Counting Sort é uma opção eficiente para ordenação quando o intervalo dos valores é conhecido e limitado. É amplamente utilizado em aplicações onde o desempenho é crítico e o intervalo dos valores permite um uso eficiente de memória.

# Árvore Binária de Busca

Uma árvore binária de busca é uma estrutura de dados em forma de árvore, na qual cada nó possui no máximo dois filhos. Ela possui a propriedade de que para cada nó, todos os valores nos filhos da subárvore esquerda são menores do que o valor do nó, e todos os valores nos filhos da subárvore direita são maiores.

## Estrutura de um Nó

Cada nó em uma árvore binária de busca contém as seguintes informações:
- Valor do nó.
- Ponteiro para o filho esquerdo.
- Ponteiro para o filho direito.

## Funções da Árvore Binária de Busca

### Add (Inserir)

A função `Add` é utilizada para inserir um novo valor na árvore. Ela percorre a árvore de forma recursiva, comparando o valor a ser inserido com o valor do nó atual. Se o valor for menor, ele é inserido na subárvore esquerda; se for maior, ele é inserido na subárvore direita. Se o valor já existir na árvore, ele não será inserido novamente.

### Search (Buscar)

A função `Search` é utilizada para buscar um valor na árvore. Ela percorre a árvore de forma recursiva, comparando o valor buscado com o valor do nó atual. Se o valor for igual, o nó é encontrado e a função retorna verdadeiro. Se o valor for menor, a busca continua na subárvore esquerda; se for maior, a busca continua na subárvore direita. Se o valor não for encontrado em nenhum nó, a função retorna falso.

### Min (Mínimo) e Max (Máximo)

As funções `Min` e `Max` são utilizadas para encontrar o valor mínimo e o valor máximo na árvore, respectivamente. O valor mínimo é encontrado percorrendo a árvore pela subárvore esquerda até encontrar um nó sem filho esquerdo. O valor máximo é encontrado percorrendo a árvore pela subárvore direita até encontrar um nó sem filho direito.

### PrintPre (Pré-ordem), PrintIn (Em-ordem) e PrintPos (Pós-ordem)

As funções `PrintPre`, `PrintIn` e `PrintPos` são utilizadas para imprimir os valores da árvore em diferentes ordens de percurso: pré-ordem, em-ordem e pós-ordem, respectivamente. 

- Pré-ordem: imprime o valor do nó atual antes de imprimir os valores dos filhos (raiz-esquerda-direita).
- Em-ordem: imprime o valor do nó atual entre a impressão dos valores dos filhos (esquerda-raiz-direita), resultando em uma impressão em ordem crescente quando a árvore está ordenada.
- Pós-ordem: imprime o valor do nó atual após a impressão dos valores dos filhos (esquerda-direita-raiz).

### PrintLevels (Impressão por níveis)

A função `PrintLevels` é utilizada para imprimir os valores da árvore por níveis, ou seja, imprime os valores dos nós de cada nível antes de passar para o próximo nível. Ela utiliza uma fila para percorrer os nós por níveis.

### Height (Altura)

A função `Height` é utilizada para calcular a altura da árvore, ou seja, o número máximo de arestas entre a raiz e as folhas. Ela percorre a árvore de forma recursiva, calculando a altura das subárvores esquerda e direita, e retorna o máximo entre essas alturas mais 1.

### Remove (Remover)

A função `Remove` é utilizada para remover um valor específico da árvore. Ela percorre a árvore de forma recursiva, buscando o nó que contém o valor a ser removido. Existem três casos possíveis:

1. O nó não tem filhos: basta remover o nó da árvore.
2. O nó tem apenas um filho: substitui o nó pelo seu único filho.
3. O nó tem dois filhos: encontra o sucessor (o menor valor na subárvore direita) ou o predecessor (o maior valor na subárvore esquerda), substitui o valor do nó a ser removido pelo valor do sucessor ou predecessor, e então remove o sucessor ou predecessor da subárvore correspondente.

## Complexidade de Tempo

A complexidade de tempo das operações em uma árvore binária de busca depende da altura da árvore. Em uma árvore balanceada, a altura é aproximadamente logarítmica em relação ao número de nós, o que resulta em um desempenho eficiente para as operações. No entanto, em uma árvore desbalanceada, a altura pode ser linear, o que leva a um desempenho pior.

## Vantagens e Desvantagens

Vantagens da Árvore Binária de Busca:
- Inserção, busca e remoção eficientes em árvores balanceadas.
- Possibilidade de imprimir os valores em diferentes ordens e por níveis.
- Estrutura de dados flexível e versátil.

Desvantagens da Árvore Binária de Busca:
- A complexidade de tempo pode ser pior caso a árvore fique desbalanceada.
- Requer a implementação correta das operações para garantir a propriedade da árvore binária de busca.

A árvore binária de busca é uma estrutura de dados fundamental em ciência da computação, utilizada em diversos algoritmos e aplicações. Sua eficiência depende da correta implementação e do balanceamento da árvore.
