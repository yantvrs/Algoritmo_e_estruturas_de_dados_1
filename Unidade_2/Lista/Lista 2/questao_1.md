# Questão 1

## Responda às seguintes perguntas:
## a) Qual a diferença entre uma árvore binária e uma árvore de busca binária (BST)?

Uma árvore binária é uma estrutura de dados composta por nós que podem ter até dois nós filhos, enquanto uma árvore de busca binária (BST) é uma árvore binária especial em que os valores dos nós são organizados de forma que todos os valores à esquerda são menores que o valor do nó em questão, e todos os valores à direita são maiores. Essa restrição faz com que a árvore de busca binária seja eficiente para operações de busca, inserção e remoção de elementos, proporcionando uma ordenação rápida e eficaz.

## b) O que significam os níveis de uma árvore?

Os níveis de uma árvore referem-se às camadas horizontais da árvore, que são contadas a partir da raiz. O nível 0 representa a raiz da árvore, e a partir dela, cada nível subsequente representa os nós que estão a uma distância de um passo da raiz. Os nós em um mesmo nível estão no mesmo nível de profundidade da árvore. Por exemplo, os nós diretamente conectados à raiz estão no nível 1, os nós conectados aos nós do nível 1 estão no nível 2, e assim por diante. Os níveis são úteis para entender a estrutura e organização da árvore, bem como para determinar a altura da árvore, que é o número máximo de níveis presentes na árvore.

## c) O que é uma BST completa?

Uma BST completa é uma árvore de busca binária em que todos os níveis, exceto possivelmente o último, estão totalmente preenchidos com nós. No último nível, os nós estão posicionados o mais à esquerda possível. Essa estrutura garante que a árvore esteja balanceada e que a ordem dos valores dos nós, quando percorridos em ordem, esteja em crescente. Além disso, uma BST completa possui uma altura mínima, proporcionando operações eficientes de busca, inserção e remoção de elementos.

## d) Como calculamos o fator de balanceamento de uma árvore?

O fator de balanceamento de uma árvore é obtido pela diferença entre as alturas das subárvores esquerda e direita de um nó. É usado para verificar se a árvore está balanceada ou não. Em uma árvore AVL, por exemplo, o fator de balanceamento de cada nó deve estar entre -1, 0 e 1. Se estiver fora desses limites, são realizadas operações de rotação ou reequilíbrio para restaurar o equilíbrio da árvore. O cálculo do fator de balanceamento é essencial para manter uma estrutura eficiente e otimizada em árvores binárias balanceadas.

## e) Qual a altura de uma árvore balanceada contendo n nós? Mostre matematicamente.

A altura de uma árvore balanceada com n nós pode ser calculada usando a fórmula da função logarítmica na base 2: altura = log₂(n + 1). Essa fórmula é aplicável para árvores balanceadas completas, onde todos os níveis, exceto o último, estão totalmente preenchidos. Em uma árvore completa, o número de nós em cada nível aumenta exponencialmente em potências de 2. No entanto, em outros tipos de árvores balanceadas, a altura pode variar dependendo da estrutura específica da árvore.

## f) Qual a complexidade de tempo das seguintes operações em uma árvore completamente desbalanceada? Add(value int), Search(value int) bool, Min() int, Max() int, PrintPre(), PrintIn(), PrintPos(), Height() int, Remove(value int) *BstNode

Em uma árvore completamente desbalanceada, onde todos os nós têm apenas um filho à direita ou à esquerda, as operações possuem complexidade de tempo que depende diretamente da altura da árvore. Vamos analisar cada operação individualmente:

- Add(value int): A complexidade de tempo para adicionar um valor em uma árvore desbalanceada é proporcional à altura da árvore, ou seja, O(h), onde h é a altura da árvore.

- Search(value int): A complexidade de tempo para buscar um valor em uma árvore desbalanceada também é proporcional à altura da árvore, ou seja, O(h).

- Min(): A complexidade de tempo para encontrar o valor mínimo em uma árvore desbalanceada é O(h), onde h é a altura da árvore.

- Max(): A complexidade de tempo para encontrar o valor máximo em uma árvore desbalanceada também é O(h), onde h é a altura da árvore.

- PrintPre(): A complexidade de tempo para imprimir os elementos da árvore em pré-ordem é proporcional ao número de nós da árvore, ou seja, O(n), onde n é o número de nós.

- PrintIn(): A complexidade de tempo para imprimir os elementos da árvore em ordem é proporcional ao número de nós da árvore, ou seja, O(n).

- PrintPos(): A complexidade de tempo para imprimir os elementos da árvore em pós-ordem é proporcional ao número de nós da árvore, ou seja, O(n).

- Height(): A complexidade de tempo para calcular a altura de uma árvore desbalanceada é O(n), onde n é o número de nós.

- Remove(value int): A complexidade de tempo para remover um valor de uma árvore desbalanceada é proporcional à altura da árvore, ou seja, O(h), onde h é a altura da árvore.

Portanto, as operações em uma árvore completamente desbalanceada têm complexidade de tempo proporcional à altura da árvore ou ao número de nós, dependendo do caso específico.

## g) Qual a complexidade de tempo das seguintes operações em uma árvore completa? Add(value int), Search(value int) bool, Min() int, Max() int, PrintPre(), PrintIn(), PrintPos(), Height() int, Remove(value int) *BstNode

Em uma árvore completa, onde todos os níveis, exceto o último, estão completamente preenchidos, as operações têm complexidade de tempo que depende da altura da árvore, representada por "h". Vamos analisar cada operação individualmente:

- Add(value int): A complexidade de tempo para adicionar um valor em uma árvore completa é O(h), onde "h" é a altura da árvore. Isso ocorre porque a inserção é realizada seguindo a estrutura da árvore completa.

- Search(value int): A complexidade de tempo para buscar um valor em uma árvore completa é O(h), onde "h" é a altura da árvore. Como a árvore completa mantém uma ordem específica, a busca pode ser realizada de forma eficiente percorrendo a árvore da raiz até o nó desejado.

- Min(): A complexidade de tempo para encontrar o valor mínimo em uma árvore completa é O(h), onde "h" é a altura da árvore. Como a árvore completa segue uma ordem específica, o valor mínimo estará localizado no nó mais à esquerda da árvore.

- Max(): A complexidade de tempo para encontrar o valor máximo em uma árvore completa também é O(h), onde "h" é a altura da árvore. O valor máximo estará localizado no nó mais à direita da árvore.

- PrintPre(): A complexidade de tempo para imprimir os elementos da árvore em pré-ordem é O(n), onde "n" é o número de nós na árvore. Isso ocorre porque cada nó precisa ser visitado uma vez.

- PrintIn(): A complexidade de tempo para imprimir os elementos da árvore em ordem é O(n), onde "n" é o número de nós na árvore. Como a árvore completa segue uma ordem específica, a impressão em ordem requer a visita de todos os nós.

- PrintPos(): A complexidade de tempo para imprimir os elementos da árvore em pós-ordem é O(n), onde "n" é o número de nós na árvore. Assim como nas operações anteriores, todos os nós precisam ser visitados para realizar a impressão em pós-ordem.

- Height(): A complexidade de tempo para calcular a altura de uma árvore completa é O(log₂(n)), onde "n" é o número de nós na árvore. A altura de uma árvore completa é limitada por logaritmo de base 2 do número de nós, devido à estrutura balanceada da árvore.

- Remove(value int): A complexidade de tempo para remover um valor de uma árvore completa é O(h), onde "h" é a altura da árvore. A remoção requer a busca pelo valor desejado e a subsequente reestruturação da árvore para manter a sua completude.

Em resumo, as operações em uma árvore completa possuem complexidade de tempo que depende da altura da árvore (O(h)), do número de nós (O(n)), ou do logaritmo de base 2 do número de nós (O(log₂(n))).

## h) Explique os possíveis casos de remoção em uma BST. Como deve-se proceder em cada caso?

Ao remover um nó de uma árvore de busca binária (BST), existem três casos possíveis, dependendo das características do nó que está sendo removido e seus filhos:

1. Caso 1: O nó não possui filhos (é uma folha).
   Nesse caso, a remoção é simples. Basta desconectar o nó da árvore, atualizando os ponteiros adequados.

2. Caso 2: O nó possui apenas um filho.
   Nesse caso, a remoção também é relativamente simples. O filho do nó que está sendo removido substitui o nó na árvore. Os ponteiros são ajustados para refletir essa mudança.

3. Caso 3: O nó possui dois filhos.
   Esse é o caso mais complexo. Ao remover um nó com dois filhos, é necessário encontrar um sucessor ou predecessor para substituí-lo. O sucessor é o menor valor maior que o nó removido, e o predecessor é o maior valor menor que o nó removido.

   - Uma abordagem comum é encontrar o sucessor. Isso pode ser feito encontrando o nó mais à esquerda na subárvore direita do nó a ser removido (ou o nó mais à direita na subárvore esquerda para encontrar o predecessor). O sucessor é então colocado no lugar do nó removido, e os ponteiros são ajustados adequadamente.
   
   - Outra abordagem é encontrar o predecessor. Nesse caso, em vez de ir para a subárvore direita, você iria para a subárvore esquerda e procuraria o nó mais à direita.

Independentemente do caso, é importante garantir que a propriedade de busca binária seja mantida após a remoção. Isso significa que os valores menores devem estar à esquerda do nó e os valores maiores devem estar à direita.

Após a substituição e ajuste dos ponteiros, é possível que a árvore precise ser rebalanceada, dependendo da estrutura específica da BST (por exemplo, em uma árvore AVL ou rubro-negra). Isso é feito para manter a árvore balanceada e preservar sua eficiência.

É importante notar que a abordagem exata para a remoção pode variar dependendo da implementação específica da BST e dos requisitos do problema em questão.

